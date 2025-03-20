package hp

import (
	"errors"
	"testing"

	E "github.com/IBM/fp-go/either"
	"github.com/stretchr/testify/assert"
)

// Mock Stats
type mockStats struct {
	hp int
}

func (m mockStats) GetHP() int {
	return m.hp
}

func TestInit(t *testing.T) {
	rock := mockStats{hp: 30}
	scissors := mockStats{hp: 20}
	paper := mockStats{hp: 10}

	hp := Init(rock, scissors, paper)

	assert.Equal(t, 60, hp.Current())
	assert.Equal(t, 60, hp.Max())
}

func TestNew(t *testing.T) {
	hp := New(30)(50)

	assert.Equal(t, 30, hp.Current())
	assert.Equal(t, 50, hp.Max())
}

func TestCurrentAndMax(t *testing.T) {
	hp := New(40)(100)

	assert.Equal(t, 40, hp.Current())
	assert.Equal(t, 100, hp.Max())
}

func TestDamaged_Success(t *testing.T) {
	hp := New(50)(100)

	result := hp.Damaged(20)

	assert.True(t, E.IsRight(result))
	assert.Equal(t, 30, result.Right().Current())
	assert.Equal(t, 100, result.Right().Max())
}

func TestDamaged_Failure(t *testing.T) {
	hp := New(20)(100)

	result := hp.Damaged(25)

	assert.True(t, E.IsLeft(result))
	assert.Equal(t, errors.New("this party lose"), result.Left())
}
