// concrete Product 2
package main 

type specialCustomerBuilder struct {
    amount int
}

func newSpecialCustomerBuilder() *specialCustomerBuilder {
    return &specialCustomerBuilder{}
}

func(b *specialCustomerBuilder) addMoney() {
    b.amount+=20
}

func(b *specialCustomerBuilder) removeMoney() {
    b.amount-=3
}

func (b* specialCustomerBuilder) getReceipt() receipt {
    return receipt {
        amount: b.amount,
    }
}