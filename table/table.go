package table

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/text"
	"io"
)

var (
	//go:embed "markup/table.html"
	markupTable []byte
)

type Table struct {
	compton.Parent
	Id        string
	ClassList []string
}

func (tbl *Table) Append(children ...compton.Component) compton.Component {
	tbl.Children = append(tbl.Children, children...)
	return tbl
}

func (t *Table) AppendHead(columns ...string) *Table {

	// assuming the first element to be thead, or create a new one
	// if table has no children

	if len(t.Children) < 1 {
		t.Append(NewHead())
	}
	thead := t.Children[0]
	for _, col := range columns {
		th := NewTh().Append(text.New(col))
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
		tr.Append(NewTd().Append(text.New(col)))
	}
	tbody.Append(tr)

	return t
}

func (tbl *Table) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupTable), w, tbl.writeTableFragment)
}

func (tbl *Table) writeTableFragment(t string, w io.Writer) error {
	switch t {
	case ".Id":
		if tbl.Id != "" {
			if err := compton.WriteId(w, tbl.Id); err != nil {
				return err
			}
		}
	case ".ClassList":
		if len(tbl.ClassList) > 0 {
			if err := compton.WriteClassList(w, tbl.ClassList...); err != nil {
				return err
			}
		}
	case ".TableContent":
		for _, child := range tbl.Children {
			if err := child.Write(w); err != nil {
				return err
			}
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func New(id string, classList ...string) *Table {
	return &Table{Id: id, ClassList: classList}
}
