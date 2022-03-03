package address

type Addresses struct {
	A_id    int    `json:"a_id"`
	Street  string `json:"street"`
	Area    string `json:"area"`
	Pincode int    `json:"pincode"`
	City    string `json:"city"`
}

type JsonResponse struct {
	Data    []Addresses
	Message string
}
