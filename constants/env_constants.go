package constants

const (
	ServerBaseUrl         = "BASE_URL"
	ServerPort            = "SERVER_PORT"
	ServerWriteTimeout    = "WRITE_TIMEOUT"
	ServerReadTimeout     = "READ_TIMEOUT"
	ServerGracefulTimeout = "GRACEFUL_TIMEOUT"

	DbHost           = "DB_HOST"
	DbName           = "DB_NAME"
	DbUserName       = "DB_USERNAME"
	DbPassword       = "DB_PASSWORD"
	DbMaxOpenConn    = "DB_MAX_OPEN_CONN"
	DbMaxIdleConn    = "DB_MAX_IDLE_CONN"
	DbMaxConLifeTime = "DB_MAX_CONN_LIFE_TIME"

	RedisHost        = "REDIS_HOST"
	RedisPassword    = "REDIS_PASSWORD"
	RedisTimeout     = "REDIS_TIMEOUT"
	RedisMaxIdleConn = "REDIS_MAX_IDLE_CONN"

	JwtIssuer            = "JWT_ISSUER"
	JwtSecret            = "JWT_SECRET"
	JwtTokenLifeTimeHour = "JWT_TOKEN_LIFE_TIME_HOUR"

	MailHost       = "MAIL_HOST"
	MailPort       = "MAIL_PORT"
	MailUsername   = "MAIL_USERNAME"
	MailPassword   = "MAIL_PASSWORD"
	MailUseTls     = "MAIL_USE_TLS"
	MailSender     = "MAIL_SENDER"
	MailMaxAttempt = "MAIL_MAX_ATTEMPTS"
)
