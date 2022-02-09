package repository

import (
	"booking-center-service/config"
	"booking-center-service/internal/common"
	"booking-center-service/internal/repository/model"
	libDB "booking-libs/db"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	BaseRepository interface {
		GetManagementsByChannel(ctx context.Context, channel string) ([]*model.Management, error)
	}

	baseRepository struct {
		cfg                  *config.Config
		db                   libDB.NoSQLDBHelper
		managementCollection *mongo.Collection
	}
)

func NewBaseRepository(cfg *config.Config,
	db libDB.NoSQLDBHelper) BaseRepository {
	managementCollection := db.Collection("management")
	return &baseRepository{
		cfg:                  cfg,
		db:                   db,
		managementCollection: managementCollection,
	}
}

func (r *baseRepository) GetManagementsByChannel(ctx context.Context, channel string) (managements []*model.Management, err error) {
	var (
		cursor         *mongo.Cursor
		listManagement []*model.Management
	)

	if channel != "" {
		cursor, err = r.managementCollection.Find(ctx, bson.D{{
			Key:   "channel",
			Value: channel,
		}})
	} else {
		cursor, err = r.managementCollection.Find(ctx, bson.M{})
	}
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &listManagement)
	if err != nil {
		return nil, err
	}

	if len(listManagement) < 1 {
		return nil, errors.New(common.ReasonNotFound.Code())
	}

	return listManagement, nil
}
