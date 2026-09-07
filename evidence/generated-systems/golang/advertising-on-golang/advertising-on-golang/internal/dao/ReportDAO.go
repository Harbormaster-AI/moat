package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ReportDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateReport - creates a new db entry
//----------------------------------------------------------------------------
func CreateReport(obj model.Report)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Report with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Report", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateReport", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetReport - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetReport(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Report

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Report with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Report using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Report using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetReport", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllReport - returns all
//----------------------------------------------------------------------------
func GetAllReport()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Report

	//----------------------------------------------------------------------------
	// Request the ORM to find all Report
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Report" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Report", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllReport", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateReport - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateReport(obj model.Report)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Report using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Report using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateReport", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteReport - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteReport(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetReport(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Report)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Report using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Report using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteReport", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a AdAccount on a Report
//----------------------------------------------------------------------------
func AssignAdAccountToReport( reportId uint64, adAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
			// assign the AdAccount	to the Report
			//----------------------------------------------------------------------------
			parentObj.AdAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Report
			//----------------------------------------------------------------------------
			return UpdateReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccount", adAccountId )
			return utils.RequestResult{false, msg, "assignAdAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AdAccount on a Report
//----------------------------------------------------------------------------
func UnassignAdAccountFromReport(reportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

		//----------------------------------------------------------------------------
		// assign an empty AdAccount to the AdAccount
		//----------------------------------------------------------------------------
		parentObj.AdAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AdAccount
		//----------------------------------------------------------------------------
		parentObj.AdAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Report
		//----------------------------------------------------------------------------
		return UpdateReport(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Campaign on a Report
//----------------------------------------------------------------------------
func AssignCampaignToReport( reportId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
			// assign the Campaign	to the Report
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the Report
			//----------------------------------------------------------------------------
			return UpdateReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a Report
//----------------------------------------------------------------------------
func UnassignCampaignFromReport(reportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the Report
		//----------------------------------------------------------------------------
		return UpdateReport(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LineItem on a Report
//----------------------------------------------------------------------------
func AssignLineItemToReport( reportId uint64, lineItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
			// assign the LineItem	to the Report
			//----------------------------------------------------------------------------
			parentObj.LineItem = &childObj

			//----------------------------------------------------------------------------
			// save the Report
			//----------------------------------------------------------------------------
			return UpdateReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItem", lineItemId )
			return utils.RequestResult{false, msg, "assignLineItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LineItem on a Report
//----------------------------------------------------------------------------
func UnassignLineItemFromReport(reportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

		//----------------------------------------------------------------------------
		// assign an empty LineItem to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItemId = nil;

		//----------------------------------------------------------------------------
		// save the Report
		//----------------------------------------------------------------------------
		return UpdateReport(parentObj)

	} else {
		return parentRequestResult
	}

}


