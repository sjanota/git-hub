package config

type PullRequest struct {
	HeadRef  string
	HeadRepo string
	Number   int
	WebURL   string
	Remote   string
}
