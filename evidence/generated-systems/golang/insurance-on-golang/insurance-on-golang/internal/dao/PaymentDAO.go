package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PaymentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePayment - creates a new db entry
//----------------------------------------------------------------------------
func CreatePayment(obj model.Payment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Payment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Payment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePayment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPayment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPayment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Payment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Payment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Payment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Payment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPayment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPayment - returns all
//----------------------------------------------------------------------------
func GetAllPayment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Payment

	//----------------------------------------------------------------------------
	// Request the ORM to find all Payment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Payment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Payment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPayment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePayment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePayment(obj model.Payment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Payment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Payment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePayment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePayment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePayment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Payment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPayment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Payment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Payment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Payment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePayment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Invoice on a Payment
//----------------------------------------------------------------------------
func AssignInvoiceToPayment( paymentId uint64, invoiceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Payment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayment(paymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Invoice

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Invoice with a
		// matching invoiceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, invoiceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Invoice	to the Payment
			//----------------------------------------------------------------------------
			parentObj.Invoice = &childObj

			//----------------------------------------------------------------------------
			// save the Payment
			//----------------------------------------------------------------------------
			return UpdatePayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Invoice", invoiceId )
			return utils.RequestResult{false, msg, "assignInvoice", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Invoice on a Payment
//----------------------------------------------------------------------------
func UnassignInvoiceFromPayment(paymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Payment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayment(paymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payment)

		//----------------------------------------------------------------------------
		// assign an empty Invoice to the Invoice
		//----------------------------------------------------------------------------
		parentObj.Invoice = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Invoice
		//----------------------------------------------------------------------------
		parentObj.InvoiceId = nil;

		//----------------------------------------------------------------------------
		// save the Payment
		//----------------------------------------------------------------------------
		return UpdatePayment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a BillingAccount on a Payment
//----------------------------------------------------------------------------
func AssignBillingAccountToPayment( paymentId uint64, billingAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Payment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayment(paymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BillingAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BillingAccount with a
		// matching billingAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, billingAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BillingAccount	to the Payment
			//----------------------------------------------------------------------------
			parentObj.BillingAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Payment
			//----------------------------------------------------------------------------
			return UpdatePayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BillingAccount", billingAccountId )
			return utils.RequestResult{false, msg, "assignBillingAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BillingAccount on a Payment
//----------------------------------------------------------------------------
func UnassignBillingAccountFromPayment(paymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Payment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayment(paymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payment)

		//----------------------------------------------------------------------------
		// assign an empty BillingAccount to the BillingAccount
		//----------------------------------------------------------------------------
		parentObj.BillingAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BillingAccount
		//----------------------------------------------------------------------------
		parentObj.BillingAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Payment
		//----------------------------------------------------------------------------
		return UpdatePayment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Policy on a Payment
//----------------------------------------------------------------------------
func AssignPolicyToPayment( paymentId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Payment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayment(paymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Policy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Policy with a
		// matching policyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, policyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Policy	to the Payment
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Payment
			//----------------------------------------------------------------------------
			return UpdatePayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Payment
//----------------------------------------------------------------------------
func UnassignPolicyFromPayment(paymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Payment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPayment(paymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Payment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Payment)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Payment
		//----------------------------------------------------------------------------
		return UpdatePayment(parentObj)

	} else {
		return parentRequestResult
	}

}


