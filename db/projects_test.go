package db

import (
	"context"
	"testing"

	"github.com/freer4an/portfolio-website/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx context.Context

func TestInsertProject(t *testing.T) {
	arg := randomProject()

	res, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)
	_, ok := res.(primitive.ObjectID)
	require.True(t, ok)
}

func TestGetProject(t *testing.T) {
	arg := randomProject()

	res1, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)
	assert.NotNil(t, res1, "ID shouldn't be zero value")

	res2, err := testStore.GetProject(ctx, arg.Name)
	require.NoError(t, err)

	require.NotZero(t, res2.Name)
	require.NotZero(t, res2.Description)
	require.NotZero(t, res2.Url)
}

func TestGetAllProjects(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := randomProject()
		_, err := testStore.CreateProject(ctx, arg)
		require.NoError(t, err)
	}

	res1, err := testStore.GetAllProjects(ctx, 5, 1)
	require.NoError(t, err)
	require.Len(t, res1, 5, "length of result should be 5")

	res2, err := testStore.GetAllProjects(ctx, 5, 2)
	require.NoError(t, err)
	require.Len(t, res2, 5, "length of result should be 5")
}

func TestDeleteProject(t *testing.T) {
	arg := randomProject()

	_, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)

	_, err = testStore.GetProject(ctx, arg.Name)
	require.NoError(t, err)

	count, err := testStore.DeleteProject(ctx, arg.Name)
	require.NoError(t, err)
	require.Equal(t, count, int64(1))

	_, err = testStore.GetProject(ctx, arg.Name)
	require.Error(t, err)
}

func TestUpdateProject(t *testing.T) {
	arg := randomProject()

	_, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)
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

	_, err = testStore.UpdateProject(ctx, arg.Name, arg2)
	require.NoError(t, err)

	res, err := testStore.GetProject(ctx, new_name)
	require.NoError(t, err)

	require.Equal(t, arg2[0].Value, res.Name)
	require.Equal(t, arg2[1].Value, res.Description)
	require.Equal(t, arg2[2].Value, res.Url)

	require.NotZero(t, res.Name)
	require.NotZero(t, res.Description)
	require.NotZero(t, res.Url)

	require.False(t, res.IsFinished, "updated to false")

}

func randomProject() Project {
	p := Project{
		Name:        util.RandomStr(6),
		Description: util.RandomStr(20),
		Url:         util.RandomStr(15),
		IsFinished:  true,
	}
	return p
}
