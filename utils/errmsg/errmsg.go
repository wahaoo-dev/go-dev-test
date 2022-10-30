package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	// 分类模块错误
	ERROR_CATEGORY_USED      = 2001
	ERROR_CATEGORY_WRONG     = 2002
	ERROR_CATEGORY_NOT_EXIST = 2003

	// 文章模块错误
	ERROR_ARTICLE_USED      = 3001
	ERROR_ARTICLE_NOT_EXIST = 3002

	// 标签模块错误
	ERROR_TAG_USED      = 4001
	ERROR_TAG_WRONG     = 4002
	ERROR_TAG_NOT_EXIST = 4003
)

var codeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",

	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "token 不存在",
	ERROR_TOKEN_RUNTIME:    "token 已过期",
	ERROR_TOKEN_WRONG:      "token 不正确",
	ERROR_TOKEN_TYPE_WRONG: "token 格式错误",
	ERROR_USER_NO_RIGHT:    "用户无权限",

	ERROR_CATEGORY_USED:      "分类已存在",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",

	ERROR_ARTICLE_NOT_EXIST: "文章不存在",

	ERROR_TAG_USED:      "标签已存在",
	ERROR_TAG_NOT_EXIST: "标签不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
