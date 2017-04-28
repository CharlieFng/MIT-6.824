package mapreduce

import (
	"fmt"
	//"log"
)

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {

	var ntasks int
	var nios int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce
	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//


	fmt.Printf("Now in %s phase, total %d tasks\n", phase, ntasks)

	var execute func(worker string, taskNum int, phase jobPhase)
	execute = func( worker string, taskNum int, phase jobPhase){
		fmt.Printf("Worker %s will execute task %d on phase %s\n", worker, taskNum, phase)
		taskArgs := new(DoTaskArgs)
		taskArgs.JobName = mr.jobName
		taskArgs.File = mr.files[taskNum]
		taskArgs.Phase = phase
		taskArgs.TaskNumber = taskNum
		taskArgs.NumOtherPhase = nios

		ok := call(worker, "Worker.DoTask", &taskArgs, new(struct{}))
		if ok {
			mr.registerChannel <- worker
			fmt.Printf("Worker %s finish %d task in %v phase\n", worker, taskNum, phase)
		}else{
			fmt.Println("Catch worker failure")
			fmt.Println("Current workers in Master", mr.workers)
			for i, w := range mr.workers {
				if w == worker {
					mr.workers = append(mr.workers[:i], mr.workers[i+1:]...)
					break
				}
			}
			fmt.Println("After,  workers in Master", mr.workers)
			var w string
			if len(mr.workers) > 0 {
				fmt.Printf("Master current workers number %d\n", len(mr.workers))
				w = mr.workers[0]

			}else{
				fmt.Printf("No workers right now, has to wait", len(mr.workers))
				w = <-mr.registerChannel
			}
			execute(w, taskNum, phase)
		}
	}

	for taskNum:=0; taskNum<ntasks; taskNum++{

		var worker string
		worker = <-mr.registerChannel
		
		go execute(worker, taskNum, phase)

	}



	fmt.Printf("Schedule: %v phase done\n", phase)
}

