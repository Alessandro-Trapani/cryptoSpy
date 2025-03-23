package ui

import (
	"encoding/json"
  "fmt"
	"cryptoSpy/HTTP"

	tea "github.com/charmbracelet/bubbletea"
)

type ViewID int

const (
	ListViewID ViewID = iota
	DetailViewID
)

type View interface {
	tea.Model
	ID() ViewID
}

type SwitchingViewMsg struct {
	ViewID ViewID
	Data   any
}

type UIController struct {
	currentView View
	views       map[ViewID]View
	width       int
	height      int
}
type Asset struct {
    AssetID        string         `json:"asset_id"`
    Name           string         `json:"name"`
    IsCrypto       int8           `json:"type_is_crypto"`
    Volume1hrUsd   float64        `json:"volume_1hrs_usd,omitempty"`
    Volume1dayUsd  float64        `json:"volume_1day_usd,omitempty"`
    Volume1mthUsd  float64        `json:"volume_1mth_usd,omitempty"`
    DataStart      string         `json:"data_start,omitempty"`
    DataEnd        string         `json:"data_end,omitempty"`
    PriceUsd       float64        `json:"price_usd,omitempty"`
    ChainAddresses json.RawMessage `json:"chain_addresses,omitempty"` 
}
func NewUiControler(initialViews map[ViewID]View, initialViewID ViewID) *UIController {
	return &UIController{
		currentView: initialViews[initialViewID],
		views:       initialViews,
	}
}

func (c *UIController) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case SwitchingViewMsg:
        if view, exists := c.views[msg.ViewID]; exists {
            if detailView, ok := view.(*DetailViewModel); ok && msg.ViewID == DetailViewID {
                if item, ok := msg.Data.(item); ok {
                    detailView.item = item
                }
            }
            
            // Clear the screen before switching views
            fmt.Print("\033[2J") // ANSI escape code to clear the screen
            fmt.Print("\033[H")  // Move cursor to home position
            
            c.currentView = view
            return c, c.currentView.Init()
        }
    // Other cases remain the same
    }
    
    updatedModel, cmd := c.currentView.Update(msg)
    if view, ok := updatedModel.(View); ok {
        c.currentView = view
    }
    return c, cmd
}
func (c *UIController) Init() tea.Cmd {
	return c.currentView.Init()
}
func (c *UIController) View() string {
	return c.currentView.View()
}

func StartUI() {
	assets := HTTP.GET("https://rest.coinapi.io/v1/assets", nil)

	assetMap := mapAssets(&assets)

	listView := NewListViewModel(assetMap, "Crypto")
	detailView := NewDetailViewModel()

	views := map[ViewID]View{
		ListViewID:   listView,
		DetailViewID: detailView,
	}

	controller := NewUiControler(views, ListViewID)
	p := tea.NewProgram(controller, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error : ", err)
	}
}

func mapAssets(jsonStr *string) map[*Asset]string {
    var assets []Asset
    err := json.Unmarshal([]byte(*jsonStr), &assets)
    if err != nil {
        fmt.Println("Error unmarshalling assets:", err)
        return nil
    }
    
    
    
    assetMap := make(map[*Asset]string)
    for i := range assets {
        // Make a deep copy to avoid pointer issues
        assetCopy := assets[i]
        
        // Create a display string with available information
        var displayInfo string
        
        if assetCopy.PriceUsd > 0 {
            displayInfo = fmt.Sprintf("Price USD: %.8f $", assetCopy.PriceUsd)
        } else {
            displayInfo = "Price: Not available"
        }
        
        // Add volume information if available
        if assetCopy.Volume1dayUsd > 0 {
            displayInfo += fmt.Sprintf(" | 24h Vol: %.2f $", assetCopy.Volume1dayUsd)
        }
        
       
        if assetCopy.IsCrypto == 1 {
            assetMap[&assetCopy] = displayInfo
        }
    }
    
        return assetMap
}
