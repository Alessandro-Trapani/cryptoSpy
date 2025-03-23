package ui

import (
	"fmt"
  "strings"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type DetailViewModel struct {
	item item
	id   ViewID
  width int
  height int
}
    
func NewDetailViewModel() *DetailViewModel {
	return &DetailViewModel{
		id: DetailViewID,
	}
}

func (m *DetailViewModel) ID() ViewID {
	return m.id
}

func (m *DetailViewModel) Init() tea.Cmd {
	return nil
}

func (m *DetailViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
       
        m.width = msg.Width
        m.height = msg.Height
        
    case SwitchingViewMsg:
        if msg.Data != nil {
            if selectedItem, ok := msg.Data.(item); ok {
                m.item = selectedItem
            }
        }
        
    case tea.KeyMsg:
        if msg.String() == "esc" {
            return m, func() tea.Msg {
                return SwitchingViewMsg{
                    ViewID: ListViewID,
                }
            }
        }
    }
    return m, nil
}

func (m *DetailViewModel) View() string {
	asset := m.item.asset

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500"))
	labelStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF"))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	divider := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Render(strings.Repeat("─", 40))
	marginStyle := lipgloss.NewStyle().Margin(1, 2)

	content := fmt.Sprintf(
		"%s\n\n%s\n\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n\n(Press ESC to go back)",
		titleStyle.Render("Asset Details"),
		divider,
		labelStyle.Render("ID:")+" "+valueStyle.Render(asset.AssetID),
		labelStyle.Render("Name:")+" "+valueStyle.Render(asset.Name),
		labelStyle.Render("Type:")+" "+valueStyle.Render(ifElse(asset.IsCrypto == 1, "Cryptocurrency", "Fiat")),
		labelStyle.Render("Price (USD):")+" "+valueStyle.Render(formatFloat(asset.PriceUsd)),
		labelStyle.Render("Volume 1 hour (USD):")+" "+valueStyle.Render(formatFloat(asset.Volume1hrUsd)),
		labelStyle.Render("Volume 1 day (USD):")+" "+valueStyle.Render(formatFloat(asset.Volume1dayUsd)),
		labelStyle.Render("Volume 1 month (USD):")+" "+valueStyle.Render(formatFloat(asset.Volume1mthUsd)),
		labelStyle.Render("Creation date")+" "+valueStyle.Render(asset.DataStart),
		labelStyle.Render("Data End:")+" "+valueStyle.Render(asset.DataEnd),
		divider,
	)

	return marginStyle.Render(content)
}

func formatFloat(value float64) string {
	if value == 0 {
		return "N/A"
	}
  if value < 1 {
    	return fmt.Sprintf("%.8f", value)
  }
	return fmt.Sprintf("%.2f", value)
}

func ifElse(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
