package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"gostart/utils/conf"
	"time"
)

func initialConfig() (key string, mix string) {
	today := time.Now()
	appKey := conf.GetString("security.app_key")
	appMessage := conf.GetString("security.app_message")
	combination := appMessage + today.Format("2006-01-02 15:04:05")

	return appKey, combination
}

func Base64() string {
	key, mix := initialConfig()
	secret := []byte(key)
	message := []byte(mix)

	hash := hmac.New(sha256.New, secret)
	hash.Write(message)

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func Hex() string {
	key, mix := initialConfig()
	secret := []byte(key)
	message := []byte(mix)

	hash := hmac.New(sha256.New, secret)
	hash.Write(message)

	return hex.EncodeToString(hash.Sum(nil))
}