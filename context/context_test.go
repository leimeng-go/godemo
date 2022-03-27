package context

import (
	"context"
	"testing"
	"time"
)

func TestContextDeline(t *testing.T) {
	rootCtx:=context.Background()
	newCtx, cancel:=context.WithTimeout(rootCtx,2*time.Second)
	defer cancel()

    for i:=0;i<5;i++{
		go func(ctx context.Context,number int) {
			ticker:=time.NewTicker(time.Duration(number+1)*time.Second)
			defer ticker.Stop()
			for{
				select {
				case <-ctx.Done():
					t.Logf("%d:超时上下文开始过了",number)
					return
				case <-ticker.C:
					t.Logf("%d:协程定时任务",number)
				}
			}

		}(newCtx,i)
	}
	select {
	}
}
