package repository

import (
	"fmt"
	"strings"
)

type Repository struct {
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

type Order struct {
	ID                 int
	Title              string
	Image              string
	Planet_description string
	Albedo_info        string
	Albedo             float64
}

type Temps_data struct {
	Title  string
	Image  string
	Albedo float64
	Distance int
	Temperature int
}

func (r *Repository) GetTempsData() ([]Temps_data, error) {
	temps_data := []Temps_data{
		{
			Title:  "Углеродная планета",
			Image:  "carbon_planet.jpg",
			Albedo: 0.1,
			Distance: 150,
			Temperature: 2600,
		},
		{
			Title:  "Ледяная планета",
			Image:  "ice_planet.png",
			Albedo: 0.6,
			Distance: 400,
			Temperature: 420,
		},
	}
	if len(temps_data) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}
	return temps_data, nil
}

func (r *Repository) GetOrders() ([]Order, error) {
	orders := []Order{
		{
			ID:                 1,
			Title:              "Горячий Юпитер",
			Image:              "hot_jupiter.png",
			Planet_description: "Горячие юпитеры — это класс экзопланет, которые похожи на Юпитер по массе и составу, но имеют очень короткие орбиты вокруг своих звезд, что приводит к высоким температурам на их поверхности.",
			Albedo_info:        "Из-за близости к звезде альбедо обычно низкое из-за поглощения света темными облаками или молекулами, такими как щелочные металлы.",
			Albedo:             0.05,
		},
		{
			ID:                 2,
			Title:              "Углеродная планета",
			Image:              "carbon_planet.jpg",
			Planet_description: "Углеродная экзопланета — тип планеты, в недрах и коре которой преобладает углерод над кислородом. Это миры, где вместо силикатных пород могут быть алмазы, графит и карбиды, а поверхность может быть покрыта углеводородной смолой или сажей.",
			Albedo_info:        "В виду темной атмосферы состоящей в основном из графита и сажи, значение альбедо крайне низко.",
			Albedo:             0.1,
		},
		{
			ID:                 3,
			Title:              "Скалистая планета",
			Image:              "rocky_planet.png",
			Planet_description: "Скалистая экзопланета — это тип планеты, которая в основном состоит из твердых материалов, таких как силикатные породы и металлы, подобно Земле, Венере и Марсу в нашей Солнечной системе. Эти планеты имеют твердую поверхность и могут иметь атмосферу.",
			Albedo_info:        "Значение альбедо может сильно варьироваться в зависимости от состава поверхности и атмосферы.",
			Albedo:             0.3,
		},
		{
			ID:                 4,
			Title:              "Газовый гигант",
			Image:              "gas_giant.jpg",
			Planet_description: "Газовый гигант — это крупная планета, состоящая преимущественно из газов, таких как водород и гелий, с возможным твердым ядром. В нашей Солнечной системе к газовым гигантам относятся Юпитер и Сатурн.",
			Albedo_info:        "Значение альбедо может варьироваться в зависимости от облаков и атмосферных условий.",
			Albedo:             0.5,
		},
		{
			ID:                 5,
			Title:              "Мини-Нептун",
			Image:              "mini_neptun.jpg",
			Planet_description: "Мини-Нептун — это тип экзопланеты, которая меньше Нептуна, но больше Земли, обычно с размером от 1,5 до 4 раз больше земного радиуса. Эти планеты имеют толстую атмосферу, богатую водородом и гелием, и могут иметь океаны из воды, аммиака или метана под атмосферой.",
			Albedo_info:        "Значение альбедо может варьироваться в зависимости от состава атмосферы и облаков.",
			Albedo:             0.4,
		},
		{
			ID:                 6,
			Title:              "Ледяная планета",
			Image:              "ice_planet.png",
			Planet_description: "Ледяная экзопланета — это тип планеты, которая находится далеко от своей звезды и имеет низкие температуры, что приводит к наличию твердых льдов на поверхности или в атмосфере. Эти планеты могут состоять из водяного льда, аммиачного льда или метанового льда.",
			Albedo_info:        "Из-за наличия льда на поверхности, значение альбедо обычно высокое.",
			Albedo:             0.6,
		},
		{
			ID:                 7,
			Title:              "Суперземля",
			Image:              "super_earth.jpg",
			Planet_description: "Суперземля — это тип экзопланеты, которая имеет массу больше земной, но значительно меньше, чем у газовых гигантов, обычно от 1 до 10 земных масс. Эти планеты могут быть скалистыми, ледяными или иметь толстую атмосферу.",
			Albedo_info:        "Значение альбедо может сильно варьироваться в зависимости от состава поверхности и атмосферы.",
			Albedo:             0.35,
		},
		{
			ID:                 8,
			Title:              "Хтоническая планета",
			Image:              "chthonic_planet.png",
			Planet_description: "Хтоническая планета — это гипотетический тип планеты, которая образовалась из остатков разрушенной планеты или луны, часто в результате столкновения. Эти планеты могут иметь необычный состав и структуру, отличающуюся от традиционных планет.",
			Albedo_info:        "Из-за необычного состава и возможного наличия большого количества металлов, значение альбедо может быть низким.",
			Albedo:             0.2,
		},
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}

	return orders, nil
}

func (r *Repository) GetOrder(id int) (Order, error) {
	orders, err := r.GetOrders()
	if err != nil {
		return Order{}, err
	}

	for _, order := range orders {
		if order.ID == id {
			return order, nil
		}
	}
	return Order{}, fmt.Errorf("заказ не найден")
}

func (r *Repository) GetOrdersByTitle(title string) ([]Order, error) {
	orders, err := r.GetOrders()
	if err != nil {
		return []Order{}, err
	}

	var result []Order
	for _, order := range orders {
		if strings.Contains(strings.ToLower(order.Title), strings.ToLower(title)) {
			result = append(result, order)
		}
	}

	return result, nil
}
