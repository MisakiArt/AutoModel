package global

import (
	"AutoModel/config"
	"gorm.io/gorm"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG    *zap.Logger
	GVA_CONFIG config.Server

	GVA_REDIS  *redis.Client
)