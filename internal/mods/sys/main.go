package sys

import (
	"context"

	"github.com/fsyabc/test/internal/mods/sys/api"
	"github.com/fsyabc/test/internal/mods/sys/schema"
	"github.com/gin-gonic/gin"
	"github.com/fsyabc/test/internal/config"
	"github.com/fsyabc/test/internal/mods/sys/api"
	"gorm.io/gorm"
)

type SYS struct {
	DB            *gorm.DB
	LoggerAPI     *api.Logger
	DictionaryAPI *api.Dictionary
}

func (a *SYS) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Dictionary))
}

func (a *SYS) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *SYS) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	logger := v1.Group("loggers")
	{
		logger.GET("", a.LoggerAPI.Query)
	}
	dictionary := v1.Group("dictionaries")
	{
		dictionary.GET("", a.DictionaryAPI.Query)
		dictionary.GET(":id", a.DictionaryAPI.Get)
		dictionary.POST("", a.DictionaryAPI.Create)
		dictionary.PUT(":id", a.DictionaryAPI.Update)
		dictionary.DELETE(":id", a.DictionaryAPI.Delete)
	}
	return nil
}

func (a *SYS) Release(ctx context.Context) error {
	return nil
}
