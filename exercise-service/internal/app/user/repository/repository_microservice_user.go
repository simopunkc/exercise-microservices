package repository

import (
	"context"
	"encoding/json"
	"exercise-service/internal/app/domain"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const getUserUrl = "internal/users/"

type RepositoryMicroserviceUser struct {
	hostname string
	username string
	password string
	client   *http.Client
}

func NewRepositoryMicroserviceUser() *RepositoryMicroserviceUser {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &RepositoryMicroserviceUser{
		hostname: os.Getenv("API_HOST"),
		username: os.Getenv("API_USERNAME"),
		password: os.Getenv("API_PASSWORD"),
		client:   &client,
	}
}

func (rmu RepositoryMicroserviceUser) IsUserExists(ctx context.Context, userID int64) bool {
	url := fmt.Sprintf("%s%s%d", rmu.hostname, getUserUrl, userID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false
	}
	req.SetBasicAuth(rmu.username, rmu.password)
	resp, err := rmu.client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	var repo domain.RepositoryUser
	err = json.Unmarshal(bodyBytes, &repo)
	if err != nil {
		log.Println(err)
		return false
	}

	return repo.User.ID > 0
}
