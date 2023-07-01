package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

type Task struct {
	task func()
}

func NewTask(t func()) Task {
	return Task{
		task: t,
	}
}

type Pool struct {
	lock       sync.Mutex
	wg         sync.WaitGroup
	workerNum  int
	taskQueue  chan Task
	entryQueue chan Task
	ctx        context.Context
	cancel     context.CancelFunc
}

func (p *Pool) AddTasks(tasks []Task) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	select {
	case <-p.ctx.Done():
		return errors.New("Pool has already close")
	default:
	}
	p.wg.Add(len(tasks))
	for i := 0; i < len(tasks); i++ {
		p.entryQueue <- tasks[i]
	}
	return nil
}

func (p *Pool) AddTask() {
	for {
		select {
		case t := <-p.entryQueue:
			p.taskQueue <- t
		case <-p.ctx.Done():
			return
		}
	}
}

func NewPool(workerNum int) *Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &Pool{
		workerNum:  workerNum,
		taskQueue:  make(chan Task),
		entryQueue: make(chan Task),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (p *Pool) Start() {
	go p.AddTask()
	for i := 0; i < p.workerNum; i++ {
		go p.worker()
	}
}

func (p *Pool) worker() {
	for {
		select {
		case t := <-p.taskQueue:
			t.task()
			p.wg.Done()
		case <-p.ctx.Done():
			return
		}
	}
}

func (p *Pool) Close() {
	p.wg.Wait()
	//防止多并发环境下 一边执行stop,一边执行addTasks
	p.lock.Lock()
	defer p.lock.Unlock()
	p.cancel()
}

func main() {
	pool := NewPool(3)
	pool.Start()
	wg := sync.WaitGroup{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tasks := make([]Task, 0)
			for i := 0; i < 100; i++ {
				task := NewTask(func() {
					fmt.Println("test")
				})
				tasks = append(tasks, task)
			}
			err := pool.AddTasks(tasks)
			if err != nil {
				log.Println(err)
			}
		}()
	}
	wg.Wait()
	pool.Close()
	tasks := make([]Task, 0)
	for i := 0; i < 100; i++ {
		task := NewTask(func() {
			fmt.Println("test")
		})
		tasks = append(tasks, task)
	}
	err := pool.AddTasks(tasks)
	if err != nil {
		log.Println(err)
	}
}
