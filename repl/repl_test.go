package repl

import testing "testing"

func ReplTest(t *testing.T) {
	if 1 != 1 {
		t.Fatalf("World gone mad innit!")
	}
}
