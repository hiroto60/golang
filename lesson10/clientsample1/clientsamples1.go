package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

func main() {
	username := "hiroto60" // GitHubのユーザー名を指定
	accessToken := ""      // アクセストークンを指定

	// GitHub APIのURLを生成
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	// HTTPリクエストを作成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// アクセストークンをAuthorizationヘッダーに設定
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// HTTPクライアントを作成し、リクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// レスポンスのステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		fmt.Println("failed to retrieve repositories, status code: ", resp.StatusCode)
		return
	}

	// レスポンスのJSONをデコード
	var repos []Repository
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// リポジトリ名を出力
	for _, repo := range repos {
		fmt.Println(repo.Name)
	}
}
