package structs

type SensitiveInformation struct {
	IdAccount
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	Contact
	Adress
}

type Adress struct {
	StreetName            string `json:"street_name"`
	NumberStreet          string `json:"number_street"`
	PostalCode            string `json:"postal_code"`
	StateName             string `json:"state_name"`
	CityName              string `json:"city_name"`
	CountryName           string `json:"country_name"`
	AdditionalInformation string `json:"additional_information"`
}

type Contact struct {
	S          string `json:"phone_number"`
	PhoneLocal string `json:"phone_local"`
}
