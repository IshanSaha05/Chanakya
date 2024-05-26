package mongodb

import (
	"github.com/IshanSaha05/pkg/models"
)

func (object *MongoDB) InsertIntoDBParliamentStateConstituencyData(datas []models.ParliamentStateConstituencyData) error {

	for _, data := range datas {
		_, err := object.collection.InsertOne(object.context, data)

		if err != nil {
			return err
		}
	}

	return nil
}
