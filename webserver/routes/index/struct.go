package index

type IndexPage struct {
	Count   int      `json:"count"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Methods []Method
