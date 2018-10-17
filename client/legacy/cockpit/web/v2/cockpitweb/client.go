// Copyright © 2018 by PACE Telematics GmbH. All rights reserved.
// Created at 2018/10/17 by Vincent Landgraf

package cockpitweb

import "context"

type Client struct{}

//
// http://docs.pacelink.net/cockpit/api-app-v2.html#user-cars-cars-get
func (c *Client) GetAllCars(ctx context.Context) ([]*Car, error) {
	// GET https://cp-1-stage.pacelink.net/api/app/v2/user/cars?uuids=00d5aed7-839e-4a0b-b4c8-536570c43990,e7d64c35-18be-469e-8ced-a48bfe7e84f0
	return nil, nil
}

type Car struct {
	// {
	// 	"cars": [
	// 	  {
	// 		"uuid": "64251556",
	// 		"vin": "WAU1234K120",
	// 		"car_model_uuid": "7d5c398d",
	// 		"name": "MY BMW M2",
	// 		"number_plate": "KA-PL2345",
	// 		"color": "Black",
	// 		"avatar_url": "https://pace.car/assets/car.jpg",
	// 		"settings": [
	// 		  {
	// 			"key": "drivelog_enabled",
	// 			"value": "yes"
	// 		  }
	// 		],
	// 		"drive_mode_settings": [
	// 		  {
	// 			"key": "drivelog_enabled",
	// 			"value": "yes"
	// 		  }
	// 		],
	// 		"status": "active",
	// 		"logbook_enabled": true,
	// 		"ecall_enabled": true,
	// 		"ecall_token": "d8028511-746c-4dec-93f0-9434086c02f7",
	// 		"updated_at": 1443100891
	// 	  }
	// 	]
	//   }
}
