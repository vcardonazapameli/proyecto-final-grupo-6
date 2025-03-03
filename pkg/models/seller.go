package models

type Seller struct {
	Id          int
	Cid         int
	CompanyName string
	Address     string
	Telephone   string
	LocalityID  int
}

type SellerDoc struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

func NewSeller(id int, cid int, companyName string, address string, telephone string, localityId int) *Seller {
	return &Seller{Id: id, Cid: cid, CompanyName: companyName, Address: address, Telephone: telephone, LocalityID: localityId}
}

func NewSellerDoc(id int, cid int, companyName string, address string, telephone string, localityId int) *SellerDoc {
	return &SellerDoc{Id: id, Cid: cid, CompanyName: companyName, Address: address, Telephone: telephone, LocalityID: localityId}
}
