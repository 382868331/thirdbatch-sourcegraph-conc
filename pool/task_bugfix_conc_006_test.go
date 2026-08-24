package pool

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc006SourceContract(t *testing.T) {
    source, err := os.ReadFile("pool.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if p.limiter == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && p.limiter == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
