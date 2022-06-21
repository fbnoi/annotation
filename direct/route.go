package direct

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"fbnoi.com/httprouter"
)

var (
	tplStr = `
{{- range .Methods -}}
router.{{ . }}("{{ $.Path }}", "{{ $.Name }}", "{{ $.HandleFunc }}")
{{ else }}
router.All("{{ $.Path }}", "{{ $.Name }}", "{{ $.HandleFunc }}")
{{- end -}}`
	tpl *template.Template
)

func Render(wr io.Writer, r *Route) error {
	if tpl == nil {
		tpl, _ = template.New(r.Name).Parse(tplStr)
	}
	for _, method := range r.Methods {
		if !httprouter.AllowMethod(method) {
			panic(fmt.Sprintf("Method %s not allowed", method))
		}
	}

	return tpl.Execute(wr, r)
}

type Route struct {
	Path       string
	Name       string
	Methods    []string
	HandleFunc string
	Filters    []*Filter
}

func (r *Route) GetName() string {
	return "Route"
}

type Filter struct {
	Path  string
	Name  string
	Func  string
	Order string
}

func (f *Filter) GetName() string {
	return f.Name
}

func Chains() httprouter.HandleFunc {
	return func(r *http.Request, w http.ResponseWriter, ps httprouter.Params) {

	}
}
