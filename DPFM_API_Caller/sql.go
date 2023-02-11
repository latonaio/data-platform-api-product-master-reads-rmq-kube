package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Output_Formatter"
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
	var general *dpfm_api_output_formatter.General
	var productDescription *dpfm_api_output_formatter.ProductDescription
	var productDescriptions []*dpfm_api_output_formatter.ProductDescription
	var productDescriptionByBusinessPartner *dpfm_api_output_formatter.ProductDescriptionByBusinessPartner
	var businessPartner *dpfm_api_output_formatter.BusinessPartner
	var bPPlant *dpfm_api_output_formatter.BPPlant
	var mrpArea *dpfm_api_output_formatter.MRPArea
	var storageLocation *dpfm_api_output_formatter.StorageLocation
	var workScheduling *dpfm_api_output_formatter.WorkScheduling
	var tax *dpfm_api_output_formatter.Tax
	var accounting *dpfm_api_output_formatter.Accounting
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
		case "ProductDescription":
			func() {
				productDescription = c.ProductDescription(mtx, input, output, errs, log)
			}()
		case "ProductDescriptions":
			func() {
				productDescription = c.ProductDescription(mtx, input, output, errs, log)
			}()
		case "BusinessPartner":
			func() {
				businessPartner = c.BusinessPartner(mtx, input, output, errs, log)
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
		General:                             general,
		ProductDescription:                  productDescription,
		ProductDescriptions:                 productDescriptions,
		ProductDescriptionByBusinessPartner: productDescriptionByBusinessPartner,
		BusinessPartner:                     businessPartner,
		BPPlant:                             bPPlant,
		MRPArea:                             mrpArea,
		StorageLocation:                     storageLocation,
		WorkScheduling:                      workScheduling,
		Tax:                                 tax,
		Accounting:                          accounting,
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
		`SELECT Product, ProductType, BaseUnit, ValidityStartDate, ProductGroup, GrossWeight, NetWeight, WeightUnit,
		InternalCapacityQuantity, InternalCapacityQuantityUnit, SizeOrDimensionText, ProductStandardID, IndustryStandardName, ItemCategory, BarcodeType,
		CountryOfOrigin, CountryOfOriginLanguage, ProductAccountAssignmentGroup, CreationDate, LastChangeDate, IsMarkedForDeletion,
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

func (c *DPFMAPICaller) ProductDescription(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductDescription {
	product := input.General.Product
	language := input.General.ProductDescription.Language

	rows, err := c.db.Query(
		`SELECT Product, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data
		WHERE (Product, Language) = (?, ?);`, product, language,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescription(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductDescriptions(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescription {
	var args []interface{}
	products := input.General.Products

	cnt := 0
	for _, v := range products {
		args = append(args, v)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data
		WHERE Product IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescriptions(input, rows)
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

func (c *DPFMAPICaller) BusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.BusinessPartner {
	product := input.General.Product
	businessPartner := input.General.BusinessPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, ValidityEndDate, ValidityStartDate, BusinessPartnerProduct, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_business_partner_data
		WHERE (Product, BusinessPartner) = (?, ?);`, product, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToBusinessPartner(input, rows)
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
) *dpfm_api_output_formatter.BPPlant {
	product := input.General.Product
	businessPartner := input.General.BusinessPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Plant, AvailabilityCheckType, ProfitCenter, MRPType, MRPController, 
		ReorderThresholdQuantity, PlanningTimeFence, MRPPlanningCalendar, SafetyStockQuantityInBaseUnit, 
		SafetyDuration, MaximumStockQuantityInBaseUnit, MinimumDeliveryQuantityInBaseUnit, MinimumDeliveryLotSizeQuantityInBaseUnit, 
		StandardDeliveryLotSizeQuantityInBaseUnit, DeliveryLotSizeRoundingQuantityInBaseUnit, MaximumDeliveryLotSizeQuantityInBaseUnit, 
		MaximumDeliveryQuantityInBaseUnit, DeliveryLotSizeIsFixed, StandardDeliveryDurationInDays, IsBatchManagementRequired, 
		BatchManagementPolicy, InventoryUnit, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner) = (?, ?);`, product, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToBPPlant(input, rows)
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
) *dpfm_api_output_formatter.MRPArea {
	product := input.General.Product
	businessPartner := input.General.BusinessPartner.BusinessPartner
	plant := input.General.BusinessPartner.BPPlant.Plant

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Plant, MRPArea, StorageLocationForMRP, MRPType, MRPController, 
		ReorderThresholdQuantity, PlanningTimeFence, MRPPlanningCalendar, SafetyStockQuantityInBaseUnit, 
		SafetyDuration, MaximumStockQuantityInBaseUnit, MinimumDeliveryQuantityInBaseUnit, MinimumDeliveryLotSizeQuantityInBaseUnit, 
		StandardDeliveryLotSizeQuantityInBaseUnit, DeliveryLotSizeRoundingQuantityInBaseUnit, MaximumDeliveryLotSizeQuantityInBaseUnit, 
		MaximumDeliveryQuantityInBaseUnit, DeliveryLotSizeIsFixed, StandardDeliveryDurationInDays, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_mrp_area_data
		WHERE (Product, BusinessPartner, Plant) = (?, ?, ?);`, product, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToMRPArea(input, rows)
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
) *dpfm_api_output_formatter.StorageLocation {
	product := input.General.Product
	businessPartner := input.General.BusinessPartner.BusinessPartner
	plant := input.General.BusinessPartner.BPPlant.Plant

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Plant, StorageLocation, CreationDate, InventoryBlockStatus, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_storage_location_data
		WHERE (Product, BusinessPartner, Plant) = (?, ?, ?);`, product, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageLocation(input, rows)
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
) *dpfm_api_output_formatter.WorkScheduling {
	product := input.General.Product
	businessPartner := input.General.BusinessPartner.BusinessPartner
	plant := input.General.BusinessPartner.BPPlant.Plant

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Plant, ProductionInvtryManagedLoc, ProductProcessingTime, ProductionSupervisor, 
		ProductProductionQuantityUnit, ProdnOrderIsBatchRequired, MatlCompIsMarkedForBackflush, ProductionSchedulingProfile, 
		IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_work_scheduling_data
		WHERE (Product, BusinessPartner, Plant) = (?, ?, ?);`, product, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToWorkScheduling(input, rows)
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
) *dpfm_api_output_formatter.Tax {
	product := input.General.Product
	businessPartner := input.General.BusinessPartner.BusinessPartner
	country := input.General.BusinessPartner.Tax.Country

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE (Product, BusinessPartner, Country) = (?, ?, ?);`, product, businessPartner, country,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToTax(input, rows)
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
) *dpfm_api_output_formatter.Accounting {
	product := input.General.Product
	businessPartner := input.General.BusinessPartner.BusinessPartner
	plant := input.General.BusinessPartner.BPPlant.Plant

	rows, err := c.db.Query(
		`SELECT Product, BusinessPartner, Plant, ValuationClass, CostingPolicy, PriceUnitQty, StandardPrice, 
		MovingAveragePrice, PriceLastChangeDate, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_accounting_data
		WHERE (Product, BusinessPartner, Plant) = (?, ?, ?);`, product, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAccounting(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
