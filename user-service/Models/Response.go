package Models

type Meta struct {
	Code    int     `json:"code"`
	Message *string `json:"message"`
}

type Response struct {
	Meta   Meta        `json:"meta"`
	Result interface{} `json:"result"`
}

type Empty struct {
}
