package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	username := "hiroto60" // GitHubのユーザー名を指定
	accessToken := ""      // アクセストークンを指定

	// GitHub APIのURLを生成
	url := fmt.Sprintf("https://api.github.com/repos/%s/golang/issues", username)

	// 作成するIssueの情報を定義
	issue := map[string]string{
		"title": "New Issue",
		"body":  "This is a new issue created using GitHub API.",
	}

	// Issueの情報をJSONに変換
	jsonData, err := json.Marshal(issue)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// リクエストヘッダーを設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// HTTPクライアントを作成し、リクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// レスポンスを解析
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Issue created successfully.")
	} else {
		fmt.Println("Failed to create issue. Status code:", resp.StatusCode)
	}
}
