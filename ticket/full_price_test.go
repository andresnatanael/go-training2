package customer

import "testing"

func TestFullPrice(t *testing.T) {

	var customer Customer = &FullPrice{}

	if customer.GetPrice() != 100 {
		t.Errorf("GOT: %d EXPECTED: %d", customer.GetPrice(), 100)
	}

}
