package models

// TemplateData struct provides different fields to accommodate various types of data that might be sent to the template.
// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string      // A map for storing string-based key-value pairs.
	IntMap    map[string]int         // A map for integer-based key-value pairs.
	FloatMap  map[string]float64     // A map for float-based key-value pairs.
	Data      map[string]interface{} // A generic map for holding data of any type.
	CSRFToken string                 // A security token for protecting against CSRF attacks.
	Flash     string                 // A message shown to the user, typically after a successful action.
	Warning   string                 // A warning message to display in the UI.
	Error     string                 // An error message to display in the UI.
}
