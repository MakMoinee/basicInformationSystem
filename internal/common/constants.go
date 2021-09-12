package common

// SERVER CONFIGURATIONS
var (
	PORT            string
	TLS_ENABLED     bool
	SERVICE_VERSION string

	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json; charset=UTF-8"
)

// Routes Resource Configuration
var (
	GET_STUDENT_INFO_BY_TEACHER_ID_R = "/get/student/{id}"
)

// MySql Configurations
var (
	MYSQL_USERNAME    string
	MYSQL_PASSWORD    string
	DB_NAME           string
	DB_DRIVER         string
	CONNECTION_STRING string
)
