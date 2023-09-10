package models

type Student struct {
	ID       uint
	NIM      string
	Fullname string
	Majority string
	Address  string
	Pict     string
}

var StudentData = []Student{
	{
		ID:       1,
		NIM:      "201106041165",
		Fullname: "Ahmad",
		Majority: "Teknik Informatika",
		Address:  "Kota Bogor",
		Pict:     "http://localhost:5000/static/img/201106041165.jpg",
	},
	{
		ID:       2,
		NIM:      "201106041166",
		Fullname: "Sidik",
		Majority: "Sistem Informasi",
		Address:  "Kota Depok",
		Pict:     "http://localhost:5000/static/img/201106041166.jpg",
	},
	{
		ID:       3,
		NIM:      "201106041167",
		Fullname: "Rudini",
		Majority: "Teknik Elektro",
		Address:  "Bumiayu",
		Pict:     "http://localhost:5000/static/img/201106041167.jpg",
	},
}
