package interfaces

type DB interface {
	Disconnect()
	InsertData(collectionName string, data any) error
	GetData(collectionName string, filter any) any
}
