package index

// IndexPage : JSON output
type IndexPage struct {
	Count   int      `json:"count"`
	Methods []Method `json:"methods"`
}

// Method : methods avalaible
type Method struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Methods array methods
type Methods []Method
