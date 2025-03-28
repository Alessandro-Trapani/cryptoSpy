# cryptoSpy - Terminal-based Cryptocurrency Information Viewer

## Overview

cryptoSpy is a terminal application built with Golang that allows users to browse and explore cryptocurrency information. Using Charm's TUI (Text User Interface) libraries, this app provides an interactive interface for exploring cryptocurrency data from the CoinAPI.

## Features

- Full list of cryptocurrencies from CoinAPI
- Interactive terminal-based user interface
- Easy navigation with keyboard controls
- Real-time search and filtering functionality
- Detailed cryptocurrency information view

## Technologies Used

- **Language**: Go (Golang)
- **TUI Libraries**: 
  - [Bubbletea](https://github.com/charmbracelet/bubbletea) - Main TUI framework
  - [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling
  - [Bubbles/List](https://github.com/charmbracelet/bubbles) - List component
- **API**: CoinAPI Rest API

## Prerequisites

- Go installed
- CoinAPI API Key

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Alessandro-Trapani/cryptoSpy.git
   cd cryptoSpy
   ```

2. Create a `config.go` file with your API key:
   ```go
   package config

   const API_KEY = "<your_api_key_here>"
   ```
## Usage

### Running the Application

```bash
go run main.go
```

Alternatively, you can build and run the executable:

```bash
go build
./cryptoSpy    # On Linux/Mac
cryptoSpy.exe  # On Windows
```

### Navigation Controls

- **Up/Down Arrows** or **j/k**: Navigate through cryptocurrency list
- **Enter**: Select cryptocurrency to view detailed information
- **/** : Open search input to filter cryptocurrencies
- **Ctrl+C**: Exit the application

## Keyboard Shortcuts

| Key | Action |
|-----|--------|
| ↑/↓ | Navigate list |
| j/k | Navigate list |
| Enter | View crypto details |
| / | Search/Filter |
| Ctrl+C | Exit |

## Detailed Information Displayed

When selecting a cryptocurrency, the app shows:
- Unique Identifier
- Name
- Type (Crypto/Fiat)
- Current Price
- 1-Hour Trading Volume
- 1-Day Trading Volume
- 1-Month Trading Volume
- Creation Date
- End Date (the current date if not applicable)
