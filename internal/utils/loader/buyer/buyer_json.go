package buyer

import (
	"encoding/json"
	"os"

	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/mappers"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

func NewBuyerJsonFile(path string) *BuyerJsonFile{
	return &BuyerJsonFile{
		path: path,
	}
}

type BuyerJsonFile struct{
	path string
}

func (ldr *BuyerJsonFile)Load()(buyers map[int]models.Buyer, err error){
	file, err := os.Open(ldr.path)
	if err != nil {
		return
	}
	defer file.Close()

	var buyersJSON []models.BuyerDocResponse
	err = json.NewDecoder(file).Decode(&buyersJSON)
	if err != nil {
		return
	}

	buyers = make(map[int]models.Buyer)
	for _, buyer := range buyersJSON {
		buyers[buyer.Id] = mappers.BuyerDocToBuyer(buyer)
	}
	return
}