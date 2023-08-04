package data

type Package struct {
	Name         string
	License      string
	Dependencies []*Package
}

type Repo struct {
	Name         string
	License      string
	Dependencies []*Package
	Type         string
}

type SBOM struct {
	GeneratedDate string
	Repos         []*Repo
}
