package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BeneficiaryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBeneficiary - creates a new db entry
//----------------------------------------------------------------------------
func CreateBeneficiary(obj model.Beneficiary)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Beneficiary with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Beneficiary", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBeneficiary", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBeneficiary - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBeneficiary(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Beneficiary

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Beneficiary with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Beneficiary using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Beneficiary using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBeneficiary", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBeneficiary - returns all
//----------------------------------------------------------------------------
func GetAllBeneficiary()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Beneficiary

	//----------------------------------------------------------------------------
	// Request the ORM to find all Beneficiary
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Beneficiary" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Beneficiary", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBeneficiary", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBeneficiary - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBeneficiary(obj model.Beneficiary)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Beneficiary using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Beneficiary using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBeneficiary", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBeneficiary - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBeneficiary(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Beneficiary with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBeneficiary(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Beneficiary so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Beneficiary)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Beneficiary using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Beneficiary using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBeneficiary", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a Beneficiary
//----------------------------------------------------------------------------
func AssignCustomerToBeneficiary( beneficiaryId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Beneficiary with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBeneficiary(beneficiaryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Beneficiary so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Beneficiary)

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
			// assign the Customer	to the Beneficiary
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Beneficiary
			//----------------------------------------------------------------------------
			return UpdateBeneficiary(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Beneficiary
//----------------------------------------------------------------------------
func UnassignCustomerFromBeneficiary(beneficiaryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Beneficiary with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBeneficiary(beneficiaryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Beneficiary so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Beneficiary)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Beneficiary
		//----------------------------------------------------------------------------
		return UpdateBeneficiary(parentObj)

	} else {
		return parentRequestResult
	}

}


