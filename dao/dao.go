package dao

var client Factory

type Factory interface {
	Users() UserDao
}

// Client 返回 Factory client
func Client() Factory {
	return client
}

// SetClient 设置 Factory client
func SetClient(factory Factory) {
	client = factory
}
