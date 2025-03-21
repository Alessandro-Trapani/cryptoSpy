package ui
import (
	"fmt"
  lipgloss "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)
type DetailViewModel struct {
    item item
    id   ViewID
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
    case SwitchingViewMsg:
   fmt.Printf("DATA : %s ",msg.Data)      
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
  
    title := m.item.Title()      
    desc := m.item.Description() 

    content := fmt.Sprintf("Title: %s\n\nDescription: %s\n\n(Press ESC to go back)", title, desc)
    
    return lipgloss.NewStyle().Margin(1, 2).Render(content)
}
