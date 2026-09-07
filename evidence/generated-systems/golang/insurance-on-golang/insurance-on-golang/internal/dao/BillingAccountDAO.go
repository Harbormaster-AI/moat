package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BillingAccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBillingAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateBillingAccount(obj model.BillingAccount)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BillingAccount with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BillingAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBillingAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBillingAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBillingAccount(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BillingAccount

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BillingAccount with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BillingAccount using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BillingAccount using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBillingAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBillingAccount - returns all
//----------------------------------------------------------------------------
func GetAllBillingAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BillingAccount

	//----------------------------------------------------------------------------
	// Request the ORM to find all BillingAccount
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BillingAccount" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BillingAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBillingAccount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBillingAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBillingAccount(obj model.BillingAccount)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BillingAccount using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BillingAccount using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBillingAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBillingAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBillingAccount(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBillingAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BillingAccount)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BillingAccount using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BillingAccount using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBillingAccount", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a BillingAccount
//----------------------------------------------------------------------------
func AssignCustomerToBillingAccount( billingAccountId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching customerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, customerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Customer	to the BillingAccount
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the BillingAccount
			//----------------------------------------------------------------------------
			return UpdateBillingAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a BillingAccount
//----------------------------------------------------------------------------
func UnassignCustomerFromBillingAccount(billingAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the BillingAccount
		//----------------------------------------------------------------------------
		return UpdateBillingAccount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a BillingAccount
//----------------------------------------------------------------------------
func AddPoliciesToBillingAccount ( billingAccountId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Policies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "unassignPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingAccount from the gorm
		//----------------------------------------------------------------------------
		return GetBillingAccount(billingAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a BillingAccount
//----------------------------------------------------------------------------
func RemovePoliciesFromBillingAccount( billingAccountId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyObj from the Policies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "removePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingAccount from the gorm
		//----------------------------------------------------------------------------
		return GetBillingAccount(billingAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more invoicesIds as a Invoices to a BillingAccount
//----------------------------------------------------------------------------
func AddInvoicesToBillingAccount ( billingAccountId uint64, invoicesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

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
		// retrieve the modified BillingAccount from the gorm
		//----------------------------------------------------------------------------
		return GetBillingAccount(billingAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more invoicesIds as a Invoices from a BillingAccount
//----------------------------------------------------------------------------
func RemoveInvoicesFromBillingAccount( billingAccountId uint64, invoicesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

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
		// retrieve the modified BillingAccount from the gorm
		//----------------------------------------------------------------------------
		return GetBillingAccount(billingAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentsIds as a Payments to a BillingAccount
//----------------------------------------------------------------------------
func AddPaymentsToBillingAccount ( billingAccountId uint64, paymentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Payment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Payment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Payments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "unassignPayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingAccount from the gorm
		//----------------------------------------------------------------------------
		return GetBillingAccount(billingAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentsIds as a Payments from a BillingAccount
//----------------------------------------------------------------------------
func RemovePaymentsFromBillingAccount( billingAccountId uint64, paymentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BillingAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingAccount(billingAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Payment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Payment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentObj from the Payments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "removePayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingAccount from the gorm
		//----------------------------------------------------------------------------
		return GetBillingAccount(billingAccountId)

	} else {
		return parentRequestResult
	}
}

