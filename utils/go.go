package utils

import (
	"fmt"

	"github.com/libretrust/compliance-check/data"
)

func analyzeGoRepo(repo *data.Repo) {
	fmt.Println("Go Repo:")
	fmt.Println(repo.Name)
}
