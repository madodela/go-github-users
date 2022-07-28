package repository_test

import (
	"testing"

	"github.com/madodela/go-github-users/repository"
)

func TestFetchAll_ErrorHandling(t *testing.T) {
	cases := []struct {
		repository *repository.GithubUser
		returnErr  error
		name       string
	}{
		{
			repository: repository.NewGithubUser("empty.csv", "../testdata/users"),
			returnErr:  nil,
			name:       "Read data from empty file",
		},
		{
			repository: repository.NewGithubUser("invalid.csv", "../testdata/users"),
			returnErr:  repository.ErrInvalidFileStruct,
			name:       "Read data from invalid file",
		},
		{
			repository: repository.NewGithubUser("valid.csv", "../testdata/users"),
			returnErr:  nil,
			name:       "Read data from valid file",
		},
		{
			repository: repository.NewGithubUser("non-existing.csv", "../testdata/users"),
			returnErr:  repository.ErrOpeningCSV,
			name:       "Read data from non-existing file",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.repository.FetchAll()

			if err != tc.returnErr {
				t.Fatalf("%v failed: Expected returnErr: %v, got: %v", tc.name, tc.returnErr, err)
			}
		})
	}
}
