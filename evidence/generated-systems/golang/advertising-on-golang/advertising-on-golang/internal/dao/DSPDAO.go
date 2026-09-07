package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DSPDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDSP - creates a new db entry
//----------------------------------------------------------------------------
func CreateDSP(obj model.DSP)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DSP with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DSP", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDSP", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDSP - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDSP(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DSP

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DSP with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DSP using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DSP using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDSP", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDSP - returns all
//----------------------------------------------------------------------------
func GetAllDSP()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DSP

	//----------------------------------------------------------------------------
	// Request the ORM to find all DSP
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DSP" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DSP", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDSP", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDSP - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDSP(obj model.DSP)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DSP using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DSP using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDSP", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDSP - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDSP(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DSP with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDSP(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DSP so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DSP)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DSP using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DSP using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDSP", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more adAccountsIds as a AdAccounts to a DSP
//----------------------------------------------------------------------------
func AddAdAccountsToDSP ( dSPId uint64, adAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DSP with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDSP(dSPId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DSP so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DSP)

		// slice the ids on comma with no spaces
		ids := strings.Split( adAccountsIds, ",")

		for _, adAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdAccount
			// with a matching adAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AdAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccounts", adAccountsId )
				return utils.RequestResult{false, msg, "unassignAdAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DSP from the gorm
		//----------------------------------------------------------------------------
		return GetDSP(dSPId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more adAccountsIds as a AdAccounts from a DSP
//----------------------------------------------------------------------------
func RemoveAdAccountsFromDSP( dSPId uint64, adAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DSP with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDSP(dSPId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DSP so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DSP)

		// slice the ids on comma with no spaces
		ids := strings.Split( adAccountsIds, ",")

		for _, adAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdAccount
			// with a matching adAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AdAccountObj from the AdAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccounts", adAccountsId )
				return utils.RequestResult{false, msg, "removeAdAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DSP from the gorm
		//----------------------------------------------------------------------------
		return GetDSP(dSPId)

	} else {
		return parentRequestResult
	}
}

