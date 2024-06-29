package hw06pipelineexecution

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	sleepPerStage = time.Millisecond * 100
	fault         = sleepPerStage / 2
)

func TestPipeline(t *testing.T) {
	// Stage generator
	g := func(_ string, f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage)
					out <- f(v)
				}
			}()
			return out
		}
	}

	stages := []Stage{
		g("Dummy", func(v interface{}) interface{} { return v }),
		g("Multiplier (* 2)", func(v interface{}) interface{} { return v.(int) * 2 }),
		g("Adder (+ 100)", func(v interface{}) interface{} { return v.(int) + 100 }),
		g("Stringifier", func(v interface{}) interface{} { return strconv.Itoa(v.(int)) }),
	}

	t.Run("simple case", func(t *testing.T) {
		in := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Equal(t, []string{"102", "104", "106", "108", "110"}, result)
		require.Less(t,
			int64(elapsed),
			// ~0.8s for processing 5 values in 4 stages (100ms every) concurrently
			int64(sleepPerStage)*int64(len(stages)+len(data)-1)+int64(fault))
	})

	t.Run("done case", func(t *testing.T) {
		in := make(Bi)
		done := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		// Abort after 200ms
		abortDur := sleepPerStage * 2
		go func() {
			<-time.After(abortDur)
			close(done)
		}()

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, done, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Len(t, result, 0)
		require.Less(t, int64(elapsed), int64(abortDur)+int64(fault))
	})
}

func TestPipe(t *testing.T) {
	// Stage generator
	g := func(f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage)
					out <- f(v)
				}
			}()
			return out
		}
	}

	stages := []Stage{
		g(func(v interface{}) interface{} { return v }),
		g(func(v interface{}) interface{} { return v.(int) * 2 }),
		g(func(v interface{}) interface{} { return v.(int) + 100 }),
		g(func(v interface{}) interface{} { return strconv.Itoa(v.(int)) }),
	}

	t.Run("empty input", func(t *testing.T) {
		in := make(Bi)
		close(in) // Close the input channel immediately

		result := make([]string, 0)
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s.(string))
		}

		require.Empty(t, result)
	})

	t.Run("midway done", func(t *testing.T) {
		in := make(Bi)
		done := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		// Send done signal after 350ms (midway through processing)
		go func() {
			time.Sleep(350 * time.Millisecond)
			close(done)
		}()

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0)
		for s := range ExecutePipeline(in, done, stages...) {
			result = append(result, s.(string))
		}

		require.True(t, len(result) < len(data))
	})

	t.Run("different data types", func(t *testing.T) {
		in := make(Bi)
		data := []interface{}{1, "two", 3.0, true}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		stages := []Stage{
			g(func(v interface{}) interface{} { return v }),
			g(func(v interface{}) interface{} {
				switch v := v.(type) {
				case int:
					return v * 2
				case string:
					return v + v
				case float64:
					return v * 2
				case bool:
					return !v
				default:
					return v
				}
			}),
		}

		result := make([]interface{}, 0)
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s)
		}

		require.Equal(t, []interface{}{2, "twotwo", 6.0, false}, result)
	})
}
