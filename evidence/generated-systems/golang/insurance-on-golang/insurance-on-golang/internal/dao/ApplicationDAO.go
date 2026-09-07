package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ApplicationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateApplication - creates a new db entry
//----------------------------------------------------------------------------
func CreateApplication(obj model.Application)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Application with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Application", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateApplication", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetApplication - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetApplication(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Application

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Application with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Application using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Application using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetApplication", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllApplication - returns all
//----------------------------------------------------------------------------
func GetAllApplication()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Application

	//----------------------------------------------------------------------------
	// Request the ORM to find all Application
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Application" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Application", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllApplication", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateApplication - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateApplication(obj model.Application)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Application using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Application using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateApplication", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteApplication - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteApplication(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetApplication(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Application using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Application using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteApplication", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a Application
//----------------------------------------------------------------------------
func AssignCustomerToApplication( applicationId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

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
			// assign the Customer	to the Application
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Application
			//----------------------------------------------------------------------------
			return UpdateApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Application
//----------------------------------------------------------------------------
func UnassignCustomerFromApplication(applicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Application
		//----------------------------------------------------------------------------
		return UpdateApplication(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a Application
//----------------------------------------------------------------------------
func AssignProductToApplication( applicationId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsuranceProduct

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsuranceProduct with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the Application
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the Application
			//----------------------------------------------------------------------------
			return UpdateApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a Application
//----------------------------------------------------------------------------
func UnassignProductFromApplication(applicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// assign an empty InsuranceProduct to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the Application
		//----------------------------------------------------------------------------
		return UpdateApplication(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Distributor on a Application
//----------------------------------------------------------------------------
func AssignDistributorToApplication( applicationId uint64, distributorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Distributor

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Distributor with a
		// matching distributorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, distributorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Distributor	to the Application
			//----------------------------------------------------------------------------
			parentObj.Distributor = &childObj

			//----------------------------------------------------------------------------
			// save the Application
			//----------------------------------------------------------------------------
			return UpdateApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Distributor", distributorId )
			return utils.RequestResult{false, msg, "assignDistributor", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Distributor on a Application
//----------------------------------------------------------------------------
func UnassignDistributorFromApplication(applicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// assign an empty Distributor to the Distributor
		//----------------------------------------------------------------------------
		parentObj.Distributor = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Distributor
		//----------------------------------------------------------------------------
		parentObj.DistributorId = nil;

		//----------------------------------------------------------------------------
		// save the Application
		//----------------------------------------------------------------------------
		return UpdateApplication(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a SelectedQuote on a Application
//----------------------------------------------------------------------------
func AssignSelectedQuoteToApplication( applicationId uint64, selectedQuoteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Quote

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Quote with a
		// matching selectedQuoteId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, selectedQuoteId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SelectedQuote	to the Application
			//----------------------------------------------------------------------------
			parentObj.SelectedQuote = &childObj

			//----------------------------------------------------------------------------
			// save the Application
			//----------------------------------------------------------------------------
			return UpdateApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SelectedQuote", selectedQuoteId )
			return utils.RequestResult{false, msg, "assignSelectedQuote", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SelectedQuote on a Application
//----------------------------------------------------------------------------
func UnassignSelectedQuoteFromApplication(applicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

		//----------------------------------------------------------------------------
		// assign an empty Quote to the SelectedQuote
		//----------------------------------------------------------------------------
		parentObj.SelectedQuote = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SelectedQuote
		//----------------------------------------------------------------------------
		parentObj.SelectedQuoteId = nil;

		//----------------------------------------------------------------------------
		// save the Application
		//----------------------------------------------------------------------------
		return UpdateApplication(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more quotesIds as a Quotes to a Application
//----------------------------------------------------------------------------
func AddQuotesToApplication ( applicationId uint64, quotesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

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
		// retrieve the modified Application from the gorm
		//----------------------------------------------------------------------------
		return GetApplication(applicationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more quotesIds as a Quotes from a Application
//----------------------------------------------------------------------------
func RemoveQuotesFromApplication( applicationId uint64, quotesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Application with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApplication(applicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Application so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Application)

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
		// retrieve the modified Application from the gorm
		//----------------------------------------------------------------------------
		return GetApplication(applicationId)

	} else {
		return parentRequestResult
	}
}

