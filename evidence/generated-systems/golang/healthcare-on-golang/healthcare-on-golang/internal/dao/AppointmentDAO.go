package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AppointmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAppointment - creates a new db entry
//----------------------------------------------------------------------------
func CreateAppointment(obj model.Appointment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Appointment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Appointment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAppointment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAppointment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAppointment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Appointment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Appointment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Appointment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Appointment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAppointment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAppointment - returns all
//----------------------------------------------------------------------------
func GetAllAppointment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Appointment

	//----------------------------------------------------------------------------
	// Request the ORM to find all Appointment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Appointment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Appointment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAppointment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAppointment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAppointment(obj model.Appointment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Appointment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Appointment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAppointment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAppointment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAppointment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAppointment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Appointment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Appointment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Appointment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAppointment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a Appointment
//----------------------------------------------------------------------------
func AssignPatientToAppointment( appointmentId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

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
			// assign the Patient	to the Appointment
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Appointment
			//----------------------------------------------------------------------------
			return UpdateAppointment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Appointment
//----------------------------------------------------------------------------
func UnassignPatientFromAppointment(appointmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Appointment
		//----------------------------------------------------------------------------
		return UpdateAppointment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Clinician on a Appointment
//----------------------------------------------------------------------------
func AssignClinicianToAppointment( appointmentId uint64, clinicianId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Clinician

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Clinician with a
		// matching clinicianId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, clinicianId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Clinician	to the Appointment
			//----------------------------------------------------------------------------
			parentObj.Clinician = &childObj

			//----------------------------------------------------------------------------
			// save the Appointment
			//----------------------------------------------------------------------------
			return UpdateAppointment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Clinician", clinicianId )
			return utils.RequestResult{false, msg, "assignClinician", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Clinician on a Appointment
//----------------------------------------------------------------------------
func UnassignClinicianFromAppointment(appointmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

		//----------------------------------------------------------------------------
		// assign an empty Clinician to the Clinician
		//----------------------------------------------------------------------------
		parentObj.Clinician = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Clinician
		//----------------------------------------------------------------------------
		parentObj.ClinicianId = nil;

		//----------------------------------------------------------------------------
		// save the Appointment
		//----------------------------------------------------------------------------
		return UpdateAppointment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Facility on a Appointment
//----------------------------------------------------------------------------
func AssignFacilityToAppointment( appointmentId uint64, facilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

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
			// assign the Facility	to the Appointment
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the Appointment
			//----------------------------------------------------------------------------
			return UpdateAppointment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a Appointment
//----------------------------------------------------------------------------
func UnassignFacilityFromAppointment(appointmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

		//----------------------------------------------------------------------------
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the Appointment
		//----------------------------------------------------------------------------
		return UpdateAppointment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Encounter on a Appointment
//----------------------------------------------------------------------------
func AssignEncounterToAppointment( appointmentId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

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
			// assign the Encounter	to the Appointment
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the Appointment
			//----------------------------------------------------------------------------
			return UpdateAppointment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a Appointment
//----------------------------------------------------------------------------
func UnassignEncounterFromAppointment(appointmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Appointment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppointment(appointmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Appointment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Appointment)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the Appointment
		//----------------------------------------------------------------------------
		return UpdateAppointment(parentObj)

	} else {
		return parentRequestResult
	}

}


