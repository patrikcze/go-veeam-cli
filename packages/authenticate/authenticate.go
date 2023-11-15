package veeamauthenticate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Issued       string `json:".issued"`
	Expires      string `json:".expires"`
}

func auth() {
	// Check if the token exists in storage
	token, err := getTokenFromStorage()
	if err != nil {
		// Token does not exist or is invalid, obtain a new one
		token, err = obtainAccessToken()
		if err != nil {
			panic(err)
		}

		// Save the token to storage
		err = saveTokenToStorage(token)
		if err != nil {
			panic(err)
		}
	}

	// Use the token for subsequent API calls
	// Your code for the "Start Entire VM Restore" REST API call goes here
	fmt.Println(token.AccessToken)
}

func getTokenFromStorage() (*TokenResponse, error) {
	// Read the token from storage (e.g., file, secure storage)
	// Implement your own logic to retrieve the token securely
	// For demonstration purposes, we'll assume the token is stored in a file named "token.json"
	filePath := "token.json"
	data, err := ioutil.ReadFile(filePath)
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

func saveTokenToStorage(token *TokenResponse) error {
	// Marshal the token into JSON data
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	// Write the token data to storage (e.g., file, secure storage)
	// Implement your own logic to store the token securely
	// For demonstration purposes, we'll assume the token is stored in a file named "token.json"
	filePath := "token.json"
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func obtainAccessToken() (*TokenResponse, error) {
	reqURL := "https://cdn.veeam.com/api/oauth2/token"
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", "string")
	data.Set("password", "pa$$word")
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
	body, err := ioutil.ReadAll(res.Body)
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