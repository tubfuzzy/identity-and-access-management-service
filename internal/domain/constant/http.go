package constant

var (
	HTTPStatus200 int = 200
	HTTPStatus400 int = 400
	HTTPStatus401 int = 401
	HTTPStatus403 int = 403
	HTTPStatus404 int = 404
	HTTPStatus500 int = 500
)

const (
	OK                string = "E0200"
	UNKNOWN           string = "E0500"
	INVALID           string = "E0400"
	DEADLINE_EXCEEDED string = "E0504"
	NOT_FOUND         string = "E0404"
	UNAUTHENTICATED   string = "E0401"
	PERMISSION_DENIED string = "E0403"
	INTERNAL_ERROR    string = "E0500"
	UNAVAILABLE       string = "E0503"
	FRAMEWORK_ERROR   string = "E0520"
)
