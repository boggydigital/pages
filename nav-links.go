package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

type NavLinksElement struct {
	BaseElement
}

func NavLinks(r Registrar) *NavLinksElement {
	navLinks := &NavLinksElement{
		BaseElement: BaseElement{
			Markup:   markup,
			TagName:  compton_atoms.NavLinks,
			Filename: compton_atoms.MarkupName(compton_atoms.NavLinks),
		},
	}

	r.RegisterStyle(compton_atoms.StyleName(compton_atoms.NavLinks), style)

	return navLinks
}

func NavLinksTargets(r Registrar, targets ...*Target) *NavLinksElement {
	nl := NavLinks(r)
	for _, t := range targets {
		appendTarget(r, nl, t)
	}
	return nl
}

func appendTarget(r Registrar, nl *NavLinksElement, t *Target) {
	li := ListItem()
	link := A(t.Href)

	if t.Icon != None {
		icon := SvgUse(r, t.Icon)
		icon.SetAttribute("title", t.Title)
		link.Append(icon)
		if t.Current {
			link.Append(SpanText(t.Title))
		}
	} else {
		link.Append(Text(t.Title))
	}
	if t.Current {
		link.AddClass("selected")
	}
	li.Append(link)
	nl.Append(li)
}
