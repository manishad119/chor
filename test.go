/**********
##############################################
####$$$$$CONCRETE IMPLEMENTATION$$$$##########
##############################################
*/

package main

import (
	"fmt"
	"strings"
	"flag"
	"github.com/manish119/chor/jobiface"
)

type JobType1 struct {
	
	name string
	val int
}

type JobType2 struct {
	
	name string
	val string
}

type IPv4Addr uint32

func (job *JobType1) DoJob() interface{} {
	sqr:=job.val*job.val
	fmt.Printf("job %s: sqr(%v)=%v\n",job.name,job.val,sqr)
	return sqr


}

func (addr1 IPv4Addr) String() string {
	addr:=uint32(addr1)
	return fmt.Sprintf("%v:%v:%v:%v",(addr>>24)&0xff,(addr>>16)&0xff,(addr>>8)&0xff,addr&0xff)
}

func (job *JobType2) DoJob() interface{} {
	splits:=strings.Split(job.val,":")
	fmt.Printf("job %s:string '%v' splitted as %v\n",job.name,job.val,splits)
	return len(splits)

}

func (job *JobType1) Name() string {
	return job.name
}

func (job *JobType2) Name() string {
	return job.name
}

func printJobResults(jobstats []*jobiface.JobStat){
	for _,jobstat:=range jobstats {
		fmt.Printf("%v stat=%v\n",jobstat.Job.Name(),jobstat.Stat)
	}
}



//P
func main(){
	//fmt.Printf("%v\n",time.Second*50)
	
	ghi:="恢复会hdj"
	runelist:=[]rune(ghi)
	fmt.Printf("len=%v ind0=%v type %T, ind3=%v type %T %s\n",len(ghi),ghi[0],ghi[0],ghi[3],ghi[3],ghi)
	fmt.Printf("rune=%v len=%v\n",runelist,len(runelist))
	fmt.Printf("addr=%v\n",IPv4Addr(0x3afe235b))

	testmap:=make(map[string]int)
	testarray:=[]int{1,2,3}
	testslice:=testarray[1:]
	fmt.Printf("%v\n",testslice)
	testslice[0]=4
	fmt.Printf("%v\n",testarray)
	func(themap map[string]int){themap["mine"]=6}(testmap)
	func(slice []int){slice[1]=6}(testslice)
	func(arrayasslice []int){arrayasslice[0]=2}(testarray)
	fmt.Printf("Map=%v slice=%v\n",testmap,testslice)
	fmt.Printf("Array=%v\n",testarray)
	//String with 0

	ghi="hjfkjnd\x00hjj"
	fmt.Printf("%s byte len=%v char len=%v\n",ghi,len(ghi),len([]rune(ghi)))



	thptr:=flag.Int("tc",3,"thread count")
	flag.Parse()
	joblist:=[]jobiface.DoableJob{&JobType1{"job1",3},&JobType2{"job2","a:bc:def"},&JobType1{"job3",7},&JobType1{"job4",-8},&JobType2{"job5","usrname:password"},&JobType1{"job6",0}}
	results:=jobiface.DoAllJobs(joblist,*thptr)
	printJobResults(results)

}

