package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"weather-api/config"
	"weather-api/internal/models"
)

type WeatherService struct{}

func (s *WeatherService) GetWeather(city string) (*models.Weather, error) {
	apiKey := config.GetAPIKey()
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Không thể kết nối tới API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API trả về lỗi")
	}

	var weather models.Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, errors.New("Lỗi giải mã JSON")
	}

	return &weather, nil
}
