package repo

import (
	"hermes/controllers/entity"
	models "hermes/models/product"

	"github.com/jmoiron/sqlx"
)

type ProductService struct {
	connection *sqlx.DB
}

func NewProductService(conn *sqlx.DB) (*ProductService) {
	return &ProductService{connection: conn}
}

func (s *ProductService) GetAll() (*models.Products, error) {
	query := `SELECT * FROM product`
	var products models.Products
	err := s.connection.Select(&products.Products, query)
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (s *ProductService) GetByPK(id int) (*models.Product, error) {
	query := `SELECT * FROM product WHERE id = $1`
	product := []models.Product{}
	err := s.connection.Select(&product, query, id)
	if err != nil {
		return nil, err
	}
	if len(product) == 0{
		return nil, nil
	}
	return &product[0], nil
}

func (s *ProductService) GetBySKU(sku string) (*models.Product, error) {
	query := `SELECT * FROM product WHERE sku = $1`
	product := []models.Product{}
	err := s.connection.Select(&product, query, sku)
	if err != nil {
		return nil, err
	}
	if len(product) == 0{
		return nil, nil
	}
	return &product[0], nil
}

func (s *ProductService) InsertOne(product entity.ProductInput) error {
	query := `INSERT INTO product(sku, name, display) VALUES($1, $2, $3)`

	tx := s.connection.MustBegin()
	tx.MustExec(query, product.Sku, product.Name, product.Display)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateOne(product entity.ProductInput) error {
	query := `UPDATE product SET sku = $1, name = $2, display = $3 WHERE id = $4`

	tx := s.connection.MustBegin()
	tx.MustExec(query, product.Sku, product.Name, product.Display, product.Id)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) DeleteOne(id int) error {
	query := `DELETE FROM product WHERE id = $1`

	tx := s.connection.MustBegin()
	tx.MustExec(query, id)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

