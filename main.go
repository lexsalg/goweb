package main

import (
	"log"
	"net/http"
	"text/template"
)

type Curso struct {
	Nombre   string
	Duracion int
}

type Usuario struct {
	Username      string
	Edad          int
	Activo        bool
	Administrador bool
	Tags          []string
	Cursos        []Curso
}

func (u Usuario) TienePermisoAdmin() bool {
	return u.Activo && u.Administrador
}

func (u Usuario) EsAdmin(llave string) bool {
	return u.Administrador && llave == "si"
}

func hola() string {
	return "hola desde uan funcion"
}

func suma(a, b int) int {
	return a + b
}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//tpl, err := template.New("hola").Parse("hola alexis tpl")
		//tpl, err := template.ParseFiles("templates/index.html") //parse files como funcion

		funciones := template.FuncMap{
			"hola": hola,
			"suma": suma,
		}

		tpl, err := template.New("index.html").Funcs(funciones).
			ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html") //parse files como metodo
		if err != nil {
			panic(err)
		}
		tags := []string{"go", "python", "c#", "c++", "java"}
		cursos := []Curso{
			{"python", 1},
			{"jave", 2},
			{"go", 3},
		}
		usuario := Usuario{
			Username:      "Alexis",
			Edad:          22,
			Activo:        true,
			Administrador: true,
			Tags:          tags,
			Cursos:        cursos,
		}

		_ = tpl.Execute(w, usuario)

	})
	log.Println("el servidor escucha:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
