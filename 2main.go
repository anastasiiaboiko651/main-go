package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Temperament string `json:"temperament"`
	Origin      string `json:"origin"`
	LifeSpan    string `json:"life_span"`
}

type CatImage struct {
	ID     string  `json:"id"`
	URL    string  `json:"url"`
	Breeds []Breed `json:"breeds"`
}

const APIKey = "Token"

func getBreedID(breedName string) (*Breed, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/breeds/search?q=%s", breedName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("live_WaBqfUo0pa8pm8MaGev4LvaO57J4vv5bGQOqAEJ8qvCHKYmalQf25lS8IddJjJBA7s", APIKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d", resp.StatusCode)
	}

	var breeds []Breed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		return nil, err
	}

	if len(breeds) == 0 {
		return nil, fmt.Errorf("breed not found")
	}

	return &breeds[0], nil
}

func getCatImage(breedID string) (*CatImage, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s", breedID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("live_WaBqfUo0pa8pm8MaGev4LvaO57J4vv5bGQOqAEJ8qvCHKYmalQf25lS8IddJjJBA", APIKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var images []CatImage
	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, fmt.Errorf("image not found")
	}

	return &images[0], nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введіть породу кота (англ, наприклад 'Bengal'): ")
	input, _ := reader.ReadString('\n')
	breedName := strings.TrimSpace(input)

	fmt.Println("Пошук породи...")
	breed, err := getBreedID(breedName)
	if err != nil {
		fmt.Println("Помилка пошуку породи:", err)
		return
	}

	fmt.Println("Завантаження фото...")
	image, err := getCatImage(breed.ID)
	if err != nil {
		fmt.Println("Помилка отримання фото:", err)
		return
	}

	fmt.Println("\n=== Результат ===")
	fmt.Printf("Порода: %s\n", breed.Name)
	fmt.Printf("Країна походження: %s\n", breed.Origin)
	fmt.Printf("Характер: %s\n", breed.Temperament)
	fmt.Printf("Опис: %s\n", breed.Description)
	fmt.Printf("Тривалість життя: %s років\n", breed.LifeSpan)
	fmt.Printf("URL фото: %s\n", image.URL)
}
