package twitchgraphapi

import "time"

const twitchApiUrl = "https://gql.twitch.tv/gql"

type (
	UsersInfo struct {
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
)

type VideoMetadata struct {
	Data struct {
		User struct {
			Id              string      `json:"id"`
			PrimaryColorHex string      `json:"primaryColorHex"`
			IsPartner       interface{} `json:"isPartner"`
			ProfileImageURL string      `json:"profileImageURL"`
			LastBroadcast   struct {
				Id        string    `json:"id"`
				StartedAt time.Time `json:"startedAt"`
				Typename  string    `json:"__typename"`
			} `json:"lastBroadcast"`
			Typename string `json:"__typename"`
		} `json:"user"`
		CurrentUser struct {
			Id       string `json:"id"`
			Typename string `json:"__typename"`
		} `json:"currentUser"`
		Video struct {
			Id                  string      `json:"id"`
			Title               string      `json:"title"`
			Description         interface{} `json:"description"`
			PreviewThumbnailURL string      `json:"previewThumbnailURL"`
			CreatedAt           time.Time   `json:"createdAt"`
			ViewCount           int         `json:"viewCount"`
			PublishedAt         time.Time   `json:"publishedAt"`
			LengthSeconds       int         `json:"lengthSeconds"`
			BroadcastType       string      `json:"broadcastType"`
			Owner               struct {
				Id          string `json:"id"`
				Login       string `json:"login"`
				DisplayName string `json:"displayName"`
				Typename    string `json:"__typename"`
			} `json:"owner"`
			Game struct {
				Id          string `json:"id"`
				Slug        string `json:"slug"`
				BoxArtURL   string `json:"boxArtURL"`
				Name        string `json:"name"`
				DisplayName string `json:"displayName"`
				Typename    string `json:"__typename"`
			} `json:"game"`
			Typename string `json:"__typename"`
		} `json:"video"`
	} `json:"data"`
	Extensions struct {
		DurationMilliseconds int    `json:"durationMilliseconds"`
		OperationName        string `json:"operationName"`
		RequestID            string `json:"requestID"`
	} `json:"extensions"`
}

const (
	version    = 1
	sha256Hash = "c25707c1e5176320ceac6b447d052480887e23bc794ca1d02becd0bcc91844fe"
)

const (
	videoMetadataOperation = "VideoMetadata"
)

type (
	Operation struct {
		OperationName string `json:"operationName"`
		Variables     struct {
			ChannelLogin string `json:"channelLogin"`
			VideoID      string `json:"videoID"`
		} `json:"variables"`
		Extensions `json:"extensions"`
	}

	Extensions struct {
		PersistedQuery `json:"persistedQuery"`
	}

	PersistedQuery struct {
		Version    int    `json:"version"`
		Sha256Hash string `json:"sha256Hash"`
	}
)
