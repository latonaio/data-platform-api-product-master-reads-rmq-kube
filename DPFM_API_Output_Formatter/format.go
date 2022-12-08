package dpfm_api_output_formatter

import (
	"data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(sdc *api_input_reader.SDC, rows *sql.Rows) (*General, error) {
	pm := &requests.General{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.ProductType,
			&pm.BaseUnit,
			&pm.ValidityStartDate,
			&pm.ProductGroup,
			&pm.Division,
			&pm.GrossWeight,
			&pm.WeightUnit,
			&pm.SizeOrDimensionText,
			&pm.IndustryStandardName,
			&pm.ProductStandardID,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.NetWeight,
			&pm.CountryOfOrigin,
			&pm.ItemCategory,
			&pm.ProductAccountAssignmentGroup,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	general := &General{
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
	return general, nil
}

func ConvertToProductDescriptionByBusinessPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*ProductDescriptionByBusinessPartner, error) {
	pm := &requests.ProductDescriptionByBusinessPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Language,
			&pm.ProductDescription,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	productDescriptionByBusinessPartner := &ProductDescriptionByBusinessPartner{
		Product:            data.Product,
		BusinessPartner:    data.BusinessPartner,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
	return productDescriptionByBusinessPartner, nil
}

// func (sdc *SDC) ConvertToBusinessPartner() *requests.BusinessPartner {
// 	dataGeneral := sdc.General
// 	data := sdc.BusinessPartner
// 	return &requests.BusinessPartner{
// 		Product:                dataGeneral.Product,
// 		BusinessPartner:        data.BusinessPartner,
// 		ValidityEndDate:        data.ValidityEndDate,
// 		ValidityStartDate:      data.ValidityStartDate,
// 		BusinessPartnerProduct: data.BusinessPartnerProduct,
// 		IsMarkedForDeletion:    data.IsMarkedForDeletion,
// 	}
// }

// func (sdc *SDC) ConvertToProcurement() *requests.Procurement {
// 	dataGeneral := sdc.General
// 	dataBusinessPartner := sdc.BusinessPartner
// 	data := sdc.Procurement
// 	return &requests.Procurement{
// 		Product:                     dataGeneral.Product,
// 		BusinessPartner:             dataBusinessPartner.BusinessPartner,
// 		Plant:                       data.Plant,
// 		Buyable:                     data.Buyable,
// 		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
// 		IsSourceListRequired:        data.IsSourceListRequired,
// 		IsMarkedForDeletion:         data.IsMarkedForDeletion,
// 	}
// }

// func (sdc *SDC) ConvertToBPPlant(num int) *requests.BPPlant {
// 	dataGeneral := sdc.General
// 	dataBusinessPartner := sdc.BusinessPartner
// 	data := sdc.BusinessPartner.BPPlant[num]
// 	return &requests.BPPlant{
// 		Product:                   dataGeneral.Product,
// 		BusinessPartner:           dataBusinessPartner.BusinessPartner,
// 		Plant:                     data.Plant,
// 		Issuable:                  data.Issuable,
// 		Receivable:                data.Receivable,
// 		IssuingStorageLocation:    data.IssuingStorageLocation,
// 		ReceivingStorageLocation:  data.ReceivingStorageLocation,
// 		AvailabilityCheckType:     data.AvailabilityCheckType,
// 		ProfitCenter:              data.ProfitCenter,
// 		MRPType:                   data.MRPType,
// 		MRPResponsible:            data.MRPResponsible,
// 		MinimumLotSizeQuantity:    data.MinimumLotSizeQuantity,
// 		MaximumLotSizeQuantity:    data.MaximumLotSizeQuantity,
// 		FixedLotSizeQuantity:      data.FixedLotSizeQuantity,
// 		IsBatchManagementRequired: data.IsBatchManagementRequired,
// 		ProcurementType:           data.ProcurementType,
// 		InventoryUnit:             data.InventoryUnit,
// 		IsMarkedForDeletion:       data.IsMarkedForDeletion,
// 	}
// }

// func (sdc *SDC) ConvertToProductDescription() *requests.ProductDescription {
// 	dataGeneral := sdc.General
// 	data := sdc.ProductDescription
// 	return &requests.ProductDescription{
// 		Product:            dataGeneral.Product,
// 		Language:           data.Language,
// 		ProductDescription: data.ProductDescription,
// 	}
// }
