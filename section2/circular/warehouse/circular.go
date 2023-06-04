package warehouse

import (
	"context"

	"github.com/Hnampk/se-with-golang/section2/circular/entity"
)

// Robot navigates the warehouse floor and fetches items for packing.
type Robot struct {
	// various fields
}

// AcquireRobot blocks until a Robot becomes available or until the
// context expires.
func AcquireRobot(ctx context.Context) *Robot { //...
	return nil
}

// Pack instructs the robot to pick up an item from its shelf and place
// it into a box that will be shipped to the customer.
func (r *Robot) Pack(item *entity.Item, to *entity.Box) error { //...
	return nil
}
