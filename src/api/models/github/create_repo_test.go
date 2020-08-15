package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golan introduction repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	// Marshal take an inpur interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.Equal(t,
		`{"name":"golang introduction","description":"a golan introduction repo","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`,
		string(bytes))

	var target CreateRepoRequest
	// Unmarshall takes an input byte array and a pointer that we are trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.Equal(t, target.Name, request.Name)
	assert.Equal(t, target.Private, request.Private)
}
