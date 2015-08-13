package datastructures

import (
	//"github.com/oleiade/lane"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const BENCH_Q_SIZE int = 1000000

func TestPriorityQueue_order_int(t *testing.T) {
	pq := NewPriorityQueue(Less(func(x, y interface{}) bool {
		return x.(int) < y.(int)
	}), 0)

	//Populate the priority queue with random ints
	var src rand.Source = rand.NewSource(0)
	var r *rand.Rand = rand.New(src)
	for i := 0; i < 10; i++ {
		assert.True(
			t,
			pq.Length() == i,
			"pq.Length() = %d; want %d", pq.Length(), i,
		)
		pq.Insert(r.Int())
	}
	var prev int = pq.PopTop().(int)
	var next int
	for pq.Length() > 0 {
		next = pq.PopTop().(int)
		assert.True(
			t,
			prev <= next,
			"%d sorted before %d; want %d sorted after %d", prev, next, prev, next,
		)
	}
}

func BenchmarkPriorityQueue_int(b *testing.B) {
	pq := NewPriorityQueue(Less(func(x, y interface{}) bool {
		return x.(int) < y.(int)
	}), BENCH_Q_SIZE)
	var src rand.Source = rand.NewSource(0)
	var r *rand.Rand = rand.New(src)
	for n := 0; n < b.N; n++ {
		//Populate the priority queue with random ints
		for i := 0; i < BENCH_Q_SIZE; i++ {
			pq.Insert(r.Int())
		}
		//Empty the priority queue
		for i := 0; i < BENCH_Q_SIZE; i++ {
			pq.PopTop()
		}
	}
}

/*
func BenchmarkLanePriorityQueue_int(b *testing.B) {
	pq := lane.NewPQueue(lane.MINPQ)
	var src rand.Source = rand.NewSource(0)
	var r *rand.Rand = rand.New(src)
	for n := 0; n < b.N; n++ {
		//Populate the priority queue with random ints
		for i := 0; i < BENCH_Q_SIZE; i++ {
			pq.Push(0, r.Int())
		}
		//Empty the priority queue
		for i := 0; i < BENCH_Q_SIZE; i++ {
			pq.Pop()
		}
	}
}
*/

func TestPriorityQueue_order_Time(t *testing.T) {
	pq := NewPriorityQueue(Less(func(x, y interface{}) bool {
		return x.(time.Time).Before(y.(time.Time))
	}), 10)

	//Populate the priority queue with random times
	var src rand.Source = rand.NewSource(0)
	var r *rand.Rand = rand.New(src)
	for i := 0; i < 10; i++ {
		assert.True(
			t,
			pq.Length() == i,
			"pq.Length() = %d; want %d", pq.Length(), i,
		)
		pq.Insert(time.Now().Add(time.Hour * time.Duration(r.Int())))
	}
	var prev time.Time = pq.PopTop().(time.Time)
	var next time.Time
	for pq.Length() > 0 {
		next = pq.PopTop().(time.Time)
		assert.True(
			t,
			prev.Before(next),
			"%s sorted before %s; want %s sorted after %s", prev, next, prev, next,
		)
	}
}

func BenchmarkPriorityQueue_Time(b *testing.B) {
	pq := NewPriorityQueue(Less(func(x, y interface{}) bool {
		return x.(time.Time).Before(y.(time.Time))
	}), BENCH_Q_SIZE)
	var src rand.Source = rand.NewSource(0)
	var r *rand.Rand = rand.New(src)
	for n := 0; n < b.N; n++ {
		//Populate the priority queue with random ints
		for i := 0; i < BENCH_Q_SIZE; i++ {
			pq.Insert(time.Now().Add(time.Hour * time.Duration(r.Int())))
		}
		//Empty the priority queue
		for i := 0; i < BENCH_Q_SIZE; i++ {
			pq.PopTop()
		}
	}
}
