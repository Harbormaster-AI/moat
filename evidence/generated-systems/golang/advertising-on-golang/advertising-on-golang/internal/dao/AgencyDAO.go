package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AgencyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAgency - creates a new db entry
//----------------------------------------------------------------------------
func CreateAgency(obj model.Agency)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Agency with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Agency", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAgency", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAgency - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAgency(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Agency

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Agency with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Agency using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Agency using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAgency", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAgency - returns all
//----------------------------------------------------------------------------
func GetAllAgency()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Agency

	//----------------------------------------------------------------------------
	// Request the ORM to find all Agency
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Agency" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Agency", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAgency", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAgency - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAgency(obj model.Agency)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Agency using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Agency using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAgency", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAgency - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAgency(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAgency(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Agency)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Agency using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Agency using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAgency", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more advertisersIds as a Advertisers to a Agency
//----------------------------------------------------------------------------
func AddAdvertisersToAgency ( agencyId uint64, advertisersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

		// slice the ids on comma with no spaces
		ids := strings.Split( advertisersIds, ",")

		for _, advertisersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Advertiser

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Advertiser
			// with a matching advertisersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , advertisersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Advertisers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Advertisers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Advertisers", advertisersId )
				return utils.RequestResult{false, msg, "unassignAdvertisers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more advertisersIds as a Advertisers from a Agency
//----------------------------------------------------------------------------
func RemoveAdvertisersFromAgency( agencyId uint64, advertisersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

		// slice the ids on comma with no spaces
		ids := strings.Split( advertisersIds, ",")

		for _, advertisersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Advertiser

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Advertiser
			// with a matching advertisersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , advertisersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AdvertiserObj from the Advertisers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Advertisers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Advertisers", advertisersId )
				return utils.RequestResult{false, msg, "removeAdvertisers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more teamsIds as a Teams to a Agency
//----------------------------------------------------------------------------
func AddTeamsToAgency ( agencyId uint64, teamsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

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
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more teamsIds as a Teams from a Agency
//----------------------------------------------------------------------------
func RemoveTeamsFromAgency( agencyId uint64, teamsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

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
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more usersIds as a Users to a Agency
//----------------------------------------------------------------------------
func AddUsersToAgency ( agencyId uint64, usersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

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
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more usersIds as a Users from a Agency
//----------------------------------------------------------------------------
func RemoveUsersFromAgency( agencyId uint64, usersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

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
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more insertionOrdersIds as a InsertionOrders to a Agency
//----------------------------------------------------------------------------
func AddInsertionOrdersToAgency ( agencyId uint64, insertionOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

		// slice the ids on comma with no spaces
		ids := strings.Split( insertionOrdersIds, ",")

		for _, insertionOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsertionOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsertionOrder
			// with a matching insertionOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insertionOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InsertionOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsertionOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsertionOrders", insertionOrdersId )
				return utils.RequestResult{false, msg, "unassignInsertionOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more insertionOrdersIds as a InsertionOrders from a Agency
//----------------------------------------------------------------------------
func RemoveInsertionOrdersFromAgency( agencyId uint64, insertionOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Agency with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgency(agencyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agency so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agency)

		// slice the ids on comma with no spaces
		ids := strings.Split( insertionOrdersIds, ",")

		for _, insertionOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsertionOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsertionOrder
			// with a matching insertionOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insertionOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsertionOrderObj from the InsertionOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsertionOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsertionOrders", insertionOrdersId )
				return utils.RequestResult{false, msg, "removeInsertionOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Agency from the gorm
		//----------------------------------------------------------------------------
		return GetAgency(agencyId)

	} else {
		return parentRequestResult
	}
}

