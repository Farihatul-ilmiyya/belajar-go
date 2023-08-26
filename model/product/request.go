package product

type RequestProduct struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Weight int    `json:"weight"`
	Brand  string `json:"brand"`
}

func (req *RequestProduct) Validate() []string {
	var msg []string
	if req.Name == "" {
		msg = append(msg, "nama product tidak boleh kosong")
	}
	if req.Price == 0 {
		msg = append(msg, "price product tidak boleh kosong")
	}
	if req.Weight == 0 {
		msg = append(msg, "weight product tidak boleh kosong")
	}
	return msg
}
