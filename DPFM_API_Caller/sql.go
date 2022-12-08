package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Output_Formatter"
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
	var general *dpfm_api_output_formatter.General
	var productDescriptionByBusinessPartner *dpfm_api_output_formatter.ProductDescriptionByBusinessPartner
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "ProductDescriptionByBusinessPartner":
			func() {
				productDescriptionByBusinessPartner = c.ProductDescriptionByBusinessPartner(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:                             general,
		ProductDescriptionByBusinessPartner: productDescriptionByBusinessPartner,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	product := input.General.Product

	rows, err := c.db.Query(
		`SELECT Product, ProductType, BaseUnit, ValidityStartDate, ProductGroup, Division, GrossWeight, WeightUnit,
		SizeOrDimensionText, IndustryStandardName, ProductStandardID, CreationDate, LastChangeDate, NetWeight,
		CountryOfOrigin, ItemCategory, ProductAccountAssignmentGroup, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data
		WHERE Product = ?;`, product,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneral(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductDescriptionByBusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductDescriptionByBusinessPartner {
	product := input.General.Product
	businessPartner := input.General.ProductDescriptionByBusinessPartner.BusinessPartner
	language := input.General.ProductDescriptionByBusinessPartner.Language

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		WHERE (Product, BusinessPartner, Language) = (?, ?, ?);`, product, businessPartner, language,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescriptionByBusinessPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
