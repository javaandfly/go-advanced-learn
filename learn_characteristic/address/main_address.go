package address

import "testing"

type Address struct {
	Address []string
}

func TestAddress(t *testing.T) {
	address := &Address{}
	addAddress(address)
}

func addAddress(address *Address) *Address {
	address.Address = append(address.Address, "12112")
	return address
}
