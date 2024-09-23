package details_summary

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/svg_use"
	"golang.org/x/net/html/atom"
	"io"
)

const (
	registrationName      = "details-sum"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "style/details-summary.css"
	styleDetailsSummary []byte
)

type DetailsSummaryElement struct {
	compton.BaseElement
	r       compton.Registrar
	details compton.Element
}

func (dse *DetailsSummaryElement) Append(children ...compton.Element) {
	dse.details.Append(children...)
}

func (dse *DetailsSummaryElement) WriteStyles(w io.Writer) error {
	if dse.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleDetailsSummary, styleRegistrationName).WriteContent(w); err != nil {
			return err
		}
	}
	return dse.details.WriteStyles(w)
}

func (dse *DetailsSummaryElement) WriteRequirements(w io.Writer) error {
	return dse.details.WriteRequirements(w)
}

func (dse *DetailsSummaryElement) WriteDeferrals(w io.Writer) error {
	return dse.details.WriteDeferrals(w)
}

func (dse *DetailsSummaryElement) SummaryMarginBlockEnd(s size.Size) *DetailsSummaryElement {
	if summaries := dse.details.GetElementsByTagName(atom.Summary); len(summaries) > 0 {
		summaries[0].AddClass(class.MarginBlockEnd(s))
	}
	return dse
}

func (dse *DetailsSummaryElement) DetailsMarginBlockEnd(s size.Size) *DetailsSummaryElement {
	dse.details.AddClass(class.MarginBlockEnd(s))
	return dse
}

func (dse *DetailsSummaryElement) BackgroundColor(c color.Color) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.BackgroundColor(c))
	}
	return dse
}

func (dse *DetailsSummaryElement) ForegroundColor(c color.Color) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.ForegroundColor(c))
	}
	return dse
}

func (dse *DetailsSummaryElement) MarkerColor(c color.Color) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.MarkerColor(c))
	}
	return dse
}

func (dse *DetailsSummaryElement) getSummary() compton.Element {
	if summaries := dse.details.GetElementsByTagName(atom.Summary); len(summaries) > 0 {
		return summaries[0]
	}
	return nil
}

func (dse *DetailsSummaryElement) WriteContent(w io.Writer) error {
	return dse.details.WriteContent(w)
}

func (dse *DetailsSummaryElement) SetId(id string) {
	dse.details.SetId(id)
}

func Closed(r compton.Registrar, summary compton.Element) *DetailsSummaryElement {
	dse := &DetailsSummaryElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.DetailsSummary,
		},
		details: els.Details(),
		r:       r,
	}

	summaryElement := els.Summary()
	summaryElement.Append(svg_use.SvgUse(r, svg_use.Plus), summary)
	dse.details.Append(summaryElement)

	return dse
}

func Open(r compton.Registrar, summary compton.Element) *DetailsSummaryElement {
	dse := Closed(r, summary)
	dse.details.SetAttribute("open", "")
	return dse
}

func Toggle(r compton.Registrar, summary compton.Element, condition bool) *DetailsSummaryElement {
	var toggle func(compton.Registrar, compton.Element) *DetailsSummaryElement
	switch condition {
	case true:
		toggle = Open
	case false:
		toggle = Closed
	}
	return toggle(r, summary)
}
