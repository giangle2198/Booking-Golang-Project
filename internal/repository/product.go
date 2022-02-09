package repository

import (
	"booking-center-service/config"
	"booking-center-service/internal/common"
	"booking-center-service/internal/dto"
	"booking-center-service/internal/repository/model"
	libMongoDB "booking-libs/db"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ProductRepository interface {
		GetListProduct(ctx context.Context, categories []string) ([]*model.Product, error)
		GetProductDetail(ctx context.Context, productID string) (*model.Product, error)
		CreateProduct(ctx context.Context, req *dto.CreateProductRequestDTO) error
		UpdateProduct(ctx context.Context, req *dto.UpdateProductRequestDTO) error
		DeleteProduct(ctx context.Context, productID string) error
	}
	productRepository struct {
		productCollection *mongo.Collection
		cfg               *config.Config
		dbHelper          libMongoDB.NoSQLDBHelper
	}
)

func NewProductRepository(cfg *config.Config,
	dbHelper libMongoDB.NoSQLDBHelper) ProductRepository {
	productCollection := dbHelper.Collection("product")
	return &productRepository{
		productCollection: productCollection,
		cfg:               cfg,
		dbHelper:          dbHelper,
	}
}

func (r *productRepository) GetListProduct(ctx context.Context, categories []string) ([]*model.Product, error) {

	var (
		cursor      *mongo.Cursor
		listProduct []*model.Product
		err         error
	)

	if len(categories) > 0 {
		cursor, err = r.productCollection.Find(ctx, bson.D{{
			Key:   "categories",
			Value: bson.D{{Key: "$in", Value: categories}},
		}})
	} else {
		cursor, err = r.productCollection.Find(ctx, bson.M{})
	}
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &listProduct)
	if err != nil {
		return nil, err
	}

	if len(listProduct) < 1 {
		return nil, errors.New(common.ReasonNotFound.Code())
	}

	return listProduct, nil
}
func (r *productRepository) GetProductDetail(ctx context.Context, productID string) (*model.Product, error) {

	var (
		productDetail *model.Product
		err           error
	)

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}
	err = r.productCollection.FindOne(ctx, bson.D{{
		Key:   "_id",
		Value: objectID,
	}}).Decode(&productDetail)

	if err != nil {
		return nil, err
	}

	return productDetail, nil
}
func (r *productRepository) CreateProduct(ctx context.Context, req *dto.CreateProductRequestDTO) error {

	_, err := r.productCollection.InsertOne(ctx, bson.M{
		"title":       req.Title,
		"description": req.Description,
		"img":         req.Img,
		"categories":  req.Categories,
		"sizes":       req.Sizes,
		"price":       req.Price,
		"color":       req.Color,
	})

	return err
}
func (r *productRepository) UpdateProduct(ctx context.Context, req *dto.UpdateProductRequestDTO) error {

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return err
	}
	_, err = r.productCollection.UpdateOne(ctx, bson.D{{
		Key:   "_id",
		Value: objectID,
	}}, bson.D{{
		Key: "$set", Value: bson.M{
			"title":       req.Title,
			"description": req.Description,
			"img":         req.Img,
			"categories":  req.Categories,
			"sizes":       req.Sizes,
			"price":       req.Price,
			"color":       req.Color,
		},
	}})

	return err
}
func (r *productRepository) DeleteProduct(ctx context.Context, productID string) error {

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}
	_, err = r.productCollection.DeleteOne(ctx, bson.D{{
		Key:   "_id",
		Value: objectID,
	}})

	return err
}
