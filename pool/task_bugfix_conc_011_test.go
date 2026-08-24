package pool

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc011SourceContract(t *testing.T) {
    source, err := os.ReadFile("result_context_pool.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "p.agg.save(idx, res, err != nil)") {
        t.Fatalf("expected source contract is missing")
    }
}
