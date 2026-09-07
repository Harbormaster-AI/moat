package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AdjusterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAdjuster - creates a new db entry
//----------------------------------------------------------------------------
func CreateAdjuster(obj model.Adjuster)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Adjuster with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Adjuster", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAdjuster", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAdjuster - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAdjuster(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Adjuster

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Adjuster with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Adjuster using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Adjuster using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAdjuster", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAdjuster - returns all
//----------------------------------------------------------------------------
func GetAllAdjuster()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Adjuster

	//----------------------------------------------------------------------------
	// Request the ORM to find all Adjuster
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Adjuster" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Adjuster", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAdjuster", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAdjuster - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAdjuster(obj model.Adjuster)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Adjuster using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Adjuster using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAdjuster", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAdjuster - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAdjuster(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Adjuster with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAdjuster(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Adjuster so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Adjuster)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Adjuster using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Adjuster using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAdjuster", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more claimsIds as a Claims to a Adjuster
//----------------------------------------------------------------------------
func AddClaimsToAdjuster ( adjusterId uint64, claimsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Adjuster with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdjuster(adjusterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Adjuster so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Adjuster)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Claims using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "unassignClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Adjuster from the gorm
		//----------------------------------------------------------------------------
		return GetAdjuster(adjusterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a Adjuster
//----------------------------------------------------------------------------
func RemoveClaimsFromAdjuster( adjusterId uint64, claimsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Adjuster with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdjuster(adjusterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Adjuster so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Adjuster)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimObj from the Claims array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "removeClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Adjuster from the gorm
		//----------------------------------------------------------------------------
		return GetAdjuster(adjusterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more serviceProvidersIds as a ServiceProviders to a Adjuster
//----------------------------------------------------------------------------
func AddServiceProvidersToAdjuster ( adjusterId uint64, serviceProvidersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Adjuster with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdjuster(adjusterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Adjuster so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Adjuster)

		// slice the ids on comma with no spaces
		ids := strings.Split( serviceProvidersIds, ",")

		for _, serviceProvidersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ServiceProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ServiceProvider
			// with a matching serviceProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serviceProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ServiceProviders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ServiceProviders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ServiceProviders", serviceProvidersId )
				return utils.RequestResult{false, msg, "unassignServiceProviders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Adjuster from the gorm
		//----------------------------------------------------------------------------
		return GetAdjuster(adjusterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serviceProvidersIds as a ServiceProviders from a Adjuster
//----------------------------------------------------------------------------
func RemoveServiceProvidersFromAdjuster( adjusterId uint64, serviceProvidersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Adjuster with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdjuster(adjusterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Adjuster so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Adjuster)

		// slice the ids on comma with no spaces
		ids := strings.Split( serviceProvidersIds, ",")

		for _, serviceProvidersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ServiceProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ServiceProvider
			// with a matching serviceProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serviceProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ServiceProviderObj from the ServiceProviders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ServiceProviders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ServiceProviders", serviceProvidersId )
				return utils.RequestResult{false, msg, "removeServiceProviders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Adjuster from the gorm
		//----------------------------------------------------------------------------
		return GetAdjuster(adjusterId)

	} else {
		return parentRequestResult
	}
}

