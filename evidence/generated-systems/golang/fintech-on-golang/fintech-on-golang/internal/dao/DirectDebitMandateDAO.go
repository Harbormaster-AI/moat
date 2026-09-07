package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DirectDebitMandateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDirectDebitMandate - creates a new db entry
//----------------------------------------------------------------------------
func CreateDirectDebitMandate(obj model.DirectDebitMandate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DirectDebitMandate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DirectDebitMandate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDirectDebitMandate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDirectDebitMandate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDirectDebitMandate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DirectDebitMandate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DirectDebitMandate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DirectDebitMandate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DirectDebitMandate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDirectDebitMandate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDirectDebitMandate - returns all
//----------------------------------------------------------------------------
func GetAllDirectDebitMandate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DirectDebitMandate

	//----------------------------------------------------------------------------
	// Request the ORM to find all DirectDebitMandate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DirectDebitMandate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DirectDebitMandate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDirectDebitMandate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDirectDebitMandate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDirectDebitMandate(obj model.DirectDebitMandate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DirectDebitMandate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DirectDebitMandate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDirectDebitMandate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDirectDebitMandate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDirectDebitMandate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DirectDebitMandate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDirectDebitMandate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DirectDebitMandate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DirectDebitMandate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DirectDebitMandate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DirectDebitMandate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDirectDebitMandate", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Account on a DirectDebitMandate
//----------------------------------------------------------------------------
func AssignAccountToDirectDebitMandate( directDebitMandateId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DirectDebitMandate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDirectDebitMandate(directDebitMandateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DirectDebitMandate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DirectDebitMandate)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching accountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, accountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Account	to the DirectDebitMandate
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the DirectDebitMandate
			//----------------------------------------------------------------------------
			return UpdateDirectDebitMandate(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a DirectDebitMandate
//----------------------------------------------------------------------------
func UnassignAccountFromDirectDebitMandate(directDebitMandateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DirectDebitMandate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDirectDebitMandate(directDebitMandateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DirectDebitMandate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DirectDebitMandate)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the DirectDebitMandate
		//----------------------------------------------------------------------------
		return UpdateDirectDebitMandate(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Creditor on a DirectDebitMandate
//----------------------------------------------------------------------------
func AssignCreditorToDirectDebitMandate( directDebitMandateId uint64, creditorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DirectDebitMandate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDirectDebitMandate(directDebitMandateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DirectDebitMandate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DirectDebitMandate)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Creditor

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Creditor with a
		// matching creditorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, creditorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Creditor	to the DirectDebitMandate
			//----------------------------------------------------------------------------
			parentObj.Creditor = &childObj

			//----------------------------------------------------------------------------
			// save the DirectDebitMandate
			//----------------------------------------------------------------------------
			return UpdateDirectDebitMandate(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Creditor", creditorId )
			return utils.RequestResult{false, msg, "assignCreditor", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Creditor on a DirectDebitMandate
//----------------------------------------------------------------------------
func UnassignCreditorFromDirectDebitMandate(directDebitMandateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DirectDebitMandate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDirectDebitMandate(directDebitMandateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DirectDebitMandate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DirectDebitMandate)

		//----------------------------------------------------------------------------
		// assign an empty Creditor to the Creditor
		//----------------------------------------------------------------------------
		parentObj.Creditor = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Creditor
		//----------------------------------------------------------------------------
		parentObj.CreditorId = nil;

		//----------------------------------------------------------------------------
		// save the DirectDebitMandate
		//----------------------------------------------------------------------------
		return UpdateDirectDebitMandate(parentObj)

	} else {
		return parentRequestResult
	}

}


