package types

type Settings struct {
	Api       string `json:"api"`       // HTTP endpoint of go-ipfs instance
	Current   string `json:"current"`   // IPFS address of current project version (e.g. /ipfs/bafybeicz...)
	OutputDir string `json:"outputDir"` // Directory with the production files (e.g. ./build)
}
