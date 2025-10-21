package rules

type Rules interface {
	GetAllRules() ([]string, error)
	GetRules(string) (map[string]any, error)
}
