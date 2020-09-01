package tmpl

import (
	"bytes"
	tmplhtml "html/template"
	"regexp"
	"strings"
	tmpltext "text/template"
)

func ExecuteTextString(data interface{}, notificationTmpl string) (string, error) {
	if notificationTmpl == "" {
		return "", nil
	}

	tmpl := tmpltext.New("").Option("missingkey=zero")
	tmpl.Funcs(tmpltext.FuncMap(DefaultFuncs))
	tpl, err := tmpl.Parse(notificationTmpl)
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	if err := tpl.ExecuteTemplate(&buf, "wechat.text", data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

type FuncMap map[string]interface{}

var DefaultFuncs = FuncMap{
	"toUpper": strings.ToUpper,
	"toLower": strings.ToLower,
	"title":   strings.Title,
	// join is equal to strings.Join but inverts the argument order
	// for easier pipelining in templates.
	"join": func(sep string, s []string) string {
		return strings.Join(s, sep)
	},
	"match": regexp.MatchString,
	"safeHtml": func(text string) tmplhtml.HTML {
		return tmplhtml.HTML(text)
	},
	"reReplaceAll": func(pattern, repl, text string) string {
		re := regexp.MustCompile(pattern)
		return re.ReplaceAllString(text, repl)
	},
	"stringSlice": func(s ...string) []string {
		return s
	},
}
