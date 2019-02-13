package scheduler

import "GoStudy/Crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workChan chan chan engine.Request

}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r

}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request){
	s.workChan <- w
}


func (s *QueuedScheduler) Run(){
	s.workChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ [] chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workQ)>0{
				activeRequest = requestQ[0]
				activeWorker = workQ[0]
			}
			select{

			case r:= <-s.requestChan:
				// send r to a ?woker
				requestQ = append(requestQ,r)
			case w:= <-s.workChan:
				// send ?next_request to w
				workQ = append(workQ,w)
			case activeWorker <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]

			}
		}
	}()
}
