package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AdmissionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAdmission - creates a new db entry
//----------------------------------------------------------------------------
func CreateAdmission(obj model.Admission)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Admission with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Admission", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAdmission", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAdmission - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAdmission(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Admission

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Admission with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Admission using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Admission using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAdmission", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAdmission - returns all
//----------------------------------------------------------------------------
func GetAllAdmission()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Admission

	//----------------------------------------------------------------------------
	// Request the ORM to find all Admission
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Admission" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Admission", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAdmission", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAdmission - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAdmission(obj model.Admission)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Admission using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Admission using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAdmission", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAdmission - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAdmission(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Admission with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAdmission(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Admission so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Admission)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Admission using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Admission using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAdmission", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Encounter on a Admission
//----------------------------------------------------------------------------
func AssignEncounterToAdmission( admissionId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Admission with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdmission(admissionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Admission so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Admission)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Encounter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Encounter with a
		// matching encounterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, encounterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Encounter	to the Admission
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the Admission
			//----------------------------------------------------------------------------
			return UpdateAdmission(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a Admission
//----------------------------------------------------------------------------
func UnassignEncounterFromAdmission(admissionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Admission with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdmission(admissionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Admission so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Admission)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the Admission
		//----------------------------------------------------------------------------
		return UpdateAdmission(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Facility on a Admission
//----------------------------------------------------------------------------
func AssignFacilityToAdmission( admissionId uint64, facilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Admission with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdmission(admissionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Admission so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Admission)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Facility

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Facility with a
		// matching facilityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, facilityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Facility	to the Admission
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the Admission
			//----------------------------------------------------------------------------
			return UpdateAdmission(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a Admission
//----------------------------------------------------------------------------
func UnassignFacilityFromAdmission(admissionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Admission with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdmission(admissionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Admission so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Admission)

		//----------------------------------------------------------------------------
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the Admission
		//----------------------------------------------------------------------------
		return UpdateAdmission(parentObj)

	} else {
		return parentRequestResult
	}

}


