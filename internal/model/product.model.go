package model

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var Products = []Product{
	{ID: 1, Name: "Indomie Goreng", Price: 6000, Stock: 100},
	{ID: 2, Name: "Indomie Ayam", Price: 10000, Stock: 200},
	{ID: 3, Name: "Es Teh", Price: 3000, Stock: 300},
}

func GetProducts() []Product {
	return Products
}
