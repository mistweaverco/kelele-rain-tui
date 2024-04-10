package main

type Package struct {
	URI       string   `json:"uri"`
	DependsOn []string `json:"dependsOn,omitempty"`
}

type SymlinkFile struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type Symlink struct {
	URI   string        `json:"uri"`
	Files []SymlinkFile `json:"files"`
}

type JsonRoot struct {
	Packages []Package `json:"packages,omitempty"`
	Symlinks []Symlink `json:"symlinks,omitempty"`
}