package config

import (
	"github.com/pkg/errors"
)

const (
	maxReconnectTries = 3
)

type UrlsInfo struct {
	Urls            []string
	CurrentUrlIndex uint32
	ReconnectTries  map[uint32]uint32
}

func NewUrlsInfo(networkConfig BaseNetworkConfig) (*UrlsInfo, error) {
	if len(networkConfig.Urls) == 0 {
		return nil, errors.New("no url provided")
	}
	newUrlsInfo := &UrlsInfo{
		Urls:            make([]string, 0),
		CurrentUrlIndex: 0,
		ReconnectTries:  make(map[uint32]uint32),
	}
	for _, url := range networkConfig.Urls {
		newUrlsInfo.Urls = append(newUrlsInfo.Urls, url)
	}
	return newUrlsInfo, nil
}

func (u *UrlsInfo) GetCurrentUrl() string {
	return u.Urls[u.CurrentUrlIndex]
}

// NextUrl This method will look until we can find an url that has less than max reconnect tries otherwise will return ""
func (u *UrlsInfo) NextUrl() string {
	// increase tries, we only call this method when we cannot connect
	u.ReconnectTries[u.CurrentUrlIndex] += 1
	newIndex := u.CurrentUrlIndex
	for {
		newIndex = (newIndex + 1) % uint32(len(u.Urls))
		if u.ReconnectTries[newIndex] >= maxReconnectTries {
			if u.CurrentUrlIndex == newIndex {
				return ""
			}
			continue
		}
		u.CurrentUrlIndex = newIndex
		break
	}
	return u.Urls[u.CurrentUrlIndex]
}

// Clear We clear all retries because in the future some might work
func (u *UrlsInfo) Clear() {
	for k, _ := range u.ReconnectTries {
		u.ReconnectTries[k] = 0
	}
}
