package storage

import (
	"fmt"
	"time"

	"github.com/bontaurean/go-test-proxy/models"
	"github.com/lithammer/shortuuid/v3"
)

type entryList map[string]*models.HistoryEntry

type historyStorage struct {
	entries entryList
}

var History = createHstoryStorage()

func createHstoryStorage() *historyStorage {
	return &historyStorage{entries: entryList{}}
}

func (s *historyStorage) Add(req models.ProxyRequest, resp models.ProxyResponse) string {
	requestId := shortuuid.New()
	s.entries[requestId] = &models.HistoryEntry{
		AddedAt:  time.Now(),
		Request:  req,
		Response: resp,
	}
	return requestId
}

func (s *historyStorage) Get(requestId string) (*models.HistoryEntry, error) {
	entry, ok := s.entries[requestId]
	if !ok {
		return nil, fmt.Errorf("request %s not known", requestId)
	}
	return entry, nil
}
