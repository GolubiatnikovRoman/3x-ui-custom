package sub

import (
	"testing"

	"github.com/mhsanaei/3x-ui/v2/database/model"
)

func TestGetDisplayNameUsesSubIDBeforeEmail(t *testing.T) {
	service := NewSubService(false, "|ie")

	client := model.Client{
		Email: "Roman_1",
		SubID: "Roman",
	}

	if got := service.getDisplayName(client); got != "Roman" {
		t.Fatalf("expected display name from subId, got %q", got)
	}
}

func TestGetDisplayNameFallsBackToEmail(t *testing.T) {
	service := NewSubService(false, "|ie")

	client := model.Client{
		Email: "Roman_1",
	}

	if got := service.getDisplayName(client); got != "Roman_1" {
		t.Fatalf("expected fallback to email, got %q", got)
	}
}
