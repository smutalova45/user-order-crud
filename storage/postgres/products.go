package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"main.go/models"
)

type productsRepo struct {
	DB *sql.DB
}

func NewProductsRepo(db *sql.DB) productsRepo {
	return productsRepo{
		DB: db,
	}
}

func (p productsRepo) Insertproduct(product models.Products) (string, error) {
	id := uuid.New()
	if _, err := p.DB.Exec(`insert into products values($1, $2, $3 )`, id, product.Name, product.Price); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (p productsRepo) GetlistProducts() ([]models.Products, error) {
	rows, err := p.DB.Query(`select * from products `)
	if err != nil {
		return nil, err
	}
	p1 := []models.Products{}
	for rows.Next() {
		pro := models.Products{}
		if err = rows.Scan(&pro.Id, &pro.Name, &pro.Price); err != nil {
			return nil, err
		}
		p1 = append(p1, pro)
	}
	return p1, nil
}

func (p productsRepo) Updateproducts(product models.Products) error {
	_, err := p.DB.Exec(`update products set price =$1 where id =$2`, product.Price, product.Id)
	if err != nil {
		return err
	}
	return nil

}

func (p productsRepo) DeleteProducts(id uuid.UUID) error {
	if _, err := p.DB.Exec("delete from orderproducts where product_id= $1", id); err != nil {
		return err
	}
	if _, err := p.DB.Query(`delete from products where id=$1`, id); err != nil {
		return err
	}
	return nil
}
