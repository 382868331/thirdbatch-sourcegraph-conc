package stream

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc001SourceContract(t *testing.T) {
    source, err := os.ReadFile("stream.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "ch <- func() {}") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "ch <=- func() {}") {
        t.Fatalf("mutated source contract is still present")
    }
}
