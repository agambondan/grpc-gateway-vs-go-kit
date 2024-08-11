package env

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"git.bluebird.id/promo/packages/zaplog"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

const (
	mongoDBURIEnv      = "MONGO_URI"
	mongoDBDatabaseEnv = "MONGO_DB"

	redisDBEnv        = "REDIS_DB"
	redisHostEnv      = "REDIS_HOST"
	redisPortEnv      = "REDIS_PORT"
	redisPasswordEnv  = "REDIS_PASSWORD"
	redisMaxIdleEnv   = "REDIS_MAX_IDLE"
	redisMaxActiveEnv = "REDIS_MAX_ACTIVE"

	promotionCacheTTLEnv = "PROMOTION_CACHE_TTL"

	grpcPortEnv = "GRPC_PORT"
	httpPortEnv = "HTTP_PORT"

	basicAuthUsernameEnv = "BASIC_AUTH_USERNAME"
	basicAuthPasswordEnv = "BASIC_AUTH_PASSWORD"

	envLevel = "ENV_LEVEL"

	singleUseLengthEnv = "PROMO_SINGLE_USE_LENGTH"

	bbOneBaseUrlEnv = "BBONE_AUTH_BASE_URL"

	upgBaseUrlEnv   = "UPG_BASE_URL"
	upgBasicAuthEnv = "UPG_BASIC_AUTH"
	upgKeyAuthEnv   = "UPG_KEY_AUTH"

	postgreDBHostEnv          = "POSTGRE_DB_HOST"
	postgreDBPortEnv          = "POSTGRE_DB_PORT"
	postgreDBUserEnv          = "POSTGRE_DB_USER"
	postgreDBPwEnv            = "POSTGRE_DB_PASS"
	postgreDBNameEnv          = "POSTGRE_DB_NAME"
	postgreDBMaxConnectionEnv = "POSTGRE_DB_MAX_CONN"
	postgreDBMaxTimeoutEnv    = "POSTGRE_DB_MAX_TIMEOUT"

	postgreDBPromoNewHostEnv          = "POSTGRE_DB_PROMO_NEW_HOST"
	postgreDBPromoNewPortEnv          = "POSTGRE_DB_PROMO_NEW_PORT"
	postgreDBPromoNewUserEnv          = "POSTGRE_DB_PROMO_NEW_USER"
	postgreDBPromoNewPwEnv            = "POSTGRE_DB_PROMO_NEW_PASS"
	postgreDBPromoNewNameEnv          = "POSTGRE_DB_PROMO_NEW_NAME"
	postgreDBPromoNewMaxConnectionEnv = "POSTGRE_DB_PROMO_NEW_MAX_CONN"
	postgreDBPromoNewMaxTimeoutEnv    = "POSTGRE_DB_PROMO_NEW_MAX_TIMEOUT"

	profilingGrpcAddrEnv   = "PROFILING_GRPC_ADDR"
	defaultDiscountNameEnv = "DEFAULT_DISCOUNT"

	googleBucketCredentials = "GOOGLE_BUCKET_CREDENTIALS"
	gcsBucketName           = "GCS_BUCKET_NAME"

	allPaymentMethodWordingEn = "ALL_PAYMENT_METHOD_WORDING_EN"
	allPaymentMethodWordingId = "ALL_PAYMENT_METHOD_WORDING_ID"

	subscriptionGRPCApiKey = "SUBSCRIPTION_GRPC_API_KEY"

	promoPortalPubSubProjectID                   = "PROMO_PORTAL_PUBSUB_PROJECT_ID"
	promoPortalPubSubCredentials                 = "PROMO_PORTAL_PUBSUB_CREDENTIALS"
	promoPortalCreatePromoTopic                  = "PROMO_PORTAL_CREATE_PROMO_TOPIC"
	promoPortalUpdatePromoTopic                  = "PROMO_PORTAL_UPDATE_PROMO_TOPIC"
	promoPortalDeactivatePromoTopic              = "PROMO_PORTAL_DEACTIVATE_PROMO_TOPIC"
	promoPortalDeactivatePromoSingleUseTopic     = "PROMO_PORTAL_DEACTIVATE_PROMO_SINGLE_USE_TOPIC"
	promoPortalDeactivatePromoBulkSingleUseTopic = "PROMO_PORTAL_DEACTIVATE_PROMO_BULK_SINGLE_USE_TOPIC"
	promoPortalReusePromoTopic                   = "PROMO_PORTAL_REUSE_PROMO_TOPIC"
	syncPromoV1                                  = "SYNC_PROMO_V1"

	promoPortalUpdateIDTopic        = "PROMO_PORTAL_UPDATE_ID_TOPIC"
	promoPortalUpdateIDSubscriberID = "PROMO_PORTAL_UPDATE_ID_SUBSCRIBER_ID"

	promoPortalCreateSingleUseCallbackTopic        = "PROMO_PORTAL_CREATE_SINGLE_USE_CALLBACK_TOPIC"
	promoPortalCreateSingleUseCallbackSubscriberID = "PROMO_PORTAL_CREATE_SINGLE_USE_CALLBACK_SUBSCRIBER_ID"

	promoPortalUpdateSingleUseUsedTopic        = "PROMO_PORTAL_UPDATE_SINGLE_USE_USED_TOPIC"
	promoPortalUpdateSingleUseUsedSubscriberID = "PROMO_PORTAL_UPDATE_SINGLE_USE_USED_SUBSCRIBER_ID"

	promoPortalUpdateUsageTopic        = "PROMO_PORTAL_UPDATE_USAGE_TOPIC"
	promoPortalUpdateUsageSubscriberID = "PROMO_PORTAL_UPDATE_USAGE_SUBSCRIBER_ID"
)

var (
	PostgresDBName    string
	PostgresHost      string
	PostgresPort      string
	PostgresUsername  string
	PostgresPassword  string
	PostgresMaxIdle   int64
	PostgresMaxActive int64

	MongoDBURI      string
	MongoDBDatabase string

	RedisDB        int64
	RedisHost      string
	RedisPort      string
	RedisPassword  string
	RedisMaxIdle   int64
	RedisMaxActive int64

	PromotionCacheTTL int

	GRPCPort string
	HTTPPort string

	BasicAuthPassword string
	BasicAuthUsername string

	EnvLevel string

	SingleUseLength int

	BBOneBaseUrl string

	UPGBaseUrl   string
	UPGBasicAuth string
	UPGKeyAuth   string

	SubscriptionGRPCApiKey string

	PostgreDBHost          string
	PostgreDBPort          string
	PostgreDBUser          string
	PostgreDBPw            string
	PostgreDBName          string
	PostgreDBMaxConnection int
	PostgreDBMaxTimeout    int

	PostgreDBPromoNewHost          string
	PostgreDBPromoNewPort          string
	PostgreDBPromoNewUser          string
	PostgreDBPromoNewPw            string
	PostgreDBPromoNewName          string
	PostgreDBPromoNewMaxConnection int
	PostgreDBPromoNewMaxTimeout    int

	ProfilingGrpcAddr   string
	DefaultDiscountName float64

	GoogleBucketCredentials string
	GCSBucketName           string

	AllPaymentMethodWordingId string
	AllPaymentMethodWordingEn string

	BluebirdWordingEn                      string
	SilverbirdWordingEn                    string
	BirdkirimWordingEn                     string
	GoldenbirdWordingEn                    string
	BluebirdSilverbirdWordingEn            string
	BluebirdBirdkirimWordingEn             string
	BluebirdGoldenbirdWordingEn            string
	SilverbirdBirdkirimWordingEn           string
	SilverbirdGoldenbirdWordingEn          string
	BirdkirimGoldenbirdWordingEn           string
	BluebirdSilverbirdBirdkirimWordingEn   string
	BluebirdSilverbirdGoldenbirdWordingEn  string
	BluebirdBirdkirimGoldenbirdWordingEn   string
	SilverbirdBirdkirimGoldenbirdWordingEn string
	AllWordingEn                           string

	BluebirdWordingId                      string
	SilverbirdWordingId                    string
	BirdkirimWordingId                     string
	GoldenbirdWordingId                    string
	BluebirdSilverbirdWordingId            string
	BluebirdBirdkirimWordingId             string
	BluebirdGoldenbirdWordingId            string
	SilverbirdBirdkirimWordingId           string
	SilverbirdGoldenbirdWordingId          string
	BirdkirimGoldenbirdWordingId           string
	BluebirdSilverbirdBirdkirimWordingId   string
	BluebirdSilverbirdGoldenbirdWordingId  string
	BluebirdBirdkirimGoldenbirdWordingId   string
	SilverbirdBirdkirimGoldenbirdWordingId string
	AllWordingId                           string

	PromoPortalPubSubProjectID                     string
	PromoPortalPubSubCredentials                   string
	PromoPortalCreatePromoTopic                    string
	PromoPortalUpdatePromoTopic                    string
	PromoPortalDeactivatePromoTopic                string
	PromoPortalDeactivatePromoSingleUseTopic       string
	PromoPortalDeactivatePromoBulkSingleUseTopic   string
	PromoPortalReusePromoTopic                     string
	PromoPortalUpdateIDTopic                       string
	PromoPortalUpdateIDSubscriberID                string
	PromoPortalCreateSingleUseCallbackTopic        string
	PromoPortalCreateSingleUseCallbackSubscriberID string
	PromoPortalUpdateSingleUseUsedTopic            string
	PromoPortalUpdateSingleUseUsedSubscriberID     string
	PromoPortalUpdateUsageTopic                    string
	PromoPortalUpdateUsageSubscriberID             string

	SyncPromoV1 bool
)

func LoadEnv() {
	logger := zaplog.WithContext(context.Background())
	var err error

	MongoDBURI, err = getEnvString(mongoDBURIEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	MongoDBDatabase, err = getEnvString(mongoDBDatabaseEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	RedisDB, err = getEnvInt64(redisDBEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	RedisHost, err = getEnvString(redisHostEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	RedisPort, err = getEnvString(redisPortEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	RedisPassword, _ = getEnvString(redisPasswordEnv)

	RedisMaxIdle, err = getEnvInt64(redisMaxIdleEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	RedisMaxActive, err = getEnvInt64(redisMaxActiveEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromotionCacheTTL, err = getEnvInt(promotionCacheTTLEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	GRPCPort, err = getEnvString(grpcPortEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	HTTPPort, err = getEnvString(httpPortEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	BasicAuthUsername, err = getEnvString(basicAuthUsernameEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}
	BasicAuthPassword, err = getEnvString(basicAuthPasswordEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	BBOneBaseUrl, err = getEnvString(bbOneBaseUrlEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	UPGBaseUrl, err = getEnvString(upgBaseUrlEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	UPGBasicAuth, err = getEnvString(upgBasicAuthEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	UPGKeyAuth, err = getEnvString(upgKeyAuthEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	SingleUseLength, err = getEnvInt(singleUseLengthEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBHost, err = getEnvString(postgreDBHostEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPort, err = getEnvString(postgreDBPortEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBUser, err = getEnvString(postgreDBUserEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPw, err = getEnvString(postgreDBPwEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBName, err = getEnvString(postgreDBNameEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBMaxConnection, err = getEnvInt(postgreDBMaxConnectionEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBMaxTimeout, err = getEnvInt(postgreDBMaxTimeoutEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPromoNewHost, err = getEnvString(postgreDBPromoNewHostEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPromoNewPort, err = getEnvString(postgreDBPromoNewPortEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPromoNewUser, err = getEnvString(postgreDBPromoNewUserEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPromoNewPw, err = getEnvString(postgreDBPromoNewPwEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPromoNewName, err = getEnvString(postgreDBPromoNewNameEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPromoNewMaxConnection, err = getEnvInt(postgreDBPromoNewMaxConnectionEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PostgreDBPromoNewMaxTimeout, err = getEnvInt(postgreDBPromoNewMaxTimeoutEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	ProfilingGrpcAddr, err = getEnvString(profilingGrpcAddrEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	EnvLevel, err = getEnvString(envLevel)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	DefaultDiscountName, err = getEnvFloat64(defaultDiscountNameEnv)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	GoogleBucketCredentials, err = getEnvString(googleBucketCredentials)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	GCSBucketName, err = getEnvString(gcsBucketName)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	AllPaymentMethodWordingEn, err = getEnvString(allPaymentMethodWordingEn)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	AllPaymentMethodWordingId, err = getEnvString(allPaymentMethodWordingId)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	SubscriptionGRPCApiKey, err = getEnvString(subscriptionGRPCApiKey)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalPubSubProjectID, err = getEnvString(promoPortalPubSubProjectID)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalPubSubCredentials, err = getEnvString(promoPortalPubSubCredentials)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalCreatePromoTopic, err = getEnvString(promoPortalCreatePromoTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalUpdatePromoTopic, err = getEnvString(promoPortalUpdatePromoTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalDeactivatePromoTopic, err = getEnvString(promoPortalDeactivatePromoTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalDeactivatePromoSingleUseTopic, err = getEnvString(promoPortalDeactivatePromoSingleUseTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalDeactivatePromoBulkSingleUseTopic, err = getEnvString(promoPortalDeactivatePromoBulkSingleUseTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	SyncPromoV1, err = getEnvBool(syncPromoV1)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalReusePromoTopic, err = getEnvString(promoPortalReusePromoTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalUpdateIDTopic, err = getEnvString(promoPortalUpdateIDTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalUpdateIDSubscriberID, err = getEnvString(promoPortalUpdateIDSubscriberID)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalCreateSingleUseCallbackTopic, err = getEnvString(promoPortalCreateSingleUseCallbackTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalCreateSingleUseCallbackSubscriberID, err = getEnvString(promoPortalCreateSingleUseCallbackSubscriberID)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalUpdateSingleUseUsedTopic, err = getEnvString(promoPortalUpdateSingleUseUsedTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalUpdateSingleUseUsedSubscriberID, err = getEnvString(promoPortalUpdateSingleUseUsedSubscriberID)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalUpdateUsageTopic, err = getEnvString(promoPortalUpdateUsageTopic)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}

	PromoPortalUpdateUsageSubscriberID, err = getEnvString(promoPortalUpdateUsageSubscriberID)
	if err != nil {
		logger.Fatal("failed loading env", zap.Error(err))
	}
}

func getEnvString(env string) (string, error) {
	res := os.Getenv(env)
	if res == "" {
		return res, fmt.Errorf("env %s is empty", env)
	}
	return res, nil
}

func getEnvInt(env string) (int, error) {
	res := os.Getenv(env)
	if res == "" {
		return 0, fmt.Errorf("env %s is empty", env)
	}

	resInt, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return resInt, nil
}

func getEnvInt64(env string) (int64, error) {
	res := os.Getenv(env)
	if res == "" {
		return 0, fmt.Errorf("env %s is empty", env)
	}

	resInt, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return int64(resInt), nil
}

func getEnvFloat64(env string) (float64, error) {
	res := os.Getenv(env)
	if res == "" {
		return 0, fmt.Errorf("env %s is empty", env)
	}

	resFloat64, err := strconv.ParseFloat(res, 64)
	if err != nil {
		return 0, err
	}

	return resFloat64, nil
}

func getEnvBool(env string) (bool, error) {
	res := os.Getenv(env)
	if res == "" {
		return false, fmt.Errorf("env %s is empty", env)
	}

	resBool, err := strconv.ParseBool(res)
	if err != nil {
		return false, err
	}

	return resBool, nil
}
