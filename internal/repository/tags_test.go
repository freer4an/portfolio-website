package repository

import (
	"math/rand"
	"testing"

	"github.com/freer4an/portfolio-website/internal/models"
	"github.com/freer4an/portfolio-website/util"
	"github.com/stretchr/testify/require"
)

func TestAddProjectTags(t *testing.T) {
	createRandomProject(t)
}

func TestDeleteProjectTags(t *testing.T) {
	project := createRandomProject(t)
	deleteTags := []string{project.Tags[0].TagName, project.Tags[1].TagName}
	err := testStore.Tag.DeleteFromProject(ctx, project.Name, deleteTags...)
	require.NoError(t, err)
	project, err = testStore.Project.GetByName(ctx, project.Name)
	require.NoError(t, err)
	require.Len(t, project.Tags, 1)
}

func randomTag(n int) []models.Tag {
	tags_types := [5]string{"technology", "language", "framework", "security", "service"}
	tags := []models.Tag{}
	for i := 0; i < n; i++ {
		tag := models.Tag{
			TagName: util.RandomStr(4),
			TagType: tags_types[rand.Intn(4)],
		}
		tags = append(tags, tag)
	}
	return tags
}
