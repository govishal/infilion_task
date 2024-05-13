package model

type Person struct {
    Name        string `json:"name"`
	Phone_number string `json:"phone_number"`
	City string `json:"city"`
	State string `json:"state"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	ZipCode string `json:"zip_code"`
}