package mongodb

import "github.com/IshanSaha05/pkg/models"

func (object *MongoDB) InsertIntoDBDistrictStateDistrictData(datas []models.DistrictStateDistrictData) error {
	for _, data := range datas {
		_, err := object.collection.InsertOne(object.context, data)

		if err != nil {
			return err
		}
	}

	return nil
}
