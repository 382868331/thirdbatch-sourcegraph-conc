package iter

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixConc007SourceContract(t *testing.T) {
    source, err := os.ReadFile("iter.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "for i := 0; i < iter.MaxGoroutines; i++ {") {
        t.Fatalf("expected source contract is missing")
    }
}
