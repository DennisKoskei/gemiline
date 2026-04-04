package session

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/denniskoskei/gemiline/internal/gemini"
	"github.com/denniskoskei/gemiline/internal/ui"
)

type Thread struct {
	Title   string
	History []string
}

type Session struct {
	Threads []Thread
	Current int
	ApiKey  string
}

func NewSession(apiKey string) *Session {
	return &Session{
		Threads: []Thread{},
		Current: -1,
		ApiKey:  apiKey,
	}
}

func (s *Session) Start(initialQuery string) {
	s.createThread(initialQuery)
	s.run()
}

func (s *Session) createThread(query string) {
	response, _ := gemini.AskGemini(s.ApiKey, query)

	thread := Thread{
		Title:   query,
		History: []string{query, response},
	}

	s.Threads = append(s.Threads, thread)
	s.Current = len(s.Threads) - 1

	ui.PrintResponse(response)
}

func (s *Session) run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("|---| T - Enter Text | M - Main Menu | Q/E - Exit |---|")
		fmt.Print(":> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		switch input {

		case "E":
			fmt.Println("Exiting program. Have a nice day!")
			return
		case "Q":
			fmt.Println("Exiting program. Have a nice day!")
			return
		case "QUIT":
			fmt.Println("Exiting program. Have a nice day!")
			return

		case "T":
			s.continueThread(reader)

		case "M":
			s.showMenu(reader)

		default:
			fmt.Println("Invalid option")
		}
	}
}

func (s *Session) continueThread(reader *bufio.Reader) {
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	response, err := gemini.AskGemini(s.ApiKey, text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	thread := &s.Threads[s.Current]
	thread.History = append(thread.History, text, response)

	ui.PrintResponse(response)
}

func (s *Session) showMenu(reader *bufio.Reader) {
	fmt.Println("\n|----------- GEMILINE -----------|")
	fmt.Println("\n|            MAIN MENU           |")
	fmt.Println("\n|                                |")

	for i, t := range s.Threads {
		fmt.Printf("%d - %s\n", i+1, t.Title)
	}
	// fmt.Println("\n|--------------------------------|")

	fmt.Print(":> ")
	var choice int
	fmt.Scanln(&choice)

	if choice > 0 && choice <= len(s.Threads) {
		s.Current = choice - 1
		s.showThread()
	}
}

func (s *Session) showThread() {
	thread := s.Threads[s.Current]

	fmt.Println("\n|----------  THREAD  ----------|")

	for i := 0; i < len(thread.History); i += 2 {
		fmt.Println("Q:", thread.History[i])
		fmt.Println("A:", thread.History[i+1])
		fmt.Println()
	}
}
