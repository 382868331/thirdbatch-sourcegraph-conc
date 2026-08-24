package pool

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisConc012SourceContract(t *testing.T) {
    source, err := os.ReadFile("pool.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if n < 1 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if n <= 1 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
