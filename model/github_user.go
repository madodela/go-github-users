package model

type GithubUser struct {
	ID          int
	Login       string
	AvatarUrl   string
	HtmlUrl     string
	Name        string
	Company     string
	Bio         string
	PublicRepos int
}
