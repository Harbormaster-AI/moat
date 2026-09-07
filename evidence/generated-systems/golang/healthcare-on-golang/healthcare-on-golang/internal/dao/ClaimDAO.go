package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ClaimDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateClaim - creates a new db entry
//----------------------------------------------------------------------------
func CreateClaim(obj model.Claim)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Claim with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Claim", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateClaim", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetClaim - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetClaim(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Claim

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Claim with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Claim using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Claim using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetClaim", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllClaim - returns all
//----------------------------------------------------------------------------
func GetAllClaim()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Claim

	//----------------------------------------------------------------------------
	// Request the ORM to find all Claim
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Claim" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Claim", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllClaim", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateClaim - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateClaim(obj model.Claim)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Claim using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Claim using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateClaim", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteClaim - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteClaim(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetClaim(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Claim using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Claim using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteClaim", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a Claim
//----------------------------------------------------------------------------
func AssignPatientToClaim( claimId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

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
			// assign the Patient	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Claim
//----------------------------------------------------------------------------
func UnassignPatientFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Coverage on a Claim
//----------------------------------------------------------------------------
func AssignCoverageToClaim( claimId uint64, coverageId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Coverage

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Coverage with a
		// matching coverageId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, coverageId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Coverage	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Coverage = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverage", coverageId )
			return utils.RequestResult{false, msg, "assignCoverage", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Coverage on a Claim
//----------------------------------------------------------------------------
func UnassignCoverageFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty Coverage to the Coverage
		//----------------------------------------------------------------------------
		parentObj.Coverage = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Coverage
		//----------------------------------------------------------------------------
		parentObj.CoverageId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Encounter on a Claim
//----------------------------------------------------------------------------
func AssignEncounterToClaim( claimId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

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
			// assign the Encounter	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a Claim
//----------------------------------------------------------------------------
func UnassignEncounterFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Payer on a Claim
//----------------------------------------------------------------------------
func AssignPayerToClaim( claimId uint64, payerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsurancePayer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsurancePayer with a
		// matching payerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, payerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Payer	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Payer = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payer", payerId )
			return utils.RequestResult{false, msg, "assignPayer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Payer on a Claim
//----------------------------------------------------------------------------
func UnassignPayerFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty InsurancePayer to the Payer
		//----------------------------------------------------------------------------
		parentObj.Payer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Payer
		//----------------------------------------------------------------------------
		parentObj.PayerId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more invoicesIds as a Invoices to a Claim
//----------------------------------------------------------------------------
func AddInvoicesToClaim ( claimId uint64, invoicesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( invoicesIds, ",")

		for _, invoicesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Invoice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Invoice
			// with a matching invoicesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , invoicesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Invoices using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Invoices").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Invoices", invoicesId )
				return utils.RequestResult{false, msg, "unassignInvoices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more invoicesIds as a Invoices from a Claim
//----------------------------------------------------------------------------
func RemoveInvoicesFromClaim( claimId uint64, invoicesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( invoicesIds, ",")

		for _, invoicesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Invoice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Invoice
			// with a matching invoicesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , invoicesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InvoiceObj from the Invoices array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Invoices").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Invoices", invoicesId )
				return utils.RequestResult{false, msg, "removeInvoices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

