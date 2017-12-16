package storage

type ContactStorage interface {
	FindAll() (ContactCollection, error)
}
