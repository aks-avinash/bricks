// Copyright © 2018 by PACE Telematics GmbH. All rights reserved.
// Created at 2018/10/17 by Vincent Landgraf

package cockpitapp

import "context"

type Client struct{}

//
// http://docs.pacelink.net/cockpit/api-web-v2.html#oauth-access-user-info-get
func (c *Client) AccessUserInfo(ctx context.Context) (*UserInfo, error) {
	// GET https://cp-1-stage.pacelink.net/api/web/v2/oauth2/me
	return nil, nil
}

type UserInfo struct {
	// {
	// 	"user": {
	// 	  "uuid": "64251556",
	// 	  "email": "user@pace.car",
	// 	  "confirmed_latest_terms": "true",
	// 	  "confirmed_latest_privacy": "true",
	// 	  "mobile_number": "+4955555555",
	// 	  "first_name": "Ben",
	// 	  "last_name": "Johnson",
	// 	  "gender": "male",
	// 	  "birthday": "1982-03-20",
	// 	  "avatar_url": "https://pace.car/assets/user.jpg",
	// 	  "created_at": 391730400,
	// 	  "onboarding_completed": true
	// 	}
	// }
}
