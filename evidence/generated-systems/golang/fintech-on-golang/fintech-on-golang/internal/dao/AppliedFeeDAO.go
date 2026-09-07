package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AppliedFeeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAppliedFee - creates a new db entry
//----------------------------------------------------------------------------
func CreateAppliedFee(obj model.AppliedFee)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AppliedFee with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AppliedFee", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAppliedFee", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAppliedFee - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAppliedFee(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AppliedFee

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AppliedFee with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AppliedFee using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AppliedFee using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAppliedFee", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAppliedFee - returns all
//----------------------------------------------------------------------------
func GetAllAppliedFee()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AppliedFee

	//----------------------------------------------------------------------------
	// Request the ORM to find all AppliedFee
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AppliedFee" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AppliedFee", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAppliedFee", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAppliedFee - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAppliedFee(obj model.AppliedFee)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AppliedFee using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AppliedFee using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAppliedFee", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAppliedFee - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAppliedFee(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AppliedFee with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAppliedFee(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AppliedFee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AppliedFee)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AppliedFee using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AppliedFee using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAppliedFee", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PaymentOrder on a AppliedFee
//----------------------------------------------------------------------------
func AssignPaymentOrderToAppliedFee( appliedFeeId uint64, paymentOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AppliedFee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppliedFee(appliedFeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AppliedFee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AppliedFee)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PaymentOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PaymentOrder with a
		// matching paymentOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, paymentOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PaymentOrder	to the AppliedFee
			//----------------------------------------------------------------------------
			parentObj.PaymentOrder = &childObj

			//----------------------------------------------------------------------------
			// save the AppliedFee
			//----------------------------------------------------------------------------
			return UpdateAppliedFee(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentOrder", paymentOrderId )
			return utils.RequestResult{false, msg, "assignPaymentOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PaymentOrder on a AppliedFee
//----------------------------------------------------------------------------
func UnassignPaymentOrderFromAppliedFee(appliedFeeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AppliedFee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppliedFee(appliedFeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AppliedFee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AppliedFee)

		//----------------------------------------------------------------------------
		// assign an empty PaymentOrder to the PaymentOrder
		//----------------------------------------------------------------------------
		parentObj.PaymentOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PaymentOrder
		//----------------------------------------------------------------------------
		parentObj.PaymentOrderId = nil;

		//----------------------------------------------------------------------------
		// save the AppliedFee
		//----------------------------------------------------------------------------
		return UpdateAppliedFee(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Transaction on a AppliedFee
//----------------------------------------------------------------------------
func AssignTransactionToAppliedFee( appliedFeeId uint64, transactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AppliedFee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppliedFee(appliedFeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AppliedFee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AppliedFee)

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
			// assign the Transaction	to the AppliedFee
			//----------------------------------------------------------------------------
			parentObj.Transaction = &childObj

			//----------------------------------------------------------------------------
			// save the AppliedFee
			//----------------------------------------------------------------------------
			return UpdateAppliedFee(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transaction", transactionId )
			return utils.RequestResult{false, msg, "assignTransaction", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Transaction on a AppliedFee
//----------------------------------------------------------------------------
func UnassignTransactionFromAppliedFee(appliedFeeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AppliedFee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAppliedFee(appliedFeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AppliedFee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AppliedFee)

		//----------------------------------------------------------------------------
		// assign an empty Transaction to the Transaction
		//----------------------------------------------------------------------------
		parentObj.Transaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Transaction
		//----------------------------------------------------------------------------
		parentObj.TransactionId = nil;

		//----------------------------------------------------------------------------
		// save the AppliedFee
		//----------------------------------------------------------------------------
		return UpdateAppliedFee(parentObj)

	} else {
		return parentRequestResult
	}

}


