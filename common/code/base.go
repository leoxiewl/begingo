package code

const (
	// SucCommon 成功通用码
	SucCommon int = 200

	// ErrCommon 失败通用码
	ErrCommon int = -1
)

const (
	// ErrUnknown 未知错误
	ErrUnknown int = iota + 10001

	// ErrBind 参数绑定错误
	ErrBind
)

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)
