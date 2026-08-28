package cake_test

import (
	"testing"
	"time"

	"github.com/yuhaowin/go-learning/ch08/cake"
)

func defaults() cake.Shop {
	return cake.Shop{
		Verbose:      testing.Verbose(),
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
}

func Benchmark(b *testing.B) {
	// Baseline: one baker, one icer, one inscriber.
	// Each step takes exactly 10ms.  No buffers.
	bakeshop := defaults()
	bakeshop.Work(b.N) // 224 ms
}

func BenchmarkBuffers(b *testing.B) {
	// Adding buffers has no effect.
	bakeshop := defaults()
	bakeshop.BakeBuf = 10
	bakeshop.IceBuf = 10
	bakeshop.Work(b.N) // 224 ms
}

func BenchmarkVariable(b *testing.B) {
	// Adding variability to rate of each step
	// increases total time due to channel delays.
	bakeshop := defaults()
	bakeshop.BakeStdDev = bakeshop.BakeTime / 4
	bakeshop.IceStdDev = bakeshop.IceTime / 4
	bakeshop.InscribeStdDev = bakeshop.InscribeTime / 4
	bakeshop.Work(b.N) // 259 ms
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Adding channel buffers reduces
	// delays resulting from variability.
	bakeshop := defaults()
	bakeshop.BakeStdDev = bakeshop.BakeTime / 4
	bakeshop.IceStdDev = bakeshop.IceTime / 4
	bakeshop.InscribeStdDev = bakeshop.InscribeTime / 4
	bakeshop.BakeBuf = 10
	bakeshop.IceBuf = 10
	bakeshop.Work(b.N) // 244 ms
}

func BenchmarkSlowIcing(b *testing.B) {
	// Making the middle stage slower
	// adds directly to the critical path.
	bakeshop := defaults()
	bakeshop.IceTime = 50 * time.Millisecond
	bakeshop.Work(b.N) // 1.032 s
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// Adding more icing cooks reduces the cost of icing
	// to its sequential component, following Amdahl's Law.
	bakeshop := defaults()
	bakeshop.IceTime = 50 * time.Millisecond
	bakeshop.NumIcers = 5
	bakeshop.Work(b.N) // 288ms
}
