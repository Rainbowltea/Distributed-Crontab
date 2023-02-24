package worker

import (
	"Distributed-Crontab/pkg"
	p1 "Distributed-Crontab/pkg/balance"
	_ "context"
	"os/exec"
	"time"
)

//任务执行器
type Executor struct {
}

var (
	G_executor *Executor
)

//执行一个任务
func (executor *Executor) ExecuteJob(info *pkg.JobExecuteInfo) {
	go func() {
		var (
			cmd     *exec.Cmd
			err     error
			output  []byte
			result  *pkg.JobExecuteResult
			jobLock *JobLock
		)
		//任务结果
		result = &pkg.JobExecuteResult{
			ExecuteInfo: info,
			Output:      make([]byte, 0),
		}

		//获取分布式锁
		//初始化分布式锁
		jobLock = G_jobMgr.CreateJobLock(info.Job.Name)
		//记录任务开始时间
		result.StartTime = time.Now()

		//上锁
		//随机睡眠(0~1s)
		//解决锁倾斜问题
		//添加权重法计算节点CPU、内存使用情况来决定自身休眠多少秒
		n := p1.GetNodeUsage()
		time.Sleep(time.Duration(n) * time.Millisecond)
		err = jobLock.TryLock()
		defer jobLock.Unlock()

		if err != nil { //上锁失败
			result.Err = err
			result.EndTime = time.Now()
		} else {
			//上锁成功后，还需要重置任务启动时间
			result.StartTime = time.Now()

			//执行shell命令
			cmd = exec.CommandContext(info.CancelCtx, "bin/bash", "-c", info.Job.Command) //Linux系统中，在windows下要改成"c:\\cygwin64\\bin\\bash.exe"

			//执行并捕获输出
			//疑问点：
			//代码在61行等待命令执行完时宕机，
			output, err = cmd.CombinedOutput()

			//记录任务结束时间
			result.EndTime = time.Now()
			result.Output = output
			result.Err = err
		}
		//任务执行后，把执行的结果返回给scheduler,Scheduler会从executingTable中删除掉执行记录
		G_scheduler.PushJobResult(result)
	}()
}

//初始化一个执行器
func InitExecutor() (err error) {
	G_executor = &Executor{}
	return
}
