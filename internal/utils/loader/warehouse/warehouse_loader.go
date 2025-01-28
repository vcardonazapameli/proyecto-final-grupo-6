package warehouse

type WarehouseLoader interface {
	Load() (any, error)
}
