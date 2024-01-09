package twitchgraphapi

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (s *Scrapper) GetVideoMetadata(username string, videoIds []string) ([]VideoMetadata, error) {

	var data []VideoMetadata

	body := make([]Operation, 0, len(videoIds))

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	s.headers.CopyTo(req)

	for _, video := range videoIds {
		body = append(body, Operation{
			OperationName: videoMetadataOperation,
			Variables: struct {
				ChannelLogin string `json:"channelLogin"`
				VideoID      string `json:"videoID"`
			}{ChannelLogin: username,
				VideoID: video,
			},
			Extensions: Extensions{
				PersistedQuery{Version: version, Sha256Hash: sha256Hash},
			},
		})
	}

	b, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	req.SetRequestURI(twitchApiUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody(b)

	err = s.httpClient.Do(req, res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
