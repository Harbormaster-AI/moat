package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PriceBookEntryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePriceBookEntry - creates a new db entry
//----------------------------------------------------------------------------
func CreatePriceBookEntry(obj model.PriceBookEntry)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PriceBookEntry with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PriceBookEntry", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePriceBookEntry", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPriceBookEntry - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPriceBookEntry(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PriceBookEntry

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PriceBookEntry with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PriceBookEntry using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PriceBookEntry using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPriceBookEntry", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPriceBookEntry - returns all
//----------------------------------------------------------------------------
func GetAllPriceBookEntry()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PriceBookEntry

	//----------------------------------------------------------------------------
	// Request the ORM to find all PriceBookEntry
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PriceBookEntry" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PriceBookEntry", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPriceBookEntry", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePriceBookEntry - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePriceBookEntry(obj model.PriceBookEntry)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PriceBookEntry using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PriceBookEntry using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePriceBookEntry", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePriceBookEntry - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePriceBookEntry(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PriceBookEntry with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPriceBookEntry(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBookEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PriceBookEntry)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PriceBookEntry using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PriceBookEntry using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePriceBookEntry", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PriceBook on a PriceBookEntry
//----------------------------------------------------------------------------
func AssignPriceBookToPriceBookEntry( priceBookEntryId uint64, priceBookId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PriceBookEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBookEntry(priceBookEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBookEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBookEntry)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PriceBook

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PriceBook with a
		// matching priceBookId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, priceBookId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PriceBook	to the PriceBookEntry
			//----------------------------------------------------------------------------
			parentObj.PriceBook = &childObj

			//----------------------------------------------------------------------------
			// save the PriceBookEntry
			//----------------------------------------------------------------------------
			return UpdatePriceBookEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBook", priceBookId )
			return utils.RequestResult{false, msg, "assignPriceBook", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PriceBook on a PriceBookEntry
//----------------------------------------------------------------------------
func UnassignPriceBookFromPriceBookEntry(priceBookEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PriceBookEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBookEntry(priceBookEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBookEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBookEntry)

		//----------------------------------------------------------------------------
		// assign an empty PriceBook to the PriceBook
		//----------------------------------------------------------------------------
		parentObj.PriceBook = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PriceBook
		//----------------------------------------------------------------------------
		parentObj.PriceBookId = nil;

		//----------------------------------------------------------------------------
		// save the PriceBookEntry
		//----------------------------------------------------------------------------
		return UpdatePriceBookEntry(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a PriceBookEntry
//----------------------------------------------------------------------------
func AssignProductToPriceBookEntry( priceBookEntryId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PriceBookEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBookEntry(priceBookEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBookEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBookEntry)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Product

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Product with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the PriceBookEntry
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the PriceBookEntry
			//----------------------------------------------------------------------------
			return UpdatePriceBookEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a PriceBookEntry
//----------------------------------------------------------------------------
func UnassignProductFromPriceBookEntry(priceBookEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PriceBookEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBookEntry(priceBookEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBookEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBookEntry)

		//----------------------------------------------------------------------------
		// assign an empty Product to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the PriceBookEntry
		//----------------------------------------------------------------------------
		return UpdatePriceBookEntry(parentObj)

	} else {
		return parentRequestResult
	}

}


