package main

type Builder interface {
    addMoney()
    removeMoney()
    getReceipt() receipt
}

func getBuilder(builderType string) Builder {
    if builderType == "new customer" {
        return &normalCustomerBuilder{}
    } 

    if builderType == "old customer" {
        return &specialCustomerBuilder{}
    }
}