package pool

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisConc004SourceContract(t *testing.T) {
    source, err := os.ReadFile("context_pool.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer func() {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "func() {") {
        t.Fatalf("mutated source contract is still present")
    }
}
