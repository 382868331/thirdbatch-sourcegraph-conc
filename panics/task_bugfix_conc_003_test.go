package panics

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc003SourceContract(t *testing.T) {
    source, err := os.ReadFile("panics.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if p == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
