package enedisconsumption

import "fmt"

func GetDailyConsumption(start string, end string) {

	response := requestGET("https://conso.boris.sh/api/daily_consumption", start, end)

	fmt.Println(string(response))

}

func GetConsumptionLoadCurve(start string, end string) {

	response := requestGET("https://conso.boris.sh/api/consumption_load_curve", start, end)

	fmt.Println(string(response))

}

func GetConsumptionMaxPower(start string, end string) {

	response := requestGET("https://conso.boris.sh/api/consumption_max_power", start, end)

	fmt.Println(string(response))

}
