package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}
	out := in
	for _, stage := range stages {
		out = executeStage(out, done, stage)
	}
	return out
}

func executeStage(in In, done In, stage Stage) Out {
	out := make(Bi, 1)

	go func() {
		defer close(out)
		stageOut := stage(in)

		for {
			select {
			case <-done:
				return
			case val, ok := <-stageOut:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case out <- val:
				}
			}
		}
	}()
	return out
}
