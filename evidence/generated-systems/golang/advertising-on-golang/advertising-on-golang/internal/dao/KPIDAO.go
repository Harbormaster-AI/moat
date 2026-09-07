package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing KPIDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateKPI - creates a new db entry
//----------------------------------------------------------------------------
func CreateKPI(obj model.KPI)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a KPI with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a KPI", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateKPI", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetKPI - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetKPI(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.KPI

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a KPI with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a KPI using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a KPI using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetKPI", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllKPI - returns all
//----------------------------------------------------------------------------
func GetAllKPI()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.KPI

	//----------------------------------------------------------------------------
	// Request the ORM to find all KPI
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all KPI" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all KPI", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllKPI", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateKPI - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateKPI(obj model.KPI)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a KPI using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a KPI using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateKPI", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteKPI - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteKPI(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the KPI with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetKPI(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KPI so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.KPI)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a KPI using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a KPI using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteKPI", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Campaign on a KPI
//----------------------------------------------------------------------------
func AssignCampaignToKPI( kPIId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the KPI with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKPI(kPIId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KPI so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KPI)

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
			// assign the Campaign	to the KPI
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the KPI
			//----------------------------------------------------------------------------
			return UpdateKPI(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a KPI
//----------------------------------------------------------------------------
func UnassignCampaignFromKPI(kPIId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the KPI with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetKPI(kPIId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.KPI so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.KPI)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the KPI
		//----------------------------------------------------------------------------
		return UpdateKPI(parentObj)

	} else {
		return parentRequestResult
	}

}


