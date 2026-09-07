package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SalesCampaignDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSalesCampaign - creates a new db entry
//----------------------------------------------------------------------------
func CreateSalesCampaign(obj model.SalesCampaign)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SalesCampaign with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SalesCampaign", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSalesCampaign", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSalesCampaign - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSalesCampaign(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SalesCampaign

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SalesCampaign with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SalesCampaign using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SalesCampaign using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSalesCampaign", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSalesCampaign - returns all
//----------------------------------------------------------------------------
func GetAllSalesCampaign()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SalesCampaign

	//----------------------------------------------------------------------------
	// Request the ORM to find all SalesCampaign
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SalesCampaign" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SalesCampaign", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSalesCampaign", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSalesCampaign - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSalesCampaign(obj model.SalesCampaign)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SalesCampaign using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SalesCampaign using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSalesCampaign", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSalesCampaign - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSalesCampaign(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SalesCampaign with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSalesCampaign(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesCampaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SalesCampaign)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SalesCampaign using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SalesCampaign using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSalesCampaign", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Region on a SalesCampaign
//----------------------------------------------------------------------------
func AssignRegionToSalesCampaign( salesCampaignId uint64, regionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SalesCampaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesCampaign(salesCampaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesCampaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesCampaign)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.SalesRegion

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a SalesRegion with a
		// matching regionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, regionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Region	to the SalesCampaign
			//----------------------------------------------------------------------------
			parentObj.Region = &childObj

			//----------------------------------------------------------------------------
			// save the SalesCampaign
			//----------------------------------------------------------------------------
			return UpdateSalesCampaign(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Region", regionId )
			return utils.RequestResult{false, msg, "assignRegion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Region on a SalesCampaign
//----------------------------------------------------------------------------
func UnassignRegionFromSalesCampaign(salesCampaignId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesCampaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesCampaign(salesCampaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesCampaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesCampaign)

		//----------------------------------------------------------------------------
		// assign an empty SalesRegion to the Region
		//----------------------------------------------------------------------------
		parentObj.Region = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Region
		//----------------------------------------------------------------------------
		parentObj.RegionId = nil;

		//----------------------------------------------------------------------------
		// save the SalesCampaign
		//----------------------------------------------------------------------------
		return UpdateSalesCampaign(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Operator on a SalesCampaign
//----------------------------------------------------------------------------
func AssignOperatorToSalesCampaign( salesCampaignId uint64, operatorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SalesCampaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesCampaign(salesCampaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesCampaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesCampaign)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Operator

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Operator with a
		// matching operatorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, operatorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Operator	to the SalesCampaign
			//----------------------------------------------------------------------------
			parentObj.Operator = &childObj

			//----------------------------------------------------------------------------
			// save the SalesCampaign
			//----------------------------------------------------------------------------
			return UpdateSalesCampaign(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Operator", operatorId )
			return utils.RequestResult{false, msg, "assignOperator", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Operator on a SalesCampaign
//----------------------------------------------------------------------------
func UnassignOperatorFromSalesCampaign(salesCampaignId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesCampaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesCampaign(salesCampaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesCampaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesCampaign)

		//----------------------------------------------------------------------------
		// assign an empty Operator to the Operator
		//----------------------------------------------------------------------------
		parentObj.Operator = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Operator
		//----------------------------------------------------------------------------
		parentObj.OperatorId = nil;

		//----------------------------------------------------------------------------
		// save the SalesCampaign
		//----------------------------------------------------------------------------
		return UpdateSalesCampaign(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more quotesIds as a Quotes to a SalesCampaign
//----------------------------------------------------------------------------
func AddQuotesToSalesCampaign ( salesCampaignId uint64, quotesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesCampaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesCampaign(salesCampaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesCampaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesCampaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Quotes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "unassignQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesCampaign from the gorm
		//----------------------------------------------------------------------------
		return GetSalesCampaign(salesCampaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more quotesIds as a Quotes from a SalesCampaign
//----------------------------------------------------------------------------
func RemoveQuotesFromSalesCampaign( salesCampaignId uint64, quotesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SalesCampaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesCampaign(salesCampaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesCampaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesCampaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QuoteObj from the Quotes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "removeQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesCampaign from the gorm
		//----------------------------------------------------------------------------
		return GetSalesCampaign(salesCampaignId)

	} else {
		return parentRequestResult
	}
}

