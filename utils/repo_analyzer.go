package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var goRepos []string
var pythonRepos []string
var nodeRepos []string
var containerfileRepos []string
var helmRepos []string

// RepositoryAnalyzers analyses the
func RepositoryAnalyzer(path string) {
	dirAnalyzer(path)
	fmt.Println("Golang Repos:")
	fmt.Println(goRepos)
	fmt.Println("Python Repos:")
	fmt.Println(pythonRepos)
	fmt.Println("Node Repos:")
	fmt.Println(nodeRepos)
	fmt.Println("Containerfile Repos:")
	fmt.Println(containerfileRepos)
	fmt.Print("Helm Repos:")
	fmt.Println(helmRepos)
}

// dirAnalyzer() analyzes directory and tries to find known types of repositories or subrepos
func dirAnalyzer(path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {

		if e.Name() == ".git" {
			continue
		}

		fmt.Println(e.Name())

		if e.IsDir() {
			dirAnalyzer(filepath.Join(path, e.Name()))
		} else {
			if tryFindGo(e.Name()) {
				goRepos = append(goRepos, path)
				continue
			}
			if tryFindNode(e.Name()) {
				nodeRepos = append(nodeRepos, path)
				continue
			}
			if tryFindPython(e.Name()) {
				pythonRepos = append(pythonRepos, path)
				continue
			}
			if tryFindContainerfile(e.Name()) {
				containerfileRepos = append(containerfileRepos, path)
				continue
			}
			if tryFindHelmChart(e.Name()) {
				helmRepos = append(helmRepos, path)
				continue
			}
		}

	}
}

func tryFindGo(name string) bool {
	return name == "go.mod"
}

func tryFindNode(name string) bool {
	return name == "package.json"
}

func tryFindPython(name string) bool {
	return name == "requirements.txt" || name == "setup.py" || name == "requirements.txt.in" || name == "requirements.in" || name == "Pipfile" || name == "Pipfile.lock"
}

func tryFindContainerfile(name string) bool {
	return name == "Dockerfile" || name == "Containerfile"
}

func tryFindHelmChart(name string) bool {
	return name == "Chart.yaml" || name == "values.yaml" || name == ".helmignore"
}
