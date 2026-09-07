package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CampaignDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCampaign - creates a new db entry
//----------------------------------------------------------------------------
func CreateCampaign(obj model.Campaign)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Campaign with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Campaign", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCampaign", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCampaign - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCampaign(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Campaign

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Campaign with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Campaign using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Campaign using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCampaign", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCampaign - returns all
//----------------------------------------------------------------------------
func GetAllCampaign()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Campaign

	//----------------------------------------------------------------------------
	// Request the ORM to find all Campaign
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Campaign" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Campaign", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCampaign", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCampaign - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCampaign(obj model.Campaign)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Campaign using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Campaign using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCampaign", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCampaign - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCampaign(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCampaign(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Campaign using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Campaign using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCampaign", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a AdAccount on a Campaign
//----------------------------------------------------------------------------
func AssignAdAccountToCampaign( campaignId uint64, adAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

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
			// assign the AdAccount	to the Campaign
			//----------------------------------------------------------------------------
			parentObj.AdAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Campaign
			//----------------------------------------------------------------------------
			return UpdateCampaign(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccount", adAccountId )
			return utils.RequestResult{false, msg, "assignAdAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AdAccount on a Campaign
//----------------------------------------------------------------------------
func UnassignAdAccountFromCampaign(campaignId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// assign an empty AdAccount to the AdAccount
		//----------------------------------------------------------------------------
		parentObj.AdAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AdAccount
		//----------------------------------------------------------------------------
		parentObj.AdAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Campaign
		//----------------------------------------------------------------------------
		return UpdateCampaign(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InsertionOrder on a Campaign
//----------------------------------------------------------------------------
func AssignInsertionOrderToCampaign( campaignId uint64, insertionOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsertionOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsertionOrder with a
		// matching insertionOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, insertionOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InsertionOrder	to the Campaign
			//----------------------------------------------------------------------------
			parentObj.InsertionOrder = &childObj

			//----------------------------------------------------------------------------
			// save the Campaign
			//----------------------------------------------------------------------------
			return UpdateCampaign(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsertionOrder", insertionOrderId )
			return utils.RequestResult{false, msg, "assignInsertionOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InsertionOrder on a Campaign
//----------------------------------------------------------------------------
func UnassignInsertionOrderFromCampaign(campaignId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// assign an empty InsertionOrder to the InsertionOrder
		//----------------------------------------------------------------------------
		parentObj.InsertionOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InsertionOrder
		//----------------------------------------------------------------------------
		parentObj.InsertionOrderId = nil;

		//----------------------------------------------------------------------------
		// save the Campaign
		//----------------------------------------------------------------------------
		return UpdateCampaign(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more lineItemsIds as a LineItems to a Campaign
//----------------------------------------------------------------------------
func AddLineItemsToCampaign ( campaignId uint64, lineItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineItem
			// with a matching lineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LineItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItems", lineItemsId )
				return utils.RequestResult{false, msg, "unassignLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more lineItemsIds as a LineItems from a Campaign
//----------------------------------------------------------------------------
func RemoveLineItemsFromCampaign( campaignId uint64, lineItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineItem
			// with a matching lineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LineItemObj from the LineItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItems", lineItemsId )
				return utils.RequestResult{false, msg, "removeLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more kpisIds as a Kpis to a Campaign
//----------------------------------------------------------------------------
func AddKpisToCampaign ( campaignId uint64, kpisIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( kpisIds, ",")

		for _, kpisId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KPI

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KPI
			// with a matching kpisId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , kpisId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Kpis using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Kpis").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Kpis", kpisId )
				return utils.RequestResult{false, msg, "unassignKpis", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more kpisIds as a Kpis from a Campaign
//----------------------------------------------------------------------------
func RemoveKpisFromCampaign( campaignId uint64, kpisIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( kpisIds, ",")

		for _, kpisId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KPI

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KPI
			// with a matching kpisId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , kpisId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove KPIObj from the Kpis array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Kpis").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Kpis", kpisId )
				return utils.RequestResult{false, msg, "removeKpis", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more trackingPixelsIds as a TrackingPixels to a Campaign
//----------------------------------------------------------------------------
func AddTrackingPixelsToCampaign ( campaignId uint64, trackingPixelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( trackingPixelsIds, ",")

		for _, trackingPixelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrackingPixel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrackingPixel
			// with a matching trackingPixelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trackingPixelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the TrackingPixels using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrackingPixels").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrackingPixels", trackingPixelsId )
				return utils.RequestResult{false, msg, "unassignTrackingPixels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more trackingPixelsIds as a TrackingPixels from a Campaign
//----------------------------------------------------------------------------
func RemoveTrackingPixelsFromCampaign( campaignId uint64, trackingPixelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( trackingPixelsIds, ",")

		for _, trackingPixelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrackingPixel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrackingPixel
			// with a matching trackingPixelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trackingPixelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TrackingPixelObj from the TrackingPixels array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrackingPixels").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrackingPixels", trackingPixelsId )
				return utils.RequestResult{false, msg, "removeTrackingPixels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more audiencesIds as a Audiences to a Campaign
//----------------------------------------------------------------------------
func AddAudiencesToCampaign ( campaignId uint64, audiencesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( audiencesIds, ",")

		for _, audiencesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AudienceSegment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AudienceSegment
			// with a matching audiencesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , audiencesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Audiences using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Audiences").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Audiences", audiencesId )
				return utils.RequestResult{false, msg, "unassignAudiences", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more audiencesIds as a Audiences from a Campaign
//----------------------------------------------------------------------------
func RemoveAudiencesFromCampaign( campaignId uint64, audiencesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( audiencesIds, ",")

		for _, audiencesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AudienceSegment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AudienceSegment
			// with a matching audiencesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , audiencesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AudienceSegmentObj from the Audiences array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Audiences").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Audiences", audiencesId )
				return utils.RequestResult{false, msg, "removeAudiences", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a Campaign
//----------------------------------------------------------------------------
func AddReportsToCampaign ( campaignId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "unassignReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a Campaign
//----------------------------------------------------------------------------
func RemoveReportsFromCampaign( campaignId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReportObj from the Reports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "removeReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

