package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TerminalDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTerminal - creates a new db entry
//----------------------------------------------------------------------------
func CreateTerminal(obj model.Terminal)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Terminal with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Terminal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTerminal", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTerminal - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTerminal(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Terminal

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Terminal with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Terminal using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Terminal using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTerminal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTerminal - returns all
//----------------------------------------------------------------------------
func GetAllTerminal()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Terminal

	//----------------------------------------------------------------------------
	// Request the ORM to find all Terminal
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Terminal" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Terminal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTerminal", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTerminal - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTerminal(obj model.Terminal)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Terminal using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Terminal using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTerminal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTerminal - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTerminal(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Terminal with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTerminal(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Terminal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Terminal)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Terminal using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Terminal using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTerminal", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Merchant on a Terminal
//----------------------------------------------------------------------------
func AssignMerchantToTerminal( terminalId uint64, merchantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Terminal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerminal(terminalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Terminal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Terminal)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Merchant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Merchant with a
		// matching merchantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, merchantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Merchant	to the Terminal
			//----------------------------------------------------------------------------
			parentObj.Merchant = &childObj

			//----------------------------------------------------------------------------
			// save the Terminal
			//----------------------------------------------------------------------------
			return UpdateTerminal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Merchant", merchantId )
			return utils.RequestResult{false, msg, "assignMerchant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Merchant on a Terminal
//----------------------------------------------------------------------------
func UnassignMerchantFromTerminal(terminalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Terminal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTerminal(terminalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Terminal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Terminal)

		//----------------------------------------------------------------------------
		// assign an empty Merchant to the Merchant
		//----------------------------------------------------------------------------
		parentObj.Merchant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Merchant
		//----------------------------------------------------------------------------
		parentObj.MerchantId = nil;

		//----------------------------------------------------------------------------
		// save the Terminal
		//----------------------------------------------------------------------------
		return UpdateTerminal(parentObj)

	} else {
		return parentRequestResult
	}

}


