package Mediator

import "testing"

func TestEstateMedium(t *testing.T) {
	medium := EstateMedium{}
	buyer := &Buyer{
		Customer: Customer{
			name:   "张三",
			id:     "1001",
			medium: &medium,
		},
		money: 800000,
	}
	seller := &Seller{
		Customer: Customer{
			name:   "李老板",
			id:     "8005",
			medium: &medium,
		},
		price:     8100,
		houseArea: 175,
	}

	medium.Register(buyer)
	medium.Register(seller)

	seller.Send("每平米涨价500！！！")
	buyer.Send("不买了。")
}
