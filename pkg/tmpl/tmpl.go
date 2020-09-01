package tmpl

import (
	"bytes"
	"text/template"
)

func ExecuteTextString(data interface{}, notificationTmpl string) (string, error) {
	if notificationTmpl == "" {
		return "", nil
	}

	tpl, err := template.New("").Option("missingkey=zero").Parse(notificationTmpl)
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	if err := tpl.ExecuteTemplate(&buf, "wechat.text", data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
