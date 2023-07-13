package dpfm_api_output_formatter

import (
	"data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.General{}

		err := rows.Scan(
			&pm.Product,
			&pm.ProductType,
			&pm.BaseUnit,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.ItemCategory,
			&pm.ProductGroup,
			&pm.GrossWeight,
			&pm.NetWeight,
			&pm.WeightUnit,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
			&pm.SizeOrDimensionText,
			&pm.ProductStandardID,
			&pm.IndustryStandardName,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.LocalRegionOfOrigin,
			&pm.LocalSubRegionOfOrigin,
			&pm.BarcodeType,
			&pm.ProductAccountAssignmentGroup,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &general, err
		}

		data := pm
		general = append(general, General{
			Product:                       data.Product,
			ProductType:                   data.ProductType,
			BaseUnit:                      data.BaseUnit,
			ValidityStartDate:             data.ValidityStartDate,
			ValidityEndDate:               data.ValidityEndDate,
			ItemCategory:                  data.ItemCategory,
			ProductGroup:                  data.ProductGroup,
			GrossWeight:                   data.GrossWeight,
			NetWeight:                     data.NetWeight,
			WeightUnit:                    data.WeightUnit,
			InternalCapacityQuantity:      data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
			SizeOrDimensionText:           data.SizeOrDimensionText,
			ProductStandardID:             data.ProductStandardID,
			IndustryStandardName:          data.IndustryStandardName,
			CountryOfOrigin:               data.CountryOfOrigin,
			CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
			LocalRegionOfOrigin:           data.LocalRegionOfOrigin,
			LocalSubRegionOfOrigin:        data.LocalSubRegionOfOrigin,
			BarcodeType:                   data.BarcodeType,
			ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &general, nil
	}

	return &general, nil
}

func ConvertToProductDescription(rows *sql.Rows) (*[]ProductDescription, error) {
	defer rows.Close()
	productDescription := make([]ProductDescription, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductDescription{}

		err := rows.Scan(
			&pm.Product,
			&pm.Language,
			&pm.ProductDescription,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productDescription, err
		}

		data := pm
		productDescription = append(productDescription, ProductDescription{
			Product:             data.Product,
			Language:            data.Language,
			ProductDescription:  data.ProductDescription,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productDescription, nil
	}

	return &productDescription, nil
}

func ConvertToProductDescByBP(rows *sql.Rows) (*[]ProductDescByBP, error) {
	defer rows.Close()
	productDescByBP := make([]ProductDescByBP, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductDescByBP{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Language,
			&pm.ProductDescription,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productDescByBP, err
		}

		data := pm
		productDescByBP = append(productDescByBP, ProductDescByBP{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Language:            data.Language,
			ProductDescription:  data.ProductDescription,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productDescByBP, nil
	}

	return &productDescByBP, nil
}

func ConvertToProductDescsByBP(rows *sql.Rows) (*[]ProductDescByBP, error) {
	defer rows.Close()
	productDescsByBP := make([]ProductDescByBP, 0)
	i := 0
	for rows.Next() {
		i++
		pm := ProductDescByBP{}
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Language,
			&pm.ProductDescription,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productDescsByBP, err
		}
		productDescsByBP = append(productDescsByBP, pm)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productDescsByBP, nil
	}

	return &productDescsByBP, nil
}

func ConvertToBusinessPartner(rows *sql.Rows) (*[]BusinessPartner, error) {
	defer rows.Close()
	businessPartner := make([]BusinessPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.BusinessPartner{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.BusinessPartnerProduct,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &businessPartner, err
		}

		data := pm
		businessPartner = append(businessPartner, BusinessPartner{
			Product:                data.Product,
			BusinessPartner:        data.BusinessPartner,
			ValidityStartDate:      data.ValidityStartDate,
			ValidityEndDate:        data.ValidityEndDate,
			BusinessPartnerProduct: data.BusinessPartnerProduct,
			CreationDate:           data.CreationDate,
			LastChangeDate:         data.LastChangeDate,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &businessPartner, nil
	}

	return &businessPartner, nil
}

func ConvertToBPPlant(rows *sql.Rows) (*[]BPPlant, error) {
	defer rows.Close()
	bPPlant := make([]BPPlant, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.BPPlant{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.MRPType,
			&pm.MRPController,
			&pm.ReorderThresholdQuantityInBaseUnit,
			&pm.PlanningTimeFenceInDays,
			&pm.MRPPlanningCalendar,
			&pm.SafetyStockQuantityInBaseUnit,
			&pm.SafetyDuration,
			&pm.SafetyDurationUnit,
			&pm.MaximumStockQuantityInBaseUnit,
			&pm.MinimumDeliveryQuantityInBaseUnit,
			&pm.MinimumDeliveryLotSizeQuantityInBaseUnit,
			&pm.StandardDeliveryQuantityInBaseUnit,
			&pm.StandardDeliveryLotSizeQuantityInBaseUnit,
			&pm.MaximumDeliveryQuantityInBaseUnit,
			&pm.MaximumDeliveryLotSizeQuantityInBaseUnit,
			&pm.DeliveryLotSizeRoundingQuantityInBaseUnit,
			&pm.DeliveryLotSizeIsFixed,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.IsBatchManagementRequired,
			&pm.BatchManagementPolicy,
			&pm.ProfitCenter,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &bPPlant, err
		}

		data := pm
		bPPlant = append(bPPlant, BPPlant{
			Product:                                   data.Product,
			BusinessPartner:                           data.BusinessPartner,
			Plant:                                     data.Plant,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
			PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			SafetyDurationUnit:                        data.SafetyDurationUnit,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
			MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDuration:                  data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
			IsBatchManagementRequired:                 data.IsBatchManagementRequired,
			BatchManagementPolicy:                     data.BatchManagementPolicy,
			ProfitCenter:                              data.ProfitCenter,
			CreationDate:                              data.CreationDate,
			LastChangeDate:                            data.LastChangeDate,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &bPPlant, nil
	}

	return &bPPlant, nil
}

func ConvertToTax(rows *sql.Rows) (*[]Tax, error) {
	defer rows.Close()
	tax := make([]Tax, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Tax{}

		err := rows.Scan(
			&pm.Product,
			&pm.Country,
			&pm.ProductTaxCategory,
			&pm.ProductTaxClassification,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &tax, err
		}

		data := pm
		tax = append(tax, Tax{
			Product:                  data.Product,
			Country:                  data.Country,
			ProductTaxCategory:       data.ProductTaxCategory,
			ProductTaxClassification: data.ProductTaxClassification,
			CreationDate:             data.CreationDate,
			LastChangeDate:           data.LastChangeDate,
			IsMarkedForDeletion:      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &tax, nil
	}

	return &tax, nil
}

func ConvertToAccounting(rows *sql.Rows) (*[]Accounting, error) {
	defer rows.Close()
	accounting := make([]Accounting, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Accounting{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.ValuationClass,
			&pm.CostingPolicy,
			&pm.PriceUnitQty,
			&pm.StandardPrice,
			&pm.MovingAveragePrice,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &accounting, err
		}

		data := pm
		accounting = append(accounting, Accounting{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Plant:               data.Plant,
			ValuationClass:      data.ValuationClass,
			CostingPolicy:       data.CostingPolicy,
			PriceUnitQty:        data.PriceUnitQty,
			StandardPrice:       data.StandardPrice,
			MovingAveragePrice:  data.MovingAveragePrice,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &accounting, nil
	}

	return &accounting, nil
}

func ConvertToStorageLocation(rows *sql.Rows) (*[]StorageLocation, error) {
	defer rows.Close()
	storageLocation := make([]StorageLocation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.StorageLocation{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.BlockStatus,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &storageLocation, err
		}

		data := pm
		storageLocation = append(storageLocation, StorageLocation{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Plant:               data.Plant,
			StorageLocation:     data.StorageLocation,
			BlockStatus:         data.BlockStatus,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &storageLocation, nil
	}

	return &storageLocation, nil
}

func ConvertToStorageBin(rows *sql.Rows) (*[]StorageBin, error) {
	defer rows.Close()
	storageBin := make([]StorageBin, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.StorageBin{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageBin,
			&pm.BlockStatus,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &storageBin, err
		}

		data := pm
		storageBin = append(storageBin, StorageBin{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Plant:               data.Plant,
			StorageLocation:     data.StorageLocation,
			StorageBin:          data.StorageBin,
			BlockStatus:         data.BlockStatus,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &storageBin, nil
	}

	return &storageBin, nil
}

func ConvertToMRPArea(rows *sql.Rows) (*[]MRPArea, error) {
	defer rows.Close()
	mRPArea := make([]MRPArea, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.MRPArea{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.MRPArea,
			&pm.StorageLocationForMRP,
			&pm.MRPType,
			&pm.MRPController,
			&pm.ReorderThresholdQuantityInBaseUnit,
			&pm.PlanningTimeFenceInDays,
			&pm.MRPPlanningCalendar,
			&pm.SafetyStockQuantityInBaseUnit,
			&pm.SafetyDuration,
			&pm.MaximumStockQuantityInBaseUnit,
			&pm.MinumumDeliveryQuantityInBaseUnit,
			&pm.MinumumDeliveryLotSizeQuantityInBaseUnit,
			&pm.StandardDeliveryLotSizeQuantityInBaseUnit,
			&pm.DeliveryLotSizeRoundingQuantityInBaseUnit,
			&pm.MaximumDeliveryLotSizeQuantityInBaseUnit,
			&pm.MaximumDeliveryQuantityInBaseUnit,
			&pm.DeliveryLotSizeIsFixed,
			&pm.StandardDeliveryDurationUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &mRPArea, err
		}

		data := pm
		mRPArea = append(mRPArea, MRPArea{
			Product:                                   data.Product,
			BusinessPartner:                           data.BusinessPartner,
			Plant:                                     data.Plant,
			MRPArea:                                   data.MRPArea,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			StorageLocationForMRP:                     data.StorageLocationForMRP,
			ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
			PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			SafetyDurationUnit:                        data.SafetyDurationUnit,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinumumDeliveryQuantityInBaseUnit:         data.MinumumDeliveryQuantityInBaseUnit,
			MinumumDeliveryLotSizeQuantityInBaseUnit:  data.MinumumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDuration:                  data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
			CreationDate:                              data.CreationDate,
			LastChangeDate:                            data.LastChangeDate,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &mRPArea, nil
	}

	return &mRPArea, nil
}

func ConvertToQuality(rows *sql.Rows) (*[]Quality, error) {
	defer rows.Close()
	quality := make([]Quality, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Quality{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.MaximumStoragePeriod,
			&pm.QualityMgmtCtrlKey,
			&pm.MatlQualityAuthorizationGroup,
			&pm.HasPostToInspectionStock,
			&pm.InspLotDocumentationIsRequired,
			&pm.SuplrQualityManagementSystem,
			&pm.RecrrgInspIntervalTimeInDays,
			&pm.ProductQualityCertificateType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &quality, err
		}

		data := pm
		quality = append(quality, Quality{
			Product:                        data.Product,
			BusinessPartner:                data.BusinessPartner,
			Plant:                          data.Plant,
			MaximumStoragePeriod:           data.MaximumStoragePeriod,
			QualityMgmtCtrlKey:             data.QualityMgmtCtrlKey,
			MatlQualityAuthorizationGroup:  data.MatlQualityAuthorizationGroup,
			HasPostToInspectionStock:       data.HasPostToInspectionStock,
			InspLotDocumentationIsRequired: data.InspLotDocumentationIsRequired,
			SuplrQualityManagementSystem:   data.SuplrQualityManagementSystem,
			RecrrgInspIntervalTimeInDays:   data.RecrrgInspIntervalTimeInDays,
			ProductQualityCertificateType:  data.ProductQualityCertificateType,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &quality, nil
	}

	return &quality, nil
}

func ConvertToProduction(rows *sql.Rows) (*[]Production, error) {
	defer rows.Close()
	production := make([]Production, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Production{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.ProductionStorageLocation,
			&pm.ProductProcessingDuration,
			&pm.ProductProductionQuantityUnit,
			&pm.MinimumProductionQuantityInBaseUnit,
			&pm.MinimumProductionLotSizeQuantityInBaseUnit,
			&pm.StandardProductionQuantityInBaseUnit,
			&pm.StandardProductionLotSizeQuantityInBaseUnit,
			&pm.MaximumProductionQuantityInBaseUnit,
			&pm.MaximumProductionLotSizeQuantityInBaseUnit,
			&pm.ProductionLotSizeRoundingQuantityInBaseUnit,
			&pm.MinimumProductionQuantityInProductionUnit,
			&pm.MinimumProductionLotSizeQuantityInProductionUnit,
			&pm.StandardProductionQuantityInProductionUnit,
			&pm.StandardProductionLotSizeQuantityInProductionUnit,
			&pm.MaximumProductionLotSizeQuantityInProductionUnit,
			&pm.MaximumProductionQuantityInProductionUnit,
			&pm.ProductionLotSizeRoundingQuantityInProductionUnit,
			&pm.ProductionLotSizeIsFixed,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.ProductIsMarkedForBackflush,
			&pm.ProductionSchedulingProfile,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &production, err
		}

		data := pm
		production = append(production, Production{
			Product:                             data.Product,
			BusinessPartner:                     data.BusinessPartner,
			Plant:                               data.Plant,
			ProductionStorageLocation:           data.ProductionStorageLocation,
			ProductProcessingDuration:           data.ProductProcessingDuration,
			ProductProductionQuantityUnit:       data.ProductProductionQuantityUnit,
			MinimumProductionQuantityInBaseUnit: data.MinimumProductionQuantityInBaseUnit,
			MinimumProductionLotSizeQuantityInBaseUnit:        data.MinimumProductionLotSizeQuantityInBaseUnit,
			StandardProductionQuantityInBaseUnit:              data.StandardProductionQuantityInBaseUnit,
			StandardProductionLotSizeQuantityInBaseUnit:       data.StandardProductionLotSizeQuantityInBaseUnit,
			MaximumProductionQuantityInBaseUnit:               data.MaximumProductionQuantityInBaseUnit,
			MaximumProductionLotSizeQuantityInBaseUnit:        data.MaximumProductionLotSizeQuantityInBaseUnit,
			ProductionLotSizeRoundingQuantityInBaseUnit:       data.ProductionLotSizeRoundingQuantityInBaseUnit,
			MinimumProductionQuantityInProductionUnit:         data.MinimumProductionQuantityInProductionUnit,
			MinimumProductionLotSizeQuantityInProductionUnit:  data.MinimumProductionLotSizeQuantityInProductionUnit,
			StandardProductionQuantityInProductionUnit:        data.StandardProductionQuantityInProductionUnit,
			StandardProductionLotSizeQuantityInProductionUnit: data.StandardProductionLotSizeQuantityInProductionUnit,
			MaximumProductionLotSizeQuantityInProductionUnit:  data.MaximumProductionLotSizeQuantityInProductionUnit,
			MaximumProductionQuantityInProductionUnit:         data.MaximumProductionQuantityInProductionUnit,
			ProductionLotSizeRoundingQuantityInProductionUnit: data.ProductionLotSizeRoundingQuantityInProductionUnit,
			ProductionLotSizeIsFixed:                          data.ProductionLotSizeIsFixed,
			ProductIsBatchManagedInProductionPlant:            data.ProductIsBatchManagedInProductionPlant,
			ProductIsMarkedForBackflush:                       data.ProductIsMarkedForBackflush,
			ProductionSchedulingProfile:                       data.ProductionSchedulingProfile,
			CreationDate:                                      data.CreationDate,
			LastChangeDate:                                    data.LastChangeDate,
			IsMarkedForDeletion:                               data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &production, nil
	}

	return &production, nil
}

func ConvertToAllergen(rows *sql.Rows) (*[]Allergen, error) {
	defer rows.Close()
	allergen := make([]Allergen, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Allergen{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Allergen,
			&pm.AllergenIsContained,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &allergen, err
		}

		data := pm
		allergen = append(allergen, Allergen{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Allergen:            data.Allergen,
			AllergenIsContained: data.AllergenIsContained,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &allergen, nil
	}

	return &allergen, nil
}

func ConvertToNutritionalInfo(rows *sql.Rows) (*[]NutritionalInfo, error) {
	defer rows.Close()
	nutritionalInfo := make([]NutritionalInfo, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.NutritionalInfo{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Nutrient,
			&pm.NutrientContent,
			&pm.NutrientContentUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &nutritionalInfo, err
		}

		data := pm
		nutritionalInfo = append(nutritionalInfo, NutritionalInfo{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Nutrient:            data.Nutrient,
			NutrientContent:     data.NutrientContent,
			NutrientContentUnit: data.NutrientContentUnit,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &nutritionalInfo, nil
	}

	return &nutritionalInfo, nil
}

func ConvertToCalories(rows *sql.Rows) (*[]Calories, error) {
	defer rows.Close()
	calories := make([]Calories, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Calories{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Calories,
			&pm.CaloryUnit,
			&pm.CaloryUnitQuantity,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &calories, err
		}

		data := pm
		calories = append(calories, Calories{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Calories:            data.Calories,
			CaloryUnit:          data.CaloryUnit,
			CaloryUnitQuantity:  data.CaloryUnitQuantity,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &calories, nil
	}

	return &calories, nil
}
