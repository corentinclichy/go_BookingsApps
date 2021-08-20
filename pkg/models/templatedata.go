package models

// @Title TemplateData
// @Description this is the template data handler
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Waring    string
	Eror      string
}
