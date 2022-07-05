package stringsvc

type (
	UpperStringRequest struct {
		S string `json:"s"`
	}
	UpperStringResponse struct {
		V   string `json:"v"`
		Err string `json:"err"`
	}
	CountRequest struct {
		S string `json:"s"`
	}
	CountResponse struct {
		V int `json:"v"`
	}
)
