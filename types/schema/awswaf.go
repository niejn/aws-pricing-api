package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Awswaf struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]Awswaf_Product
	Terms		map[string]map[string]Awswaf_Term
}
type Awswaf_Product struct {	Sku	string
	ProductFamily	string
	Attributes	Awswaf_Product_Attributes
}
type Awswaf_Product_Attributes struct {	Servicecode	string
	Location	string
	LocationType	string
	Group	string
	GroupDescription	string
	Usagetype	string
	Operation	string
}

type Awswaf_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions Awswaf_Term_PriceDimensions
	TermAttributes Awswaf_Term_TermAttributes
}

type Awswaf_Term_PriceDimensions struct {
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	Awswaf_Term_PricePerUnit
	AppliesTo	[]interface{}
}

type Awswaf_Term_PricePerUnit struct {
	USD	string
}

type Awswaf_Term_TermAttributes struct {

}
func (a *Awswaf) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/awswaf/current/index.json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, a)
	if err != nil {
		return err
	}

	return nil
}