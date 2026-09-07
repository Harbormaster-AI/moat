package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing APIClientDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAPIClient - creates a new db entry
//----------------------------------------------------------------------------
func CreateAPIClient(obj model.APIClient)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a APIClient with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a APIClient", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAPIClient", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAPIClient - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAPIClient(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.APIClient

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a APIClient with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a APIClient using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a APIClient using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAPIClient", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAPIClient - returns all
//----------------------------------------------------------------------------
func GetAllAPIClient()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.APIClient

	//----------------------------------------------------------------------------
	// Request the ORM to find all APIClient
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all APIClient" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all APIClient", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAPIClient", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAPIClient - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAPIClient(obj model.APIClient)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a APIClient using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a APIClient using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAPIClient", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAPIClient - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAPIClient(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the APIClient with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAPIClient(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APIClient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.APIClient)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a APIClient using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a APIClient using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAPIClient", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a APIClient
//----------------------------------------------------------------------------
func AddConsentsToAPIClient ( aPIClientId uint64, consentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the APIClient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAPIClient(aPIClientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APIClient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.APIClient)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Consents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "unassignConsents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified APIClient from the gorm
		//----------------------------------------------------------------------------
		return GetAPIClient(aPIClientId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more consentsIds as a Consents from a APIClient
//----------------------------------------------------------------------------
func RemoveConsentsFromAPIClient( aPIClientId uint64, consentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the APIClient with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAPIClient(aPIClientId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APIClient so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.APIClient)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConsentObj from the Consents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "removeConsents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified APIClient from the gorm
		//----------------------------------------------------------------------------
		return GetAPIClient(aPIClientId)

	} else {
		return parentRequestResult
	}
}

