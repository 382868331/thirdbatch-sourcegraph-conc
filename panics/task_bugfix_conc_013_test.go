package panics

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc013SourceContract(t *testing.T) {
    source, err := os.ReadFile("panics.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer p.tryRecover()") {
        t.Fatalf("expected source contract is missing")
    }
}
