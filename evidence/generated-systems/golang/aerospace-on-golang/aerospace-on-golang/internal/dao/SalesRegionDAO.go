package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SalesRegionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSalesRegion - creates a new db entry
//----------------------------------------------------------------------------
func CreateSalesRegion(obj model.SalesRegion)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a SalesRegion with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SalesRegion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSalesRegion", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSalesRegion - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSalesRegion(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SalesRegion

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SalesRegion with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SalesRegion using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SalesRegion using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSalesRegion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSalesRegion - returns all
//----------------------------------------------------------------------------
func GetAllSalesRegion()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SalesRegion

	//----------------------------------------------------------------------------
	// Request the ORM to find all SalesRegion
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SalesRegion" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SalesRegion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSalesRegion", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSalesRegion - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSalesRegion(obj model.SalesRegion)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a SalesRegion using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SalesRegion using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSalesRegion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSalesRegion - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSalesRegion(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SalesRegion with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSalesRegion(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SalesRegion)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SalesRegion using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SalesRegion using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSalesRegion", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more operatorsIds as a Operators to a SalesRegion
//----------------------------------------------------------------------------
func AddOperatorsToSalesRegion ( salesRegionId uint64, operatorsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesRegion(salesRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesRegion)

		// slice the ids on comma with no spaces
		ids := strings.Split( operatorsIds, ",")

		for _, operatorsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Operator

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Operator
			// with a matching operatorsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , operatorsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Operators using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Operators").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Operators", operatorsId )
				return utils.RequestResult{false, msg, "unassignOperators", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesRegion from the gorm
		//----------------------------------------------------------------------------
		return GetSalesRegion(salesRegionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more operatorsIds as a Operators from a SalesRegion
//----------------------------------------------------------------------------
func RemoveOperatorsFromSalesRegion( salesRegionId uint64, operatorsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SalesRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesRegion(salesRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesRegion)

		// slice the ids on comma with no spaces
		ids := strings.Split( operatorsIds, ",")

		for _, operatorsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Operator

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Operator
			// with a matching operatorsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , operatorsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OperatorObj from the Operators array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Operators").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Operators", operatorsId )
				return utils.RequestResult{false, msg, "removeOperators", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesRegion from the gorm
		//----------------------------------------------------------------------------
		return GetSalesRegion(salesRegionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more salesCampaignsIds as a SalesCampaigns to a SalesRegion
//----------------------------------------------------------------------------
func AddSalesCampaignsToSalesRegion ( salesRegionId uint64, salesCampaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesRegion(salesRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesRegion)

		// slice the ids on comma with no spaces
		ids := strings.Split( salesCampaignsIds, ",")

		for _, salesCampaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalesCampaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalesCampaign
			// with a matching salesCampaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , salesCampaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SalesCampaigns using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SalesCampaigns").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesCampaigns", salesCampaignsId )
				return utils.RequestResult{false, msg, "unassignSalesCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesRegion from the gorm
		//----------------------------------------------------------------------------
		return GetSalesRegion(salesRegionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more salesCampaignsIds as a SalesCampaigns from a SalesRegion
//----------------------------------------------------------------------------
func RemoveSalesCampaignsFromSalesRegion( salesRegionId uint64, salesCampaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SalesRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesRegion(salesRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesRegion)

		// slice the ids on comma with no spaces
		ids := strings.Split( salesCampaignsIds, ",")

		for _, salesCampaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalesCampaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalesCampaign
			// with a matching salesCampaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , salesCampaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SalesCampaignObj from the SalesCampaigns array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SalesCampaigns").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesCampaigns", salesCampaignsId )
				return utils.RequestResult{false, msg, "removeSalesCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesRegion from the gorm
		//----------------------------------------------------------------------------
		return GetSalesRegion(salesRegionId)

	} else {
		return parentRequestResult
	}
}

