package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {

			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurency in dirents
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil //cancelled
	}
	defer func() { <-sema }() // release token

	// read directory
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

type ByteSize float64

type size struct {
	scale      string
	scaledSize ByteSize
}

func transform(nbytes int64) size {
	mp := []string{"KB", "MB", "GB", "TB", "PB"}
	var scale string
	var scaledSize ByteSize
	for i, k := range mp {
		shift := 1 << (10 * uint(i+1))
		value := ByteSize(nbytes) / ByteSize(shift)
		if value > 1.0 {
			scale = k
			scaledSize = value
		}
	}
	return size{scale: scale, scaledSize: scaledSize}
}

func printDiskUsage(nfiles, nbytes int64) {
	s := transform(nbytes)
	fmt.Printf("%d files %.1f %s\n", nfiles, s.scaledSize, s.scale)
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	verbose := flag.Bool("v", false, "show verbose progress messages")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree
	fileSizes := make(chan int64)
	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Cancel traversal when input is detected
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Print the results periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish
			for range fileSizes {
				// do nothing
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	// final totals
	printDiskUsage(nfiles, nbytes)
}
