package product

type ProductLoader interface {
	Load() (any, error)
}
