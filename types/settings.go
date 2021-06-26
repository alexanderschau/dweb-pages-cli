package types

type Settings struct {
	Api         string `json:"api"`         // HTTP endpoint of go-ipfs instance
	Current     string `json:"current"`     // IPFS address of current project version (e.g. /ipfs/bafybeicz...)
	CurrentIPNS bool   `json:"currentIPNS"` // If set to `true`, the current record in the settings file will not be updated.
	OutputDir   string `json:"outputDir"`   // Directory with the production files (e.g. ./build)
	Then        string `json:"then"`        // Command, which will be executed after a new CID is generated
}
