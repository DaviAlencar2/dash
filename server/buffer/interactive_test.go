package buffer

import (
	"bufio"
	"dash/consts"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var interactive = flag.Bool("interactive", false, "abrir o teste manual do buffer")

// Run: go test ./server/buffer -v -count=1 -args -interactive
func TestBufferInteractive(t *testing.T) {
	if !*interactive {
		t.Skip("use -args -interactive para abrir o teste manual")
	}
	terminal, err := os.Open("/dev/tty")
	if err != nil {
		t.Fatalf("não foi possível abrir o terminal: %v", err)
	}
	defer terminal.Close()

	b := &Buffer[string]{}
	help := func() {
		fmt.Println("add <texto> | pop | show | fill <1-100> | reset | help | quit")
	}
	itemLabel := func(item *Item[string]) string {
		if item == nil {
			return "nil"
		}
		return fmt.Sprintf("%q (%p)", item.value, item)
	}
	show := func() {
		fmt.Printf("head=%s | tail=%s | contador=%d | capacidade=%d\n",
			itemLabel(b.head), itemLabel(b.tail), b.len, consts.BUFFER_SIZE)
		seen := make(map[*Item[string]]bool)
		values := []string{}
		for item := b.head; item != nil; item = item.next {
			if seen[item] {
				fmt.Printf("ALERTA: ciclo detectado em %s\n", itemLabel(item))
				break
			}
			seen[item] = true
			values = append(values, fmt.Sprintf("%q", item.value))
		}
		fmt.Printf("Cadeia: [%s] | itens alcançáveis=%d\n", strings.Join(values, " -> "), len(seen))
		if b.len != len(seen) {
			fmt.Println("ALERTA: contador diferente da quantidade de itens alcançáveis")
		}
		if b.len > consts.BUFFER_SIZE {
			fmt.Println("ALERTA: contador acima da capacidade configurada")
		}
	}
	// A broken buffer may panic. Keep the terminal available for inspection/reset.
	operate := func(action func()) {
		defer func() {
			if problem := recover(); problem != nil {
				fmt.Printf("ALERTA: operação causou panic: %v\n", problem)
			}
			show()
		}()
		action()
	}

	fmt.Println("Buffer interativo: observando a implementação atual, sem corrigir inconsistências.")
	help()
	show()
	scanner := bufio.NewScanner(terminal)
	sequence := 0
	for {
		fmt.Print("buffer> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		command, argument := parts[0], ""
		if len(parts) == 2 {
			argument = strings.TrimSpace(parts[1])
		}
		if command != "add" && command != "fill" && argument != "" {
			fmt.Println("Este comando não recebe argumentos.")
			continue
		}
		switch command {
		case "add":
			if argument == "" {
				fmt.Println("Uso: add <texto>")
				continue
			}
			operate(func() { b.Queue(NewItem(argument)) })
		case "pop":
			operate(func() {
				fmt.Printf("Removido: %s\n", itemLabel(b.Dequeue()))
			})
		case "show":
			show()
		case "fill":
			count, err := strconv.Atoi(argument)
			if err != nil || count < 1 || count > 100 {
				fmt.Println("Uso: fill <1-100>")
				continue
			}
			for range count {
				sequence++
				value := fmt.Sprintf("item-%d", sequence)
				fmt.Printf("Adicionando %s\n", value)
				operate(func() { b.Queue(NewItem(value)) })
			}
		case "reset":
			b = &Buffer[string]{}
			sequence = 0
			show()
		case "help":
			help()
		case "quit":
			return
		default:
			fmt.Println("Comando desconhecido.")
			help()
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("erro lendo o terminal: %v", err)
	}
	fmt.Println("\nSessão encerrada (EOF).")
}
