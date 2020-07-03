package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"text/template"
)

const templateFile = "../dir-index.html"

// Copied from go-ipfs/core/corehttp/gateway_indexPage.go
type listingTemplateData struct {
	Listing  []directoryItem
	Size     string
	Path     string
	Breadcrumbs []breadcrumb
	BackLink string
	Hash     string
}

type directoryItem struct {
	Size string
	Name string
	Path string
}

type breadcrumb struct {
	Name string
	Path string
}

var testPath = "/ipfs/QmFooBarQXB2mzChmMeKY47C43LxUdg1NDJ5MWcKMKxDu7/a/b/c"
var testData = listingTemplateData{
	Listing: []directoryItem{{
		Size: "25 MiB",
		Name: "short-film.mov",
		Path: testPath + "/short-film.mov",
	}, {
		Size: "1 KiB",
		Name: "this-piece-of-papers-got-47-words-37-sentences-58-words-we-wanna-know.txt",
		Path: testPath + "/this-piece-of-papers-got-47-words-37-sentences-58-words-we-wanna-know.txt",
	}},
	Size: "25 MiB",
	Path: testPath,
	Breadcrumbs: []breadcrumb{{
		Name: "ipfs",
	}, {
		Name: "QmFooBarQXB2mzChmMeKY47C43LxUdg1NDJ5MWcKMKxDu7",
		Path: testPath + "/../../..",
	}, {
		Name: "a",
		Path: testPath + "/../..",
	}, {
		Name: "b",
		Path: testPath + "/..",
	}, {
		Name: "c",
		Path: testPath,
	}},
	BackLink: testPath + "/..",
	Hash:     "QmFooBazBar2mzChmMeKY47C43LxUdg1NDJ5MWcKMKxDu7",
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Ha-ha, tricked you! There are no files here!", http.StatusNotFound)
			return
		}
		listingTemplate, err := template.New("dir-index.html").Funcs(template.FuncMap{
			"iconFromExt": func(name string) string {
				return "ipfs-_blank" // place-holder
			},
			"urlEscape": func(rawUrl string) string {
				pathUrl := url.URL{Path: rawUrl}
				return pathUrl.String()
			},
		}).ParseFiles(templateFile)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse template file: %s", err), http.StatusInternalServerError)
			return
		}
		err = listingTemplate.Execute(w, &testData)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to execute template: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	if _, err := os.Stat(templateFile); err != nil {
		wd, _ := os.Getwd()
		fmt.Printf("could not open template file %q, relative to %q: %s\n", templateFile, wd, err)
		os.Exit(1)
	}
	fmt.Printf("listening on localhost:3000\n")
	http.ListenAndServe("localhost:3000", mux)
}
