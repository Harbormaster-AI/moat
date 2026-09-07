package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ContractDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateContract - creates a new db entry
//----------------------------------------------------------------------------
func CreateContract(obj model.Contract)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Contract with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Contract", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateContract", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetContract - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetContract(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Contract

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Contract with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Contract using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Contract using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetContract", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllContract - returns all
//----------------------------------------------------------------------------
func GetAllContract()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Contract

	//----------------------------------------------------------------------------
	// Request the ORM to find all Contract
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Contract" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Contract", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllContract", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateContract - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateContract(obj model.Contract)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Contract using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Contract using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateContract", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteContract - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteContract(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetContract(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Contract using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Contract using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteContract", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Contract
//----------------------------------------------------------------------------
func AssignOrganizationToContract( contractId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

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
			// assign the Organization	to the Contract
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Contract
			//----------------------------------------------------------------------------
			return UpdateContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Contract
//----------------------------------------------------------------------------
func UnassignOrganizationFromContract(contractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Contract
		//----------------------------------------------------------------------------
		return UpdateContract(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Contract
//----------------------------------------------------------------------------
func AssignAccountToContract( contractId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching accountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, accountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Account	to the Contract
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Contract
			//----------------------------------------------------------------------------
			return UpdateContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Contract
//----------------------------------------------------------------------------
func UnassignAccountFromContract(contractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Contract
		//----------------------------------------------------------------------------
		return UpdateContract(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Contract
//----------------------------------------------------------------------------
func AssignOwnerToContract( contractId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.User

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a User with a
		// matching ownerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, ownerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Owner	to the Contract
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Contract
			//----------------------------------------------------------------------------
			return UpdateContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Contract
//----------------------------------------------------------------------------
func UnassignOwnerFromContract(contractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Contract
		//----------------------------------------------------------------------------
		return UpdateContract(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more ordersIds as a Orders to a Contract
//----------------------------------------------------------------------------
func AddOrdersToContract ( contractId uint64, ordersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

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
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ordersIds as a Orders from a Contract
//----------------------------------------------------------------------------
func RemoveOrdersFromContract( contractId uint64, ordersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

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
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more casesIds as a Cases to a Contract
//----------------------------------------------------------------------------
func AddCasesToContract ( contractId uint64, casesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		// slice the ids on comma with no spaces
		ids := strings.Split( casesIds, ",")

		for _, casesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Case_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Case_
			// with a matching casesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , casesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Cases using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Cases").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cases", casesId )
				return utils.RequestResult{false, msg, "unassignCases", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more casesIds as a Cases from a Contract
//----------------------------------------------------------------------------
func RemoveCasesFromContract( contractId uint64, casesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		// slice the ids on comma with no spaces
		ids := strings.Split( casesIds, ",")

		for _, casesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Case_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Case_
			// with a matching casesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , casesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Case_Obj from the Cases array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Cases").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cases", casesId )
				return utils.RequestResult{false, msg, "removeCases", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

