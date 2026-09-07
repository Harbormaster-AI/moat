package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PriceBookDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePriceBook - creates a new db entry
//----------------------------------------------------------------------------
func CreatePriceBook(obj model.PriceBook)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PriceBook with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PriceBook", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePriceBook", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPriceBook - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPriceBook(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PriceBook

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PriceBook with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PriceBook using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PriceBook using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPriceBook", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPriceBook - returns all
//----------------------------------------------------------------------------
func GetAllPriceBook()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PriceBook

	//----------------------------------------------------------------------------
	// Request the ORM to find all PriceBook
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PriceBook" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PriceBook", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPriceBook", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePriceBook - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePriceBook(obj model.PriceBook)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PriceBook using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PriceBook using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePriceBook", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePriceBook - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePriceBook(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPriceBook(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PriceBook)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PriceBook using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PriceBook using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePriceBook", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a PriceBook
//----------------------------------------------------------------------------
func AssignOrganizationToPriceBook( priceBookId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the PriceBook
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the PriceBook
			//----------------------------------------------------------------------------
			return UpdatePriceBook(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a PriceBook
//----------------------------------------------------------------------------
func UnassignOrganizationFromPriceBook(priceBookId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the PriceBook
		//----------------------------------------------------------------------------
		return UpdatePriceBook(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more entriesIds as a Entries to a PriceBook
//----------------------------------------------------------------------------
func AddEntriesToPriceBook ( priceBookId uint64, entriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		// slice the ids on comma with no spaces
		ids := strings.Split( entriesIds, ",")

		for _, entriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PriceBookEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PriceBookEntry
			// with a matching entriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , entriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Entries using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Entries").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Entries", entriesId )
				return utils.RequestResult{false, msg, "unassignEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PriceBook from the gorm
		//----------------------------------------------------------------------------
		return GetPriceBook(priceBookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more entriesIds as a Entries from a PriceBook
//----------------------------------------------------------------------------
func RemoveEntriesFromPriceBook( priceBookId uint64, entriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		// slice the ids on comma with no spaces
		ids := strings.Split( entriesIds, ",")

		for _, entriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PriceBookEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PriceBookEntry
			// with a matching entriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , entriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PriceBookEntryObj from the Entries array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Entries").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Entries", entriesId )
				return utils.RequestResult{false, msg, "removeEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PriceBook from the gorm
		//----------------------------------------------------------------------------
		return GetPriceBook(priceBookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more quotesIds as a Quotes to a PriceBook
//----------------------------------------------------------------------------
func AddQuotesToPriceBook ( priceBookId uint64, quotesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		// slice the ids on comma with no spaces
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Quotes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "unassignQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PriceBook from the gorm
		//----------------------------------------------------------------------------
		return GetPriceBook(priceBookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more quotesIds as a Quotes from a PriceBook
//----------------------------------------------------------------------------
func RemoveQuotesFromPriceBook( priceBookId uint64, quotesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		// slice the ids on comma with no spaces
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QuoteObj from the Quotes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "removeQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PriceBook from the gorm
		//----------------------------------------------------------------------------
		return GetPriceBook(priceBookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ordersIds as a Orders to a PriceBook
//----------------------------------------------------------------------------
func AddOrdersToPriceBook ( priceBookId uint64, ordersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Order

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Order
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Orders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "unassignOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PriceBook from the gorm
		//----------------------------------------------------------------------------
		return GetPriceBook(priceBookId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ordersIds as a Orders from a PriceBook
//----------------------------------------------------------------------------
func RemoveOrdersFromPriceBook( priceBookId uint64, ordersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PriceBook with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPriceBook(priceBookId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PriceBook so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PriceBook)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Order

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Order
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OrderObj from the Orders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "removeOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PriceBook from the gorm
		//----------------------------------------------------------------------------
		return GetPriceBook(priceBookId)

	} else {
		return parentRequestResult
	}
}

