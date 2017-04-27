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
	for taskNum:=0; taskNum<ntasks; taskNum++{

		var worker string
		worker = <-mr.registerChannel
		go func( worker string, taskNum int, phase jobPhase){
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
			}
		}(worker,taskNum, phase)
	}



	fmt.Printf("Schedule: %v phase done\n", phase)
}

