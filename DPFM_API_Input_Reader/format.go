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
		ItemCategory:                  data.ItemCategory,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToGeneralPDF() *requests.GeneralPDF {
	dataGeneral := sdc.General
	data := sdc.General.GeneralPDF
	return &requests.GeneralPDF{
		Product:      dataGeneral.Product,
		DocType:      data.DocType,
		DocVersionID: data.DocVersionID,
		DocID:        data.DocID,
		FileName:     data.FileName,
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
	}
}

func (sdc *SDC) ConvertToProcurement() *requests.Procurement {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	dataBPPlant := sdc.General.BusinessPartner.BPPlant
	data := sdc.General.BusinessPartner.BPPlant.Procurement
	return &requests.Procurement{
		Product:                     dataGeneral.Product,
		BusinessPartner:             dataBusinessPartner.BusinessPartner,
		Plant:                       dataBPPlant.Plant,
		Buyable:                     data.Buyable,
		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
		IsSourceListRequired:        data.IsSourceListRequired,
		IsMarkedForDeletion:         data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToBPPlant() *requests.BPPlant {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	data := sdc.General.BusinessPartner.BPPlant
	return &requests.BPPlant{
		Product:                   dataGeneral.Product,
		BusinessPartner:           dataBusinessPartner.BusinessPartner,
		Plant:                     data.Plant,
		Issuable:                  data.Issuable,
		Receivable:                data.Receivable,
		IssuingStorageLocation:    data.IssuingStorageLocation,
		ReceivingStorageLocation:  data.ReceivingStorageLocation,
		AvailabilityCheckType:     data.AvailabilityCheckType,
		ProfitCenter:              data.ProfitCenter,
		MRPType:                   data.MRPType,
		MRPResponsible:            data.MRPResponsible,
		MinimumLotSizeQuantity:    data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:    data.MaximumLotSizeQuantity,
		FixedLotSizeQuantity:      data.FixedLotSizeQuantity,
		IsBatchManagementRequired: data.IsBatchManagementRequired,
		ProcurementType:           data.ProcurementType,
		InventoryUnit:             data.InventoryUnit,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToProductDescription() *requests.ProductDescription {
	dataGeneral := sdc.General
	data := sdc.General.ProductDescription
	return &requests.ProductDescription{
		Product:            dataGeneral.Product,
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
		BusinessPartner:          dataBusinessPartner.BusinessPartner,
		Country:                  data.Country,
		TaxCategory:              data.TaxCategory,
		ProductTaxClassification: data.ProductTaxClassification,
	}
}

func (sdc *SDC) ConvertToSales() *requests.Sales {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	data := sdc.General.BusinessPartner.Sales
	return &requests.Sales{
		Product:             dataGeneral.Product,
		BusinessPartner:     dataBusinessPartner.BusinessPartner,
		Sellable:            data.Sellable,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
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

func (sdc *SDC) ConvertToMRPArea() *requests.MRPArea {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	dataBPPlant := sdc.General.BusinessPartner.BPPlant
	data := sdc.General.BusinessPartner.BPPlant.MRPArea
	return &requests.MRPArea{
		Product:                       dataGeneral.Product,
		BusinessPartner:               dataBusinessPartner.BusinessPartner,
		Plant:                         dataBPPlant.Plant,
		MRPArea:                       data.MRPArea,
		MRPType:                       data.MRPType,
		MRPResponsible:                data.MRPResponsible,
		MRPGroup:                      data.MRPGroup,
		ReorderThresholdQuantity:      data.ReorderThresholdQuantity,
		PlanningTimeFence:             data.PlanningTimeFence,
		LotSizeRoundingQuantity:       data.LotSizeRoundingQuantity,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		MaximumStockQuantity:          data.MaximumStockQuantity,
		DfltStorageLocationExtProcmt:  data.DfltStorageLocationExtProcmt,
		MRPPlanningCalendar:           data.MRPPlanningCalendar,
		SafetyStockQuantity:           data.SafetyStockQuantity,
		SafetyDuration:                data.SafetyDuration,
		FixedLotSizeQuantity:          data.FixedLotSizeQuantity,
		PlannedDeliveryDurationInDays: data.PlannedDeliveryDurationInDays,
		StorageLocation:               data.StorageLocation,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToStorageLocation() *requests.StorageLocation {
	dataGeneral := sdc.General
	dataBusinessPartner := sdc.General.BusinessPartner
	dataBPPlant := sdc.General.BusinessPartner.BPPlant
	data := sdc.General.BusinessPartner.BPPlant.StorageLocation
	return &requests.StorageLocation{
		Product:              dataGeneral.Product,
		BusinessPartner:      dataBusinessPartner.BusinessPartner,
		Plant:                dataBPPlant.Plant,
		StorageLocation:      data.StorageLocation,
		CreationDate:         data.CreationDate,
		InventoryBlockStatus: data.InventoryBlockStatus,
		IsMarkedForDeletion:  data.IsMarkedForDeletion,
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
