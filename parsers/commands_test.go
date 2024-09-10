package parsers

import (
	"testing"
)

func TestCreateCommandContext(t *testing.T) {
	ctx := CreateCommandContext()

	if ctx.controller != nil || ctx.gman != nil {
		t.Errorf("expected %v got %v", nil, ctx.controller)
	}
}
