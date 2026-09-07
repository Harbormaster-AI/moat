package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ScreeningDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateScreening - creates a new db entry
//----------------------------------------------------------------------------
func CreateScreening(obj model.Screening)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Screening with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Screening", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateScreening", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetScreening - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetScreening(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Screening

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Screening with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Screening using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Screening using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetScreening", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllScreening - returns all
//----------------------------------------------------------------------------
func GetAllScreening()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Screening

	//----------------------------------------------------------------------------
	// Request the ORM to find all Screening
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Screening" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Screening", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllScreening", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateScreening - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateScreening(obj model.Screening)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Screening using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Screening using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateScreening", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteScreening - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteScreening(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetScreening(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Screening)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Screening using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Screening using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteScreening", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a KycProfile on a Screening
//----------------------------------------------------------------------------
func AssignKycProfileToScreening( screeningId uint64, kycProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreening(screeningId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Screening)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.KYCProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a KYCProfile with a
		// matching kycProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, kycProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the KycProfile	to the Screening
			//----------------------------------------------------------------------------
			parentObj.KycProfile = &childObj

			//----------------------------------------------------------------------------
			// save the Screening
			//----------------------------------------------------------------------------
			return UpdateScreening(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfile", kycProfileId )
			return utils.RequestResult{false, msg, "assignKycProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a KycProfile on a Screening
//----------------------------------------------------------------------------
func UnassignKycProfileFromScreening(screeningId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreening(screeningId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Screening)

		//----------------------------------------------------------------------------
		// assign an empty KYCProfile to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfileId = nil;

		//----------------------------------------------------------------------------
		// save the Screening
		//----------------------------------------------------------------------------
		return UpdateScreening(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more alertsIds as a Alerts to a Screening
//----------------------------------------------------------------------------
func AddAlertsToScreening ( screeningId uint64, alertsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreening(screeningId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Screening)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceAlert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceAlert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Alerts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "unassignAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Screening from the gorm
		//----------------------------------------------------------------------------
		return GetScreening(screeningId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more alertsIds as a Alerts from a Screening
//----------------------------------------------------------------------------
func RemoveAlertsFromScreening( screeningId uint64, alertsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Screening with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScreening(screeningId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Screening so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Screening)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceAlert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceAlert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ComplianceAlertObj from the Alerts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "removeAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Screening from the gorm
		//----------------------------------------------------------------------------
		return GetScreening(screeningId)

	} else {
		return parentRequestResult
	}
}

