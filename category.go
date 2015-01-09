package dipso

type CategoryResponse struct {
	Status		Status			`json:"Status"`
	Categories	[]Category		`json:"Categories"`
}

type Category struct {
	Description		string			`json:"Description"`
	Id				int				`json:"Id"`
	Name			string			`json:"Name"`
	Refinements		[]Refinement	`json:"Refinements"`
}

type Refinement struct {
	Description		string		`json:"Description"`
	Id				int			`json:"Id"`
	Name			string		`json:"Name"`
	Url				string		`json:"Url"`
}


