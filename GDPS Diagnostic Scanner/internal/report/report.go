package report

import "fmt"

type Report struct {
    entries []string
}

func New() *Report {
    return &Report{entries: []string{}}
}

func (r *Report) Add(entry string) {
    r.entries = append(r.entries, entry)
}

func (r *Report) Print() {
    fmt.Println("=== GDPS Diagnostic Report ===")
    for _, e := range r.entries {
        fmt.Println("- " + e)
    }
}
