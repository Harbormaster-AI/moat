package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ServiceProviderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateServiceProvider - creates a new db entry
//----------------------------------------------------------------------------
func CreateServiceProvider(obj model.ServiceProvider)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ServiceProvider with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ServiceProvider", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateServiceProvider", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetServiceProvider - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetServiceProvider(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ServiceProvider

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ServiceProvider with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ServiceProvider using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ServiceProvider using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetServiceProvider", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllServiceProvider - returns all
//----------------------------------------------------------------------------
func GetAllServiceProvider()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ServiceProvider

	//----------------------------------------------------------------------------
	// Request the ORM to find all ServiceProvider
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ServiceProvider" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ServiceProvider", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllServiceProvider", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateServiceProvider - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateServiceProvider(obj model.ServiceProvider)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ServiceProvider using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ServiceProvider using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateServiceProvider", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteServiceProvider - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteServiceProvider(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ServiceProvider with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetServiceProvider(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ServiceProvider)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ServiceProvider using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ServiceProvider using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteServiceProvider", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more claimsIds as a Claims to a ServiceProvider
//----------------------------------------------------------------------------
func AddClaimsToServiceProvider ( serviceProviderId uint64, claimsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ServiceProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetServiceProvider(serviceProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ServiceProvider)

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
		// retrieve the modified ServiceProvider from the gorm
		//----------------------------------------------------------------------------
		return GetServiceProvider(serviceProviderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a ServiceProvider
//----------------------------------------------------------------------------
func RemoveClaimsFromServiceProvider( serviceProviderId uint64, claimsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ServiceProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetServiceProvider(serviceProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ServiceProvider)

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
		// retrieve the modified ServiceProvider from the gorm
		//----------------------------------------------------------------------------
		return GetServiceProvider(serviceProviderId)

	} else {
		return parentRequestResult
	}
}

