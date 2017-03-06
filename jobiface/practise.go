package jobiface
	
import (
	"time"
	
)
/*****
####################################
###########$$INTERFACES$$##########
#####################################
**/

/*A performable task to do for any thread
It returns one value of arbitrary type
If you need  additional information, do it by modifying
your concrete DoableJob object preferably
a struct pointer. The concrete job object should
be the complete job description itself and it does
not take any external parameters*/
/*IMPORTANT! DoableJob must be a pointer type and
not a struct type to prevent runtime error*/
type DoableJob interface {
	DoJob() interface{} //Do the specific job
	Name() string  //get the Job name

}

type JobStat struct {
	Job DoableJob
	Stat interface{}

}
/*Do all jobs in the list by delegating them
to goroutines*/
func DoAllJobs(joblist []DoableJob,thcnt int) []*JobStat{
	//Make a channel and add all jobs to it
	ch:=make(chan DoableJob,len(joblist))
	//Make an empty channel to get statuses
	och:=make(chan JobStat,len(joblist))
	oput:=make([]*JobStat,0)




	for _,v:=range joblist {
		ch<-v
	}

	
	for i:=0;i<thcnt;i++{
		go startJobs(ch,och,i)
	}
	//Wait for all the outputs
	deadthcnt:=0

	//Wait for all threads to signal that they have completed
	//the job

	for deadthcnt<thcnt {
		jobstat:=<-och
		if nil!=jobstat.Job{
			oput=append(oput,&jobstat)
		} else {
			deadthcnt++
			//fmt.Printf("deadthread=%v\n",deadthcnt)

		}
		time.Sleep(20*time.Millisecond)		
	}

	return oput


	

}

func startJobs(ch chan DoableJob,och chan JobStat,thid int){
	//fmt.Printf("Goroutine %v started\n",thid)
	for {
		select {
		case job:=<-ch:
			//Sleep for 30 millisecond before 
			//starting the job
			time.Sleep(20*time.Millisecond)
			//fmt.Printf("goroutine %v doing job %s\n",thid,job.Name())
			stat:=job.DoJob()
			och<-JobStat{job,stat}
		default:
			//If no remaining job, signal to main thread and return
			//Main thread quits after it receives signals
			//from all the goroutines. nil Job is the signal
			//fmt.Printf("goroutine %v exiting\n",thid)
			och <-JobStat{nil,0}
			return
		}

	}
}

