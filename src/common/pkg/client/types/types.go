package types

type Money struct {
	CurrencyCode string `json:"currencyCode,omitempty"`
	Units        int64  `json:"units,omitempty"`
	Nanos        int32  `json:"nanos,omitempty"`
}
type Product struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Picture     string   `json:"picture,omitempty"`
	PriceUsd    Money    `json:"priceUsd,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}
