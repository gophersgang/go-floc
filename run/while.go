package run

import floc "github.com/workanator/go-floc"

/*
While repeats running the job while the condition is met.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE

Diagram:
                    YES
    +-------[JOB]<------+
    |                   |
    V                   | NO
  ----(CONDITION MET?)--+---->
*/
func While(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		for predicate(state) && !flow.IsFinished() {
			job(flow, state, update)
		}
	}
}
