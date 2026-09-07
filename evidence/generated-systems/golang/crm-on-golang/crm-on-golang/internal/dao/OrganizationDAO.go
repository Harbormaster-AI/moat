package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OrganizationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOrganization - creates a new db entry
//----------------------------------------------------------------------------
func CreateOrganization(obj model.Organization)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Organization with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Organization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOrganization", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOrganization - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOrganization(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Organization

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Organization with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Organization using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Organization using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOrganization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOrganization - returns all
//----------------------------------------------------------------------------
func GetAllOrganization()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Organization

	//----------------------------------------------------------------------------
	// Request the ORM to find all Organization
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Organization" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Organization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOrganization", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOrganization - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOrganization(obj model.Organization)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Organization using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Organization using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOrganization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOrganization - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOrganization(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOrganization(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Organization)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Organization using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Organization using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOrganization", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more usersIds as a Users to a Organization
//----------------------------------------------------------------------------
func AddUsersToOrganization ( organizationId uint64, usersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( usersIds, ",")

		for _, usersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.User

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a User
			// with a matching usersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , usersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Users using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Users").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Users", usersId )
				return utils.RequestResult{false, msg, "unassignUsers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more usersIds as a Users from a Organization
//----------------------------------------------------------------------------
func RemoveUsersFromOrganization( organizationId uint64, usersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( usersIds, ",")

		for _, usersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.User

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a User
			// with a matching usersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , usersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove UserObj from the Users array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Users").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Users", usersId )
				return utils.RequestResult{false, msg, "removeUsers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Organization
//----------------------------------------------------------------------------
func AddAccountsToOrganization ( organizationId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Accounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a Organization
//----------------------------------------------------------------------------
func RemoveAccountsFromOrganization( organizationId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the Accounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more teamsIds as a Teams to a Organization
//----------------------------------------------------------------------------
func AddTeamsToOrganization ( organizationId uint64, teamsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( teamsIds, ",")

		for _, teamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Team

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Team
			// with a matching teamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , teamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Teams using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Teams").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Teams", teamsId )
				return utils.RequestResult{false, msg, "unassignTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more teamsIds as a Teams from a Organization
//----------------------------------------------------------------------------
func RemoveTeamsFromOrganization( organizationId uint64, teamsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( teamsIds, ",")

		for _, teamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Team

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Team
			// with a matching teamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , teamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TeamObj from the Teams array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Teams").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Teams", teamsId )
				return utils.RequestResult{false, msg, "removeTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more territoriesIds as a Territories to a Organization
//----------------------------------------------------------------------------
func AddTerritoriesToOrganization ( organizationId uint64, territoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( territoriesIds, ",")

		for _, territoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Territory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Territory
			// with a matching territoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , territoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Territories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Territories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Territories", territoriesId )
				return utils.RequestResult{false, msg, "unassignTerritories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more territoriesIds as a Territories from a Organization
//----------------------------------------------------------------------------
func RemoveTerritoriesFromOrganization( organizationId uint64, territoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( territoriesIds, ",")

		for _, territoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Territory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Territory
			// with a matching territoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , territoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TerritoryObj from the Territories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Territories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Territories", territoriesId )
				return utils.RequestResult{false, msg, "removeTerritories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more productsIds as a Products to a Organization
//----------------------------------------------------------------------------
func AddProductsToOrganization ( organizationId uint64, productsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( productsIds, ",")

		for _, productsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Product

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Product
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Products using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Products").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )
				return utils.RequestResult{false, msg, "unassignProducts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productsIds as a Products from a Organization
//----------------------------------------------------------------------------
func RemoveProductsFromOrganization( organizationId uint64, productsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( productsIds, ",")

		for _, productsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Product

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Product
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProductObj from the Products array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Products").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )
				return utils.RequestResult{false, msg, "removeProducts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more priceBooksIds as a PriceBooks to a Organization
//----------------------------------------------------------------------------
func AddPriceBooksToOrganization ( organizationId uint64, priceBooksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( priceBooksIds, ",")

		for _, priceBooksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PriceBook

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PriceBook
			// with a matching priceBooksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , priceBooksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PriceBooks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PriceBooks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBooks", priceBooksId )
				return utils.RequestResult{false, msg, "unassignPriceBooks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more priceBooksIds as a PriceBooks from a Organization
//----------------------------------------------------------------------------
func RemovePriceBooksFromOrganization( organizationId uint64, priceBooksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( priceBooksIds, ",")

		for _, priceBooksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PriceBook

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PriceBook
			// with a matching priceBooksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , priceBooksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PriceBookObj from the PriceBooks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PriceBooks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBooks", priceBooksId )
				return utils.RequestResult{false, msg, "removePriceBooks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a Organization
//----------------------------------------------------------------------------
func AddCampaignsToOrganization ( organizationId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

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
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a Organization
//----------------------------------------------------------------------------
func RemoveCampaignsFromOrganization( organizationId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

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
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

