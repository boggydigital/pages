package details_toggle

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/colors"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/shared"
	"io"
)

const (
	elementNameTemplate = "details-"
	summaryMarginAttr   = "data-summary-margin"
	backgroundColorAttr = "data-background-color"
	foregroundColorAttr = "data-foreground-color"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/details-open.html"
	markupDetailsOpen []byte
	//go:embed "markup/details-closed.html"
	markupDetailsClosed []byte
)

type Details struct {
	compton.BaseElement
	wcr     compton.Registrar
	summary string
	open    bool
}

func openClosed(o bool) string {
	switch o {
	case true:
		return "open"
	case false:
		return "closed"
	}
	return ""
}

func (d *Details) Register(w io.Writer) error {
	openClosedName := elementNameTemplate + openClosed(d.open)
	if d.wcr.RequiresRegistration(openClosedName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(openClosedName)); err != nil {
			return err
		}
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, d.templateFragmentWriter); err != nil {
			return err
		}
	}
	return d.BaseElement.Register(w)
}

func (d *Details) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".OpenClosed":
		if _, err := io.WriteString(w, openClosed(d.open)); err != nil {
			return err
		}
	case ".HostBackgroundColor":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostBackgroundColor)); err != nil {
			return err
		}
	case ".HostForegroundColor":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostForegroundColor)); err != nil {
			return err
		}
	case ".HostSummaryMargin":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostSummaryMargin)); err != nil {
			return err
		}
	}
	return nil
}

func (d *Details) SetSummaryMargin(amount measures.Unit) *Details {
	d.SetAttr(summaryMarginAttr, amount.String())
	return d
}

func (d *Details) SetBackgroundColor(color colors.Color) *Details {
	d.SetAttr(backgroundColorAttr, color.String())
	return d
}

func (d *Details) SetForegroundColor(color colors.Color) *Details {
	d.SetAttr(foregroundColorAttr, color.String())
	return d
}

func (d *Details) Write(w io.Writer) error {
	markup := markupDetailsClosed
	if d.open {
		markup = markupDetailsOpen
	}
	return compton.WriteContents(bytes.NewReader(markup), w, d.elementFragmentWriter)
}

func (d *Details) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Summary":
		if _, err := io.WriteString(w, d.summary); err != nil {
			return err
		}
	case ".OpenClosed":
		if _, err := io.WriteString(w, openClosed(d.open)); err != nil {
			return err
		}
	case compton.ContentToken:
		fallthrough
	case compton.AttributesToken:
		if err := d.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func NewClosed(wcr compton.Registrar, summary string) *Details {
	return &Details{
		BaseElement: compton.BaseElement{
			Markup:  markupDetailsClosed,
			TagName: compton_atoms.DetailsClosed,
		},
		wcr:     wcr,
		open:    false,
		summary: summary,
	}
}

func NewOpen(wcr compton.Registrar, summary string) *Details {
	return &Details{
		BaseElement: compton.BaseElement{
			Markup:  markupDetailsOpen,
			TagName: compton_atoms.DetailsOpen,
		},
		wcr:     wcr,
		open:    true,
		summary: summary,
	}
}
