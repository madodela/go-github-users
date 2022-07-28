// Package repository contains functionality to handle GithubUsers model information
package repository

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/madodela/go-github-users/model"
)

// GithubUser struct will contain the information of the datasource
type GithubUser struct {
	csvFileName string
	csvFilePath string
}

type IGithubUserRepository interface {
	FetchAll() ([]model.GithubUser, error)
}

// NewGithubUser provides an instance of
func NewGithubUser(fileName, filePath string) *GithubUser {
	return &GithubUser{
		fileName,
		filePath,
	}
}

// FetchAll returns an array of all the available users, or any error if it happens
func (g *GithubUser) FetchAll() ([]model.GithubUser, error) {
	githubUsers, err := g.csvToGithubUsers()
	if err != nil {
		return nil, err
	}
	return githubUsers, nil
}

// csvToGithubUsers function is private, it takes care of reading and providing the info from the CSV file if available
func (g *GithubUser) csvToGithubUsers() ([]model.GithubUser, error) {
	file, err := os.Open(g.csvFilePath + "/" + g.csvFileName)
	if err != nil {
		return nil, ErrOpeningCSV
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	var githubUsers []model.GithubUser

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return []model.GithubUser{}, ErrReadingLineCSV
		}

		if len(record) < 8 {
			return []model.GithubUser{}, ErrInvalidFileStruct
		}

		publicRepos, err := strconv.Atoi(record[7])
		if err != nil {
			publicRepos = 0
		}

		if ID, err := strconv.Atoi(record[0]); err == nil {
			githubUsers = append(githubUsers, model.GithubUser{
				ID:          ID,
				Login:       record[1],
				AvatarUrl:   record[2],
				HtmlUrl:     record[3],
				Name:        record[4],
				Company:     record[5],
				Bio:         record[6],
				PublicRepos: publicRepos,
			})
		}
	}
	return githubUsers, nil
}
