package models

type Url struct{
	Name string `json:"name"`
	Url string `json:"url"`
}
type UrlPortal struct {
	Title string `json:"title"`
	Urls *[]Url `json:"urls"`
}

type PortalData struct {
	Data []UrlPortal `json:"data"`
}
