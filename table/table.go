package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/text"
)

var (
	//go:embed "markup/table.html"
	markupTable []byte
)

type Table struct {
	compton.BaseElement
}

func (t *Table) AppendHead(columns ...string) *Table {

	// assuming the first element to be thead, or create a new one
	// if table has no children

	if len(t.Children) < 1 {
		t.Append(NewHead())
	}
	thead := t.Children[0]
	for _, col := range columns {
		th := NewTh()
		th.Append(text.New(col))
		thead.Append(th)
	}

	return t
}

func (t *Table) AppendRow(data ...string) *Table {

	// assuming the second element to be tbody, or create a new one
	// if table has fewer than 2 children

	if len(t.Children) < 2 {
		t.Append(NewBody())
	}
	tbody := t.Children[len(t.Children)-1]
	tr := NewTr()
	for _, col := range data {
		td := NewTd()
		td.Append(text.New(col))
		tr.Append(td)
	}
	tbody.Append(tr)

	return t
}

func New() *Table {
	return &Table{compton.BaseElement{Markup: markupTable}}
}
