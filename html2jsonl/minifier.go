package main

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/xml"
)

type Minifier struct {
	M *minify.M
}

func NewMinifier() *Minifier {
	m := minify.New()
	m.AddFunc("text/xml", xml.Minify)
	m.Add("text/xml", &xml.Minifier{
		KeepWhitespace: false,
	})
	return &Minifier{m}
}

func (m *Minifier) Minify(htmlstr string) (string, error) {
	return m.M.String("text/xml", htmlstr)
}
