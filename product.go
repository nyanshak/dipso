package dipso

type ProductResponse struct {
	Status		Status			`json:"Status"`
	Products	ProductList		`json:"Products"`
}

type ProductList struct {
	List		[]Product	`json:"List"`
	Offset		int			`json:"Offset"`
	Total		int			`json:"Total"`
	Url			string		`json:"Url"`
}


type Product struct {
	Id					int					`json:"Id"`
	Name				string				`json:"Name"`
	Url					string				`json:"Url"`
	Appelation			Appellation			`json:"Appellation"`
	Labels				[]Label				`json:"Labels"`
	Type				string				`json:"Type"`
	Varietal			Varietal			`json:"Varietal"`
	Vineyard			Vineyard			`json:"Vineyard"`
	Vintage				string				`json:"Vintage"`
	Community			Community			`json:"Community"`
	Description			string				`json:"Description"` // Only provided with partner level access, otherwise blank
	GeoLocation			GeoLocation			`json:"GeoLocation"`
	PriceMax			float64				`json:"PriceMax"`
	PriceMin			float64				`json:"PriceMin"`
	PriceRetail			float64				`json:"PriceRetail"`
	ProductAttributes	[]ProductAttribute	`json:"ProductAttributes"`
	Ratings				Ratings				`json:"Ratings"`
	Retail				Retail				`json:"Retail"`
	Vintages			Vintages			`json:"Vintages"` // Only provided with partner level access, otherwise blank
}

type Appellation struct {
	Id		int			`json:"Id"`
	Name	string		`json:"Name"`
	Url		string		`json:"Url"`
	Region	Region		`json:"Region"`
}

type Region struct {
	Id		int			`json:"Id"`
	Name	string		`json:"Name"`
	Url		string		`json:"Url"`
	Area	string		`json:"Area"`
}

type Area struct {
	Id		int			`json:"Id"`
	Name	string		`json:"Name"`
	Url		string		`json:"Url"`
	Parent	*Area		`json:"Parent"`
}

type Label struct {
	Id		string		`json:"Id"`
	Name	string		`json:"Name"`
	Url		string		`json:"Url"`
}

type Varietal struct {
	Id			int			`json:"Id"`
	Name		string		`json:"Name"`
	Url			string		`json:"Url"`
	WineType	WineType	`json:"WineType"`
}

type WineType struct {
	Id		int			`json:"Id"`
	Name	string		`json:"Name"`
	Url		string		`json:"Url"`
}

type Vineyard struct {
	Id			int			`json:"Id"`
	Name		string		`json:"Name"`
	Url			string		`json:"Url"`
	ImageUrl	string		`json:"ImageUrl"`
	GeoLocation	GeoLocation	`json:"GeoLocation"`
}

// This is only provided with Wine.com partner access (shows as -360, -360 otherwise)
type GeoLocation struct {
	Latitude	int			`json:"Latitude"`
	Longitude	int			`json:"Longitude"`
	Url			string		`json:"Url"`
}

type Community struct {
	Reviews		Reviews				`json:"Reviews"`
	Url			string				`json:"Url"`
}

type Reviews struct {
	HighestScore	int		`json:"HighestScore"`
	List			[]int	`json:"List"`
	Url				string	`json:"Url"`
}

type ProductAttribute struct {
	Id			int			`json:"Id"`
	Name		string		`json:"Name"`
	Url			string		`json:"Url"`
	ImageUrl	string		`json:"ImageUrl"`
}

type Ratings struct {
	HighestScore	int		`json:"HighestScore"`
	List			[]int	`json:"List"` // only provided with partner access
}

type Retail struct {
	InStock	bool	`json:"InStock"`
	Price	float64	`json:"Price"`
	Sku		string	`json:"Sku"`
	State	string	`json:"State"`
	Url		string	`json:"Url"`
}

type Vintages struct {
	List	[]string	`json:"List"`
}
