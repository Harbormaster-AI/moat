package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LineItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLineItem - creates a new db entry
//----------------------------------------------------------------------------
func CreateLineItem(obj model.LineItem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LineItem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LineItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLineItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLineItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLineItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LineItem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LineItem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LineItem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LineItem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLineItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLineItem - returns all
//----------------------------------------------------------------------------
func GetAllLineItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LineItem

	//----------------------------------------------------------------------------
	// Request the ORM to find all LineItem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LineItem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LineItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLineItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLineItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLineItem(obj model.LineItem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LineItem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LineItem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLineItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLineItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLineItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLineItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LineItem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LineItem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LineItem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLineItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Campaign on a LineItem
//----------------------------------------------------------------------------
func AssignCampaignToLineItem( lineItemId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

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
			// assign the Campaign	to the LineItem
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the LineItem
			//----------------------------------------------------------------------------
			return UpdateLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a LineItem
//----------------------------------------------------------------------------
func UnassignCampaignFromLineItem(lineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the LineItem
		//----------------------------------------------------------------------------
		return UpdateLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a TargetingProfile on a LineItem
//----------------------------------------------------------------------------
func AssignTargetingProfileToLineItem( lineItemId uint64, targetingProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.TargetingProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TargetingProfile with a
		// matching targetingProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, targetingProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the TargetingProfile	to the LineItem
			//----------------------------------------------------------------------------
			parentObj.TargetingProfile = &childObj

			//----------------------------------------------------------------------------
			// save the LineItem
			//----------------------------------------------------------------------------
			return UpdateLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TargetingProfile", targetingProfileId )
			return utils.RequestResult{false, msg, "assignTargetingProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TargetingProfile on a LineItem
//----------------------------------------------------------------------------
func UnassignTargetingProfileFromLineItem(lineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		//----------------------------------------------------------------------------
		// assign an empty TargetingProfile to the TargetingProfile
		//----------------------------------------------------------------------------
		parentObj.TargetingProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TargetingProfile
		//----------------------------------------------------------------------------
		parentObj.TargetingProfileId = nil;

		//----------------------------------------------------------------------------
		// save the LineItem
		//----------------------------------------------------------------------------
		return UpdateLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Deal on a LineItem
//----------------------------------------------------------------------------
func AssignDealToLineItem( lineItemId uint64, dealId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Deal

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Deal with a
		// matching dealId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, dealId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Deal	to the LineItem
			//----------------------------------------------------------------------------
			parentObj.Deal = &childObj

			//----------------------------------------------------------------------------
			// save the LineItem
			//----------------------------------------------------------------------------
			return UpdateLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deal", dealId )
			return utils.RequestResult{false, msg, "assignDeal", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Deal on a LineItem
//----------------------------------------------------------------------------
func UnassignDealFromLineItem(lineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		//----------------------------------------------------------------------------
		// assign an empty Deal to the Deal
		//----------------------------------------------------------------------------
		parentObj.Deal = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Deal
		//----------------------------------------------------------------------------
		parentObj.DealId = nil;

		//----------------------------------------------------------------------------
		// save the LineItem
		//----------------------------------------------------------------------------
		return UpdateLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more placementsIds as a Placements to a LineItem
//----------------------------------------------------------------------------
func AddPlacementsToLineItem ( lineItemId uint64, placementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		// slice the ids on comma with no spaces
		ids := strings.Split( placementsIds, ",")

		for _, placementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Placement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Placement
			// with a matching placementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , placementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Placements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Placements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Placements", placementsId )
				return utils.RequestResult{false, msg, "unassignPlacements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineItem from the gorm
		//----------------------------------------------------------------------------
		return GetLineItem(lineItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more placementsIds as a Placements from a LineItem
//----------------------------------------------------------------------------
func RemovePlacementsFromLineItem( lineItemId uint64, placementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		// slice the ids on comma with no spaces
		ids := strings.Split( placementsIds, ",")

		for _, placementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Placement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Placement
			// with a matching placementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , placementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PlacementObj from the Placements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Placements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Placements", placementsId )
				return utils.RequestResult{false, msg, "removePlacements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineItem from the gorm
		//----------------------------------------------------------------------------
		return GetLineItem(lineItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more creativesIds as a Creatives to a LineItem
//----------------------------------------------------------------------------
func AddCreativesToLineItem ( lineItemId uint64, creativesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		// slice the ids on comma with no spaces
		ids := strings.Split( creativesIds, ",")

		for _, creativesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeAsset

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeAsset
			// with a matching creativesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , creativesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Creatives using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Creatives").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Creatives", creativesId )
				return utils.RequestResult{false, msg, "unassignCreatives", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineItem from the gorm
		//----------------------------------------------------------------------------
		return GetLineItem(lineItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more creativesIds as a Creatives from a LineItem
//----------------------------------------------------------------------------
func RemoveCreativesFromLineItem( lineItemId uint64, creativesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		// slice the ids on comma with no spaces
		ids := strings.Split( creativesIds, ",")

		for _, creativesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeAsset

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeAsset
			// with a matching creativesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , creativesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CreativeAssetObj from the Creatives array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Creatives").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Creatives", creativesId )
				return utils.RequestResult{false, msg, "removeCreatives", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineItem from the gorm
		//----------------------------------------------------------------------------
		return GetLineItem(lineItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more performanceMetricsIds as a PerformanceMetrics to a LineItem
//----------------------------------------------------------------------------
func AddPerformanceMetricsToLineItem ( lineItemId uint64, performanceMetricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		// slice the ids on comma with no spaces
		ids := strings.Split( performanceMetricsIds, ",")

		for _, performanceMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceMetric
			// with a matching performanceMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , performanceMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PerformanceMetrics using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PerformanceMetrics").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PerformanceMetrics", performanceMetricsId )
				return utils.RequestResult{false, msg, "unassignPerformanceMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineItem from the gorm
		//----------------------------------------------------------------------------
		return GetLineItem(lineItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more performanceMetricsIds as a PerformanceMetrics from a LineItem
//----------------------------------------------------------------------------
func RemovePerformanceMetricsFromLineItem( lineItemId uint64, performanceMetricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineItem(lineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineItem)

		// slice the ids on comma with no spaces
		ids := strings.Split( performanceMetricsIds, ",")

		for _, performanceMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceMetric
			// with a matching performanceMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , performanceMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PerformanceMetricObj from the PerformanceMetrics array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PerformanceMetrics").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PerformanceMetrics", performanceMetricsId )
				return utils.RequestResult{false, msg, "removePerformanceMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineItem from the gorm
		//----------------------------------------------------------------------------
		return GetLineItem(lineItemId)

	} else {
		return parentRequestResult
	}
}

