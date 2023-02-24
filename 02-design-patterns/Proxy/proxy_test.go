package Proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := RentHouseProxy{rentInCity: &RentHouseCity{}}
	proxy.Rent(3000)
}
