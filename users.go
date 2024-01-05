package twitchgraphapi

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (s *Scrapper) GetUserData(username string) (UsersInfo, error) {
	var users UsersInfo

	data := fmt.Sprintf(
		`[{"operationName":"ChannelRoot_AboutPanel","variables":{"channelLogin":"%s","skipSchedule":true},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"6089531acef6c09ece01b440c41978f4c8dc60cb4fa0124c9a9d3f896709b6c6"}}}]`,
		username)

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	s.headers.CopyTo(req)

	req.SetBody([]byte(data))

	err := s.httpClient.Do(req, res)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(res.Body(), &users)
	if err != nil {
		return users, err
	}
	return users, nil
}
