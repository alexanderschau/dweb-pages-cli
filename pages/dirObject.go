package pages

import (
	shell "github.com/ipfs/go-ipfs-api"
)

func ReadDirObject(dir *shell.IpfsObject) map[string]string {
	res := map[string]string{}
	for _, link := range dir.Links {
		res[link.Name] = link.Hash
	}
	return res
}

func ToDirObject(links map[string]string) []shell.ObjectLink {
	var dirLinks []shell.ObjectLink

	for name, hash := range links {
		dirLinks = append(dirLinks, shell.ObjectLink{
			Name: name,
			Hash: hash,
		})
	}

	return dirLinks
}
