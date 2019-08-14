package repositories

// IRepositoryBase is a
type IRepositoryBase interface {
	Get()
	GetAll()
	Update()
	Delete()
	Create()
}
