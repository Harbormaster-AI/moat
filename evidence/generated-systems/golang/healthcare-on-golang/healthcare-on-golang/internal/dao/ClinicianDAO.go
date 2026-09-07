package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ClinicianDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateClinician - creates a new db entry
//----------------------------------------------------------------------------
func CreateClinician(obj model.Clinician)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Clinician with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Clinician", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateClinician", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetClinician - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetClinician(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Clinician

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Clinician with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Clinician using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Clinician using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetClinician", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllClinician - returns all
//----------------------------------------------------------------------------
func GetAllClinician()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Clinician

	//----------------------------------------------------------------------------
	// Request the ORM to find all Clinician
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Clinician" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Clinician", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllClinician", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateClinician - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateClinician(obj model.Clinician)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Clinician using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Clinician using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateClinician", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteClinician - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteClinician(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetClinician(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Clinician)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Clinician using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Clinician using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteClinician", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more careTeamsIds as a CareTeams to a Clinician
//----------------------------------------------------------------------------
func AddCareTeamsToClinician ( clinicianId uint64, careTeamsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( careTeamsIds, ",")

		for _, careTeamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CareTeam

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CareTeam
			// with a matching careTeamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , careTeamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CareTeams using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CareTeams").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CareTeams", careTeamsId )
				return utils.RequestResult{false, msg, "unassignCareTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more careTeamsIds as a CareTeams from a Clinician
//----------------------------------------------------------------------------
func RemoveCareTeamsFromClinician( clinicianId uint64, careTeamsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( careTeamsIds, ",")

		for _, careTeamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CareTeam

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CareTeam
			// with a matching careTeamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , careTeamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CareTeamObj from the CareTeams array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CareTeams").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CareTeams", careTeamsId )
				return utils.RequestResult{false, msg, "removeCareTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more appointmentsIds as a Appointments to a Clinician
//----------------------------------------------------------------------------
func AddAppointmentsToClinician ( clinicianId uint64, appointmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( appointmentsIds, ",")

		for _, appointmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Appointment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Appointment
			// with a matching appointmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , appointmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Appointments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Appointments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Appointments", appointmentsId )
				return utils.RequestResult{false, msg, "unassignAppointments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more appointmentsIds as a Appointments from a Clinician
//----------------------------------------------------------------------------
func RemoveAppointmentsFromClinician( clinicianId uint64, appointmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( appointmentsIds, ",")

		for _, appointmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Appointment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Appointment
			// with a matching appointmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , appointmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AppointmentObj from the Appointments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Appointments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Appointments", appointmentsId )
				return utils.RequestResult{false, msg, "removeAppointments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more encountersIds as a Encounters to a Clinician
//----------------------------------------------------------------------------
func AddEncountersToClinician ( clinicianId uint64, encountersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( encountersIds, ",")

		for _, encountersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Encounter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Encounter
			// with a matching encountersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , encountersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Encounters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Encounters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounters", encountersId )
				return utils.RequestResult{false, msg, "unassignEncounters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more encountersIds as a Encounters from a Clinician
//----------------------------------------------------------------------------
func RemoveEncountersFromClinician( clinicianId uint64, encountersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( encountersIds, ",")

		for _, encountersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Encounter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Encounter
			// with a matching encountersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , encountersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EncounterObj from the Encounters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Encounters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounters", encountersId )
				return utils.RequestResult{false, msg, "removeEncounters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more proceduresIds as a Procedures to a Clinician
//----------------------------------------------------------------------------
func AddProceduresToClinician ( clinicianId uint64, proceduresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

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
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more proceduresIds as a Procedures from a Clinician
//----------------------------------------------------------------------------
func RemoveProceduresFromClinician( clinicianId uint64, proceduresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

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
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more imagingReportsIds as a ImagingReports to a Clinician
//----------------------------------------------------------------------------
func AddImagingReportsToClinician ( clinicianId uint64, imagingReportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingReportsIds, ",")

		for _, imagingReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingReport

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingReport
			// with a matching imagingReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ImagingReports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingReports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingReports", imagingReportsId )
				return utils.RequestResult{false, msg, "unassignImagingReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more imagingReportsIds as a ImagingReports from a Clinician
//----------------------------------------------------------------------------
func RemoveImagingReportsFromClinician( clinicianId uint64, imagingReportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Clinician with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClinician(clinicianId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Clinician so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Clinician)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingReportsIds, ",")

		for _, imagingReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingReport

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingReport
			// with a matching imagingReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ImagingReportObj from the ImagingReports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingReports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingReports", imagingReportsId )
				return utils.RequestResult{false, msg, "removeImagingReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Clinician from the gorm
		//----------------------------------------------------------------------------
		return GetClinician(clinicianId)

	} else {
		return parentRequestResult
	}
}

