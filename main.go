package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Ttype представляет задачу с временем создания, завершения и результатом
type Ttype struct {
	id         int
	cT         string // время создания
	fT         string // время завершения
	taskRESULT string // результат задачи
}

func main() {
	// 1. Логгирование:
	// В новом коде добавлено логгирование в файл app.log, что позволяет сохранять все логи в файл для последующего анализа.
	// В старом коде логгирование отсутствовало.
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла логов: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	fmt.Println("Файл журнала создан: app.log")

	var wg sync.WaitGroup
	taskChan := make(chan Ttype, 10)
	doneTasks := make(chan Ttype, 10)
	undoneTasks := make(chan Ttype, 10)

	// 3. Генерация задач:
	// В новом коде функция generateTasks генерирует задачи в течение 10 секунд и отправляет их в канал taskChan.
	// В старом коде генерация задач была реализована в анонимной функции, что усложняло понимание кода.
	go generateTasks(taskChan)
	// 4. Обработка задач:
	// В новом коде функция processTasks обрабатывает задачи из канала taskChan и сортирует их в каналы doneTasks и undoneTasks.
	// В старом коде обработка задач была реализована в анонимной функции, что усложняло понимание кода.
	go processTasks(taskChan, doneTasks, undoneTasks, &wg)
	// 6. Периодическая печать результатов:
	// В новом коде функция periodicResultPrinting периодически печатает результаты задач каждые 3 секунды.
	// В старом коде периодическая печать результатов отсутствовала.
	go periodicResultPrinting(doneTasks, undoneTasks)

	// 5. Асинхронная обработка:
	// В новом коде используется sync.WaitGroup для ожидания завершения всех задач.
	// В старом коде ожидание завершения задач было реализовано через time.Sleep, что не гарантировало корректное завершение всех задач.
	wg.Wait()
	log.Println("Все задачи завершены")
	time.Sleep(1 * time.Second) // гарантируем, что все результаты будут напечатаны перед выходом

	// 7. Завершение программы:
	// В новом коде добавлено ожидание нажатия любой клавиши пользователем для грациозного завершения программы.
	// В старом коде завершение программы было реализовано через time.Sleep, что не гарантировало корректное завершение всех задач.
	fmt.Println("\033[1;32mНажмите клавишу Enter для завершения программы...\033[0m")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// generateTasks генерирует задачи и отправляет их в канал taskChan
func generateTasks(taskChan chan Ttype) {
	defer close(taskChan)
	start := time.Now()
	for time.Since(start) < 10*time.Second {
		ft := time.Now().Format(time.RFC3339)
		if time.Now().Nanosecond()%2 > 0 {
			ft = "Some error occurred"
		}
		task := Ttype{cT: ft, id: int(time.Now().Unix())}
		log.Printf("Создана задача: %+v\n", task)
		taskChan <- task
		time.Sleep(100 * time.Millisecond) // имитация интервала генерации задач
	}
}

// processTasks обрабатывает задачи из канала taskChan и сортирует их в doneTasks и undoneTasks
func processTasks(taskChan, doneTasks, undoneTasks chan Ttype, wg *sync.WaitGroup) {
	for task := range taskChan {
		wg.Add(1)
		go func(t Ttype) {
			defer wg.Done()
			processedTask := taskWorker(t)
			taskSorter(processedTask, doneTasks, undoneTasks)
		}(task)
	}
	wg.Wait()
	close(doneTasks)
	close(undoneTasks)
	log.Println("Каналы doneTasks и undoneTasks закрыты")
}

// taskWorker обрабатывает одну задачу и возвращает её с результатом
func taskWorker(task Ttype) Ttype {
	tt, err := time.Parse(time.RFC3339, task.cT)
	if err != nil || tt.After(time.Now().Add(-20*time.Second)) {
		task.taskRESULT = "task has been successful"
	} else {
		task.taskRESULT = "something went wrong"
	}
	task.fT = time.Now().Format(time.RFC3339Nano)
	time.Sleep(150 * time.Millisecond) // имитация времени обработки задачи
	log.Printf("Обработана задача: %+v\n", task)
	return task
}

// taskSorter сортирует задачи по результату и отправляет их в соответствующие каналы
func taskSorter(task Ttype, doneTasks, undoneTasks chan Ttype) {
	if task.taskRESULT == "task has been successful" {
		doneTasks <- task
	} else {
		undoneTasks <- task
	}
}

// periodicResultPrinting периодически печатает результаты задач
func periodicResultPrinting(doneTasks, undoneTasks chan Ttype) {
	for {
		time.Sleep(3 * time.Second)
		log.Println("Периодическая печать результатов")
		printResults(doneTasks, undoneTasks)
	}
}

// printResults печатает результаты выполненных и невыполненных задач
func printResults(doneTasks, undoneTasks chan Ttype) {
	fmt.Println("Ошибки:")
	for len(undoneTasks) > 0 {
		task := <-undoneTasks
		fmt.Printf("Задача id %d время %s, ошибка %s\n", task.id, task.cT, task.taskRESULT)
	}

	fmt.Println("Выполненные задачи:")
	for len(doneTasks) > 0 {
		task := <-doneTasks
		fmt.Printf("Задача id %d время %s, результат %s\n", task.id, task.cT, task.taskRESULT)
	}
}
