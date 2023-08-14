package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/libretrust/compliance-check/data"
	"gopkg.in/yaml.v2"
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
	sbom := &data.SBOM{
		GeneratedDate: time.Now().Format(time.RFC3339),
	}

	helms := make(map[string]*data.Repo)
	containers := make(map[string]*data.Repo)
	node := make(map[string]*data.Repo)
	python := make(map[string]*data.Repo)
	goReposMap := make(map[string]*data.Repo)

	for _, repo := range helmRepos {
		helms[repo] = &data.Repo{
			Name: repo,
			Type: "helm",
		}
	}

	for _, repo := range containerfileRepos {
		containers[repo] = &data.Repo{
			Name: repo,
			Type: "containerfile",
		}
	}

	for _, repo := range nodeRepos {
		node[repo] = &data.Repo{
			Name: repo,
			Type: "node",
		}
	}

	for _, repo := range pythonRepos {
		python[repo] = &data.Repo{
			Name: repo,
			Type: "python",
		}
	}

	for _, repo := range goRepos {
		goReposMap[repo] = &data.Repo{
			Name: repo,
			Type: "go",
		}
	}

	for _, repo := range helms {
		analyzeHelmRepo(repo)
		sbom.Repos = append(sbom.Repos, repo)
	}

	for _, repo := range containers {
		analyzeContainerfileRepo(repo)
		sbom.Repos = append(sbom.Repos, repo)
	}

	for _, repo := range node {
		analyzeNodeRepo(repo)
		sbom.Repos = append(sbom.Repos, repo)
	}

	for _, repo := range python {
		analyzePythonRepo(repo)
		sbom.Repos = append(sbom.Repos, repo)
	}

	for _, repo := range goReposMap {
		analyzeGoRepo(repo)
		sbom.Repos = append(sbom.Repos, repo)
	}

	fmt.Println("SBOM:")

	yamldata, err := yaml.Marshal(sbom)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(yamldata))

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
