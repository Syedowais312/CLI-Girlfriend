package cmd

import (
	"fmt"
	"github.com/Syedowais312/CLI-Girlfriend/internal"
	"time"
    "strings"
	"github.com/spf13/cobra"
)

var persona string

var chatCmd = &cobra.Command{
	Use:   "chat [Input]",
	Short: "Chat with your AI assistant ",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return fmt.Errorf("please provide a prompt: my-girlfriend chat \"your question here\"")
		}

		prompt := args[0]

		// Choose persona
		systemPrompt := getPersonaPrompt(persona)

		// Spinner animation
		frames := []string{"⠋","⠙","⠹","⠸","⠼","⠴","⠦","⠧","⠇","⠏"}

		thinking := true
		go func() {
			i := 0
			for thinking {
				fmt.Printf("\r%s %s",
					internal.Color(internal.Yellow, "Thinking"),
					internal.Color(internal.Cyan, frames[i%len(frames)]),
				)
				time.Sleep(120 * time.Millisecond)
				i++
			}
			fmt.Print("\r                          \r") // clear line
		}()
		// Query Gemini
		resp, err := internal.QueryChatbotAPI(systemPrompt, prompt)
		thinking = false

		if err != nil {
			return err
		}
		role := "Assistant"
		switch strings.ToLower(persona) {
		case "engineer":
			role = "engineer"
		case "girlfriend":
			role = "💖 Girlfriend"
		default:
			role = "Assistant"
		}
		fmt.Println(internal.Color(internal.Magenta,"\n"+role))
		fmt.Println(internal.Color(internal.Green," "+resp))

		return nil
	},
}

func getPersonaPrompt(p string) string {
	switch p {
	case "girlfriend":
		return "You are a sweet, loving, supportive girlfriend who replies gently, emotionally, and warmly. Use cute emojis. Never be inappropriate."

	case "engineer":
		return "You are a senior software engineer. Be logical, detailed, and technical. Explain concepts clearly with examples."

	default:
		return "You are a helpful AI assistant."
	}
}

func init() {
	rootCmd.AddCommand(chatCmd)

	chatCmd.Flags().StringVarP(
		&persona,
		"persona",
		"p",
		"girlfriend",
		"Choose persona: girlfriend, engineer",
	)
}
