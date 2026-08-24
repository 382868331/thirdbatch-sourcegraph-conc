package pool

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc010SourceContract(t *testing.T) {
    source, err := os.ReadFile("pool.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer func() { p.initOnce = sync.Once{} }()") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "func() { p.initOnce = sync.Once{} }()") {
        t.Fatalf("mutated source contract is still present")
    }
}
