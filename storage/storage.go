package storage

type Entry struct {
    Site     string `json:"site"`     // Capitalized to export for JSON
    Username string `json:"username"` // Capitalized to export
    Password string `json:"password"` // Capitalized to export
}

type Vault struct {
    Entries []Entry `json:"entries"`
}