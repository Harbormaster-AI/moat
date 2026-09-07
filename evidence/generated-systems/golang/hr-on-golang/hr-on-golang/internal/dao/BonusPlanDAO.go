package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BonusPlanDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBonusPlan - creates a new db entry
//----------------------------------------------------------------------------
func CreateBonusPlan(obj model.BonusPlan)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BonusPlan with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BonusPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBonusPlan", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBonusPlan - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBonusPlan(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BonusPlan

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BonusPlan with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BonusPlan using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BonusPlan using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBonusPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBonusPlan - returns all
//----------------------------------------------------------------------------
func GetAllBonusPlan()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BonusPlan

	//----------------------------------------------------------------------------
	// Request the ORM to find all BonusPlan
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BonusPlan" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BonusPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBonusPlan", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBonusPlan - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBonusPlan(obj model.BonusPlan)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BonusPlan using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BonusPlan using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBonusPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBonusPlan - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBonusPlan(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BonusPlan with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBonusPlan(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BonusPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BonusPlan)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BonusPlan using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BonusPlan using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBonusPlan", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more compensationPackagesIds as a CompensationPackages to a BonusPlan
//----------------------------------------------------------------------------
func AddCompensationPackagesToBonusPlan ( bonusPlanId uint64, compensationPackagesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BonusPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBonusPlan(bonusPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BonusPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BonusPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( compensationPackagesIds, ",")

		for _, compensationPackagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompensationPackage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompensationPackage
			// with a matching compensationPackagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , compensationPackagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CompensationPackages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompensationPackages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompensationPackages", compensationPackagesId )
				return utils.RequestResult{false, msg, "unassignCompensationPackages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BonusPlan from the gorm
		//----------------------------------------------------------------------------
		return GetBonusPlan(bonusPlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more compensationPackagesIds as a CompensationPackages from a BonusPlan
//----------------------------------------------------------------------------
func RemoveCompensationPackagesFromBonusPlan( bonusPlanId uint64, compensationPackagesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BonusPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBonusPlan(bonusPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BonusPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BonusPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( compensationPackagesIds, ",")

		for _, compensationPackagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompensationPackage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompensationPackage
			// with a matching compensationPackagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , compensationPackagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CompensationPackageObj from the CompensationPackages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompensationPackages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompensationPackages", compensationPackagesId )
				return utils.RequestResult{false, msg, "removeCompensationPackages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BonusPlan from the gorm
		//----------------------------------------------------------------------------
		return GetBonusPlan(bonusPlanId)

	} else {
		return parentRequestResult
	}
}

