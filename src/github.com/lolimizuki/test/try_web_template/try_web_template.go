package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"
	// "net/http"

)

type (
	Person struct {
		Name  string
		Mails []string
	}
)

func main() {
	// output()
	// if_else_if()
	// template_func()
	// embed_templates()
	my_sub_template_test()
}

func output() {
	t := template.New("hello_template")
	ts, err := t.Parse(`Hello ~~~ {{.Name}}, 
your mails are 
{{range .Mails}} #: {{.}} 
{{end}}
    `)

	if err != nil {
		fmt.Println(err)
		return
	}

	p := Person{
		Name:  "Mizuki",
		Mails: []string{"lolimizuki@gmail.com", "inabamizuki@hotmail.com", "inabamizuki@umail.hinet.net"},
	}

	ts.Execute(os.Stdout, p)
}

func if_else_if() {
	tWitoutValue := template.Must(template.New("try_if_no_value").Parse("There is no value\n{{if ``}}no ouput{{end}}\n"))
	tWitoutValue.Execute(os.Stdout, nil)

	tWithValue := template.Must(template.New("try_if_has_value").Parse("There is value\n{{if `ohhhhh`}}ouput{{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	rand.Seed(int64(time.Now().Nanosecond()))
	choice := rand.Int31n(2)

	var b bool
	if choice == 0 {
		b = true
	} else {
		b = false
	}

	tWithBool := template.Must(template.New("try_if_bool").Parse(`There is bool\n
{{if .}}YES!!! Get
{{else}}No No ... No Value
{{end}}
v={{.}}`))
	tWithBool.Execute(os.Stdout, b)
}

func AtToAt(argv ...interface{}) string {
	argv_0 := argv[0].(string)
	stringComponent := strings.Split(argv_0, "@")
	return stringComponent[0] + " at " + stringComponent[1]
}

func template_func() {
	mytemplate := template.New("template_func")
	mytemplate = mytemplate.Funcs(template.FuncMap{"atToAt": AtToAt})

	mytemplate = template.Must(mytemplate.Parse(`{{.|atToAt}}`))

	value := "lolimizuki@gmail.com"
	mytemplate.Execute(os.Stdout, value)
}

func embed_templates() {
	subTmpl, err := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	subTmpl.ExecuteTemplate(os.Stdout, "header", nil)
	// fmt.Println()
	subTmpl.ExecuteTemplate(os.Stdout, "content", nil)
	// fmt.Println()
	subTmpl.ExecuteTemplate(os.Stdout, "footer", nil)
	subTmpl.Execute(os.Stdout, nil)
}

func my_sub_template_test() {
	subTempl, _ := template.ParseFiles("maincontent.tmpl", "separate.tmpl")
	subTempl.ExecuteTemplate(os.Stdout, "maincontent", nil)
	subTempl.ExecuteTemplate(os.Stdout, "separate", nil)
}
