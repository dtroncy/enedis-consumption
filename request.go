package enedisconsumption

import (
	"fmt"
	"log"
	"net/http"
)

func requestGET(url string, start string, end string) (*http.Response, error) {

	configFile := "conf.json"

	config, err := loadConfig(configFile)

	if err != nil {
		fmt.Println("Erreur lors du chargement de la configuration:", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("prm", config.PRM)
	q.Add("start", start)
	q.Add("end", end)

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+config.Token)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête")
	}

	defer response.Body.Close()

	return response, err

}
