package result

type Result struct {
	Subdomain string   `csv:"subdomain"`
	Answers   []string `csv:"answers"`
}
