package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ImagingReportDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateImagingReport - creates a new db entry
//----------------------------------------------------------------------------
func CreateImagingReport(obj model.ImagingReport)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ImagingReport with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ImagingReport", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateImagingReport", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetImagingReport - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetImagingReport(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ImagingReport

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ImagingReport with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ImagingReport using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ImagingReport using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetImagingReport", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllImagingReport - returns all
//----------------------------------------------------------------------------
func GetAllImagingReport()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ImagingReport

	//----------------------------------------------------------------------------
	// Request the ORM to find all ImagingReport
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ImagingReport" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ImagingReport", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllImagingReport", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateImagingReport - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateImagingReport(obj model.ImagingReport)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ImagingReport using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ImagingReport using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateImagingReport", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteImagingReport - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteImagingReport(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetImagingReport(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ImagingReport)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ImagingReport using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ImagingReport using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteImagingReport", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ImagingOrder on a ImagingReport
//----------------------------------------------------------------------------
func AssignImagingOrderToImagingReport( imagingReportId uint64, imagingOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ImagingOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ImagingOrder with a
		// matching imagingOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, imagingOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ImagingOrder	to the ImagingReport
			//----------------------------------------------------------------------------
			parentObj.ImagingOrder = &childObj

			//----------------------------------------------------------------------------
			// save the ImagingReport
			//----------------------------------------------------------------------------
			return UpdateImagingReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingOrder", imagingOrderId )
			return utils.RequestResult{false, msg, "assignImagingOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ImagingOrder on a ImagingReport
//----------------------------------------------------------------------------
func UnassignImagingOrderFromImagingReport(imagingReportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

		//----------------------------------------------------------------------------
		// assign an empty ImagingOrder to the ImagingOrder
		//----------------------------------------------------------------------------
		parentObj.ImagingOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ImagingOrder
		//----------------------------------------------------------------------------
		parentObj.ImagingOrderId = nil;

		//----------------------------------------------------------------------------
		// save the ImagingReport
		//----------------------------------------------------------------------------
		return UpdateImagingReport(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Clinician on a ImagingReport
//----------------------------------------------------------------------------
func AssignClinicianToImagingReport( imagingReportId uint64, clinicianId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

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
			// assign the Clinician	to the ImagingReport
			//----------------------------------------------------------------------------
			parentObj.Clinician = &childObj

			//----------------------------------------------------------------------------
			// save the ImagingReport
			//----------------------------------------------------------------------------
			return UpdateImagingReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Clinician", clinicianId )
			return utils.RequestResult{false, msg, "assignClinician", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Clinician on a ImagingReport
//----------------------------------------------------------------------------
func UnassignClinicianFromImagingReport(imagingReportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

		//----------------------------------------------------------------------------
		// assign an empty Clinician to the Clinician
		//----------------------------------------------------------------------------
		parentObj.Clinician = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Clinician
		//----------------------------------------------------------------------------
		parentObj.ClinicianId = nil;

		//----------------------------------------------------------------------------
		// save the ImagingReport
		//----------------------------------------------------------------------------
		return UpdateImagingReport(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Encounter on a ImagingReport
//----------------------------------------------------------------------------
func AssignEncounterToImagingReport( imagingReportId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

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
			// assign the Encounter	to the ImagingReport
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the ImagingReport
			//----------------------------------------------------------------------------
			return UpdateImagingReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a ImagingReport
//----------------------------------------------------------------------------
func UnassignEncounterFromImagingReport(imagingReportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the ImagingReport
		//----------------------------------------------------------------------------
		return UpdateImagingReport(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ImagingCenter on a ImagingReport
//----------------------------------------------------------------------------
func AssignImagingCenterToImagingReport( imagingReportId uint64, imagingCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ImagingCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ImagingCenter with a
		// matching imagingCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, imagingCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ImagingCenter	to the ImagingReport
			//----------------------------------------------------------------------------
			parentObj.ImagingCenter = &childObj

			//----------------------------------------------------------------------------
			// save the ImagingReport
			//----------------------------------------------------------------------------
			return UpdateImagingReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingCenter", imagingCenterId )
			return utils.RequestResult{false, msg, "assignImagingCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ImagingCenter on a ImagingReport
//----------------------------------------------------------------------------
func UnassignImagingCenterFromImagingReport(imagingReportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingReport with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingReport(imagingReportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingReport so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingReport)

		//----------------------------------------------------------------------------
		// assign an empty ImagingCenter to the ImagingCenter
		//----------------------------------------------------------------------------
		parentObj.ImagingCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ImagingCenter
		//----------------------------------------------------------------------------
		parentObj.ImagingCenterId = nil;

		//----------------------------------------------------------------------------
		// save the ImagingReport
		//----------------------------------------------------------------------------
		return UpdateImagingReport(parentObj)

	} else {
		return parentRequestResult
	}

}


