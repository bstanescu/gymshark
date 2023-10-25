package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_closestMinPacks(t *testing.T) {
	Packs := []int{250, 500, 1000, 2000, 5000}
	tests := []struct {
		packs        []int
		target       int
		wantSolution map[int]int
	}{
		{
			packs:        Packs,
			target:       0,
			wantSolution: map[int]int{},
		},
		{
			packs:        Packs,
			target:       1,
			wantSolution: map[int]int{250: 1},
		},
		{
			packs:        Packs,
			target:       250,
			wantSolution: map[int]int{250: 1},
		},
		{
			packs:        Packs,
			target:       251,
			wantSolution: map[int]int{500: 1},
		},
		{
			packs:        Packs,
			target:       12001,
			wantSolution: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Set: %v, %d", tt.packs, tt.target)
		t.Run(name, func(t *testing.T) {
			solution := closestMinPacks(tt.packs, tt.target)
			if !reflect.DeepEqual(solution, tt.wantSolution) {
				t.Errorf("closestMinPacks() got = %v, want %v", solution, tt.wantSolution)
			}
		})
	}
}
