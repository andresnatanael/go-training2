package customer

const (
	value int = 100
)

//FullPrice structure
type FullPrice struct {
	name string
}

//GetPrice implementation
func (c *FullPrice) GetPrice() int {
	return value
}
