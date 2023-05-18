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
	var quality *[]dpfm_api_output_formatter.Quality
	var storageLocation *[]dpfm_api_output_formatter.StorageLocation
	var storageBin *[]dpfm_api_output_formatter.StorageBin
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
				productDescByBP = c.ProductDescByBP(mtx, input, output, errs, log)
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
		case "BPPlants":
			func() {
				bPPlant = c.BPPlants(mtx, input, output, errs, log)
			}()
		case "MRPArea":
			func() {
				mrpArea = c.MRPArea(mtx, input, output, errs, log)
			}()
		case "MRPAreas":
			func() {
				mrpArea = c.MRPAreas(mtx, input, output, errs, log)
			}()
		case "Quality":
			func() {
				quality = c.Quality(mtx, input, output, errs, log)
			}()
		case "Qualities":
			func() {
				quality = c.Qualities(mtx, input, output, errs, log)
			}()
		case "StorageLocation":
			func() {
				storageLocation = c.StorageLocation(mtx, input, output, errs, log)
			}()
		case "StorageLocations":
			func() {
				storageLocation = c.StorageLocations(mtx, input, output, errs, log)
			}()
		case "StorageBin":
			func() {
				storageBin = c.StorageBin(mtx, input, output, errs, log)
			}()
		case "StorageBins":
			func() {
				storageBin = c.StorageBins(mtx, input, output, errs, log)
			}()
		case "WorkScheduling":
			func() {
				workScheduling = c.WorkScheduling(mtx, input, output, errs, log)
			}()
		case "WorkSchedulings":
			func() {
				workScheduling = c.WorkSchedulings(mtx, input, output, errs, log)
			}()
		case "Tax":
			func() {
				tax = c.Tax(mtx, input, output, errs, log)
			}()
		case "Taxes":
			func() {
				tax = c.Taxes(mtx, input, output, errs, log)
			}()
		case "Accounting":
			func() {
				accounting = c.Accounting(mtx, input, output, errs, log)
			}()
		case "Accountings":
			func() {
				accounting = c.Accountings(mtx, input, output, errs, log)
			}()
		case "GeneralsByBP":
			func() {
				general = c.GeneralsByBP(mtx, input, output, errs, log)
			}()
		case "ProductDescsByBP":
			func() {
				productDescByBP = c.ProductDescsByBP(mtx, input, output, errs, log)
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
		Quality:            quality,
		StorageLocation:    storageLocation,
		StorageBin:         storageBin,
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
	productDescription := input.General.BusinessPartner[0].ProductDescription

	cnt := 0
	for _, v := range productDescription {
		args = append(args, product, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data
		WHERE (Product, Language) IN ( `+repeat+` ) 
		ORDER BY Product DESC;`, args...,
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

func (c *DPFMAPICaller) ProductDescByBP(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescByBP {
	in := ""
	for i, v := range input.General.BusinessPartner[0].ProductDescription[0].ProductDescByBP {
		if i == 0 {
			in = fmt.Sprintf("( %d, '%s', '%s' )", v.BusinessPartner, v.Language, v.Product)
			continue
		}
		in = fmt.Sprintf("%s ,( %d, '%s', '%s' )", in, v.BusinessPartner, v.Language, v.Product)
	}

	where := fmt.Sprintf(" WHERE ( BusinessPartner, Language, Product ) IN ( %s ) ", in)
	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		` + where + ` ORDER BY Product DESC, BusinessPartner DESC;`,
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
		WHERE (Product, BusinessPartner) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, args...,
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
	allergen := input.General.BusinessPartner[0].Allergen

	cnt := 0
	for _, v := range allergen {
		args = append(args, product, v.BusinessPartner, v.Allergen)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_allergen_data
		WHERE (Product, BusinessPartner, Allergen) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, args...,
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
	allergen := input.General.BusinessPartner[0].Allergen[0]

	where := fmt.Sprintf("WHERE pma.Product = '%s' AND pma.BusinessPartner = %d", allergen.Product, allergen.BusinessPartner)

	rows, err := c.db.Query(
		`SELECT pma.Product, pma.BusinessPartner, aa.AllergenName, pma.AllergenIsContained, pma.IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_allergen_data pma
		INNER JOIN data_platform_allergen_allergen_text_data as aa ON pma.Allergen = aa.Allergen
		` + where + ` ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`,
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
	product := input.General.Product
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_nutritional_info_data
		WHERE Product = ? 
		ORDER BY Product DESC, BusinessPartner DESC;`, product,
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
	product := input.General.Product
	bp := input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_calories_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY Product DESC, BusinessPartner DESC;`, product, *bp,
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
	bPPlant := input.General.BPPlant

	cnt := 0
	for _, v := range bPPlant {
		args = append(args, product, v.BusinessPartner, v.Plant)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, args...,
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

func (c *DPFMAPICaller) BPPlants(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BPPlant {
	product := input.General.Product
	businessPartner := input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, product, *businessPartner,
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
	bPPlant := input.General.BPPlant

	cnt := 0
	for _, v := range bPPlant {
		mrpArea := v.MRPArea
		for _, w := range mrpArea {
			args = append(args, product, v.BusinessPartner, v.Plant, w.MRPArea)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_mrp_area_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, args...,
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

func (c *DPFMAPICaller) MRPAreas(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.MRPArea {
	product := input.General.Product
	businessPartner := *input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_mrp_area_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, product, businessPartner,
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

func (c *DPFMAPICaller) Quality(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Quality {
	var args []interface{}
	product := input.General.Product
	bPPlant := input.General.BPPlant

	cnt := 0
	for _, v := range bPPlant {
		args = append(args, product, v.BusinessPartner, v.Plant)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_quality_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToQuality(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Qualities(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Quality {
	product := input.General.Product
	businessPartner := *input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_quality_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, product, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToQuality(rows)
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
	bPPlant := input.General.BPPlant

	cnt := 0
	for _, v := range bPPlant {
		storageLocation := v.StorageLocation
		for _, w := range storageLocation {
			args = append(args, product, v.BusinessPartner, v.Plant, w.StorageLocation)
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

func (c *DPFMAPICaller) StorageLocations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageLocation {
	product := input.General.Product
	businessPartner := input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_location_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, InventoryBlockStatus DESC, BusinessPartner DESC;`, product, *businessPartner,
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

func (c *DPFMAPICaller) StorageBin(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageBin {
	var args []interface{}
	product := input.General.Product
	bPPlant := input.General.BPPlant

	cnt := 0
	for _, v := range bPPlant {
		storageLocation := v.StorageLocation
		for _, w := range storageLocation {
			args = append(args, product, v.BusinessPartner, v.Plant, w.StorageLocation, w.StorageBin)
		}
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_bin_data
		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, InventoryBlockStatus DESC, BusinessPartner DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStorageBin(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) StorageBins(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageBin {
	product := input.General.Product
	businessPartner := input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_bin_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, InventoryBlockStatus DESC, BusinessPartner DESC;`, product, *businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStorageBin(rows)
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
	bPPlant := input.General.BPPlant

	cnt := 0
	for _, v := range bPPlant {
		args = append(args, product, v.BusinessPartner, v.Plant)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_work_scheduling_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, args...,
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

func (c *DPFMAPICaller) WorkSchedulings(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.WorkScheduling {
	product := input.General.Product
	businessPartner := input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_work_scheduling_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC;`, product, *businessPartner,
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
		`SELECT Product, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE (Product, Country, ProductTaxCategory) IN ( `+repeat+` ) 
		ORDER BY Product DESC;`, args...,
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

func (c *DPFMAPICaller) Taxes(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Tax {
	product := input.General.Product

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE Product = ? 
		ORDER BY Product DESC;`, product,
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
	bPPlant := input.General.BPPlant

	cnt := 0
	for _, v := range bPPlant {
		args = append(args, product, v.BusinessPartner, v.Plant)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_accounting_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` ) 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC, Plant DESC;`, args...,
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

func (c *DPFMAPICaller) Accountings(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	product := input.General.Product
	businessPartner := input.BusinessPartnerID

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_accounting_data
		WHERE Product = ? AND BusinessPartner = ? 
		ORDER BY IsMarkedForDeletion DESC, Product DESC, BusinessPartner DESC, Plant DESC;`, product, *businessPartner,
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

func (c *DPFMAPICaller) GeneralsByBP(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	where := fmt.Sprintf("WHERE bp.BusinessPartner = %d ", *input.BusinessPartnerID)
	if input.General.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s AND g.IsMarkedForDeletion = %v", where, *input.General.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT g.Product, g.ProductGroup, g.BaseUnit, g.ValidityStartDate, g.IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_business_partner_data as bp
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data as g
		ON bp.Product = g.Product
		` + where + ` ORDER BY bp.IsMarkedForDeletion DESC, bp.Product DESC, bp.BusinessPartner DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneralsByBP(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductDescsByBP(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescByBP {
	where := fmt.Sprintf("WHERE BusinessPartner = %d ", *input.BusinessPartnerID)

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		` + where + ` ORDER BY Product DESC, BusinessPartner DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductDescsByBP(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
