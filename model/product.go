package model

type Product struct {
	Id    int
	Name  string
	Price float32
	Stock int
}

func (p Product) StatusStock() string {
	var status string

	if p.Stock < 3 {
		status = "Stock hampir habis"
	} else if p.Stock < 10 {
		status = "Stock terbatas"
	}
	return status
}
