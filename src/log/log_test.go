package log

import (
	"context"
	"testing"

	"github.com/Kidsunbo/kie_toolbox_go/logs"
)

func TestLog(t *testing.T) {
	logs.CtxInfo(context.TODO(), "A = %d", 10)
}
