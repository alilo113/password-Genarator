package commands

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// --------------------
// Data model
// --------------------

type Entry struct {
	Site     string `json:"site"`
	Username string `json:"username"`
	Password string `json:"password"` // encrypted
}

// ~/.pwman/store.json
var storageFile = func() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".pwman", "store.json")
}()

// --------------------
// Storage helpers
// --------------------

func ensureStorageDir() error {
	dir := filepath.Dir(storageFile)
	return os.MkdirAll(dir, 0700)
}

func loadStorage() ([]Entry, error) {
	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		return []Entry{}, nil
	}

	data, err := os.ReadFile(storageFile)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}

func saveStorage(entries []Entry) error {
	if err := ensureStorageDir(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, data, 0600)
}

// --------------------
// Crypto
// --------------------

func deriveKey(master string) []byte {
	sum := sha256.Sum256([]byte(master))
	return sum[:]
}

func encryptPassword(master, password string) (string, error) {
	key := deriveKey(master)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(password), nil)
	return hex.EncodeToString(ciphertext), nil
}

func decryptPassword(master, encryptedHex string) (string, error) {
	key := deriveKey(master)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// --------------------
// Commands
// --------------------

func runAdd(providedPassword string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Site: ")
	site, _ := reader.ReadString('\n')
	site = strings.TrimSpace(site)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	var password string
	if providedPassword != "" {
		password = providedPassword
		fmt.Println("Using generated password:", password)
	} else {
		fmt.Print("Password (leave empty to generate): ")
		pwInput, _ := reader.ReadString('\n')
		pwInput = strings.TrimSpace(pwInput)

		if pwInput == "" {
			password = generatePassword(16)
			fmt.Println("Generated password:", password)
		} else {
			password = pwInput
		}
	}

	fmt.Print("Master password: ")
	master, _ := reader.ReadString('\n')
	master = strings.TrimSpace(master)

	encPassword, err := encryptPassword(master, password)
	if err != nil {
		fmt.Println("Encryption failed:", err)
		return
	}

	entries, err := loadStorage()
	if err != nil {
		fmt.Println("Failed to load storage:", err)
		return
	}

	entries = append(entries, Entry{
		Site:     site,
		Username: username,
		Password: encPassword,
	})

	if err := saveStorage(entries); err != nil {
		fmt.Println("Failed to save storage:", err)
		return
	}

	fmt.Printf("Saved password for %s (%s)!\n", site, username)
}

func runGet() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Site: ")
	site, _ := reader.ReadString('\n')
	site = strings.TrimSpace(site)

	fmt.Print("Master password: ")
	master, _ := reader.ReadString('\n')
	master = strings.TrimSpace(master)

	entries, err := loadStorage()
	if err != nil {
		fmt.Println("Failed to load storage:", err)
		return
	}

	for _, e := range entries {
		if e.Site == site {
			password, err := decryptPassword(master, e.Password)
			if err != nil {
				fmt.Println("Wrong master password or corrupted entry")
				return
			}

			fmt.Printf("Site: %s\nUsername: %s\nPassword: %s\n", e.Site, e.Username, password)
			return
		}
	}

	fmt.Println("No entry found for site:", site)
}