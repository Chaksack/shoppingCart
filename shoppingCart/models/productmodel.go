package models

import (
	"trainingProject/shoppingCart/config"
	"trainingProject/shoppingCart/entities"
)

type ProductModel struct {

}

func (*ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where id=?", id)
		if err2 != nil {
			return entities.Product{}, err2
		}else{
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)
			}
			return product, nil
		}
	}
}
