# Pokedex CLI

A small interactive REPL written in Go that talks to the [PokeAPI](https://pokeapi.co/).
Browse the Pokémon world, explore location areas, catch Pokémon, and inspect the ones
you've caught — all from your terminal.

## Features

- Interactive read-eval-print loop (REPL)
- Paginated browsing of location areas with `map` / `mapb`
- Explore a location area to see which Pokémon appear there
- Catch Pokémon with a chance-based Pokéball throw
- Inspect the stats and types of anything you've caught
- In-memory response cache with periodic eviction (5-minute TTL)
- Graceful handling of empty input, EOF, and HTTP errors

## Requirements

- Go 1.27 or newer (an internet connection is needed — data comes from PokeAPI)

## Build & Run

```bash
go build -o pokedexcli .
./pokedexcli
```

Or run without building:

```bash
go run .
```

## Usage

Once running you'll see a prompt:

```
Pokedex >
```

Type a command and press Enter. Press `Ctrl+D` (EOF) or type `exit` to quit.

## Commands

| Command | Description |
| --- | --- |
| `help` | Show the available commands |
| `map` | List the next 20 location areas |
| `mapb` | List the previous 20 location areas |
| `explore <area>` | Show the Pokémon found in a location area |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | Show stats and types of a caught Pokémon |
| `pokedex` | List all Pokémon you've caught |
| `exit` | Exit the Pokedex |

### Example session

```
Pokedex > map
canalave-city-area
eterna-city-area
...

Pokedex > explore canalave-city-area
Pokemon:
 - tentacool
 - staryu
...

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > inspect tentacool
Name: tentacool
Height: 9
Weight: 455
Stats:
  -hp: 40
  -attack: 40
  -defense: 35
  ...
Types:
  - water
  - poison
```

## Project layout

| File / dir | Purpose |
| --- | --- |
| `main.go` | Entry point, command registry, shared config and HTTP client |
| `repl.go` | REPL loop and input parsing |
| `command_*.go` | One file per command (`help`, `exit`, `map`, `explore`, `catch`, `inspect`, `pokedex`) |
| `type_pokemon.go` | Structs mirroring the PokeAPI Pokémon resource |
| `cache/pokecache.go` | Thread-safe TTL cache used to avoid repeat API calls |

## Testing

```bash
go test ./...
```

The cache package has tests covering expiry, concurrency, copy semantics, and clean
shutdown. Input parsing (`cleanInput`) is covered in `repl_test.go`.

## Notes

- Responses are cached in memory for 5 minutes; cache keys are full request URLs so
  different endpoints never collide.
- HTTP requests use a 10-second timeout.
- This is a learning project built while working through the boot.dev Go course.
