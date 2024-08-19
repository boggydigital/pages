package nav_links

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/els"
	"github.com/boggydigital/compton/svg_icons"
	"io"
)

const (
	navLinksElementName = "nav-links"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/nav-links.html"
	markupNavLinks []byte
)

type NavLinks struct {
	compton.BaseElement
	wcr compton.Registrar
}

func (nl *NavLinks) Register(w io.Writer) error {
	if nl.wcr.RequiresRegistration(navLinksElementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(navLinksElementName)); err != nil {
			return err
		}
		if _, err := io.Copy(w, bytes.NewReader(markupTemplate)); err != nil {
			return err
		}
	}
	return nl.BaseElement.Register(w)
}

func New(wcr compton.Registrar) *NavLinks {
	return &NavLinks{
		BaseElement: compton.BaseElement{
			Markup:  markupNavLinks,
			TagName: compton_atoms.NavLinks,
		},
		wcr: wcr,
	}
}

func NewLinks(wcr compton.Registrar, targets ...*Target) *NavLinks {
	nl := New(wcr)

	for _, t := range targets {
		appendTarget(nl, t)

	}

	return nl
}

func appendTarget(nl *NavLinks, t *Target) {
	link := els.NewA(t.Href)

	if t.Icon != svg_icons.None {
		icon := svg_icons.NewIcon(t.Icon)
		icon.SetClass("icon")
		icon.SetAttr("title", t.Title)
		link.Append(icon)
		if t.Current {
			link.Append(els.NewSpanText(t.Title))
		}
	} else {
		link.Append(els.NewText(t.Title))
	}
	if t.Current {
		link.SetClass("selected")
	}
	nl.Append(link)
}
