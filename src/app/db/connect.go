package db

type DAO interface {
	Create(item interface{}) error
	Read(input map[string]interface{}) (interface{}, error)
	Update(item interface{}) error
	Delete(input interface{}) error
}
