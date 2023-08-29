package db

import (
	"context"
	"testing"

	"github.com/freer4an/portfolio-website/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx context.Context

func TestInsertProject(t *testing.T) {
	arg := randomProject()

	res, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, res.Name, "must be not empty")
	require.NotZero(t, res.Description)
	require.NotZero(t, res.Url)
	require.True(t, res.IsFinished, "Should be true")
	assert.NotZero(t, res.ID, "ID shouldn't be zero value")
}

func TestGetProject(t *testing.T) {
	arg := randomProject()

	res1, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)
	assert.NotZero(t, res1.ID, "ID shouldn't be zero value")

	res2, err := testStore.GetProject(ctx, res1.ID)
	require.NoError(t, err)

	require.Equal(t, res1.ID, res2.ID)
	require.Equal(t, res1.Name, res2.Name)
	require.Equal(t, res1.Description, res2.Description)
	require.Equal(t, res1.Url, res2.Url)
}

func TestDeleteProject(t *testing.T) {
	arg := randomProject()

	res1, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)
	assert.NotZero(t, res1.ID, "ID shouldn't be zero value")

	res2, err := testStore.GetProject(ctx, res1.ID)
	require.NoError(t, err)
	require.Equal(t, res1.ID, res2.ID)

	err = testStore.DeleteProject(ctx, res2.ID)
	require.NoError(t, err)

	res2, err = testStore.GetProject(ctx, res1.ID)
	require.Error(t, err)
	assert.Zero(t, res2.ID, "deleted ID should be zero value")
}

func TestUpdateProject(t *testing.T) {
	arg := randomProject()

	res1, err := testStore.CreateProject(ctx, arg)
	require.NoError(t, err)
	assert.NotZero(t, res1.ID, "ID shouldn't be zero value")

	arg2 := updateProject{
		Name:        "Updated Name",
		Description: "Updated Description",
		IsFinished:  false,
	}

	res2, err := testStore.UpdateProject(ctx, res1.ID, arg2)
	require.NoError(t, err)
	assert.NotZero(t, res2.ID, "ID shouldn't be zero value")
	require.Equal(t, arg2.Name, res2.Name)
	require.Equal(t, arg2.Description, res2.Description)
	require.False(t, res2.IsFinished, "updated to false")

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
