package pool

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisConc008SourceContract(t *testing.T) {
    source, err := os.ReadFile("result_error_pool.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "p.agg.save(idx, res, err != nil)") {
        t.Fatalf("expected source contract is missing")
    }
}
