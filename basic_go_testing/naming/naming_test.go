package naming

import "testing"

func TestDog_Bark_muzzled(t *testing.T) {
	// Test case to check if a muzzled dog can bark

}
func TestDog_Bark_unmuzzled(t *testing.T) {
	// Test case to check if not muzzled

}

func TestColor(t *testing.T) {
	// Test case to check if the color is valid
	arg := "blue"
	want := "#0000FF"
	got := Color(arg)
	if got != want {
		t.Errorf("Color (%q)Expected %s but got %s", arg, want, got)
	}
}
