package cache

import (
	"strings"
	"sync"

	"github.com/ONSdigital/dp-topic-api/models"
)

// SubtopicsIDs contains a list of subtopics in map form with addition to mutex locking
// The subtopicsMap is used to keep a record of subtopics to be later used to generate the subtopics id `query` for a topic
// and to check if the subtopic id given by a user exists
type SubtopicsIDs struct {
	mutex        *sync.RWMutex
	subtopicsMap map[string]*models.Topic
}

// NewSubTopicsMap creates a new subtopics id map to store subtopic ids with addition to mutex locking
func NewSubTopicsMap() *SubtopicsIDs {
	return &SubtopicsIDs{
		mutex:        &sync.RWMutex{},
		subtopicsMap: make(map[string]*models.Topic),
	}
}

// Get returns a bool value for the given key (id) to inform that the subtopic id exists
func (t *SubtopicsIDs) Get(key string) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.subtopicsMap[key] != nil {
		return true
	}

	return false
}

// GetSubtopicItems returns a list of subtopics for given topic
func (t *SubtopicsIDs) GetSubtopicItems() map[string]*models.Topic {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	subtopics := make(map[string]*models.Topic)

	for id, value := range t.subtopicsMap {
		subtopics[id] = value
	}

	return subtopics
}

// GetSubtopicsIDsQuery gets the subtopics ID query for a topic
func (t *SubtopicsIDs) GetSubtopicsIDsQuery() string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	ids := make([]string, 0, len(t.subtopicsMap))

	for id := range t.subtopicsMap {
		ids = append(ids, id)
	}

	return strings.Join(ids, ",")
}

// AppendSubtopicID appends the subtopic id to the map stored in SubtopicsIDs with consideration to mutex locking
func (t *SubtopicsIDs) AppendSubtopicItems(topic *models.Topic) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.subtopicsMap == nil {
		t.subtopicsMap = make(map[string]*models.Topic)
	}

	t.subtopicsMap[topic.ID] = topic
}
