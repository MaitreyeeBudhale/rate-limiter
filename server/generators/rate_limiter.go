package generators

// generate the request as per stated by client side ie. simulation
import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/MaitreyeeBudhale/rate-limiter.git/config"
)

func createWorker(totalRequest int, nextRequest *atomic.Int64, accepted *atomic.Int64, rejected *atomic.Int64, limiter config.Limiter, wg *sync.WaitGroup, w int) {
	defer wg.Done()

	for {
		id := nextRequest.Add(1)

		if id > int64(totalRequest) {
			return
		}

		if limiter.Allow() {
			fmt.Println(id, "accepted by wroker", w)
			accepted.Add(1)
		} else {
			rejected.Add(1)
			fmt.Println(id, "rejected by worker", w)
		}
	}
}

func RateLimiter(totalRequests int, workers int, limiter config.Limiter) config.Stats {
	var wg sync.WaitGroup
	var nextRequest atomic.Int64 //shared counter for the requests processed

	// shared counter for accept or reject status
	var accepted atomic.Int64
	var rejected atomic.Int64

	for w := 0; w < workers; w++ {
		wg.Add(1)

		// create new go routine for each worker
		go createWorker(totalRequests, &nextRequest, &accepted, &rejected, limiter, &wg, w)
	}
	// main function waits for the workers to generate all the requests
	wg.Wait()

	return config.Stats{
		Accepted: accepted.Load(),
		Rejected: rejected.Load(),
	}
}
