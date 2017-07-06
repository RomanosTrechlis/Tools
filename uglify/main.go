package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"flag"
	"os"
	"regexp"
)


var (
	input string
	output string
	verbose bool
)

func init() {
	flag.StringVar(&input, "i", "style.css", "input file")
	flag.StringVar(&output, "o", "style.min.css", "output file")
	flag.BoolVar(&verbose, "verbose", false, "be verbose")
}

func main() {
	flag.Parse()
	if strings.HasSuffix(input, ".css") {
		content, _ := ioutil.ReadFile(input)
		r := cssUglify(string(content))
		if verbose {
			fmt.Print(string(r))
		}
		ioutil.WriteFile(output, r , os.ModePerm)
	} else {
		fmt.Print("Only css is supported for now.")
	}
}


func cssUglify(c string) []byte {
	n := strings.Count(c, ";\r\n}")
	c = strings.Replace(c, ";\r\n}", "}", n)

	n = strings.Count(c, "\r\n")
	c = strings.Replace(c, "\r\n", "", n)

	re := regexp.MustCompile("[ ]\\{")
	c = re.ReplaceAllString(c, "{")

	re = regexp.MustCompile("/\\*[A-Za-z0-9;:\\-%\\.<>!@#$^&\\*()\\[\\] ]*\\*/")
	c = re.ReplaceAllString(c, "")

	re = regexp.MustCompile("\\{[ ]*")
	c = re.ReplaceAllString(c, "{")

	re = regexp.MustCompile("\\}[ ]*")
	c = re.ReplaceAllString(c, "}")

	re = regexp.MustCompile(";[ ]*")
	c = re.ReplaceAllString(c, ";")

	re = regexp.MustCompile(",[ ]*")
	c = re.ReplaceAllString(c, ",")

	re = regexp.MustCompile(":[ ]*")
	c = re.ReplaceAllString(c, ":")

	return []byte(c)
}
