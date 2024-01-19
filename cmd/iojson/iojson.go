package iojson

///  era_lab11@mil.ru
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"os"
	"time"
)

const JSON_SAVE_PATH = "C:/programms/Go/Goprojects/src/gogogo/ui/static/json/weather.json"

type WeatherData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Localtime string  `json:"localtime"`
}

type Current struct {
	Temperature float64 `json:"temp_c"`
	Humidity    float64 `json:"humidity"`
}

func getWeatherData(api_url string) (WeatherData, error) {
	var weatherData WeatherData
	response, err := http.Get(api_url)
	if err != nil {
		fmt.Println("Ошибка при выполнении Get-запроса:", err)
		return weatherData, err
	}
	defer response.Body.Close()

	//Чтение тела ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении данных из ответа:", err)
		return weatherData, err
	}

	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return weatherData, err
	}
	return weatherData, nil
	// fmt.Println("Latitude:", weatherData.Location.Latitude)
	// fmt.Println("Longitude:", weatherData.Location.Longitude)
	// fmt.Println("Name:", weatherData.Location.Name)
	// fmt.Println("Region:", weatherData.Location.Region)
	// fmt.Println("Country:", weatherData.Location.Country)
	// fmt.Println("Localtime:", weatherData.Location.Localtime)
	// fmt.Println("Temperature:", weatherData.Current.Temperature)
	// fmt.Println("Humidity:", weatherData.Current.Humidity)

}

func Start() {

	start := time.Now()
	citys := []string{
		"Voronezh",
		"Abakan",
		"Anapskaya",
		"Arkhangelsk",
		"Anadyr",
		"Ali-Yurt",
		"Astrakhan",
		"Barnaul",
		"Belgorod",
		"Blagoveshchensk",
		"Birobidzhan",
		"Bryansk",
		"Vladivostok",
		"Vladikavkaz",
		"Vladimir",
		"Vorkuta",
		"Vologda",
		//"58.52681103974211, 31.27466909975728",///Великий новгород
		"Volgograd",
		"Grozny",
		"Gorno-Altaysk",
		"Yekaterinburg",
		"Izhevsk",
		"Irkutsk",
		"Ivanovo",
		"Yoshkar-Ola",
		"Kaluga",
		"Kaliningrad",
		"Kazan",
		"Kemerovo",
		"Kostroma",
		"Krasnoyarsk",
		"Krasnodar",
		"Kirov",
		"Kurgan",
		"Kursk",
		//"51.751041558041074, 94.44148701653755", ///Kyzyl
		"Lipetsk",
		"Moscow",
		"Makhachkala",
		"Magadan",
		"Maykop",
		"Murmansk",
		"Naryan-Mar",
		"Nalchik",
		"Novgorod",
		"Novosibirsk",
		"Omsk",
		"Orenburg",
		"Orel",
		"Petrozavodsk",
		"Petropavlovsk-Kamchatskiy",
		"Perm",
		"Penza",
		"Pskov",
		"Rostov-na-Donu",
		"Ryazan",
		"Saransk",
		"Samara",
		"Salekhard",
		"Smolensk",
		"Stavropol",
		"Syktyvkar",
		"Sochi",
		"Saint-Petersburg",
		"Saratov",
		"Tambov",
		"Tver",
		"Tomsk",
		"Tyumen",
		"Tula",
		"Ulyanovsk",
		"Ufa",
		"Ulan-Ude",
		"Khanty-Mansiysk",
		"Khabarovsk",
		"Chelyabinsk",
		"Cheboksary",
		"Chita",
		"Cherkessk",
		"Elista",
		"Yakutsk",
		"Yaroslavl",
	}
	var messages []WeatherData
	arrLen := len(citys)
	for i := 0; i < arrLen; i++ {
		// log.Println(citys[i])
		api_url := "http://api.weatherapi.com/v1/current.json?key=b2055c02917c44aa9f581406232512&q=" + citys[i] + "&aqi=no" //Ключ для обращения к weatherAPI
		message, err := getWeatherData(api_url)
		if err != nil {
			fmt.Println("Ошибка сохранения структуры")
			return
		}
		messages = append(messages, message)

	}
	jsonData, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("Ошибка при кодировании JSON:", err)
		return
	}

	file, err := os.Create(JSON_SAVE_PATH)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Массив структур сохранен в файл!")

	duration := time.Since(start)

	fmt.Println("Время выполнения для ", arrLen, "городов =", duration)
}
