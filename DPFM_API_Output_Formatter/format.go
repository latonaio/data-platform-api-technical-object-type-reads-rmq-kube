package dpfm_api_output_formatter

import (
	"data-platform-api-technical-object-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToTechnicalObjectType(rows *sql.Rows) (*[]TechnicalObjectType, error) {
	defer rows.Close()
	technicalObjectType := make([]TechnicalObjectType, 0)

	i := 0
	for rows.Next() {
		pm := &requests.TechnicalObjectType{}
		i++

		err := rows.Scan(
			&pm.TechnicalObjectType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		technicalObjectType = append(technicalObjectType, TechnicalObjectType{
			TechnicalObjectType: 	data.TechnicalObjectType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &technicalObjectType, nil
}

func ConvertToTechnicalObjectTypeText(rows *sql.Rows) (*[]TechnicalObjectTypeText, error) {
	defer rows.Close()
	technicalObjectTypeText := make([]TechnicalObjectTypeText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := requests.TechnicalObjectTypeText{}

		err := rows.Scan(
			&pm.TechnicalObjectType,
			&pm.Language,
			&pm.TechnicalObjectTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &technicalObjectTypeText, err
		}

		data := pm
		technicalObjectTypeText = append(technicalObjectTypeText, TechnicalObjectTypeText{
			TechnicalObjectType:     data.TechnicalObjectType,
			Language:          		 data.Language,
			TechnicalObjectTypeName: data.TechnicalObjectTypeName,
			CreationDate:			 data.CreationDate,
			LastChangeDate:			 data.LastChangeDate,
			IsMarkedForDeletion:	 data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &technicalObjectTypeText, nil
	}

	return &technicalObjectTypeText, nil
}
