package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EncounterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEncounter - creates a new db entry
//----------------------------------------------------------------------------
func CreateEncounter(obj model.Encounter)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Encounter with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Encounter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEncounter", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEncounter - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEncounter(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Encounter

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Encounter with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Encounter using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Encounter using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEncounter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEncounter - returns all
//----------------------------------------------------------------------------
func GetAllEncounter()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Encounter

	//----------------------------------------------------------------------------
	// Request the ORM to find all Encounter
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Encounter" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Encounter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEncounter", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEncounter - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEncounter(obj model.Encounter)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Encounter using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Encounter using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEncounter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEncounter - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEncounter(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEncounter(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Encounter using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Encounter using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEncounter", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a Encounter
//----------------------------------------------------------------------------
func AssignPatientToEncounter( encounterId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

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
			// assign the Patient	to the Encounter
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Encounter
			//----------------------------------------------------------------------------
			return UpdateEncounter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Encounter
//----------------------------------------------------------------------------
func UnassignPatientFromEncounter(encounterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Encounter
		//----------------------------------------------------------------------------
		return UpdateEncounter(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Clinician on a Encounter
//----------------------------------------------------------------------------
func AssignClinicianToEncounter( encounterId uint64, clinicianId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

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
			// assign the Clinician	to the Encounter
			//----------------------------------------------------------------------------
			parentObj.Clinician = &childObj

			//----------------------------------------------------------------------------
			// save the Encounter
			//----------------------------------------------------------------------------
			return UpdateEncounter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Clinician", clinicianId )
			return utils.RequestResult{false, msg, "assignClinician", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Clinician on a Encounter
//----------------------------------------------------------------------------
func UnassignClinicianFromEncounter(encounterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// assign an empty Clinician to the Clinician
		//----------------------------------------------------------------------------
		parentObj.Clinician = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Clinician
		//----------------------------------------------------------------------------
		parentObj.ClinicianId = nil;

		//----------------------------------------------------------------------------
		// save the Encounter
		//----------------------------------------------------------------------------
		return UpdateEncounter(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Facility on a Encounter
//----------------------------------------------------------------------------
func AssignFacilityToEncounter( encounterId uint64, facilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

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
			// assign the Facility	to the Encounter
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the Encounter
			//----------------------------------------------------------------------------
			return UpdateEncounter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a Encounter
//----------------------------------------------------------------------------
func UnassignFacilityFromEncounter(encounterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the Encounter
		//----------------------------------------------------------------------------
		return UpdateEncounter(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Appointment on a Encounter
//----------------------------------------------------------------------------
func AssignAppointmentToEncounter( encounterId uint64, appointmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Appointment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Appointment with a
		// matching appointmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, appointmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Appointment	to the Encounter
			//----------------------------------------------------------------------------
			parentObj.Appointment = &childObj

			//----------------------------------------------------------------------------
			// save the Encounter
			//----------------------------------------------------------------------------
			return UpdateEncounter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Appointment", appointmentId )
			return utils.RequestResult{false, msg, "assignAppointment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Appointment on a Encounter
//----------------------------------------------------------------------------
func UnassignAppointmentFromEncounter(encounterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// assign an empty Appointment to the Appointment
		//----------------------------------------------------------------------------
		parentObj.Appointment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Appointment
		//----------------------------------------------------------------------------
		parentObj.AppointmentId = nil;

		//----------------------------------------------------------------------------
		// save the Encounter
		//----------------------------------------------------------------------------
		return UpdateEncounter(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Admission on a Encounter
//----------------------------------------------------------------------------
func AssignAdmissionToEncounter( encounterId uint64, admissionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Admission

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Admission with a
		// matching admissionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, admissionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Admission	to the Encounter
			//----------------------------------------------------------------------------
			parentObj.Admission = &childObj

			//----------------------------------------------------------------------------
			// save the Encounter
			//----------------------------------------------------------------------------
			return UpdateEncounter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Admission", admissionId )
			return utils.RequestResult{false, msg, "assignAdmission", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Admission on a Encounter
//----------------------------------------------------------------------------
func UnassignAdmissionFromEncounter(encounterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// assign an empty Admission to the Admission
		//----------------------------------------------------------------------------
		parentObj.Admission = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Admission
		//----------------------------------------------------------------------------
		parentObj.AdmissionId = nil;

		//----------------------------------------------------------------------------
		// save the Encounter
		//----------------------------------------------------------------------------
		return UpdateEncounter(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Discharge on a Encounter
//----------------------------------------------------------------------------
func AssignDischargeToEncounter( encounterId uint64, dischargeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Discharge

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Discharge with a
		// matching dischargeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, dischargeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Discharge	to the Encounter
			//----------------------------------------------------------------------------
			parentObj.Discharge = &childObj

			//----------------------------------------------------------------------------
			// save the Encounter
			//----------------------------------------------------------------------------
			return UpdateEncounter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Discharge", dischargeId )
			return utils.RequestResult{false, msg, "assignDischarge", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Discharge on a Encounter
//----------------------------------------------------------------------------
func UnassignDischargeFromEncounter(encounterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		//----------------------------------------------------------------------------
		// assign an empty Discharge to the Discharge
		//----------------------------------------------------------------------------
		parentObj.Discharge = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Discharge
		//----------------------------------------------------------------------------
		parentObj.DischargeId = nil;

		//----------------------------------------------------------------------------
		// save the Encounter
		//----------------------------------------------------------------------------
		return UpdateEncounter(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more diagnosesIds as a Diagnoses to a Encounter
//----------------------------------------------------------------------------
func AddDiagnosesToEncounter ( encounterId uint64, diagnosesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( diagnosesIds, ",")

		for _, diagnosesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Diagnosis

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Diagnosis
			// with a matching diagnosesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , diagnosesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Diagnoses using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Diagnoses").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Diagnoses", diagnosesId )
				return utils.RequestResult{false, msg, "unassignDiagnoses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more diagnosesIds as a Diagnoses from a Encounter
//----------------------------------------------------------------------------
func RemoveDiagnosesFromEncounter( encounterId uint64, diagnosesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( diagnosesIds, ",")

		for _, diagnosesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Diagnosis

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Diagnosis
			// with a matching diagnosesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , diagnosesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DiagnosisObj from the Diagnoses array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Diagnoses").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Diagnoses", diagnosesId )
				return utils.RequestResult{false, msg, "removeDiagnoses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more proceduresIds as a Procedures to a Encounter
//----------------------------------------------------------------------------
func AddProceduresToEncounter ( encounterId uint64, proceduresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( proceduresIds, ",")

		for _, proceduresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Procedure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Procedure
			// with a matching proceduresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , proceduresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Procedures using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Procedures").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Procedures", proceduresId )
				return utils.RequestResult{false, msg, "unassignProcedures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more proceduresIds as a Procedures from a Encounter
//----------------------------------------------------------------------------
func RemoveProceduresFromEncounter( encounterId uint64, proceduresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( proceduresIds, ",")

		for _, proceduresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Procedure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Procedure
			// with a matching proceduresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , proceduresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProcedureObj from the Procedures array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Procedures").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Procedures", proceduresId )
				return utils.RequestResult{false, msg, "removeProcedures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more observationsIds as a Observations to a Encounter
//----------------------------------------------------------------------------
func AddObservationsToEncounter ( encounterId uint64, observationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( observationsIds, ",")

		for _, observationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Observation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Observation
			// with a matching observationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , observationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Observations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Observations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Observations", observationsId )
				return utils.RequestResult{false, msg, "unassignObservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more observationsIds as a Observations from a Encounter
//----------------------------------------------------------------------------
func RemoveObservationsFromEncounter( encounterId uint64, observationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( observationsIds, ",")

		for _, observationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Observation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Observation
			// with a matching observationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , observationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ObservationObj from the Observations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Observations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Observations", observationsId )
				return utils.RequestResult{false, msg, "removeObservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ordersIds as a Orders to a Encounter
//----------------------------------------------------------------------------
func AddOrdersToEncounter ( encounterId uint64, ordersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClinicalOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClinicalOrder
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Orders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "unassignOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ordersIds as a Orders from a Encounter
//----------------------------------------------------------------------------
func RemoveOrdersFromEncounter( encounterId uint64, ordersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Encounter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEncounter(encounterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Encounter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Encounter)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClinicalOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClinicalOrder
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClinicalOrderObj from the Orders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "removeOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Encounter from the gorm
		//----------------------------------------------------------------------------
		return GetEncounter(encounterId)

	} else {
		return parentRequestResult
	}
}

