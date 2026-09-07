package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DiagnosisDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDiagnosis - creates a new db entry
//----------------------------------------------------------------------------
func CreateDiagnosis(obj model.Diagnosis)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Diagnosis with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Diagnosis", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDiagnosis", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDiagnosis - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDiagnosis(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Diagnosis

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Diagnosis with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Diagnosis using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Diagnosis using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDiagnosis", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDiagnosis - returns all
//----------------------------------------------------------------------------
func GetAllDiagnosis()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Diagnosis

	//----------------------------------------------------------------------------
	// Request the ORM to find all Diagnosis
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Diagnosis" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Diagnosis", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDiagnosis", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDiagnosis - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDiagnosis(obj model.Diagnosis)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Diagnosis using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Diagnosis using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDiagnosis", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDiagnosis - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDiagnosis(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Diagnosis with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDiagnosis(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Diagnosis so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Diagnosis)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Diagnosis using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Diagnosis using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDiagnosis", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Encounter on a Diagnosis
//----------------------------------------------------------------------------
func AssignEncounterToDiagnosis( diagnosisId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Diagnosis with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDiagnosis(diagnosisId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Diagnosis so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Diagnosis)

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
			// assign the Encounter	to the Diagnosis
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the Diagnosis
			//----------------------------------------------------------------------------
			return UpdateDiagnosis(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a Diagnosis
//----------------------------------------------------------------------------
func UnassignEncounterFromDiagnosis(diagnosisId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Diagnosis with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDiagnosis(diagnosisId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Diagnosis so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Diagnosis)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the Diagnosis
		//----------------------------------------------------------------------------
		return UpdateDiagnosis(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Patient on a Diagnosis
//----------------------------------------------------------------------------
func AssignPatientToDiagnosis( diagnosisId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Diagnosis with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDiagnosis(diagnosisId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Diagnosis so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Diagnosis)

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
			// assign the Patient	to the Diagnosis
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Diagnosis
			//----------------------------------------------------------------------------
			return UpdateDiagnosis(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Diagnosis
//----------------------------------------------------------------------------
func UnassignPatientFromDiagnosis(diagnosisId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Diagnosis with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDiagnosis(diagnosisId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Diagnosis so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Diagnosis)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Diagnosis
		//----------------------------------------------------------------------------
		return UpdateDiagnosis(parentObj)

	} else {
		return parentRequestResult
	}

}


