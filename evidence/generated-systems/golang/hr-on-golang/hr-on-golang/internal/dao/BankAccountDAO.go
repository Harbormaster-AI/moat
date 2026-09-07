package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BankAccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBankAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateBankAccount(obj model.BankAccount)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BankAccount with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BankAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBankAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBankAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBankAccount(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BankAccount

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BankAccount with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BankAccount using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BankAccount using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBankAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBankAccount - returns all
//----------------------------------------------------------------------------
func GetAllBankAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BankAccount

	//----------------------------------------------------------------------------
	// Request the ORM to find all BankAccount
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BankAccount" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BankAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBankAccount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBankAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBankAccount(obj model.BankAccount)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BankAccount using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BankAccount using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBankAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBankAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBankAccount(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BankAccount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBankAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BankAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BankAccount)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BankAccount using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BankAccount using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBankAccount", requestResult.Data}

	}

	return requestResult
}



