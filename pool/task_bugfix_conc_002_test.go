package pool

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc002SourceContract(t *testing.T) {
    source, err := os.ReadFile("error_pool.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(errs) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
