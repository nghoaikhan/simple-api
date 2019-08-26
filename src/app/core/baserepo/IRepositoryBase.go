package baserepo

// IRepositoryBase is a
type IRepositoryBase interface {
	Find()
	FindOne()
	FindById()
	Update()
	UpdateById()
	Delete()
	Create(dto interface{})
}
