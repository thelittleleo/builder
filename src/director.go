// Director 
package main

type director struct {
    builder Builder
}

func newDirector(b Builder) *director {
    return &director{
        builder: b,
    }
}

func (d *director) setBuilder(b iBuilder) {
    d.builder = b
}

func (d *director) buildReceiptForNew() receipt {
    d.builder.addMoney()
    d.builder.removeMoney()
    d.builder.addMoney()
    return d.builder.getReceipt()
}

func (d *director) buildReceiptForOld() receipt {
    d.builder.addMoney()
    d.builder.removeMoney()
    return d.builder.getReceipt()
}