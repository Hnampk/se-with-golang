package circular

import (
	"context"

	"github.com/Hnampk/se-with-golang/section2/circular/entity"
	"github.com/Hnampk/se-with-golang/section2/circular/warehouse"
)

func wireComponents() {
	entity.AcquirePacker = func(ctx context.Context) entity.Packer {
		return warehouse.AcquireRobot(ctx)
	}
}
