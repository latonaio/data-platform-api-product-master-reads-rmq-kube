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
			return &general, err
		}

		data := pm
		general = append(general, General{

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
			CountryOfOrigin:               data.CountryOfOrigin,
			CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
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
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productDescByBP, err
		}

		data := pm
		productDescByBP = append(productDescByBP, ProductDescByBP{
			Product:            data.Product,
			BusinessPartner:    data.BusinessPartner,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productDescByBP, nil
	}

	return &productDescByBP, nil
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
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.BusinessPartnerProduct,
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
			ValidityEndDate:        data.ValidityEndDate,
			ValidityStartDate:      data.ValidityStartDate,
			BusinessPartnerProduct: data.BusinessPartnerProduct,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &businessPartner, nil
	}

	return &businessPartner, nil
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
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productDescription, err
		}

		data := pm
		productDescription = append(productDescription, ProductDescription{
			Product:            data.Product,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productDescription, nil
	}

	return &productDescription, nil
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
			&pm.CaloryUnitQuantity,
			&pm.Calories,
			&pm.CaloryUnit,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &calories, err
		}

		data := pm
		calories = append(calories, Calories{
			Product:            data.Product,
			BusinessPartner:    data.BusinessPartner,
			CaloryUnitQuantity: data.CaloryUnitQuantity,
			Calories:           data.Calories,
			CaloryUnit:         data.CaloryUnit,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &calories, nil
	}

	return &calories, nil
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
			&pm.AvailabilityCheckType,
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
			&pm.ProfitCenter,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &bPPlant, err
		}

		data := pm
		bPPlant = append(bPPlant, BPPlant{
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
			&pm.PriceLastChangeDate,
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
			PriceLastChangeDate: data.PriceLastChangeDate,
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
			&pm.CreationDate,
			&pm.InventoryBlockStatus,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &storageLocation, err
		}

		data := pm
		storageLocation = append(storageLocation, StorageLocation{
			Product:              data.Product,
			BusinessPartner:      data.BusinessPartner,
			Plant:                data.Plant,
			StorageLocation:      data.StorageLocation,
			CreationDate:         data.CreationDate,
			InventoryBlockStatus: data.InventoryBlockStatus,
			IsMarkedForDeletion:  data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &storageLocation, nil
	}

	return &storageLocation, nil
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
			&pm.ReorderThresholdQuantity,
			&pm.PlanningTimeFence,
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
			&pm.StandardDeliveryDurationInDays,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &mRPArea, err
		}

		data := pm
		mRPArea = append(mRPArea, MRPArea{
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
			MinumumDeliveryQuantityInBaseUnit:        data.MinumumDeliveryQuantityInBaseUnit,
			MinumumDeliveryLotSizeQuantityInBaseUnit: data.MinumumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &mRPArea, nil
	}

	return &mRPArea, nil
}

func ConvertToWorkScheduling(rows *sql.Rows) (*[]WorkScheduling, error) {
	defer rows.Close()
	workScheduling := make([]WorkScheduling, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.WorkScheduling{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.ProductionInvtryManagedLoc,
			&pm.ProductProcessingTime,
			&pm.ProductionSupervisor,
			&pm.ProductProductionQuantityUnit,
			&pm.ProdnOrderIsBatchRequired,
			&pm.PDTCompIsMarkedForBackflush,
			&pm.ProductionSchedulingProfile,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &workScheduling, err
		}

		data := pm
		workScheduling = append(workScheduling, WorkScheduling{
			Product:                       data.Product,
			BusinessPartner:               data.BusinessPartner,
			Plant:                         data.Plant,
			ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
			ProductProcessingTime:         data.ProductProcessingTime,
			ProductionSupervisor:          data.ProductionSupervisor,
			ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
			ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
			PDTCompIsMarkedForBackflush:   data.PDTCompIsMarkedForBackflush,
			ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &workScheduling, nil
	}

	return &workScheduling, nil
}
