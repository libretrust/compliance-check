package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/libretrust/compliance-check/data"
)

type PyPiResponse struct {
	Info struct {
		License string `json:"license"`
	} `json:"info"`
}

func analyzePythonRepo(repo *data.Repo) {
	path := repo.Name
	packages := make(map[string]string)
	scanRequirements(path+"/requirements.txt", packages)
	scanRequirements(path+"/requirements-dev.txt", packages)
	scanRequirements(path+"/requirements.in", packages)
	scanRequirements(path+"/requirements-test.in", packages)
	scanRequirements(path+"/requirements-test.txt", packages)
	for k, v := range packages {
		fmt.Println(k, v)
		fmt.Println("Querying PyPi for license... " + k)
		license := queryPypi(k)
		repo.Dependencies = append(repo.Dependencies, &data.Package{Name: k, License: license})
	}
}

func scanRequirements(filename string, packages map[string]string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Pattern to match package name (before any version specifiers like '==')
	pattern := regexp.MustCompile(`^([a-zA-Z0-9-_]+)([=><!~]+.*)?$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindStringSubmatch(line)
		if matches != nil {
			packages[matches[1]] = ""
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func queryPypi(packageName string) string {

	url := fmt.Sprintf("https://pypi.org/pypi/%s/json", packageName)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return "error getting license"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "error getting license"
	}

	var data PyPiResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return "error getting license"
	}

	return data.Info.License
}
