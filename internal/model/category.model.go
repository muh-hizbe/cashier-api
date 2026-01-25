package model

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Makanan"},
	{ID: 2, Name: "Snack", Description: "Snack"},
	{ID: 3, Name: "Minuman", Description: "Minumam"},
}

func GetCategories() []Category {
	return Categories
}
