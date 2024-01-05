package attachments

import (
	"testing"
)

func TestNew(t *testing.T) {
	attachments := New()

	// Test that Static field is set correctly
	expectedStatic := StaticFiles
	if attachments.Static != expectedStatic {
		t.Errorf("Expected Static field to be %v, but got %v", expectedStatic, attachments.Static)
	}

	// Test that Templates field is set correctly
	expectedTemplates := TemplatesFiles
	if attachments.Templates != expectedTemplates {
		t.Errorf("Expected Templates field to be %v, but got %v", expectedTemplates, attachments.Templates)
	}
}
