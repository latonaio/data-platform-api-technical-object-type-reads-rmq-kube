package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-technical-object-type-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-technical-object-type-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var technicalObjectType *[]dpfm_api_output_formatter.TechnicalObjectType
	var technicalObjectTypeText *[]dpfm_api_output_formatter.TechnicalObjectTypeText
	for _, fn := range accepter {
		switch fn {
		case "TechnicalObjectType":
			func() {
				technical-objectType = c.TechnicalObjectType(mtx, input, output, errs, log)
			}()
		case "TechnicalObjectTypeText":
			func() {
				technical-objectTypeText = c.TechnicalObjectTypeText(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		TechnicalObjectType:     technical-objectType,
		TechnicalObjectTypeText: technical-objectTypeText,
	}

	return data
}

func (c *DPFMAPICaller) TechnicalObjectType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.TechnicalObjectType {
	technical-objectType := input.TechnicalObjectType[0].TechnicalObjectType

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_technical_object_type_technical_object_type_data
		WHERE TechnicalObjectType = ?;`, technicalObjectType,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToTechnicalObjectType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) TechnicalObjectTypeText(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.TechnicalObjectTypeText {
	where := "WHERE  (TechnicalObjectType, Language) IN "
	in := ""
	for _, v := range input.TechnicalObjectType {
		for _, vv := range v.TechnicalObjectTypeText {
			in = fmt.Sprintf("%s ( '%s', '%s' ), ", in, v.TechnicalObjectType, vv.Language)
		}
	}

	where = fmt.Sprintf("%s ( %s )", where, in[:len(in)-2])
	c.l.Info(where)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_technical_object_type_technical-object_type_text_data
		` + where + ` ;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToTechnicalObjectTypeText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
