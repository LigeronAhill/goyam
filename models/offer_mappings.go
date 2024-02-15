package models

type OfferMappingsResponse struct {
	Status string                    `json:"status"`
	Result GetOfferMappingsResultDTO `json:"result"`
}
type GetOfferMappingsResultDTO struct {
	Paging        ScrollingPagerDTO    `json:"paging"`
	OfferMappings []GetOfferMappingDTO `json:"offerMappings"`
}
type ScrollingPagerDTO struct {
	NextPageToken string `json:"nextPageToken"`
	PrevPageTokev string `json:"prevPageToken"`
}
type GetOfferMappingDTO struct {
	Offer   GetOfferDTO   `json:"offer"`
	Mapping GetMappingDTO `json:"mapping"`
}
type GetOfferDTO struct {
	OfferId               string                   `json:"offerId"`
	Name                  string                   `json:"name"`
	Category              string                   `json:"category"`
	Pictures              []string                 `json:"pictures"`
	Videos                []string                 `json:"videos"`
	Vendor                string                   `json:"vendor"`
	Barcodes              []string                 `json:"barcodes"`
	Description           string                   `json:"description"`
	ManufacturerCountries []string                 `json:"manufacturerCountries"`
	WeightDimensions      OfferWeightDimensionsDTO `json:"weightDimensions"`
	VendorCode            string                   `json:"vendorCode"`
	Tags                  []string                 `json:"tags"`
	ShelfLife             TimePeriodDTO            `json:"shelfLife"`
	LifeTime              TimePeriodDTO            `json:"lifeTime"`
	GuaranteePeriod       TimePeriodDTO            `json:"guaranteePeriod"`
	CustomsCommodityCode  string                   `json:"customsCommodityCode"`
	Certificates          []string                 `json:"certificates"`
	BoxCount              int                      `json:"boxCount"`
	Condition             OfferConditionDTO        `json:"condition"`
	Type                  string                   `json:"type"`
	Downloadable          bool                     `json:"downloadable"`
	Adult                 bool                     `json:"adult"`
	Age                   AgeDTO                   `json:"age"`
	Params                []OfferParamDTO          `json:"params"`
	BasicPrice            GetPriceWithDiscountDTO  `json:"basicPrice"`
	PurchasePrice         GetPriceDTO              `json:"purchasePrice"`
	AdditionalExpenses    GetPriceDTO              `json:"additionalExpenses"`
	CofinancePrice        GetPriceDTO              `json:"cofinancePrice"`
	CardStatus            string                   `json:"cardStatus"`
	Campaigns             []OfferCampaignStatusDTO `json:"campaigns"`
	SellingPrograms       []OfferSellingProgramDTO `json:"sellingPrograms"`
	Archived              bool                     `json:"archived"`
}
type GetMappingDTO struct {
	MarketSku          int    `json:"marketSku"`
	MarketSkuName      string `json:"marketSkuName"`
	MarketModelId      int    `json:"marketModelId"`
	MarketModelName    string `json:"marketModelName"`
	MarketCategoryId   int    `json:"marketCategoryId"`
	MarketCategoryName string `json:"marketCategoryName"`
}
type OfferWeightDimensionsDTO struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}
type TimePeriodDTO struct {
	TimePeriod int    `json:"timePeriod"`
	TimeUnit   string `json:"timeUnit"`
	Comment    string `json:"comment"`
}
type OfferConditionDTO struct {
	Type    string `json:"type"`
	Quality string `json:"quality"`
	Reason  string `json:"reason"`
}
type AgeDTO struct {
	Value   int    `json:"value"`
	AgeUnit string `json:"ageUnit"`
}
type OfferParamDTO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type GetPriceWithDiscountDTO struct {
	Value        float64 `json:"value"`
	CurrencyId   string  `json:"currencyId"`
	DiscountBase float64 `json:"discountBase"`
	UpdatedAt    string  `json:"updatedAt"`
}
type GetPriceDTO struct {
	Value      float64 `json:"value"`
	CurrencyId string  `json:"currencyId"`
	UpdatedAt  string  `json:"updatedAt"`
}

type OfferCampaignStatusDTO struct {
	CampaignId int    `json:"campaignId"`
	Status     string `json:"status"`
}
type OfferSellingProgramDTO struct {
	SellingProgram string `json:"sellingProgram"`
	Status         string `json:"status"`
}
