package main

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"testing"
	"time"
)

func TestProgram(t *testing.T) {
	// Захватываем вывод программы
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Создаем канал для завершения программы
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		main()
	}()

	// Отправляем сигнал остановки программе
	go func() {
		time.Sleep(2 * time.Second)
		close(exit)
	}()

	// Ждем, пока программа завершится
	select {
	case <-exit:
	case <-time.After(5 * time.Second):
		t.Fatal("program did not exit in a timely manner")
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("failed to read program output: %v", err)
	}

	output := buf.String()
	expectedPattern := `NTP Time: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+\d{2}:\d{2}`

	// Компилируем регулярное выражение
	re, err := regexp.Compile(expectedPattern)
	if err != nil {
		t.Fatalf("failed to compile regex pattern: %v", err)
	}

	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		line := scanner.Text()
		if !re.MatchString(line) {
			t.Errorf("output line does not match expected pattern: %s", line)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("failed to read program output: %v", err)
	}

	t.Log("Program output:", output)
}
