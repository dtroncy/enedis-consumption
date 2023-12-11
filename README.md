# enedis-consumption
Module Go récupérant votre consommation électrique côté Enedis en passant par [Conso API](https://conso.boris.sh/)

Importer le module :
```
import (
	conso "github.com/dtroncy/enedis-consumption"
)
```
Ajouter un fichier de configuration ``conf.json`` à la racine de votre projet :
```
{
    "token": "token",
    "prm": "prm"
}
```
Pour obtenir ses informations, il faut s'inscrire sur le site suivant : https://conso.boris.sh/

Il vous permettra d'obtenir un token pour pouvoir récupérer les données côté Enedis.


Utiliser le module pour récupérer sa consommation journalière :
```
conso.GetDailyConsumption("date de début", "date de fin")
```
Exemple :
```
conso.GetDailyConsumption("2023-12-08", "2023-12-11")
```

Le retour est le suivant :
```
{
    "usage_point_id": "11111111111111",
    "start": "2023-12-08",
    "end": "2023-12-11",
    "quality": "BRUT",
    "reading_type": {
        "unit": "Wh",
        "measurement_kind": "energy",
        "aggregate": "sum",
        "measuring_period": "P1D"
    },
    "interval_reading": [
        {
            "value": "21547",
            "date": "2023-12-08"
        },
        {
            "value": "14528",
            "date": "2023-12-09"
        },
        {
            "value": "10258",
            "date": "2023-12-10"
        }
    ]
}
```

Utiliser le module pour récupérer la courbe de consommation (consommation par tranche de 30 min). Attention uniquement si activé dans votre compte Enedis :
```
conso.GetConsumptionLoadCurve("date de début", "date de fin")
```
Exemple :
```
conso.GetConsumptionLoadCurve("2023-12-08", "2023-12-11")
```

Le retour est le suivant :
```
{
    "usage_point_id": "11111111111111",
    "start": "2023-12-08",
    "end": "2023-12-11",
    "quality": "BRUT",
    "reading_type": {
        "unit": "W",
        "measurement_kind": "power",
        "aggregate": "average"
    },
    "interval_reading": [
        {
            "value": "5466",
            "date": "2023-12-08 00:30:00",
            "interval_length": "PT30M",
            "measure_type": "B"
        },
        {
            "value": "4156",
            "date": "2023-12-08 01:00:00",
            "interval_length": "PT30M",
            "measure_type": "B"
        },
        {
            "value": "3616",
            "date": "2023-12-08 01:30:00",
            "interval_length": "PT30M",
            "measure_type": "B"
        },
        {
            "value": "3630",
            "date": "2023-12-08 02:00:00",
            "interval_length": "PT30M",
            "measure_type": "B"
        },
[.............]
        {
            "value": "198",
            "date": "2023-12-11 00:00:00",
            "interval_length": "PT30M",
            "measure_type": "B"
        }
    ]
}
```

Utiliser le module pour récupérer la puissance max utilisée:
```
conso.GetConsumptionMaxPower("date de début", "date de fin")
```
Exemple :
```
conso.GetConsumptionMaxPower("2023-12-08", "2023-12-11")
```

Le retour est le suivant :
```
{
    "usage_point_id": "11111111111111",
    "start": "2023-12-08",
    "end": "2023-12-11",
    "quality": "BRUT",
    "reading_type": {
        "unit": "VA",
        "measurement_kind": "power",
        "aggregate": "maximum",
        "measuring_period": "P1D"
    },
    "interval_reading": [
        {
            "value": "7450",
            "date": "2023-12-08 11:43:55"
        },
        {
            "value": "7550",
            "date": "2023-12-09 09:14:43"
        },
        {
            "value": "6790",
            "date": "2023-12-10 01:28:44"
        }
    ]
}
```