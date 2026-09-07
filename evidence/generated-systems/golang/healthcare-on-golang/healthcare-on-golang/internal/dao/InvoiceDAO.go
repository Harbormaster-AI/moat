package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InvoiceDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInvoice - creates a new db entry
//----------------------------------------------------------------------------
func CreateInvoice(obj model.Invoice)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Invoice with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Invoice", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInvoice", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInvoice - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInvoice(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Invoice

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Invoice with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Invoice using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Invoice using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInvoice", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInvoice - returns all
//----------------------------------------------------------------------------
func GetAllInvoice()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Invoice

	//----------------------------------------------------------------------------
	// Request the ORM to find all Invoice
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Invoice" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Invoice", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInvoice", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInvoice - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInvoice(obj model.Invoice)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Invoice using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Invoice using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInvoice", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInvoice - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInvoice(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Invoice with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInvoice(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Invoice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Invoice)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Invoice using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Invoice using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInvoice", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a Invoice
//----------------------------------------------------------------------------
func AssignPatientToInvoice( invoiceId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Invoice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvoice(invoiceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Invoice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Invoice)

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
			// assign the Patient	to the Invoice
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Invoice
			//----------------------------------------------------------------------------
			return UpdateInvoice(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Invoice
//----------------------------------------------------------------------------
func UnassignPatientFromInvoice(invoiceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Invoice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvoice(invoiceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Invoice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Invoice)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Invoice
		//----------------------------------------------------------------------------
		return UpdateInvoice(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Claim on a Invoice
//----------------------------------------------------------------------------
func AssignClaimToInvoice( invoiceId uint64, claimId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Invoice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvoice(invoiceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Invoice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Invoice)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Claim

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Claim with a
		// matching claimId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, claimId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Claim	to the Invoice
			//----------------------------------------------------------------------------
			parentObj.Claim = &childObj

			//----------------------------------------------------------------------------
			// save the Invoice
			//----------------------------------------------------------------------------
			return UpdateInvoice(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claim", claimId )
			return utils.RequestResult{false, msg, "assignClaim", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Claim on a Invoice
//----------------------------------------------------------------------------
func UnassignClaimFromInvoice(invoiceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Invoice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvoice(invoiceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Invoice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Invoice)

		//----------------------------------------------------------------------------
		// assign an empty Claim to the Claim
		//----------------------------------------------------------------------------
		parentObj.Claim = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Claim
		//----------------------------------------------------------------------------
		parentObj.ClaimId = nil;

		//----------------------------------------------------------------------------
		// save the Invoice
		//----------------------------------------------------------------------------
		return UpdateInvoice(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more paymentsIds as a Payments to a Invoice
//----------------------------------------------------------------------------
func AddPaymentsToInvoice ( invoiceId uint64, paymentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Invoice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvoice(invoiceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Invoice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Invoice)

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
		// retrieve the modified Invoice from the gorm
		//----------------------------------------------------------------------------
		return GetInvoice(invoiceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentsIds as a Payments from a Invoice
//----------------------------------------------------------------------------
func RemovePaymentsFromInvoice( invoiceId uint64, paymentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Invoice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvoice(invoiceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Invoice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Invoice)

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
		// retrieve the modified Invoice from the gorm
		//----------------------------------------------------------------------------
		return GetInvoice(invoiceId)

	} else {
		return parentRequestResult
	}
}

