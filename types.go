package twitchgraphapi

type UsersInfo struct {
	Data struct {
		CurrentUser interface{} `json:"currentUser"`
		User        struct {
			Id              string      `json:"id"`
			Description     string      `json:"description"`
			DisplayName     string      `json:"displayName"`
			IsPartner       interface{} `json:"isPartner"`
			PrimaryColorHex interface{} `json:"primaryColorHex"`
			ProfileImageURL string      `json:"profileImageURL"`
			Followers       struct {
				TotalCount int    `json:"totalCount"`
				Typename   string `json:"__typename"`
			} `json:"followers"`
			Channel struct {
				Id           string `json:"id"`
				SocialMedias []struct {
					Id       string `json:"id"`
					Name     string `json:"name"`
					Title    string `json:"title"`
					Url      string `json:"url"`
					Typename string `json:"__typename"`
				} `json:"socialMedias"`
				Typename string `json:"__typename"`
			} `json:"channel"`
			LastBroadcast struct {
				Id   string `json:"id"`
				Game struct {
					Id          string `json:"id"`
					DisplayName string `json:"displayName"`
					Typename    string `json:"__typename"`
				} `json:"game"`
				Typename string `json:"__typename"`
			} `json:"lastBroadcast"`
			PrimaryTeam interface{} `json:"primaryTeam"`
			Videos      struct {
				Edges []struct {
					Node struct {
						Id   string `json:"id"`
						Game struct {
							Id          string `json:"id"`
							DisplayName string `json:"displayName"`
							Typename    string `json:"__typename"`
						} `json:"game"`
						Status   string `json:"status"`
						Typename string `json:"__typename"`
					} `json:"node"`
					Typename string `json:"__typename"`
				} `json:"edges"`
				Typename string `json:"__typename"`
			} `json:"videos"`
			Typename string `json:"__typename"`
		} `json:"user"`
	} `json:"data"`
	Extensions struct {
		DurationMilliseconds int    `json:"durationMilliseconds"`
		OperationName        string `json:"operationName"`
		RequestID            string `json:"requestID"`
	} `json:"extensions"`
}
