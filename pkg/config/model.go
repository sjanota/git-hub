package config

type PullRequest struct {
	HeadRef  string
	HeadRepo string
	BaseRef  string
	BaseRepo string
	Number   int
	WebURL   string
	Remote   string
}
