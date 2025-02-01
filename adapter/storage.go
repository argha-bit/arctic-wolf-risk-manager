package adapter

import (
	"arctic-wolf-risk-manager/models"
	"log"
	"sync"

	"github.com/matryer/resync"
)

type RiskStorage struct {
	data  map[string]models.Risk
	mutex sync.RWMutex
}

var StorageInstance *RiskStorage
var onceStorageInstance resync.Once

func GetInstance() *RiskStorage {
	onceStorageInstance.Do(
		func() {
			StorageInstance = &RiskStorage{
				data: make(map[string]models.Risk),
			}
		})
	return StorageInstance
}

func (r *RiskStorage) Set(riskModel *models.Risk) {
	r.mutex.Lock()
	r.data[riskModel.Id] = *riskModel
	r.mutex.Unlock()
}
func (r *RiskStorage) Get(id string) []models.Risk {
	data := []models.Risk{}
	if id == "" {
		log.Println("no id recieved, returning entire list")
		for _, risk := range r.data {
			data = append(data, risk)
		}
		return data
	} else {
		if record, ok := r.data[id]; !ok {
			log.Println("no record found for given id", id)
			return data
		} else {
			data = append(data, record)
		}
		return data
	}
}
