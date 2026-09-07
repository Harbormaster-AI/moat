package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PerformanceMetricDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePerformanceMetric - creates a new db entry
//----------------------------------------------------------------------------
func CreatePerformanceMetric(obj model.PerformanceMetric)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PerformanceMetric with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PerformanceMetric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePerformanceMetric", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPerformanceMetric - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPerformanceMetric(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PerformanceMetric

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PerformanceMetric with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PerformanceMetric using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PerformanceMetric using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPerformanceMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPerformanceMetric - returns all
//----------------------------------------------------------------------------
func GetAllPerformanceMetric()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PerformanceMetric

	//----------------------------------------------------------------------------
	// Request the ORM to find all PerformanceMetric
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PerformanceMetric" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PerformanceMetric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPerformanceMetric", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePerformanceMetric - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePerformanceMetric(obj model.PerformanceMetric)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PerformanceMetric using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PerformanceMetric using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePerformanceMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePerformanceMetric - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePerformanceMetric(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPerformanceMetric(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PerformanceMetric using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PerformanceMetric using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePerformanceMetric", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a AdAccount on a PerformanceMetric
//----------------------------------------------------------------------------
func AssignAdAccountToPerformanceMetric( performanceMetricId uint64, adAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AdAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AdAccount with a
		// matching adAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, adAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AdAccount	to the PerformanceMetric
			//----------------------------------------------------------------------------
			parentObj.AdAccount = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceMetric
			//----------------------------------------------------------------------------
			return UpdatePerformanceMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccount", adAccountId )
			return utils.RequestResult{false, msg, "assignAdAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AdAccount on a PerformanceMetric
//----------------------------------------------------------------------------
func UnassignAdAccountFromPerformanceMetric(performanceMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// assign an empty AdAccount to the AdAccount
		//----------------------------------------------------------------------------
		parentObj.AdAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AdAccount
		//----------------------------------------------------------------------------
		parentObj.AdAccountId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceMetric
		//----------------------------------------------------------------------------
		return UpdatePerformanceMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Campaign on a PerformanceMetric
//----------------------------------------------------------------------------
func AssignCampaignToPerformanceMetric( performanceMetricId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Campaign

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Campaign with a
		// matching campaignId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, campaignId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Campaign	to the PerformanceMetric
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceMetric
			//----------------------------------------------------------------------------
			return UpdatePerformanceMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a PerformanceMetric
//----------------------------------------------------------------------------
func UnassignCampaignFromPerformanceMetric(performanceMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceMetric
		//----------------------------------------------------------------------------
		return UpdatePerformanceMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LineItem on a PerformanceMetric
//----------------------------------------------------------------------------
func AssignLineItemToPerformanceMetric( performanceMetricId uint64, lineItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LineItem

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LineItem with a
		// matching lineItemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lineItemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LineItem	to the PerformanceMetric
			//----------------------------------------------------------------------------
			parentObj.LineItem = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceMetric
			//----------------------------------------------------------------------------
			return UpdatePerformanceMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItem", lineItemId )
			return utils.RequestResult{false, msg, "assignLineItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LineItem on a PerformanceMetric
//----------------------------------------------------------------------------
func UnassignLineItemFromPerformanceMetric(performanceMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// assign an empty LineItem to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItemId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceMetric
		//----------------------------------------------------------------------------
		return UpdatePerformanceMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Placement on a PerformanceMetric
//----------------------------------------------------------------------------
func AssignPlacementToPerformanceMetric( performanceMetricId uint64, placementId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Placement

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Placement with a
		// matching placementId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, placementId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Placement	to the PerformanceMetric
			//----------------------------------------------------------------------------
			parentObj.Placement = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceMetric
			//----------------------------------------------------------------------------
			return UpdatePerformanceMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Placement", placementId )
			return utils.RequestResult{false, msg, "assignPlacement", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Placement on a PerformanceMetric
//----------------------------------------------------------------------------
func UnassignPlacementFromPerformanceMetric(performanceMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// assign an empty Placement to the Placement
		//----------------------------------------------------------------------------
		parentObj.Placement = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Placement
		//----------------------------------------------------------------------------
		parentObj.PlacementId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceMetric
		//----------------------------------------------------------------------------
		return UpdatePerformanceMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CreativeAsset on a PerformanceMetric
//----------------------------------------------------------------------------
func AssignCreativeAssetToPerformanceMetric( performanceMetricId uint64, creativeAssetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CreativeAsset

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CreativeAsset with a
		// matching creativeAssetId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, creativeAssetId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CreativeAsset	to the PerformanceMetric
			//----------------------------------------------------------------------------
			parentObj.CreativeAsset = &childObj

			//----------------------------------------------------------------------------
			// save the PerformanceMetric
			//----------------------------------------------------------------------------
			return UpdatePerformanceMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CreativeAsset", creativeAssetId )
			return utils.RequestResult{false, msg, "assignCreativeAsset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CreativeAsset on a PerformanceMetric
//----------------------------------------------------------------------------
func UnassignCreativeAssetFromPerformanceMetric(performanceMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PerformanceMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPerformanceMetric(performanceMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PerformanceMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PerformanceMetric)

		//----------------------------------------------------------------------------
		// assign an empty CreativeAsset to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAsset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CreativeAsset
		//----------------------------------------------------------------------------
		parentObj.CreativeAssetId = nil;

		//----------------------------------------------------------------------------
		// save the PerformanceMetric
		//----------------------------------------------------------------------------
		return UpdatePerformanceMetric(parentObj)

	} else {
		return parentRequestResult
	}

}


