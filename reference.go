package dipso

type ReferenceResponse struct {
	Status		Status		`json:"Status"`
	Blogs		[]string	`json:"Blog"`
	Books		[]Book		`json:"Books"`
	Vineyards	[]Vineyard	`json:"Vineyards"`
}

type Book struct {
	Id			string		`json:"Id"`
	Title		string		`json:"Title"`
	Articles	[]Article	`json:"Articles"`
}

type Article struct {
	Id			string		`json:"Id"`
	Title		string		`json:"Title"`
	Abstract	string		`json:"Abstract"`
	Url			string		`json:"Url"`
	Footnotes	[]Footnote	`json:"Footnotes"`
}

type Footnote struct {
	Id			string		`json:"Id"`
	Title		string		`json:"Title"`
	Url			string		`json:"Url"`
}
