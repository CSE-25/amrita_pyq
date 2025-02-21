package configs

const (
	BASE_URL        = "http://dspace.amritanet.edu:8080"
	COURSE_URL      = BASE_URL + "/xmlui/handle/123456789/"
	COURSE_LIST_URL = COURSE_URL + "16"
)

const (
	CONFIGS_PATH        = "cmd/configs"
	HELPERS_PATH        = "cmd/helpers"
	LOGO_PATH           = "cmd/logo"
	MODEL_PATH          = "cmd/model"
	REQUEST_CLIENT_PATH = "cmd/requestClient"
	ROOT_PATH           = "cmd/root"
	STACK_PATH          = "cmd/stack"
	VERSION_PATH        = "cmd/version"
)

var TestPaths = [...]string{
	CONFIGS_PATH,
	HELPERS_PATH,
	LOGO_PATH,
	MODEL_PATH,
	REQUEST_CLIENT_PATH,
	ROOT_PATH,
	STACK_PATH,
	VERSION_PATH,
}
