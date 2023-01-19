package dpfm_api_input_reader

import (
	"data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.General
	return &requests.General{
		Product:                       data.Product,
		ProductType:                   data.ProductType,
		BaseUnit:                      data.BaseUnit,
		ValidityStartDate:             data.ValidityStartDate,
		ProductGroup:                  data.ProductGroup,
		Division:                      data.Division,
		GrossWeight:                   data.GrossWeight,
		WeightUnit:                    data.WeightUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		IndustryStandardName:          data.IndustryStandardName,
		ProductStandardID:             data.ProductStandardID,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
		NetWeight:                     data.NetWeight,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ItemCategory:                  data.ItemCategory,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
		BusinessPartner:               data.BusinessPartner,
		ProductDescription:            data.ProductDescription,
		Tax:                           data.Tax,
	}
}

func (sdc *SDC) ConvertToBusinessPartner() *requests.BusinessPartner {
	dataGeneral := sdc.General
	data := sdc.General.BusinessPartner
	return &requests.BusinessPartner{
		Product:                dataGeneral.Product,
		BusinessPartner:        data.BusinessPartner,
		ValidityEndDate:        data.ValidityEndDate,
		ValidityStartDate:      data.ValidityStartDate,
		BusinessPartnerProduct: data.BusinessPartnerProduct,
		IsMarkedForDeletion:    data.IsMarkedForDeletion,
		BPPlant:                data.BPPlant,
		ProductDescByBP:        data.ProductDescByBP,
	}
}

func (sdc *SDC) ConvertToBPPlant() *requests.BPPlant {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	data := sdc.General.BusinessPartner.BPPlant
	return &requests.BPPlant{
		Product:                                   data.Product,
		BusinessPartner:                           data.BusinessPartner,
		Plant:                                     data.Plant,
		AvailabilityCheckType:                     data.AvailabilityCheckType,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
		PlanningTimeFence:                         data.PlanningTimeFence,
		MRPPlanningCalender:                       data.MRPPlanningCalender,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
		MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
		IsBatchManagementRequired:                 data.IsBatchManagementRequired,
		BatchManagementPolicy:                     data.BatchManagementPolicy,
		InventoryUnit:                             data.InventoryUnit,
		ProfitCenter:                              data.ProfitCenter,
		IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		StorageLocation:                           data.StorageLocation,
		MRPArea:                                   data.MRPArea,
		WorkScheduling:                            data.WorkScheduling,
		Accounting:                                data.Accounting,
	}
}

func (sdc *SDC) ConvertToStorageLocation() *requests.StorageLocation {
	dataGeneral := sdc.General
	data := sdc.General.StorageLocation
	return &requests.StorageLocation{
		Product:              dataGeneral.Product,
		BusinessPartner:      data.BusinessPartner,
		Plant:                data.Plant,
		StorageLocation:      data.StorageLocation,
		CreationDate:         data.CreationDate,
		InventoryBlockStatus: data.InventoryBlockStatus,
		IsMarkedForDeletion:  data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToMRPArea() *requests.MRPArea {
	dataGeneral := sdc.General
	data := sdc.General.MRPArea
	return &requests.MRPArea{
		Product:                                  dataGeneral.Product,
		BusinessPartner:                          dataBusinessPartner.BusinessPartner,
		Plant:                                    dataBPPlant.Plant,
		MRPArea:                                  data.MRPArea,
		StorageLocationForMRP:                    data.StorageLocationForMRP,
		MRPType:                                  data.MRPType,
		MRPController:                            data.MRPController,
		ReorderThresholdQuantity:                 data.ReorderThresholdQuantity,
		PlanningTimeFence:                        data.PlanningTimeFence,
		MRPPlanningCalendar:                      data.MRPPlanningCalendar,
		SafetyStockQuantityInBaseUnit:            data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                           data.SafetyDuration,
		MaximumStockQuantityInBaseUnit:           data.MaximumStockQuantityInBaseUnit,
		MinimumDeliveryQuantityInBaseUnit:        data.MinimumDeliveryQuantityInBaseUnit,
		MinimumDeliveryLotSizeQuantityInBaseUnit: data.MinimumDeliveryLotSizeQuantityInBaseUnit,
		StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
		IsMarkedForDeletion:                       data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToWorkScheduling() *requests.WorkScheduling {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	dataBPPlant := sdc.General.BusinessPartner.BPPlant
	data := sdc.General.BusinessPartner.BPPlant.WorkScheduling
	return &requests.WorkScheduling{
		Product:                       dataGeneral.Product,
		BusinessPartner:               dataBusinessPartner.BusinessPartner,
		Plant:                         dataBPPlant.Plant,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	dataBPPlant := sdc.General.BusinessPartner.BPPlant
	data := sdc.General.BusinessPartner.BPPlant.Accounting
	return &requests.Accounting{
		Product:             dataGeneral.Product,
		BusinessPartner:     dataBusinessPartner.BusinessPartner,
		Plant:               dataBPPlant.Plant,
		ValuationClass:      data.ValuationClass,
		CostingPolicy:       data.CostingPolicy,
		PriceUnitQty:        data.PriceUnitQty,
		StandardPrice:       data.StandardPrice,
		MovingAveragePrice:  data.MovingAveragePrice,
		PriceLastChangeDate: data.PriceLastChangeDate,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToProductDescription() *requests.ProductDescription {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	data := sdc.General.ProductDescription
	return &requests.ProductDescription{
		Product:            dataGeneral.Product,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
}

func (sdc *SDC) ConvertToProductDescriptionByBusinessPartner() *requests.ProductDescriptionByBusinessPartner {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	data := sdc.General.ProductDescriptionByBusinessPartner
	return &requests.ProductDescriptionByBusinessPartner{
		Product:            dataGeneral.Product,
		BusinessPartner:    dataBusinessPartner.BusinessPartner,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
}

func (sdc *SDC) ConvertToTax() *requests.Tax {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	data := sdc.General.BusinessPartner.Tax
	return &requests.Tax{
		Product:                  dataGeneral.Product,
		Country:                  data.Country,
		ProductTaxCategory:       data.ProductTaxCategory,
		ProductTaxClassification: data.ProductTaxClassification,
	}
}
