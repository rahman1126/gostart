package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func SetConfigFile(filepath, filename, filetype string) {
	viper.SetConfigName(filename)
	viper.SetConfigType(filetype)
	viper.AddConfigPath(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetString(str string) string {
	return viper.GetString(fmt.Sprintf("%v", str))
}

func GetAppName() string {
	return viper.GetString("name")
}

func GetAddr() string {
	return fmt.Sprintf("%v", viper.GetString("host.address"))
}

func GetPort() string {
	return fmt.Sprintf("%v", viper.GetString("host.port"))
}

func GetFullAddr() string {
	return fmt.Sprintf("%v:%v", viper.GetString("host.address"), viper.GetString("host.port"))
}

func GetCtxTimeout() time.Duration {
	timeout := viper.GetInt("context.timeout")
	return time.Duration(timeout) * time.Second
}

func GetDBDriver() string {
	return viper.GetString("database.driver")
}

func GetDBName(driver string) string {
	return viper.GetString(fmt.Sprintf("database.%v.dbname", driver))
}

func GetDBUser(driver string) string {
	return viper.GetString(fmt.Sprintf("database.%v.username", driver))
}

func GetDBPass(driver string) string {
	return viper.GetString(fmt.Sprintf("database.%v.password", driver))
}

func GetDBHost(driver string) string {
	return viper.GetString(fmt.Sprintf("database.%v.host", driver))
}

func GetDBPort(driver string) string {
	return viper.GetString(fmt.Sprintf("database.%v.port", driver))
}

func IsUsingRedis() bool {
	return viper.GetBool("redis.caching")
}

func GetRedisAddr() string {
	return fmt.Sprintf("%v:%v", viper.GetString("redis.address"), viper.GetString("redis.port"))
}