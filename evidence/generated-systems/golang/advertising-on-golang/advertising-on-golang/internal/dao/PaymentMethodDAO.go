package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PaymentMethodDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePaymentMethod - creates a new db entry
//----------------------------------------------------------------------------
func CreatePaymentMethod(obj model.PaymentMethod)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PaymentMethod with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PaymentMethod", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePaymentMethod", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPaymentMethod - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPaymentMethod(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PaymentMethod

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PaymentMethod with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PaymentMethod using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PaymentMethod using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPaymentMethod", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPaymentMethod - returns all
//----------------------------------------------------------------------------
func GetAllPaymentMethod()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PaymentMethod

	//----------------------------------------------------------------------------
	// Request the ORM to find all PaymentMethod
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PaymentMethod" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PaymentMethod", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPaymentMethod", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePaymentMethod - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePaymentMethod(obj model.PaymentMethod)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PaymentMethod using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PaymentMethod using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePaymentMethod", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePaymentMethod - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePaymentMethod(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PaymentMethod with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPaymentMethod(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentMethod so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PaymentMethod)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PaymentMethod using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PaymentMethod using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePaymentMethod", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a BillingProfile on a PaymentMethod
//----------------------------------------------------------------------------
func AssignBillingProfileToPaymentMethod( paymentMethodId uint64, billingProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PaymentMethod with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentMethod(paymentMethodId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentMethod so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentMethod)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BillingProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BillingProfile with a
		// matching billingProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, billingProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BillingProfile	to the PaymentMethod
			//----------------------------------------------------------------------------
			parentObj.BillingProfile = &childObj

			//----------------------------------------------------------------------------
			// save the PaymentMethod
			//----------------------------------------------------------------------------
			return UpdatePaymentMethod(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BillingProfile", billingProfileId )
			return utils.RequestResult{false, msg, "assignBillingProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BillingProfile on a PaymentMethod
//----------------------------------------------------------------------------
func UnassignBillingProfileFromPaymentMethod(paymentMethodId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PaymentMethod with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPaymentMethod(paymentMethodId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PaymentMethod so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PaymentMethod)

		//----------------------------------------------------------------------------
		// assign an empty BillingProfile to the BillingProfile
		//----------------------------------------------------------------------------
		parentObj.BillingProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BillingProfile
		//----------------------------------------------------------------------------
		parentObj.BillingProfileId = nil;

		//----------------------------------------------------------------------------
		// save the PaymentMethod
		//----------------------------------------------------------------------------
		return UpdatePaymentMethod(parentObj)

	} else {
		return parentRequestResult
	}

}


