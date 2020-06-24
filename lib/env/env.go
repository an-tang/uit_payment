package env

import (
	"fmt"
	"os"
	"strconv"
)

func getEnv(env string) string {
	if value := os.Getenv(env); value == "" {
		panic(fmt.Sprintf("ENV %s is empty", env))
	} else {
		return value
	}
}

func GetDBHost() string {
	return getEnv("DB_HOST")
}

func GetDBPort() string {
	return getEnv("DB_PORT")
}

func GetDBSSLMode() string {
	return getEnv("DB_SSL_MODE")
}

func GetDBUserName() string {
	return getEnv("DB_USERNAME")
}

func GetDBPassword() string {
	return getEnv("DB_PASSWORD")
}

func GetDBName() string {
	return getEnv("DB_NAME")
}

func GetMaxOpenConns() int {
	if v, err := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS")); err == nil {
		return v
	}

	return 0
}

func GetMaxIdleConns() int {
	if v, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS")); err == nil {
		return v
	}

	return 0
}

func GetLogAge() string {
	return getEnv("LOG_AGE")
}

func GetGraylogURL() string {
	return getEnv("GRAYLOG_ADDRESS")
}

func GetEnvironment() string {
	return getEnv("ENVIRONMENT")
}

func GetFacility() string {
	return getEnv("FACILITY")
}

func GetLogLevel() string {
	return getEnv("LOG_LEVEL")
}

func GetMomoQRCodeURL() string {
	return getEnv("MOMO_QRCODE_URL")
}

func GetMomoConfirmURL() string {
	return getEnv("MOMO_CONFIRM_URL")
}

func GetMomoGetPaymentURL() string {
	return getEnv("MOMO_GETPAYMENT_URL")
}

func GetMomoRefundURL() string {
	return getEnv("MOMO_REFUND_URL")
}

func GetMomoPartnerCode() string {
	return getEnv("MOMO_PARTNERCODE")
}

func GetMomoAccessKey() string {
	return getEnv("MOMO_ACCESSKEY")
}

func GetMomoSecretKey() string {
	return getEnv("MOMO_SECRETKEY")
}

func GetMomoPublicKey() string {
	return getEnv("MOMO_PUBLICKEY")
}

func GetMomoVersion() float32 {
	version := getEnv("MOMO_VERSION")
	value, err := strconv.ParseFloat(version, 32)
	if err != nil {
		// do something sensible
	}
	float := float32(value)
	return float
}

func UitTravelURL() string {
	return getEnv("UIT_TRAVEL_URL")
}
