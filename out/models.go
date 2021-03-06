package out

import "github.com/crdant/cf-route-resource"

type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	Application string   `json:"application"`
	Create      []string `json:"create"`
	RandomPort  bool     `json:"randomPort"`
	Map         []string `json:"map"`
	Unmap       []string `json:"unmap"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
