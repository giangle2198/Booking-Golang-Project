package registry

import (
	"booking-center-service/config"
	"booking-center-service/internal/api"
	"booking-center-service/internal/helper"
	"booking-center-service/internal/repository"
	"booking-center-service/internal/usecase"
	"fmt"

	libDB "booking-libs/db"
	libRegistry "booking-libs/registry"

	bimClient "booking-identity-management/cmd/client"

	"github.com/sarulabs/di"
)

var (
	// Config
	ConfigDIName = "Config"
	// API
	APIDIName = "API"
	// Usecase
	BaseUsecaseDIName    = "BaseUsecase"
	ProductUsecaseDIName = "ProductUsecase"
	// Repository
	BaseRepositoryDIName    = "BaseRepository"
	ProductRepositoryDIName = "ProductRepository"
	// Helper
	// DBHelperDIName       = "DBHelper"
	PbConverterDIName    = "PbConverter"
	ModelConverterDIName = "ModelConverterDIName"
	MongoHelperDIName    = "MongoHelper"

	// Adapter
	BIMClientDIName = "bimClientAdapter"
)

// BuildDIContainer Wrapper for lib registry BuildDIContainer
func BuildDIContainer() {
	initBuilder()
	libRegistry.BuildDIContainer()
}

// GetDependency Wrapper for lib registry GetDependency
func GetDependency(name string) interface{} {
	return libRegistry.GetDependency(name)
}

func initBuilder() {
	// init for configs builder
	libRegistry.ConfigsBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ConfigDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg, err := config.Load()
				return cfg, err
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for helpers builder
	libRegistry.HelpersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  MongoHelperDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return libDB.NewMongoDBHelper(cfg.Mongo.URL, cfg.Mongo.Database), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
			// }, di.Def{
			// 	Name:  DBHelperDIName,
			// 	Scope: di.App,
			// 	Build: func(di di.Container) (interface{}, error) {
			// 		cfg := di.Get(ConfigDIName).(*config.Config)
			// 		return libDB.NewMySQLDBHelper(cfg.MySQL.Host, cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Database, cfg.MySQL.Port), nil
			// 	},
			// 	Close: func(obj interface{}) error {
			// 		return nil
			// 	},
		}, di.Def{
			Name:  PbConverterDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				return helper.NewPbConverter(), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  ModelConverterDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				return helper.NewModelConverter(), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for apdeters builder
	libRegistry.AdaptersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  BIMClientDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return bimClient.NewBIMClient(fmt.Sprintf("%s:%v", cfg.BIMClient.Host, cfg.BIMClient.Port))
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for usecases builder
	libRegistry.UsecasesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  BaseUsecaseDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				baseRepository := di.Get(BaseRepositoryDIName).(repository.BaseRepository)
				return usecase.NewBaseUsecase(cfg, baseRepository), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  ProductUsecaseDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				productRepository := di.Get(ProductRepositoryDIName).(repository.ProductRepository)
				return usecase.NewProductUsecase(cfg, productRepository), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})

		return arr
	}

	// init for repositories builder
	libRegistry.RepositoriesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ProductRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				dbHelper := di.Get(MongoHelperDIName).(libDB.NoSQLDBHelper)
				return repository.NewProductRepository(cfg, dbHelper), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  BaseRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				dbHelper := di.Get(MongoHelperDIName).(libDB.NoSQLDBHelper)
				return repository.NewBaseRepository(cfg, dbHelper), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for apis builder
	libRegistry.APIsBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  APIDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				baseUsecase := di.Get(BaseUsecaseDIName).(usecase.BaseUsecase)
				productUsecase := di.Get(ProductUsecaseDIName).(usecase.ProductUsecase)
				pbConvert := di.Get(PbConverterDIName).(helper.PbConverter)
				return api.NewAPI(cfg, pbConvert, baseUsecase, productUsecase), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}
}
