package minmaxstack

import (
	"testing"
)

type action struct {
	isPush bool
	val    int
}

func createStackFromActions(actions []action) *MinmaxStack {
	stack := New()
	for _, act := range actions {
		if act.isPush {
			stack.Push(act.val)
		} else {
			stack.Pop()
		}
	}
	return stack
}

func TestMinmaxStack_Len(t *testing.T) {

	tests := []struct {
		name    string
		actions []action
		want    int
	}{
		{"", []action{}, 0},
		{"", []action{
			{true, 0},
		}, 1},
		{"", []action{
			{true, 0},
		}, 1},
		{"", []action{
			{true, 0},
			{true, 3},
		}, 2},
		{"", []action{
			{true, 0},
			{true, 1},
			{false, 0},
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := createStackFromActions(tt.actions)
			if got := ms.Len(); got != tt.want {
				t.Errorf("MinmaxStack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinmaxStack_Integration(t *testing.T) {

	tests := []struct {
		name    string
		actions []action
		want    *Minmax
	}{
		{"", []action{}, nil},
		{"", []action{
			{true, 0},
		}, &Minmax{0, 0}},
		{"", []action{
			{true, 1},
			{true, 2},
			{true, 0},
			{true, 4},
			{true, 2},
		}, &Minmax{0, 4}},
		{"", []action{
			{true, 1},
			{true, 3},
			{true, 0},
			{false, 0},
			{true, 2},
		}, &Minmax{1, 3}},
		{"", []action{
			{true, 6},
			{true, 5},
			{true, 4},
			{true, 3},
			{false, 0},
			{true, 7},
			{true, 8},
			{true, 9},
			{false, 0},
		}, &Minmax{4, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := createStackFromActions(tt.actions)
			got, ok := ms.GetMinmax()
			if (!ok && tt.want != nil) || (tt.want != nil && got != *tt.want) {
				t.Errorf("MinmaxStack.GetMinmax() = %v, want %v", got, tt.want)
			}
		})
	}
}
