package repository

import (
	"context"
	"testing"

	"github.com/freer4an/portfolio-website/internal/models"
	"github.com/freer4an/portfolio-website/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx context.Context

func createRandomProject(t *testing.T) models.Project {
	project := randomProject()

	res, err := testStore.Project.Create(ctx, project)
	require.NoError(t, err)

	projectU, err := testStore.Project.GetByName(ctx, project.Name)
	require.NoError(t, err)
	project.ID = res

	return projectU
}
func TestInsertProject(t *testing.T) {
	createRandomProject(t)
}

func TestGetProject(t *testing.T) {
	arg := createRandomProject(t)

	res2, err := testStore.Project.GetByName(ctx, arg.Name)
	require.NoError(t, err)

	require.NotZero(t, res2.Name)
	require.NotZero(t, res2.Description)
	require.NotZero(t, res2.Url)
}

func TestGetAllProjects(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProject(t)
	}

	res1, err := testStore.Project.GetAll(ctx, 5, 1)
	require.NoError(t, err)
	require.Len(t, res1, 5, "length of result should be 5")

	res2, err := testStore.Project.GetAll(ctx, 5, 2)
	require.NoError(t, err)
	require.Len(t, res2, 5, "length of result should be 5")
}

func TestDeleteProject(t *testing.T) {
	arg := createRandomProject(t)

	count, err := testStore.Project.Delete(ctx, arg.Name)
	require.NoError(t, err)
	require.Equal(t, count, int64(1))

	_, err = testStore.Project.GetByName(ctx, arg.Name)
	require.Error(t, err)
}

func TestUpdateProject(t *testing.T) {
	arg := createRandomProject(t)

	new_name := util.RandomStr(6)
	arg2 := bson.D{
		{
			Key:   "name",
			Value: new_name,
		},
		{
			Key:   "description",
			Value: "New Description",
		},
		{
			Key:   "url",
			Value: "New Url",
		},
		{
			Key:   "finished",
			Value: false,
		},
	}

	_, err := testStore.Project.Update(ctx, arg.Name, arg2)
	require.NoError(t, err)

	res, err := testStore.Project.GetByName(ctx, new_name)
	require.NoError(t, err)

	require.Equal(t, arg2[0].Value, res.Name)
	require.Equal(t, arg2[1].Value, res.Description)
	require.Equal(t, arg2[2].Value, res.Url)

	require.NotZero(t, res.Name)
	require.NotZero(t, res.Description)
	require.NotZero(t, res.Url)

	require.False(t, res.Finished, "updated to false")

}

func randomProject() models.Project {
	p := models.Project{
		Name:        util.RandomStr(6),
		Abstract:    util.RandomStr(25),
		Description: util.RandomStr(50),
		Url:         util.RandomStr(12),
		Finished:    true,
	}
	return p
}
