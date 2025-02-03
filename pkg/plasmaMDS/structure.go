package plasmaMDS

type PlasmaMDS struct {
	ID          string       `json:"id,omitempty"`
	URL         string       `json:"url,omitempty"`
	Source      PlasmaSource `json:"source,omitempty"`
	Medium      PlasmaMedium `json:"medium,omitempty"`
	Target      PlasmaTarget `json:"target,omitempty"`
	Diagnostics Diagnostics  `json:"diagnostics,omitempty"`
	Resource    []Resource   `json:"resource,omitempty"`
}

type PlasmaSource struct {
	Name          string `json:"name,omitempty"`
	Application   string `json:"application,omitempty"`
	Specification string `json:"specification,omitempty"`
	Properties    string `json:"properties,omitempty"`
	Procedure     string `json:"procedure,omitempty"`
}

type PlasmaMedium struct {
	Name       string `json:"name,omitempty"`
	Properties string `json:"properties,omitempty"`
	Procedure  string `json:"procedure,omitempty"`
}

type PlasmaTarget struct {
	Name       string `json:"name,omitempty"`
	Properties string `json:"properties,omitempty"`
	Procedure  string `json:"procedure,omitempty"`
}

type Diagnostics struct {
	Name       string `json:"name,omitempty"`
	Properties string `json:"properties,omitempty"`
	Procedure  string `json:"procedure,omitempty"`
}

type Resource struct {
	ID       string  `json:"id,omitempty"`
	URL      string  `json:"url,omitempty"`
	Filetype string  `json:"filetype,omitempty"`
	Datatype string  `json:"datatype,omitempty"`
	Range    string  `json:"range,omitempty"`
	Quality  Quality `json:"quality,omitempty"`
}
