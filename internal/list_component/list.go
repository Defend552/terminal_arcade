package list_component

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	bubbles "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item string
type ItemDelegate struct{}

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

func (i Item) FilterValue() string {
	return string(i)
}

func (d ItemDelegate) Height() int {
	return 1
}

func (d ItemDelegate) Spacing() int {
	return 0
}

func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func New(title string, items []string, width int) bubbles.Model {
	l := bubbles.New(
		toItems(items),
		ItemDelegate{},
		width,
		listHeight,
	)

	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle

	return l
}

func toItems(strs []string) []bubbles.Item {
	items := make([]bubbles.Item, len(strs))
	for i, s := range strs {
		items[i] = Item(s)
	}
	return items
}

func SelectedString(l list.Model) string {
	selected, ok := l.SelectedItem().(Item)
	if !ok {
		return ""
	}
	return string(selected)
}
