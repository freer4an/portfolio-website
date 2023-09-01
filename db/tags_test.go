package db

import (
	"math/rand"
	"testing"

	"github.com/freer4an/portfolio-website/util"
	"github.com/stretchr/testify/require"
)

func randomProjectWithTags(t *testing.T) Project {
	project := createRandomProject(t)
	tags := randomTag(3)
	err := testStore.AddProjectTags(ctx, project.Name, tags...)
	require.NoError(t, err)
	projectU, err := testStore.GetProject(ctx, project.Name)
	require.NoError(t, err)
	require.Len(t, projectU.Tags, 3)
	return projectU
}
func TestAddProjectTags(t *testing.T) {
	randomProjectWithTags(t)
}

func TestDeleteProjectTags(t *testing.T) {
	project := randomProjectWithTags(t)
	deleteTags := []string{project.Tags[0].TagName, project.Tags[1].TagName}
	err := testStore.DeleteProjectTags(ctx, project.Name, deleteTags...)
	require.NoError(t, err)
	project, err = testStore.GetProject(ctx, project.Name)
	require.NoError(t, err)
	require.Len(t, project.Tags, 1)
}

func randomTag(n int) []Tag {
	tags_types := [5]string{"technology", "language", "framework", "security", "service"}
	tags := []Tag{}
	for i := 0; i < n; i++ {
		tag := Tag{
			TagName: util.RandomStr(4),
			TagType: tags_types[rand.Intn(4)],
		}
		tags = append(tags, tag)
	}
	return tags
}
