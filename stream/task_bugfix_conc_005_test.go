package stream

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc005SourceContract(t *testing.T) {
    source, err := os.ReadFile("stream.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer func() {") {
        t.Fatalf("expected source contract is missing")
    }
}
