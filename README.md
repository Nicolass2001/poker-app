# 🃏 Multiplayer Poker Game

A small project I built to learn WebSockets and practice Go! It's a fully functional Texas Hold'em poker game where you can play with friends online in real-time.

## What I Built

Just wanted to challenge myself with real-time multiplayer gaming, so I implemented:

- **Live multiplayer** using WebSockets (no page refreshes needed!)
- **Complete poker engine** with all the game logic from scratch
- **Room-based system** so friends can create private games
- **Proper hand evaluation** - handles all poker hands from high card to royal flush
- **Full betting mechanics** with blinds, raises, calls, folds, all-ins

## Tech Stack & Skills Demonstrated

**Backend (Go)**

- **Gorilla WebSocket** for real-time bidirectional communication
- **Concurrent programming** with goroutines and mutexes for thread-safe game state
- **HTTP server** built with Go's standard `net/http` package
- **Clean architecture** with separated concerns (handlers, game logic, models)
- **Custom poker engine** implementing complex game rules and hand evaluation algorithms

**Frontend (Vanilla JS/HTML/CSS)**

- **WebSocket client** handling real-time game updates
- **Event-driven architecture** for responsive UI interactions
- **CSS Grid/Flexbox** for poker table layout
- **ES6 modules** for organized JavaScript code

**System Design**

- **Multi-room architecture** supporting concurrent games
- **State management** for complex poker game flows
- **RESTful endpoints** for room creation/joining + WebSocket for gameplay
- **Error handling** and input validation throughout the stack

## Project Structure

```
poker-app/
├── main.go                    # Main server entry point
├── cmd/main.go                # Command-line poker game demo
├── handlers/                  # HTTP and WebSocket handlers
├── poker/                     # Core poker game logic
├── static/                    # Frontend assets
└── templates/                 # HTML templates
```

## Quick Start

Need Go 1.24+ and a modern browser.

```bash
git clone <repo-url>
cd poker-app
go run main.go
```

Then open `http://localhost:8080` and you're good to go!

Want to see the poker logic in action? Try the CLI demo:

```bash
go run cmd/main.go
```

## How to Play

### Creating a Room

1. Go to the home page (`http://localhost:8080`)
2. Fill in the "Create Room" form:
   - **Small Blind**: The small blind amount (default: 100)
   - **Big Blind**: The big blind amount (default: 200)
   - **Starting Chips**: Initial chip count for all players (default: 1500)
   - **Nickname**: Your player name
3. Click "Create" to create the room and join as the first player

### Joining a Room

1. Get the room code from the room creator
2. Fill in the "Join Room" form:
   - **Room Code**: 4-character room code
   - **Nickname**: Your unique player name
3. Click "Join" to enter the room

### Gameplay

- **Starting the Game**: Once all players have joined, click "Start Game"
- **Player Actions**:
  - **Check**: Pass without betting (when no bet is required)
  - **Call**: Match the current highest bet
  - **Raise**: Increase the bet amount
  - **Fold**: Surrender your hand
- **Game Flow**: Pre-flop → Flop → Turn → River → Showdown
- **Winning**: The player with the best hand (or last remaining player) wins the pot

## Game Rules

This application implements standard Texas Hold'em poker rules:

- Each player receives 2 hole cards
- 5 community cards are dealt in stages (3 on flop, 1 on turn, 1 on river)
- Players make the best 5-card hand using any combination of hole and community cards
- Standard poker hand rankings apply (Royal Flush, Straight Flush, Four of a Kind, etc.)

## What I Learned Building This

**Go Backend Skills:**

- **Gorilla WebSocket package** - Real-time bidirectional communication
- **Concurrency patterns** with goroutines and mutexes for thread-safe operations
- **HTTP routing** and middleware using standard library
- **Clean code architecture** separating game logic from handlers
- **Complex algorithm implementation** (hand evaluation, betting logic)

**Frontend Development:**

- **WebSocket client implementation** in vanilla JavaScript
- **Event-driven programming** for real-time UI updates
- **CSS Grid/Flexbox** for responsive poker table layout
- **Module-based JavaScript** for maintainable code structure

**System Architecture:**

- **Multi-room concurrent server** handling multiple games simultaneously
- **State synchronization** between server and multiple clients
- **RESTful API design** combined with WebSocket real-time features
- **Error handling strategies** throughout the full stack

## Cool Technical Bits

Some interesting problems I solved:

- **Hand evaluation algorithm** that finds the best 5-card combination from 7 cards
- **Betting round state machine** handling complex poker betting rules
- **Thread-safe game state** with proper mutex usage for concurrent access
- **WebSocket message routing** with JSON-based action/response system
- **Poker table UI** built with pure CSS (no frameworks!)

The CLI demo in `cmd/main.go` is pretty neat - shows off the core poker engine without any UI.
