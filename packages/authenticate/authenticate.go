package veeamauthenticate

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/patrikcze/go-veeam-cli/packages/encryption"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Issued       string `json:".issued"`
	Expires      string `json:".expires"`
}

func Authenticate(servername, username, password string, port int, key []byte) (*TokenResponse, error) {
	// Check if the token exists in storage
	token, err := GetTokenFromStorage(key)
	if err != nil {
		// Token does not exist or is invalid, obtain a new one
		token, err = obtainAccessToken(servername, username, password, port)
		if err != nil {
			return nil, err
		}

		// Save the token to storage
		err = saveTokenToStorage(token, key)
		if err != nil {
			return nil, err
		}
	}

	// Use the token for subsequent API calls
	// fmt.Println(token.AccessToken)

	return token, nil
}

func GetTokenFromStorage(key []byte) (*TokenResponse, error) {
	// Read the encrypted token from storage
	filePath := "token.json"
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Decrypt the token
	data, err := encryption.Decrypt(encryptedData, key)
	if err != nil {
		return nil, err
	}

	// Unmarshal the token JSON data into a TokenResponse struct
	var token TokenResponse
	err = json.Unmarshal(data, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// This function saves gathered RestAPI Token into JSON File.
func saveTokenToStorage(token *TokenResponse, key []byte) error {
	// Marshal the token into JSON data
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	// Encrypt the token
	encryptedData, err := encryption.Encrypt(data, key)
	if err != nil {
		return err
	}

	// Write the encrypted token data to storage
	filePath := "token.json"
	err = os.WriteFile(filePath, encryptedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Function will Obtain Authorization Token from Veeam B&R RestAPI call for provided user and password!
func obtainAccessToken(servername, username, password string, port int) (*TokenResponse, error) {
	// Before making the HTTP request, disable certificate verification
	// Please note that this approach should only be used for testing or development purposes.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	reqURL := fmt.Sprintf("https://%s:%d/api/oauth2/token", servername, port)
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("refresh_token", "string")
	data.Set("code", "string")
	data.Set("use_short_term_refresh", "true")
	data.Set("vbr_token", "string")
	req, err := http.NewRequest("POST", reqURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("x-api-version", "1.1-rev0")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Parse the token response
	var token TokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
