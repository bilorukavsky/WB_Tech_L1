/*
Реализовать все возможные способы остановки выполнения горутины.
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина начала работу")
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Горутина завершила работу")
	case <-ctx.Done():
		fmt.Println("Горутина была отменена")
	}
}

func task_6_2() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker2(ctx, &wg)

	// Через какое-то время отменяем выполнение горутины
	time.Sleep(1 * time.Second)
	cancel()

	wg.Wait()
}
