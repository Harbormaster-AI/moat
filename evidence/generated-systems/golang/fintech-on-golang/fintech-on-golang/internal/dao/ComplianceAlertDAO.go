package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ComplianceAlertDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateComplianceAlert - creates a new db entry
//----------------------------------------------------------------------------
func CreateComplianceAlert(obj model.ComplianceAlert)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ComplianceAlert with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ComplianceAlert", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateComplianceAlert", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetComplianceAlert - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetComplianceAlert(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ComplianceAlert

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ComplianceAlert with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ComplianceAlert using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ComplianceAlert using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetComplianceAlert", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllComplianceAlert - returns all
//----------------------------------------------------------------------------
func GetAllComplianceAlert()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ComplianceAlert

	//----------------------------------------------------------------------------
	// Request the ORM to find all ComplianceAlert
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ComplianceAlert" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ComplianceAlert", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllComplianceAlert", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateComplianceAlert - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateComplianceAlert(obj model.ComplianceAlert)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ComplianceAlert using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ComplianceAlert using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateComplianceAlert", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteComplianceAlert - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteComplianceAlert(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ComplianceAlert with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetComplianceAlert(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ComplianceAlert)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ComplianceAlert using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ComplianceAlert using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteComplianceAlert", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Screening on a ComplianceAlert
//----------------------------------------------------------------------------
func AssignScreeningToComplianceAlert( complianceAlertId uint64, screeningId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ComplianceAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceAlert(complianceAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceAlert)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Screening

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Screening with a
		// matching screeningId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, screeningId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Screening	to the ComplianceAlert
			//----------------------------------------------------------------------------
			parentObj.Screening = &childObj

			//----------------------------------------------------------------------------
			// save the ComplianceAlert
			//----------------------------------------------------------------------------
			return UpdateComplianceAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Screening", screeningId )
			return utils.RequestResult{false, msg, "assignScreening", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Screening on a ComplianceAlert
//----------------------------------------------------------------------------
func UnassignScreeningFromComplianceAlert(complianceAlertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceAlert(complianceAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceAlert)

		//----------------------------------------------------------------------------
		// assign an empty Screening to the Screening
		//----------------------------------------------------------------------------
		parentObj.Screening = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Screening
		//----------------------------------------------------------------------------
		parentObj.ScreeningId = nil;

		//----------------------------------------------------------------------------
		// save the ComplianceAlert
		//----------------------------------------------------------------------------
		return UpdateComplianceAlert(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Transaction on a ComplianceAlert
//----------------------------------------------------------------------------
func AssignTransactionToComplianceAlert( complianceAlertId uint64, transactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ComplianceAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceAlert(complianceAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceAlert)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Transaction

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Transaction with a
		// matching transactionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, transactionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Transaction	to the ComplianceAlert
			//----------------------------------------------------------------------------
			parentObj.Transaction = &childObj

			//----------------------------------------------------------------------------
			// save the ComplianceAlert
			//----------------------------------------------------------------------------
			return UpdateComplianceAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )
			return utils.RequestResult{false, msg, "assignTransaction", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a ComplianceAlert
//----------------------------------------------------------------------------
func UnassignTransactionFromComplianceAlert(complianceAlertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceAlert(complianceAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceAlert)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		parentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		parentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the ComplianceAlert
		//----------------------------------------------------------------------------
		return UpdateComplianceAlert(parentObj)

	} else {
		return parentRequestResult
	}

}


