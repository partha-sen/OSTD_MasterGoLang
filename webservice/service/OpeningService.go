package service

import (
	"encoding/json"
	"fmt"

	"github.com/partha-sen/ostd/webservice/dao"
	"github.com/partha-sen/ostd/webservice/model"
	"github.com/pkg/errors"
)

func GetOpeningById(id int) (model.Any, error) {

	data, err := dao.GetOpeningById(id)
	if err != nil {
		return data, errors.Wrap(err, fmt.Sprintf("Couldn't retrieve Opening for id  %d", id))
	}
	return data, nil
}

func GetAllOpening() ([]model.Any, error) {

	data, err := dao.GetAllOpening()

	if err != nil {
		return data, errors.Wrap(err, "Couldn't retrieve Opening from database")
	}
	return data, nil

}

func jsonUnmarshel(by []byte) (model.Opening, error) {
	var o model.Opening
	err := json.Unmarshal(by, &o)
	return o, err
}

func SaveOpening(by []byte) (int64, error) {
	obj, err := jsonUnmarshel(by)

	if err != nil {
		return 0, err
	}

	id, err := dao.InsertOpening(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't save record into database")
	}
	return id, nil
}

func UpdateOpening(by []byte, pathId int) (int64, error) {

	obj, err := jsonUnmarshel(by)

	if err != nil {
		return 0, err
	}
	obj.Id = pathId
	id, err := dao.UpdateOpening(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't update record into database")
	}
	return id, nil
}

func DeleteOpening(id int) (int64, error) {

	count, err := dao.DeleteOpening(id)

	if err != nil {
		return count, errors.Wrap(err, "Couldn't delete record")
	}
	return count, nil
}
