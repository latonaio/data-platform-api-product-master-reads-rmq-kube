package dpfm_api_input_reader

import (
	"data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.General
	return &requests.General{
		Product:                       data.Product,
		ProductType:                   &data.ProductType,
		BaseUnit:                      &data.BaseUnit,
		ValidityStartDate:             &data.ValidityStartDate,
		ProductGroup:                  &data.ProductGroup,
		Division:                      &data.Division,
		GrossWeight:                   data.GrossWeight,
		WeightUnit:                    &data.WeightUnit,
		SizeOrDimensionText:           &data.SizeOrDimensionText,
		IndustryStandardName:          &data.IndustryStandardName,
		ProductStandardID:             &data.ProductStandardID,
		CreationDate:                  &data.CreationDate,
		LastChangeDate:                &data.LastChangeDate,
		NetWeight:                     data.NetWeight,
		CountryOfOrigin:               &data.CountryOfOrigin,
		ItemCategory:                  &data.ItemCategory,
		ProductAccountAssignmentGroup: &data.ProductAccountAssignmentGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToBusinessPartner() *requests.BusinessPartner {
	dataGeneral := sdc.General
	data := sdc.General.BusinessPartner
	return &requests.BusinessPartner{
		Product:             dataGeneral.Product,
		BusinessPartner:     data.BusinessPartner,
		ValidityEndDate:     data.ValidityEndDate,
		ValidityStartDate:   data.ValidityStartDate,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
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
