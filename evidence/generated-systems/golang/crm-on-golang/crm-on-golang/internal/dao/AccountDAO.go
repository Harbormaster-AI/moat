package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateAccount(obj model.Account)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Account with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Account", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAccount(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Account

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Account with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Account using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Account using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAccount - returns all
//----------------------------------------------------------------------------
func GetAllAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Account

	//----------------------------------------------------------------------------
	// Request the ORM to find all Account
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Account" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Account", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAccount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAccount(obj model.Account)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Account using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Account using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAccount(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Account using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Account using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAccount", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Account
//----------------------------------------------------------------------------
func AssignOrganizationToAccount( accountId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
			// assign the Organization	to the Account
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Account
			//----------------------------------------------------------------------------
			return UpdateAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Account
//----------------------------------------------------------------------------
func UnassignOrganizationFromAccount(accountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Account
		//----------------------------------------------------------------------------
		return UpdateAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ParentAccount on a Account
//----------------------------------------------------------------------------
func AssignParentAccountToAccount( accountId uint64, parentAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching parentAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, parentAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ParentAccount	to the Account
			//----------------------------------------------------------------------------
			parentObj.ParentAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Account
			//----------------------------------------------------------------------------
			return UpdateAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ParentAccount", parentAccountId )
			return utils.RequestResult{false, msg, "assignParentAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ParentAccount on a Account
//----------------------------------------------------------------------------
func UnassignParentAccountFromAccount(accountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// assign an empty Account to the ParentAccount
		//----------------------------------------------------------------------------
		parentObj.ParentAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ParentAccount
		//----------------------------------------------------------------------------
		parentObj.ParentAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Account
		//----------------------------------------------------------------------------
		return UpdateAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Account
//----------------------------------------------------------------------------
func AssignOwnerToAccount( accountId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
			// assign the Owner	to the Account
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Account
			//----------------------------------------------------------------------------
			return UpdateAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Account
//----------------------------------------------------------------------------
func UnassignOwnerFromAccount(accountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Account
		//----------------------------------------------------------------------------
		return UpdateAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Territory on a Account
//----------------------------------------------------------------------------
func AssignTerritoryToAccount( accountId uint64, territoryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Territory

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Territory with a
		// matching territoryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, territoryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Territory	to the Account
			//----------------------------------------------------------------------------
			parentObj.Territory = &childObj

			//----------------------------------------------------------------------------
			// save the Account
			//----------------------------------------------------------------------------
			return UpdateAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Territory", territoryId )
			return utils.RequestResult{false, msg, "assignTerritory", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Territory on a Account
//----------------------------------------------------------------------------
func UnassignTerritoryFromAccount(accountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		//----------------------------------------------------------------------------
		// assign an empty Territory to the Territory
		//----------------------------------------------------------------------------
		parentObj.Territory = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Territory
		//----------------------------------------------------------------------------
		parentObj.TerritoryId = nil;

		//----------------------------------------------------------------------------
		// save the Account
		//----------------------------------------------------------------------------
		return UpdateAccount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more childAccountsIds as a ChildAccounts to a Account
//----------------------------------------------------------------------------
func AddChildAccountsToAccount ( accountId uint64, childAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( childAccountsIds, ",")

		for _, childAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching childAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ChildAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildAccounts", childAccountsId )
				return utils.RequestResult{false, msg, "unassignChildAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more childAccountsIds as a ChildAccounts from a Account
//----------------------------------------------------------------------------
func RemoveChildAccountsFromAccount( accountId uint64, childAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( childAccountsIds, ",")

		for _, childAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching childAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the ChildAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildAccounts", childAccountsId )
				return utils.RequestResult{false, msg, "removeChildAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contactsIds as a Contacts to a Account
//----------------------------------------------------------------------------
func AddContactsToAccount ( accountId uint64, contactsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( contactsIds, ",")

		for _, contactsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contact

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contact
			// with a matching contactsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contactsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Contacts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contacts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contacts", contactsId )
				return utils.RequestResult{false, msg, "unassignContacts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contactsIds as a Contacts from a Account
//----------------------------------------------------------------------------
func RemoveContactsFromAccount( accountId uint64, contactsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( contactsIds, ",")

		for _, contactsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contact

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contact
			// with a matching contactsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contactsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ContactObj from the Contacts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contacts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contacts", contactsId )
				return utils.RequestResult{false, msg, "removeContacts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more opportunitiesIds as a Opportunities to a Account
//----------------------------------------------------------------------------
func AddOpportunitiesToAccount ( accountId uint64, opportunitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( opportunitiesIds, ",")

		for _, opportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching opportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , opportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Opportunities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Opportunities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunities", opportunitiesId )
				return utils.RequestResult{false, msg, "unassignOpportunities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more opportunitiesIds as a Opportunities from a Account
//----------------------------------------------------------------------------
func RemoveOpportunitiesFromAccount( accountId uint64, opportunitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( opportunitiesIds, ",")

		for _, opportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching opportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , opportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OpportunityObj from the Opportunities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Opportunities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunities", opportunitiesId )
				return utils.RequestResult{false, msg, "removeOpportunities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more casesIds as a Cases to a Account
//----------------------------------------------------------------------------
func AddCasesToAccount ( accountId uint64, casesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more casesIds as a Cases from a Account
//----------------------------------------------------------------------------
func RemoveCasesFromAccount( accountId uint64, casesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more activitiesIds as a Activities to a Account
//----------------------------------------------------------------------------
func AddActivitiesToAccount ( accountId uint64, activitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Activities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "unassignActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more activitiesIds as a Activities from a Account
//----------------------------------------------------------------------------
func RemoveActivitiesFromAccount( accountId uint64, activitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ActivityObj from the Activities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "removeActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a Account
//----------------------------------------------------------------------------
func AddCampaignsToAccount ( accountId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( campaignsIds, ",")

		for _, campaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Campaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Campaign
			// with a matching campaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , campaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Campaigns using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Campaigns").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaigns", campaignsId )
				return utils.RequestResult{false, msg, "unassignCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a Account
//----------------------------------------------------------------------------
func RemoveCampaignsFromAccount( accountId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( campaignsIds, ",")

		for _, campaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Campaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Campaign
			// with a matching campaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , campaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CampaignObj from the Campaigns array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Campaigns").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaigns", campaignsId )
				return utils.RequestResult{false, msg, "removeCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more quotesIds as a Quotes to a Account
//----------------------------------------------------------------------------
func AddQuotesToAccount ( accountId uint64, quotesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more quotesIds as a Quotes from a Account
//----------------------------------------------------------------------------
func RemoveQuotesFromAccount( accountId uint64, quotesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ordersIds as a Orders to a Account
//----------------------------------------------------------------------------
func AddOrdersToAccount ( accountId uint64, ordersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ordersIds as a Orders from a Account
//----------------------------------------------------------------------------
func RemoveOrdersFromAccount( accountId uint64, ordersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

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
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contractsIds as a Contracts to a Account
//----------------------------------------------------------------------------
func AddContractsToAccount ( accountId uint64, contractsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Contracts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "unassignContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contractsIds as a Contracts from a Account
//----------------------------------------------------------------------------
func RemoveContractsFromAccount( accountId uint64, contractsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ContractObj from the Contracts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "removeContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more notesIds as a Notes to a Account
//----------------------------------------------------------------------------
func AddNotesToAccount ( accountId uint64, notesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( notesIds, ",")

		for _, notesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Note

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Note
			// with a matching notesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Notes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notes", notesId )
				return utils.RequestResult{false, msg, "unassignNotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more notesIds as a Notes from a Account
//----------------------------------------------------------------------------
func RemoveNotesFromAccount( accountId uint64, notesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( notesIds, ",")

		for _, notesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Note

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Note
			// with a matching notesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove NoteObj from the Notes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notes", notesId )
				return utils.RequestResult{false, msg, "removeNotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more emailMessagesIds as a EmailMessages to a Account
//----------------------------------------------------------------------------
func AddEmailMessagesToAccount ( accountId uint64, emailMessagesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( emailMessagesIds, ",")

		for _, emailMessagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailMessagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailMessagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EmailMessages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmailMessages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmailMessages", emailMessagesId )
				return utils.RequestResult{false, msg, "unassignEmailMessages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more emailMessagesIds as a EmailMessages from a Account
//----------------------------------------------------------------------------
func RemoveEmailMessagesFromAccount( accountId uint64, emailMessagesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Account with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccount(accountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Account so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Account)

		// slice the ids on comma with no spaces
		ids := strings.Split( emailMessagesIds, ",")

		for _, emailMessagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailMessagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailMessagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmailMessageObj from the EmailMessages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmailMessages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmailMessages", emailMessagesId )
				return utils.RequestResult{false, msg, "removeEmailMessages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Account from the gorm
		//----------------------------------------------------------------------------
		return GetAccount(accountId)

	} else {
		return parentRequestResult
	}
}

