{
	"device": {
		"name": "RoboCar Station",
		"description": "Car services offered by robot",
		"producer": {
			"services": [{
				"serviceID": 1,
				"name": "RoboAir",
				"description": "Car tyre pressure checked and topped up by robot",
				"prices": [{
						"priceID": 1,
						"priceDescription": "Measure and adjust pressure",
						"pricePerUnit": {
							"amount": 25,
							"currencyCode": "GBP"
						},
						"unitID": 1,
						"unitDescription": "Tyre"
					}, {
						"priceID": 2,
						"priceDescription": "Measure and adjust pressure - four tyres for the price of three",
						"pricePerUnit": {
							"amount": 75,
							"currencyCode": "GBP"
						},
						"unitID": 2,
						"unitDescription": "4 Tyre"
					}

				]
			}, {
				"serviceID": 2,
				"name": "RoboWash",
				"description": "Car washed by robot",
				"prices": [{
						"priceID": 1,
						"priceDescription": "Car wash",
						"pricePerUnit": {
							"amount": 500,
							"currencyCode": "GBP"
						},
						"unitID": 1,
						"unitDescription": "Single wash"
					}, {
						"priceID": 2,
						"priceDescription": "SUV Wash",
						"pricePerUnit": {
							"amount": 650,
							"currencyCode": "GBP"
						},
						"unitID": 1,
						"unitDescription": "Single wash"
					}

				]
			}],
			"config": {
				"pspMerchantServiceKey": "T_S_f50ecb46-ca82-44a7-9c40-421818af5996",
				"pspMerchantClientKey": "T_C_03eaa1d3-4642-4079-b030-b543ee04b5af"
			}
		},
		"consumer": null,
		"config": null
	}
}