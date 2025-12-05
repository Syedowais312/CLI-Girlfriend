package cmd

import (
	"fmt"
	"github.com/Syedowais312/CLI-Girlfriend/internal"
	"github.com/spf13/cobra"
	"strings"
	
)

var persona string
var pixelCols int
var pixelRows int

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

		// NOTE: static pre-render removed. Animation will render frames including
		// the first frame to avoid races between two concurrent renderers.

	thinking := true

	go internal.RenderPixelAnimation([]string{
		"assets/image.png",
		"assets/image_2.png",
		"assets/image_3.png",
		"assets/image_4.png",
		"assets/image_5.png",
		}, pixelCols, pixelRows, &thinking, 120)

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
		fmt.Println(internal.Color(internal.Magenta, "\n"+role))
		fmt.Println(internal.Color(internal.Green, " "+resp))

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

	// Image rendering size (columns x rows). Lower cols/rows => larger "pixels" in terminal.
	chatCmd.Flags().IntVarP(&pixelCols, "cols", "c", 100, "Image columns for terminal rendering (higher = clearer image)")
	chatCmd.Flags().IntVarP(&pixelRows, "rows", "r", 80, "Image rows for terminal rendering (higher = clearer image)")
}
