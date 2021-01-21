// concrete builder 1
package main 

type normalCustomerBuilder struct {
    amount int
}

func newNormalCustomerBuilder() *normalCustomerBuilder {
    return &normalCustomerBuilder{}
}

func(b *normalCustomerBuilder) addMoney() {
    b.amount+=10
}

func(b *normalCustomerBuilder) removeMoney() {
    b.amount-=2
}

func (b *normalCustomerBuilder) getReceipt() receipt {
    return receipt {
        amount: b.amount,
    }
}
