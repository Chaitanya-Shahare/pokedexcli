# Pokedex CLI

Pokedex CLI is a command-line interface application written in Go. It allows
users to interact with the Pokemon API to explore the map, catch and inspect Pokemon.

## Features

### Commands

- **Map Locations**: The application allows users to explore the map using the
  `map` & `mapb` command. This will display the list of Pokemon available in the map.
  - `map` command displays next 20 locations in the map.
  - `mapb` command displays the previous 20 locations in the map.
- **Explore Location**: Users can explore the location for Pokemons using the `explore`
  command followed by the location's name.
- **Catch Pokemon**: The application allows users to catch Pokemon using the
  `catch` command followed by the Pokemon's name. This depends on Pokemon's base level.
- **Inspect Pokemon**: Users can inspect the caught Pokemon using the `inspect`
  command followed by the Pokemon's name.
- **Pokedex**: Users can view the list of Pokemon they have caught using the
  `pokedex` command.

### Implementation

- **API**: The application interacts with the [Pokemon API](https://pokeapi.co/).
- **Cache**: The application caches the Pokemon data to reduce the number of API
  calls.

## Installation

To install the application, you need to have Go installed on your machine. You
can download it from [here](https://golang.org/dl/). Once Go is installed, you
can clone this repository and build the application.

```bash
git clone https://github.com/Chaitanya-Shahare/pokedexcli.git
cd pokedexcli
go build
```

## Usage

To run the application, you can use the following command:

```bash
./pokedexcli
```

### Map Locations

```bash
pokedex > map
```

### Explore Location

```bash
pokedex > explore location_name
```

### Catch Pokemon

```bash
pokedex > catch pokemon_name
```

### Inspect Pokemon

```bash
pokedex > inspect pokemon_name
```

### Pokedex

```bash
pokedex > pokedex
```

### Exit

```bash
pokedex > exit
```
