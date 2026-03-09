package assets

import (
	"Copycat/internal/web/environment"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

//go:embed templates/*
var TemplateFS embed.FS

//go:embed static/*
var StaticFS embed.FS

func GetTemplates(path string) *template.Template {
	if *environment.DevMode == true {
		tplPath := getAssetsDir("templates")

		if path != "" {
			tplPath += "/" + path
		}

		tpl, err := template.ParseGlob(filepath.Join(tplPath, "*.html"))

		if err != nil {
			panic(err)
		}

		return tpl
	}

	tplPath := "templates/*.html"

	if path != "" {
		tplPath = "templates/" + path + "/*.html"
	}

	tpl, err := template.ParseFS(TemplateFS, tplPath)

	if err != nil {
		panic(err)
	}

	return tpl
}

func GetStaticFS() http.FileSystem {
	if *environment.DevMode == true {
		staticPath := getAssetsDir("static")
		fmt.Println("Serving static files from:", staticPath)

		if _, err := os.Stat(staticPath); err != nil {
			panic(fmt.Sprintf("Static directory not found: %s", staticPath))
		}

		return http.Dir(staticPath)
	}

	sub, err := fs.Sub(StaticFS, "static")

	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}

func getAssetsDir(path string) string {
	_, filename, _, _ := runtime.Caller(0)
	absPath, err := filepath.Abs(filepath.Join(filepath.Dir(filename), path))

	if err != nil {
		panic(err)
	}

	return absPath
}
