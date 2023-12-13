package enedisconsumption

import "net/http"

func GetDailyConsumption(start string, end string) (*http.Response, error) {

	return requestGET("https://conso.boris.sh/api/daily_consumption", start, end)

}

func GetConsumptionLoadCurve(start string, end string) (*http.Response, error) {

	return requestGET("https://conso.boris.sh/api/consumption_load_curve", start, end)

}

func GetConsumptionMaxPower(start string, end string) (*http.Response, error) {

	return requestGET("https://conso.boris.sh/api/consumption_max_power", start, end)

}
