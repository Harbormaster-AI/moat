package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MedicationDispenseDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMedicationDispense - creates a new db entry
//----------------------------------------------------------------------------
func CreateMedicationDispense(obj model.MedicationDispense)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MedicationDispense with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MedicationDispense", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMedicationDispense", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMedicationDispense - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMedicationDispense(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MedicationDispense

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MedicationDispense with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MedicationDispense using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MedicationDispense using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMedicationDispense", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMedicationDispense - returns all
//----------------------------------------------------------------------------
func GetAllMedicationDispense()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MedicationDispense

	//----------------------------------------------------------------------------
	// Request the ORM to find all MedicationDispense
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MedicationDispense" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MedicationDispense", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMedicationDispense", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMedicationDispense - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMedicationDispense(obj model.MedicationDispense)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MedicationDispense using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MedicationDispense using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMedicationDispense", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMedicationDispense - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMedicationDispense(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MedicationDispense with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMedicationDispense(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationDispense so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MedicationDispense)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MedicationDispense using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MedicationDispense using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMedicationDispense", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a MedicationOrder on a MedicationDispense
//----------------------------------------------------------------------------
func AssignMedicationOrderToMedicationDispense( medicationDispenseId uint64, medicationOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MedicationDispense with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationDispense(medicationDispenseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationDispense so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationDispense)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.MedicationOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a MedicationOrder with a
		// matching medicationOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, medicationOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the MedicationOrder	to the MedicationDispense
			//----------------------------------------------------------------------------
			parentObj.MedicationOrder = &childObj

			//----------------------------------------------------------------------------
			// save the MedicationDispense
			//----------------------------------------------------------------------------
			return UpdateMedicationDispense(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MedicationOrder", medicationOrderId )
			return utils.RequestResult{false, msg, "assignMedicationOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a MedicationOrder on a MedicationDispense
//----------------------------------------------------------------------------
func UnassignMedicationOrderFromMedicationDispense(medicationDispenseId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicationDispense with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationDispense(medicationDispenseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationDispense so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationDispense)

		//----------------------------------------------------------------------------
		// assign an empty MedicationOrder to the MedicationOrder
		//----------------------------------------------------------------------------
		parentObj.MedicationOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the MedicationOrder
		//----------------------------------------------------------------------------
		parentObj.MedicationOrderId = nil;

		//----------------------------------------------------------------------------
		// save the MedicationDispense
		//----------------------------------------------------------------------------
		return UpdateMedicationDispense(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Pharmacy on a MedicationDispense
//----------------------------------------------------------------------------
func AssignPharmacyToMedicationDispense( medicationDispenseId uint64, pharmacyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MedicationDispense with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationDispense(medicationDispenseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationDispense so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationDispense)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Pharmacy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Pharmacy with a
		// matching pharmacyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, pharmacyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Pharmacy	to the MedicationDispense
			//----------------------------------------------------------------------------
			parentObj.Pharmacy = &childObj

			//----------------------------------------------------------------------------
			// save the MedicationDispense
			//----------------------------------------------------------------------------
			return UpdateMedicationDispense(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pharmacy", pharmacyId )
			return utils.RequestResult{false, msg, "assignPharmacy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Pharmacy on a MedicationDispense
//----------------------------------------------------------------------------
func UnassignPharmacyFromMedicationDispense(medicationDispenseId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicationDispense with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationDispense(medicationDispenseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationDispense so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationDispense)

		//----------------------------------------------------------------------------
		// assign an empty Pharmacy to the Pharmacy
		//----------------------------------------------------------------------------
		parentObj.Pharmacy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Pharmacy
		//----------------------------------------------------------------------------
		parentObj.PharmacyId = nil;

		//----------------------------------------------------------------------------
		// save the MedicationDispense
		//----------------------------------------------------------------------------
		return UpdateMedicationDispense(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Patient on a MedicationDispense
//----------------------------------------------------------------------------
func AssignPatientToMedicationDispense( medicationDispenseId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MedicationDispense with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationDispense(medicationDispenseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationDispense so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationDispense)

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
			// assign the Patient	to the MedicationDispense
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the MedicationDispense
			//----------------------------------------------------------------------------
			return UpdateMedicationDispense(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a MedicationDispense
//----------------------------------------------------------------------------
func UnassignPatientFromMedicationDispense(medicationDispenseId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicationDispense with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicationDispense(medicationDispenseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicationDispense so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicationDispense)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the MedicationDispense
		//----------------------------------------------------------------------------
		return UpdateMedicationDispense(parentObj)

	} else {
		return parentRequestResult
	}

}


