package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// InitConfig init the configuration
func InitConfig() {
	// enable config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Use config file:", viper.ConfigFileUsed())
	}

	// Set default configuration
	// System
	viper.SetDefault("ENV", "development") // environment: development/staging/production

	// HTTP
	viper.SetDefault("HTTP_PORT", "80")
	viper.SetDefault("HTTP_HEADER_LANG", "X-lang")

	// GRPC
	viper.SetDefault("GRPC_SERVER_PORT", "8080")
	viper.SetDefault("GRPC_CONN_TIMEOUT", 10)          // GRPC 连接超时时间，单位为秒
	viper.SetDefault("GRPC_HOST_EXAMPLE", "127.0.0.1") // Example 服务的 GRPC 地址，其他服务需要添加自己的 key
	viper.SetDefault("GRPC_PORT_EXAMPLE", "8080")      // Example 服务的 GRPC 端口，其他服务需要添加自己的 key

	// DB
	viper.SetDefault("DB_ENGINE", "mysql")
	viper.SetDefault("DB_HOST", "127.0.0.1")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_NAME", "amf")
	viper.SetDefault("DB_CONN_TIMEOUT", 10) // DB 连接超时时间，单位为秒。各 DB 通用

	// MongoDB
	viper.SetDefault("MONGODB_USER", "")
	viper.SetDefault("MONGODB_PASSWORD", "")
	viper.SetDefault("MONGODB_DBNAME", "amf")
	viper.SetDefault("MONGODB_HOST", "127.0.0.1")
	viper.SetDefault("MONGODB_PORT", "27017")
	viper.SetDefault("MONGODB_SSL", "false")
	viper.SetDefault("MONGODB_CONN_TIMEOUT", 10) // MongoDB 连接超时时间，单位为秒
	viper.SetDefault("MONGODB_OP_TIMEOUT", 5)    // MongoDB 操作超时时间，单位为秒

	// Kubernetes
	viper.SetDefault("K8S_CONN_TIMEOUT", 10)        // K8S 连接超时时间，单位为秒
	viper.SetDefault("K8S_NAMESPACE", "amf-system") // 系统默认使用的 namespace

	// Logic
	viper.SetDefault("PAGE_LIMIT", 20)

	// enable env variables
	viper.AutomaticEnv()
}
