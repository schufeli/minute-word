package word

type Word struct {
	Word        string   `json:"word"`
	Grammar     string   `json:"grammar"`
	Explanation string   `json:"explanation"`
	Example     string   `json:"example"`
	Usages      []string `json:"usages"`
}
