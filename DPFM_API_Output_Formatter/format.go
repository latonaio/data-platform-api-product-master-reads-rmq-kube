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
			&pm.GrossWeight,
			&pm.NetWeight,
			&pm.WeightUnit,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.SizeOrDimensionText,
			&pm.ProductStandardID,
			&pm.IndustryStandardName,
			&pm.ItemCategory,
			&pm.BarcodeType,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.ProductAccountAssignmentGroup,
			&pm.CreationDate,
			&pm.LastChangeDate,
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
		GrossWeight:                   data.GrossWeight,
		NetWeight:                     data.NetWeight,
		WeightUnit:                    data.WeightUnit,
		InternalCapacityQuantity:      data.InternalCapacityQuantity,
		InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		ProductStandardID:             data.ProductStandardID,
		IndustryStandardName:          data.IndustryStandardName,
		ItemCategory:                  data.ItemCategory,
		BarcodeType:                   data.BarcodeType,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
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

func ConvertToBusinessPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*BusinessPartner, error) {
	pm := &requests.BusinessPartner{}

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
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.BusinessPartnerProduct,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	businessPartner := &BusinessPartner{
		Product:                data.Product,
		BusinessPartner:        data.BusinessPartner,
		ValidityEndDate:        data.ValidityEndDate,
		ValidityStartDate:      data.ValidityStartDate,
		BusinessPartnerProduct: data.BusinessPartnerProduct,
		IsMarkedForDeletion:    data.IsMarkedForDeletion,
	}
	return businessPartner, nil
}

func ConvertToProductDescription(sdc *api_input_reader.SDC, rows *sql.Rows) (*ProductDescription, error) {
	pm := &requests.ProductDescription{}

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
			&pm.Language,
			&pm.ProductDescription,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	productDescription := &ProductDescription{
		Product:            data.Product,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
	return productDescription, nil
}

func ConvertToProductDescriptions(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ProductDescription, error) {
	pds := make([]ProductDescription, 0, len(sdc.General.Products))
	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		pd := requests.ProductDescription{}
		err := rows.Scan(
			&pd.Product,
			&pd.Language,
			&pd.ProductDescription,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
		pds = append(pds, ProductDescription{
			Product:            pd.Product,
			Language:           pd.Language,
			ProductDescription: pd.ProductDescription,
		})
	}

	return &pds, nil
}

func ConvertToBPPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*BPPlant, error) {
	pm := &requests.BPPlant{}

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
			&pm.Plant,
			&pm.AvailabilityCheckType,
			&pm.ProfitCenter,
			&pm.MRPType,
			&pm.MRPController,
			&pm.ReorderThresholdQuantity,
			&pm.PlanningTimeFence,
			&pm.MRPPlanningCalendar,
			&pm.SafetyStockQuantityInBaseUnit,
			&pm.SafetyDuration,
			&pm.MaximumStockQuantityInBaseUnit,
			&pm.MinimumDeliveryQuantityInBaseUnit,
			&pm.MinimumDeliveryLotSizeQuantityInBaseUnit,
			&pm.StandardDeliveryLotSizeQuantityInBaseUnit,
			&pm.DeliveryLotSizeRoundingQuantityInBaseUnit,
			&pm.MaximumDeliveryLotSizeQuantityInBaseUnit,
			&pm.MaximumDeliveryQuantityInBaseUnit,
			&pm.DeliveryLotSizeIsFixed,
			&pm.StandardDeliveryDurationInDays,
			&pm.IsBatchManagementRequired,
			&pm.BatchManagementPolicy,
			&pm.InventoryUnit,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	bPPlant := &BPPlant{
		Product:                                  data.Product,
		BusinessPartner:                          data.BusinessPartner,
		Plant:                                    data.Plant,
		AvailabilityCheckType:                    data.AvailabilityCheckType,
		ProfitCenter:                             data.ProfitCenter,
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
		IsBatchManagementRequired:                 data.IsBatchManagementRequired,
		BatchManagementPolicy:                     data.BatchManagementPolicy,
		InventoryUnit:                             data.InventoryUnit,
		IsMarkedForDeletion:                       data.IsMarkedForDeletion,
	}
	return bPPlant, nil
}

func ConvertToTax(sdc *api_input_reader.SDC, rows *sql.Rows) (*Tax, error) {
	pm := &requests.Tax{}

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
			&pm.Country,
			&pm.ProductTaxCategory,
			&pm.ProductTaxClassification,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	tax := &Tax{
		Product:                  data.Product,
		Country:                  data.Country,
		ProductTaxCategory:       data.ProductTaxCategory,
		ProductTaxClassification: data.ProductTaxClassification,
	}
	return tax, nil
}

func ConvertToAccounting(sdc *api_input_reader.SDC, rows *sql.Rows) (*Accounting, error) {
	pm := &requests.Accounting{}

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
			&pm.Plant,
			&pm.ValuationClass,
			&pm.CostingPolicy,
			&pm.PriceUnitQty,
			&pm.StandardPrice,
			&pm.MovingAveragePrice,
			&pm.PriceLastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	accounting := &Accounting{
		Product:             data.Product,
		BusinessPartner:     data.BusinessPartner,
		Plant:               data.Plant,
		ValuationClass:      data.ValuationClass,
		CostingPolicy:       data.CostingPolicy,
		PriceUnitQty:        data.PriceUnitQty,
		StandardPrice:       data.StandardPrice,
		MovingAveragePrice:  data.MovingAveragePrice,
		PriceLastChangeDate: data.PriceLastChangeDate,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
	return accounting, nil
}

func ConvertToStorageLocation(sdc *api_input_reader.SDC, rows *sql.Rows) (*StorageLocation, error) {
	pm := &requests.StorageLocation{}

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
			&pm.Plant,
			&pm.StorageLocation,
			&pm.CreationDate,
			&pm.InventoryBlockStatus,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	storageLocation := &StorageLocation{
		Product:              data.Product,
		BusinessPartner:      data.BusinessPartner,
		Plant:                data.Plant,
		StorageLocation:      data.StorageLocation,
		CreationDate:         data.CreationDate,
		InventoryBlockStatus: data.InventoryBlockStatus,
		IsMarkedForDeletion:  data.IsMarkedForDeletion,
	}
	return storageLocation, nil
}

func ConvertToMRPArea(sdc *api_input_reader.SDC, rows *sql.Rows) (*MRPArea, error) {
	pm := &requests.MRPArea{}

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
			&pm.Plant,
			&pm.MRPArea,
			&pm.StorageLocationForMRP,
			&pm.MRPType,
			&pm.MRPController,
			&pm.ReorderThresholdQuantity,
			&pm.PlanningTimeFence,
			&pm.MRPPlanningCalendar,
			&pm.SafetyStockQuantityInBaseUnit,
			&pm.SafetyDuration,
			&pm.MaximumStockQuantityInBaseUnit,
			&pm.MinimumDeliveryQuantityInBaseUnit,
			&pm.MinimumDeliveryLotSizeQuantityInBaseUnit,
			&pm.StandardDeliveryLotSizeQuantityInBaseUnit,
			&pm.DeliveryLotSizeRoundingQuantityInBaseUnit,
			&pm.MaximumDeliveryLotSizeQuantityInBaseUnit,
			&pm.MaximumDeliveryQuantityInBaseUnit,
			&pm.DeliveryLotSizeIsFixed,
			&pm.StandardDeliveryDurationInDays,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	mRPArea := &MRPArea{
		Product:                                  data.Product,
		BusinessPartner:                          data.BusinessPartner,
		Plant:                                    data.Plant,
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
	return mRPArea, nil
}

func ConvertToWorkScheduling(sdc *api_input_reader.SDC, rows *sql.Rows) (*WorkScheduling, error) {
	pm := &requests.WorkScheduling{}

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
			&pm.Plant,
			&pm.ProductionInvtryManagedLoc,
			&pm.ProductProcessingTime,
			&pm.ProductionSupervisor,
			&pm.ProductProductionQuantityUnit,
			&pm.ProdnOrderIsBatchRequired,
			&pm.MatlCompIsMarkedForBackflush,
			&pm.ProductionSchedulingProfile,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	workScheduling := &WorkScheduling{
		Product:                       data.Product,
		BusinessPartner:               data.BusinessPartner,
		Plant:                         data.Plant,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
	return workScheduling, nil
}
