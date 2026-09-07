package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AllergyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAllergy - creates a new db entry
//----------------------------------------------------------------------------
func CreateAllergy(obj model.Allergy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Allergy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Allergy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAllergy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAllergy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAllergy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Allergy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Allergy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Allergy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Allergy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAllergy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAllergy - returns all
//----------------------------------------------------------------------------
func GetAllAllergy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Allergy

	//----------------------------------------------------------------------------
	// Request the ORM to find all Allergy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Allergy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Allergy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAllergy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAllergy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAllergy(obj model.Allergy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Allergy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Allergy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAllergy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAllergy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAllergy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Allergy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAllergy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Allergy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Allergy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Allergy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Allergy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAllergy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a Allergy
//----------------------------------------------------------------------------
func AssignPatientToAllergy( allergyId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Allergy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAllergy(allergyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Allergy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Allergy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Patient

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Patient with a
		// matching patientId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, patientId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Patient	to the Allergy
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Allergy
			//----------------------------------------------------------------------------
			return UpdateAllergy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Allergy
//----------------------------------------------------------------------------
func UnassignPatientFromAllergy(allergyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Allergy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAllergy(allergyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Allergy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Allergy)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Allergy
		//----------------------------------------------------------------------------
		return UpdateAllergy(parentObj)

	} else {
		return parentRequestResult
	}

}


