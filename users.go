package twitchgraphapi

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (s *Scrapper) GetUserData(username string) ([]UsersInfo, error) {
	var users []UsersInfo

	data := fmt.Sprintf(
		`[{"operationName":"ChannelRoot_AboutPanel","variables":{"channelLogin":"%s","skipSchedule":true},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"6089531acef6c09ece01b440c41978f4c8dc60cb4fa0124c9a9d3f896709b6c6"}}},{"operationName":"HomeTrackQuery","variables":{"channelLogin":"stray228"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"129cbf14117ce8e95f01bd2043154089058146664df866d0314e84355ffb4e05"}}}]`,
		username)

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	s.headers.CopyTo(req)

	req.SetBody([]byte(data))
	req.SetRequestURI(twitchApiUrl)
	req.Header.SetMethod(fasthttp.MethodPost)

	err := s.httpClient.Do(req, res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Body(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Scrapper) GetUserCategories(username string) (UsersCategories, error) {
	var categories UsersCategories

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	s.headers.CopyTo(req)

	data := fmt.Sprintf(
		`{"operationName":"HomeShelfGames","variables":{"channelLogin":"%s"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"cb7711739c2b520ebf89f3027863c0f985e8094df91cc5ef28896d57375a9700"}}}`,
		username)

	req.SetBody([]byte(data))
	req.SetRequestURI(twitchApiUrl)
	req.Header.SetMethod(fasthttp.MethodPost)

	err := s.httpClient.Do(req, res)
	if err != nil {
		return categories, err
	}
	err = json.Unmarshal(res.Body(), &categories)
	if err != nil {
		return categories, err
	}

	return categories, nil

}
