package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
}

/*
Run ...
*/
func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	// Channels
	in := make(chan Request)
	out := make(chan ParseResult)
	for i :=0; i < e.WorkerCount; i ++{
		createWorker(in, out)

	}

	
	for {
		result := <- out
		for _ , item : range result.Items {
			fmt.Printf("Got item: %v", item)
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker(in chan Request, out chan ParseResult){
	go func(){
		for {
			request := <- in
			result, err := worder(request)
			if err != nil{
				continue
			}
			out <-result
		}()
	}
}
