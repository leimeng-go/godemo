package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const (
	shortDuration = 1 * time.Millisecond //例子中合理的阻塞持续时间
)

func TestContextDeline(t *testing.T) {
	rootCtx := context.Background()
	newCtx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go func(ctx context.Context, number int) {
			ticker := time.NewTicker(time.Duration(number+1) * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					t.Logf("%d:超时上下文开始过了", number)
					return
				case <-ticker.C:
					t.Logf("%d:协程定时任务", number)
				}
			}

		}(newCtx, i)
	}
	select {}
}

func TestContextWithCancel(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 返回时goroutine不会泄漏
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		t.Log(n)
		if n == 6 {
			break
		}
	}
}

// TestContextWithDeadline 在某一个时刻取消
func TestContextWithDeadline(t *testing.T) {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}
}

func TestContextWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("overslept")
	case <-ctx.Done():
		t.Log(ctx.Err())
	}
}

func TestContextWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			t.Logf("found value: %+v", v)
			return
		}
		t.Logf("key not found: %s", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
