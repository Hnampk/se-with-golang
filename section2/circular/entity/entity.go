package entity

import (
	"context"
	// "github.com/Hnampk/se-with-golang/section2/circular/warehouse"
)

// Packer is implemented by objects that can pack an Item into a Box.
type Packer interface {
	Pack(*Item, *Box) error
}

type Item struct {
}

// Box contains a list of items that are shipped to the customer.
type Box struct {
	// various fields
}

// AcquirePacker returns a Packer instance.
var AcquirePacker func(context.Context) Packer

// Pack qty items of type i into the box.
func (b *Box) Pack(i *Item, qty int) error {
	// robot := warehouse.Acquire() // compile error: import cycle detected
	// ...
	p := AcquirePacker(context.Background())

	for j := 0; j < qty; j++ {
		if err := p.Pack(i, b); err != nil {
			return err
		}
	}
	return nil
}
