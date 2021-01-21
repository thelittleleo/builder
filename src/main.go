// Client code
package main

import "fmt"

func main() {
    normalCustomerBuilder := getBuilder("new customer")
    specialCustomerBuilder := getBuilder("old customer")

    director := newDirector(normalCustomerBuilder)
    normalReceipt := director.buildReceiptForNew()
    fmt.Printf("Total amount: %d\n", normalReceipt.amount)

    director.setBuilder(specialCustomerBuilder)
    specialReceipt := director.buildReceiptForOld()
    fmt.Printf("Normal House Door Type: %s\n", specialReceipt.amount)

}