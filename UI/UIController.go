package ui

import (
	"encoding/json"
	"fmt"
	"strconv"

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

func NewUiControler( initialViews map[ViewID]View, initialViewID ViewID) *UIController{
  return &UIController{
    currentView: initialViews[initialViewID],
    views: initialViews,
  }
}



func (c *UIController) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case SwitchingViewMsg:
       if view, exists := c.views[msg.ViewID]; exists {
      c.currentView = view
            return c, c.currentView.Init()
    }
  case tea.WindowSizeMsg:
    c.width = msg.Width
    c.height = msg.Height
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
    ListViewID: listView,
    DetailViewID: detailView,
  }
  
  controller := NewUiControler(views, ListViewID)
  p := tea.NewProgram(controller, tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Println("Error : ", err)
  }
}

func mapAssets(jsonStr *string) map[string]string {
	type Asset struct {
		AssetID    string  `json:"asset_id"`
		Name       string  `json:"name"`
		IsCrypto   int8    `json:"type_is_crypto"`
		Volume_1hr float32 `json:"volume_1hrs_usd"`
	}


	var assets []Asset
	err := json.Unmarshal([]byte(*jsonStr), &assets)
	if err != nil {
		fmt.Println("Error:", err)
	}


	assetMap := make(map[string]string)

	for _, asset := range assets {
		if asset.IsCrypto == 0 {
			formattedKey := fmt.Sprintf("%-8s |    %s", asset.AssetID, asset.Name)
			assetMap[formattedKey] = "VOL 1h : " + strconv.FormatFloat(float64(asset.Volume_1hr), 'f', 0, 64) + " USD"
		}
	}
	return assetMap
}
