package main

import (
	"context"
	"testing"

	pb "task_1/proto"

	"github.com/stretchr/testify/assert"
)

func TestCheckTable(t *testing.T) {
	s := &server{}

	tests := []struct {
		name      string
		tableID   int32
		wantState bool
	}{
		{"Table 5 is available", 5, true},
		{"Table 10 is available", 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &pb.TableRequest{Id: tt.tableID}

			res, err := s.CheckTable(context.Background(), req)

			assert.NoError(t, err)
			assert.Equal(t, tt.wantState, res.Available)
		})
	}
}
