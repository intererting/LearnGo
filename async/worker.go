// Jason Waldrip协助完成了这个示例
// work包管理一个goroutine池来完成工作
package main

import "sync"

// Worker必须满足接口类型，
// 才能使用工作池
type Worker interface {
	Task()
}

// Pool提供一个goroutine池，这个池可以完成
// 任何已提交的Worker任务
type WorkerPool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New创建一个新工作池
func NewWorkPool(maxGoroutines int) *WorkerPool {
	p := WorkerPool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run提交工作到工作池
func (p *WorkerPool) Run(w Worker) {
	p.work <- w
}

// Shutdown等待所有goroutine停止工作
func (p *WorkerPool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
