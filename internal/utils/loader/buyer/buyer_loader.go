package buyer

type BuyerLoader interface {
	Load()(any, error)
}