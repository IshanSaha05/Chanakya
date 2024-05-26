package mongodb

import (
	"github.com/IshanSaha05/pkg/models"
)

func (object *MongoDB) InsertIntoDBAssemblyStateData(datas []models.AssemblyStateData) error {
	for _, data := range datas {
		_, err := object.collection.InsertOne(object.context, data)

		if err != nil {
			return err
		}
	}

	return nil
}

func (object *MongoDB) InsertIntoDBAssemblyStateConstituencyData(datas []models.AssemblyStateConstituencyData) error {
	for _, data := range datas {
		_, err := object.collection.InsertOne(object.context, data)

		if err != nil {
			return err
		}
	}

	return nil
}
