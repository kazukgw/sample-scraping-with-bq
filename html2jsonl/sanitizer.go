package main

import "github.com/microcosm-cc/bluemonday"

var defaultSanitizer *bluemonday.Policy

type Sanitizer struct {
	*bluemonday.Policy
}

func NewSanitizer() *Sanitizer {
	p := bluemonday.NewPolicy()
	for _, t := range tags() {
		p = p.AllowAttrs(attrs()...).OnElements(t)
	}
	return &Sanitizer{p}
}

func (s *Sanitizer) Sanitize(html string) string {
	return s.Policy.Sanitize(html)
}

func attrs() []string {
	return []string{
		"alt",
		"class",
		"code",
		"codebase",
		"color",
		"content",
		"contenteditable",
		"contextmenu",
		"controls",
		"coords",
		"data",
		"data-*",
		"dirname",
		"enctype",
		"for",
		"form",
		"href",
		"hreflang",
		"http-equiv",
		"id",
		"itemprop",
		"keytype",
		"kind",
		"label",
		"lang",
		"language",
		"list",
		"manifest",
		"media",
		"method",
		"min",
		"multiple",
		"muted",
		"name",
		"novalidate",
		"open",
		"optimum",
		"pattern",
		"ping",
		"placeholder",
		"poster",
		"preload",
		"radiogroup",
		"readonly",
		"rel",
		"required",
		"reversed",
		"rows",
		"rowspan",
		"sandbox",
		"scope",
		"scoped",
		"seamless",
		"selected",
		"shape",
		"size",
		"sizes",
		"span",
		"spellcheck",
		"src",
		"start",
		"step",
		// "style",
		"summary",
		"tabindex",
		"target",
		"title",
		"type",
		"usemap",
		"value",
		"wrap",
	}
}

func tags() []string {
	return []string{"a",
		"address",
		"area",
		"article",
		"aside",
		"b",
		"base",
		"basefont",
		"blockquote",
		"body",
		"br",
		"button",
		"cite",
		"code",
		"col",
		"colgroup",
		"datalist",
		"dd",
		"del",
		"details",
		"dfn",
		"dialog",
		"dir",
		"div",
		"dl",
		"dt",
		"em",
		"embed",
		"fieldset",
		"figure",
		"font",
		"footer",
		"form",
		"h1",
		"h2",
		"h3",
		"h4",
		"h5",
		"h6",
		"head",
		"header",
		"hr",
		"html",
		// "i",
		// "img",
		"input",
		"ins",
		"kbd",
		"keygen",
		"label",
		"legend",
		"li",
		// "link",
		"main",
		"map",
		"mark",
		"menu",
		"menuitem",
		// "meta",
		"meter",
		"nav",
		"ol",
		"optgroup",
		"option",
		"output",
		"p",
		"param",
		"pre",
		"progress",
		"q",
		"rp",
		"rt",
		"ruby",
		"s",
		"samp",
		// "script",
		"section",
		"select",
		"small",
		"source",
		"span",
		"strike",
		"strong",
		// "style",
		"sub",
		"summary",
		"sup",
		"table",
		"tbody",
		"td",
		"textarea",
		"tfoot",
		"th",
		"thead",
		"time",
		"title",
		"tr",
		"track",
		"tt",
		"u",
		"ul",
		"var",
	}
}
