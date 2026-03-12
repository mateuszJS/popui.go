package props

import "github.com/a-h/templ"

// SidebarVariant constants
const (
	SidebarVariantDark string = "dark"
)

// Sidebar Templ component props
type Sidebar struct {
	Title      string
	Class      string
	Variant    string // "dark" for dark background variant, default is light
	Attributes templ.Attributes
}

// SidebarSection Templ component props
type SidebarSection struct {
	Title      string
	Class      string
	Attributes templ.Attributes
}

// SidebarItem defines the property for a single sidebar navigation item.
type SidebarItem struct {
	ID         string
	Class      string
	Selected   bool
	Attributes templ.Attributes
	Href       templ.SafeURL
}
