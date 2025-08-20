package bigo

import "testing"

// benchmark find nemo using time module
// func BenchmarkFindNemoByTime(arr []string) {

// 	start := time.Now()

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == "nemo" {
// 			fmt.Println("Found...")
// 		}
// 	}

// 	elapsed := time.Since(start)

// 	fmt.Println("took", elapsed.Nanoseconds())
// }

// benchmark find nemo using benchmark module
func BenchmarkFindNemo(b *testing.B) {
	// var fish = []string{"nemo"}
	var lotsOfFish = make([]string, 10000)

	for i, _ := range lotsOfFish {
		lotsOfFish[i] = "dory"
	}
	lotsOfFish[9999] = "nemo"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindNemo(lotsOfFish)
	}
}

func BenchmarkGetFirstFish(b *testing.B) {

	arr := make([]string, 1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetFirstFish(arr)

	}

}

/**
ns/op → average time per call

B/op → bytes allocated per call

allocs/op → allocations per call

Test 1: 1000 entries
374284              6646 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/Crn-devs/dsa-go/bigo 2.534s

Test 2: 10,000 entries
54511             22511 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/Crn-devs/dsa-go/bigo 1.456s


**/
