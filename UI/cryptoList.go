package ui

import (


	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }



type ListViewModel struct {
	items []item
	List  list.Model
	id    ViewID
}


func NewListViewModel(items map[string]string, title string) *ListViewModel {

    listItems := []list.Item{}
    for key, value := range items {
        listItems = append(listItems, item{title: key, desc: value})
    }

    m := &ListViewModel{
        List: list.New(listItems, list.NewDefaultDelegate(), 0, 0),
        id:   ListViewID,
    }
    m.List.Title = title
    
    return m
}
func (m *ListViewModel) ID() ViewID {
	return ListViewID
}

func (m ListViewModel) Init() tea.Cmd {
	return nil
}

func (m *ListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

switch msg := msg.(type) { 
  case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
    
		} else if m.List.FilterState() != list.Filtering && msg.String() == "enter" {
			selected := m.List.SelectedItem()
      return m, func() tea.Msg {
       
				return SwitchingViewMsg{
					ViewID: DetailViewID,
					Data:   selected,
				}
			}

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)

	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m ListViewModel) View() string {
	return docStyle.Render(m.List.View())
}


