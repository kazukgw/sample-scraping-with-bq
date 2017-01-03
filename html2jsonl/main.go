package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("==> start")
	var outputFileName, inputFileName string
	var encCount int
	flag.StringVar(&outputFileName, "o", "output.jsonl", "")
	flag.StringVar(&inputFileName, "i", "index.html", "")
	flag.IntVar(&encCount, "c", 100, "")
	flag.Parse()

	fmt.Printf(
		"input: %s\noutput: %s\ncount: %d\n",
		inputFileName,
		outputFileName,
		encCount,
	)

	fmt.Println("==> create sanitizer")
	sanit := NewSanitizer()

	fmt.Println("==> create minifier")
	mini := NewMinifier()

	fmt.Println("==> init output file")
	outputfile := InitOutFile(outputFileName)
	defer outputfile.Close()

	fmt.Println("==> init sample html")
	sample := NewHTMLSample(inputFileName, sanit, mini)

	fmt.Println("==> create json encoder")
	enc := json.NewEncoder(outputfile)
	enc.SetEscapeHTML(false)

	fmt.Println("==> create encode ...")
	for i := 0; i < encCount; i++ {
		if err := enc.Encode(sample); err != nil {
			Must(err)
		}
	}
	fmt.Println("==> done !")
}

type HTMLSample struct {
	HTML string `json:"html"`
}

func NewHTMLSample(
	filename string,
	sanit *Sanitizer,
	mini *Minifier,
) *HTMLSample {
	dat, err := ioutil.ReadFile(filename)
	Must(err)
	html := sanit.Sanitize(string(dat))
	html, err = mini.Minify(html)
	Must(err)
	return &HTMLSample{html}
}

func InitOutFile(outputfile string) *os.File {
	if Exists(outputfile) {
		err := os.Remove(outputfile)
		Must(err)
	}
	outfile, err := os.OpenFile(
		outputfile,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0755,
	)
	Must(err)
	return outfile
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
