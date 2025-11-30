# my-girlfriend 

A conversational CLI tool powered by **Google Gemini AI** that lets you chat with different AI personas. Choose from a supportive girlfriend, a senior engineer, or a helpful assistant—each with their own personality and communication style.

## Features

- **🎭 Multiple Personas**: Switch between different AI personalities
  - **Girlfriend** (default): Sweet, loving, and emotionally supportive
  - **Engineer**: Logical, technical, and detailed explanations
  - **Assistant**: General-purpose helpful assistant
- **💬 Conversation Memory**: Automatically saves and loads chat history for continuous conversations
- **⚡ Real-time Responses**: Powered by Google Gemini 2.0-flash for fast, intelligent responses
- **🎨 Interactive UI**: Animated thinking spinner while waiting for responses
- **📁 Persistent Storage**: Chat history stored in `~/.my-girlfriend/history.json`

## Project Structure

```
my-girlfriend/
├── main.go                 # Entry point
├── go.mod / go.sum        # Go dependencies (Gemini, Cobra CLI)
├── cmd/
│   ├── root.go            # Root CLI command setup (Cobra)
│   └── chat.go            # Chat command with persona support
├── internal/
│   ├── client.go          # Gemini API client and query logic
│   └── memory.go          # Chat history persistence utilities
├── model/
│   └── chatbot.go         # Data models (ChatMessage, ChatHistory)
├── .env                   # API keys (GEMINI_API_KEY)
└── RUN_ME.md             # Quick start guide
```

## Prerequisites

- **Go** >= 1.25.4
- **Google Gemini API Key** (free tier available at [Google AI Studio](https://aistudio.google.com/app/apikey))

## Installation & Setup

### 1. Clone or navigate to the project
```bash
cd /path/to/my-girlfriend
```

### 2. Get a Gemini API Key
- Visit [Google AI Studio](https://aistudio.google.com/app/apikey)
- Create a new API key (free tier available)
- Copy the key

### 3. Set the API Key
```bash
export GEMINI_API_KEY="your-gemini-api-key-here"
```

Or add it to your `.env` file:
```properties
GEMINI_API_KEY=your-gemini-api-key-here
```

Then source it:
```bash
source .env
```

### 4. Build the Project
```bash
go build -o my-girlfriend
```

## Usage

### Basic Chat
```bash
./my-girlfriend chat "Hello, how are you?"
```

### With Persona Flag
```bash
# Chat with girlfriend persona (default)
./my-girlfriend chat -p girlfriend "Tell me something sweet"

# Chat with engineer persona
./my-girlfriend chat -p engineer "Explain REST APIs"

# Chat with general assistant
./my-girlfriend chat -p girlfriend "What's 2 + 2?"
```

### One-liner (with API key)
```bash
GEMINI_API_KEY="your-key" ./my-girlfriend chat "Hello there!"
```

### Run without building (using `go run`)
```bash
go run . -- chat "Hello!"
```

### Input handling

- The `chat` command accepts a single prompt argument. If your message contains spaces, pass it as a single quoted string:
  ```bash
  ./my-girlfriend chat "How are you today?"
  ```
- If you omit quotes, only the first space-separated token will be used (e.g. `./my-girlfriend chat hello there` will use `hello`).

## Example Conversations

### Girlfriend Persona
```bash
 $ ./my-girlfriend chat -p girlfriend "I had a rough day"
 Thinking ⠋
 💖 Girlfriend:
 Oh no, I'm so sorry to hear that! 💔 Remember, I'm here for you. Want to talk about it? Whatever you're going through, we'll get through it together. 💕
```

### Engineer Persona
```bash
$ ./my-girlfriend chat -p engineer "How does REST work?"
Thinking | 
Response:
REST (Representational State Transfer) is an architectural style for building web APIs.
It uses HTTP methods (GET, POST, PUT, DELETE) to perform CRUD operations on resources...
```

## Conversation Memory

- Chat history is automatically saved to `~/.my-girlfriend/history.json`
- Each conversation includes the user's message and the AI's response
- History is loaded on startup, so the AI can provide context-aware responses
- To reset history, delete the file:
  ```bash
  rm ~/.my-girlfriend/history.json
  ```
 - Or use the helper function (if you add a command) to clear memory programmatically.

## API & Technology Stack

- **[Google Generative AI (Gemini)](https://ai.google.dev/)**: State-of-the-art language model (gemini-2.0-flash)
- **[Cobra](https://cobra.dev/)**: CLI framework for Go
- **Go 1.25.4**: Programming language

## Troubleshooting

### Error: "GEMINI_API_KEY not set"
- Ensure your API key is exported: `export GEMINI_API_KEY="your-key"`
- Verify it in the `.env` file and source it

### Error: "please provide a prompt"
- Provide a question after `chat` command:
  ```bash
  ./my-girlfriend chat "Your question here"
  ```

### Slow Responses
- First response may be slower (API initialization)
- Subsequent requests should be faster
- Check your internet connection

### History File Issues
- If history becomes corrupted, delete it:
  ```bash
  rm ~/.my-girlfriend/history.json
  ```
- A fresh history file will be created on next chat

## Building & Development

### Build from source
```bash
go build -o my-girlfriend
```

### Run in development mode
```bash
go run . -- chat "test message"
```

### View available commands
```bash
./my-girlfriend --help
./my-girlfriend chat --help
```

### Check installed dependencies
```bash
go list -m all
```

## Cost

- **Free**: Google Gemini API offers a free tier with rate limits
- Check [Google AI Pricing](https://ai.google.dev/pricing) for details
- Monitor your usage in [Google Cloud Console](https://console.cloud.google.com/)

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `GEMINI_API_KEY` | Yes | Google Gemini API Key for authentication |
| `HUGGING_FACE_API` | Optional | (Present in `.env` for alternate integrations) Not currently used by the Gemini client but kept for future features.

## License

See `LICENSE` file for details.

## Contributing

Feel free to fork, modify, and extend this project! Some ideas:
- Add more personas (motivational coach, therapist, comedian, etc.)
- Implement streaming responses
- Add conversation export (to PDF, CSV)
- Web UI frontend
- Multi-user support

## Feedback & Issues

If you encounter bugs or have feature requests, please document them clearly with:
- Steps to reproduce
- Expected behavior
- Actual behavior
- Your Go version and OS

---

**Made with ❤️ using Go and Google Gemini AI**
