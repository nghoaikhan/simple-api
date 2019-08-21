package repositories

// IRepositoryBase is a
type IRepositoryBase interface {
	Find()
	FindOne()
	FindById()
	Update()
	UpdateById()
	Delete()
	Create()
}
