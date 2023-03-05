package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
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
	var general *[]dpfm_api_output_formatter.General
	var productDescription *[]dpfm_api_output_formatter.ProductDescription
	var productDescByBP *[]dpfm_api_output_formatter.ProductDescByBP
	var businessPartner *[]dpfm_api_output_formatter.BusinessPartner
	var allergen *[]dpfm_api_output_formatter.Allergen
	var nutritionalInfo *[]dpfm_api_output_formatter.NutritionalInfo
	var calories *[]dpfm_api_output_formatter.Calories
	var bPPlant *[]dpfm_api_output_formatter.BPPlant
	var mrpArea *[]dpfm_api_output_formatter.MRPArea
	var storageLocation *[]dpfm_api_output_formatter.StorageLocation
	var workScheduling *[]dpfm_api_output_formatter.WorkScheduling
	var tax *[]dpfm_api_output_formatter.Tax
	var accounting *[]dpfm_api_output_formatter.Accounting
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "ProductDescription":
			func() {
				productDescription = c.ProductDescription(mtx, input, output, errs, log)
			}()
		case "ProductDescByBP":
			func() {
				productDescByBP = c.ProductDescByBusinessPartner(mtx, input, output, errs, log)
			}()
		case "BusinessPartner":
			func() {
				businessPartner = c.BusinessPartner(mtx, input, output, errs, log)
			}()
		case "Allergen":
			func() {
				allergen = c.Allergen(mtx, input, output, errs, log)
			}()
		case "Allergens":
			func() {
				allergen = c.Allergens(mtx, input, output, errs, log)
			}()
		case "NutritionalInfo":
			func() {
				nutritionalInfo = c.NutritionalInfo(mtx, input, output, errs, log)
			}()
		case "Calories":
			func() {
				calories = c.Calories(mtx, input, output, errs, log)
			}()
		case "BPPlant":
			func() {
				bPPlant = c.BPPlant(mtx, input, output, errs, log)
			}()
		case "MRPArea":
			func() {
				mrpArea = c.MRPArea(mtx, input, output, errs, log)
			}()
		case "StorageLocation":
			func() {
				storageLocation = c.StorageLocation(mtx, input, output, errs, log)
			}()
		case "WorkScheduling":
			func() {
				workScheduling = c.WorkScheduling(mtx, input, output, errs, log)
			}()
		case "Tax":
			func() {
				tax = c.Tax(mtx, input, output, errs, log)
			}()
		case "Accounting":
			func() {
				accounting = c.Accounting(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:            general,
		ProductDescription: productDescription,
		ProductDescByBP:    productDescByBP,
		BusinessPartner:    businessPartner,
		Allergen:           allergen,
		NutritionalInfo:    nutritionalInfo,
		Calories:           calories,
		BPPlant:            bPPlant,
		MRPArea:            mrpArea,
		StorageLocation:    storageLocation,
		WorkScheduling:     workScheduling,
		Tax:                tax,
		Accounting:         accounting,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	product := input.General.Product

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data
		WHERE Product = ?;`, product,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductDescription(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescription {
	var args []interface{}
	product := input.General.Product
	productDescription := input.General.ProductDescription

	cnt := 0
	for _, v := range productDescription {
		args = append(args, product, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data
		WHERE (Product, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductDescription(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductDescByBusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescByBP {
	in := ""
	for i, v := range input.General.ProductDescByBP {
		if i == 0 {
			in = fmt.Sprintf("( %d, '%s', '%s' )", v.BusinessPartner, v.Language, v.Product)
			continue
		}
		in = fmt.Sprintf("%s ,( %d, '%s', '%s' )", in, v.BusinessPartner, v.Language, v.Product)
	}

	where := fmt.Sprintf("WHERE ( BusinessPartner, Language, Product ) IN ( %s )", in)
	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescByBP(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BusinessPartner {
	var args []interface{}
	product := input.General.Product
	businessPartner := input.General.BusinessPartner

	cnt := 0
	for _, v := range businessPartner {
		args = append(args, product, v.BusinessPartner)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_business_partner_data
		WHERE (Product, BusinessPartner) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToBusinessPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	return data
}

// func (c *DPFMAPICaller) BusinessPartner(
// 	mtx *sync.Mutex,
// 	input *dpfm_api_input_reader.SDC,
// 	output *dpfm_api_output_formatter.SDC,
// 	errs *[]error,
// 	log *logger.Logger,
// ) *[]dpfm_api_output_formatter.BusinessPartner {
// 	var args []interface{}
// 	product := input.General.Product
// 	businessPartner := input.General.BusinessPartner

// 	cnt := 0
// 	for _, v := range businessPartner {
// 		args = append(args, product, v.BusinessPartner, v.ValidityEndDate, v.ValidityStartDate)
// 		cnt++
// 	}

// 	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"
// 	rows, err := c.db.Query(
// 		`SELECT *
// 		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_business_partner_data
// 		WHERE (Product, BusinessPartner, ValidityEndDate, ValidityStartDate) IN ( `+repeat+` );`, args...,
// 	)
// 	if err != nil {
// 		*errs = append(*errs, err)
// 		return nil
// 	}

// 	data, err := dpfm_api_output_formatter.ConvertToBusinessPartner(rows)
// 	if err != nil {
// 		*errs = append(*errs, err)
// 		return nil
// 	}
// 	defer rows.Close()

// 	return data
// }

func (c *DPFMAPICaller) Allergen(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Allergen {
	var args []interface{}
	product := input.General.Product
	allergen := input.General.Allergen

	cnt := 0
	for _, v := range allergen {
		args = append(args, product, v.BusinessPartner, v.Allergen)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_allergen_data
		WHERE (Product, BusinessPartner, Allergen) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAllergen(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Allergens(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Allergen {
	allergen := input.General.Allergen[0]

	where := fmt.Sprintf("WHERE pma.Product = '%s' AND pma.BusinessPartener = %d AND aa.Language = '%s'", allergen.Product, allergen.BusinessPartner, allergen.Language)

	rows, err := c.db.Query(
		`SELECT pma.Product, pma.BusinessPartener, aa.AllergenName, pma.AllergenIsContained
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_allergen_data pma
		INNER JOIN data_platform_allergen_allergen_text_data as aa ON pma.Allergen = aa.Allergen
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAllergen(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) NutritionalInfo(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.NutritionalInfo {
	var args []interface{}
	product := input.General.Product
	nutritionalInfo := input.General.NutritionalInfo

	cnt := 0
	for _, v := range nutritionalInfo {
		args = append(args, product, v.BusinessPartner, v.Nutrient)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_nutritional_info_data
		WHERE (Product, BusinessPartner, Nutrient) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToNutritionalInfo(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Calories(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Calories {
	var args []interface{}
	product := input.General.Product
	calories := input.General.Calories

	cnt := 0
	for _, v := range calories {
		args = append(args, product, v.BusinessPartner, v.Calories)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_calories_data
		WHERE (Product, BusinessPartner, Calories) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCalories(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BPPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BPPlant {
	var args []interface{}
	product := input.General.Product
	businessPartner := input.General.BusinessPartner

	cnt := 0
	for _, v := range businessPartner {
		bPPlant := v.BPPlant
		for _, w := range bPPlant {
			args = append(args, product, v.BusinessPartner, w.Plant)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToBPPlant(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) MRPArea(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.MRPArea {
	var args []interface{}
	product := input.General.Product
	businessPartner := input.General.BusinessPartner

	cnt := 0
	for _, v := range businessPartner {
		bPPlant := v.BPPlant
		for _, w := range bPPlant {
			mrpArea := w.MRPArea
			for _, x := range mrpArea {
				args = append(args, product, v.BusinessPartner, w.Plant, x.MRPArea)
			}
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_mrp_area_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToMRPArea(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) StorageLocation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageLocation {
	var args []interface{}
	product := input.General.Product
	businessPartner := input.General.BusinessPartner

	cnt := 0
	for _, v := range businessPartner {
		bPPlant := v.BPPlant
		for _, w := range bPPlant {
			storageLocation := w.StorageLocation
			for _, x := range storageLocation {
				args = append(args, product, v.BusinessPartner, w.Plant, x.StorageLocation)
			}
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_location_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStorageLocation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) WorkScheduling(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.WorkScheduling {
	var args []interface{}
	product := input.General.Product
	businessPartner := input.General.BusinessPartner

	cnt := 0
	for _, v := range businessPartner {
		bPPlant := v.BPPlant
		for _, w := range bPPlant {
			args = append(args, product, v.BusinessPartner, w.Plant)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_work_scheduling_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToWorkScheduling(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Tax(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Tax {
	var args []interface{}
	product := input.General.Product
	tax := input.General.Tax

	cnt := 0
	for _, v := range tax {
		args = append(args, product, v.Country, v.ProductTaxCategory)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE (Product, BusinessPartner, Country) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToTax(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Accounting(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	var args []interface{}
	product := input.General.Product
	businessPartner := input.General.BusinessPartner

	cnt := 0
	for _, v := range businessPartner {
		bPPlant := v.BPPlant
		for _, w := range bPPlant {
			args = append(args, product, v.BusinessPartner, w.Plant)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_accounting_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAccounting(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
