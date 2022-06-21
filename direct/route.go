package direct

import (
	"fmt"
	"io"
	"text/template"

	"fbnoi.com/httprouter"
)

const ROUTE_TPL = `
{{ range method .Methods }}
router.{{ method }}("{{ .Path }}", "{{ .Name }}", "{{ .HandleFunc }}")
{{ else }}
router.All("{{ .Path }}", "{{ .Name }}", "{{ .HandleFunc }}")
{{ end }}
`

type Route struct {
	Path       string
	Name       string
	Methods    []string
	HandleFunc string
}

func (r *Route) GetName() string {
	return "Route"
}

func (r *Route) Render(wr io.Writer) error {
	for _, method := range r.Methods {
		if !httprouter.AllowMethod(method) {
			panic(fmt.Sprintf("Method %s not allowed", method))
		}
	}
	tpl, err := template.New(r.Name).Parse(ROUTE_TPL)
	if err != nil {
		return err
	}

	return tpl.Execute(wr, r)
}
