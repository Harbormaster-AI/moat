package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BrandSafetyPolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBrandSafetyPolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreateBrandSafetyPolicy(obj model.BrandSafetyPolicy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BrandSafetyPolicy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BrandSafetyPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBrandSafetyPolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBrandSafetyPolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBrandSafetyPolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BrandSafetyPolicy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BrandSafetyPolicy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BrandSafetyPolicy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BrandSafetyPolicy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBrandSafetyPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBrandSafetyPolicy - returns all
//----------------------------------------------------------------------------
func GetAllBrandSafetyPolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BrandSafetyPolicy

	//----------------------------------------------------------------------------
	// Request the ORM to find all BrandSafetyPolicy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BrandSafetyPolicy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BrandSafetyPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBrandSafetyPolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBrandSafetyPolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBrandSafetyPolicy(obj model.BrandSafetyPolicy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BrandSafetyPolicy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BrandSafetyPolicy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBrandSafetyPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBrandSafetyPolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBrandSafetyPolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BrandSafetyPolicy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBrandSafetyPolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BrandSafetyPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BrandSafetyPolicy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BrandSafetyPolicy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BrandSafetyPolicy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBrandSafetyPolicy", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more targetingProfilesIds as a TargetingProfiles to a BrandSafetyPolicy
//----------------------------------------------------------------------------
func AddTargetingProfilesToBrandSafetyPolicy ( brandSafetyPolicyId uint64, targetingProfilesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BrandSafetyPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBrandSafetyPolicy(brandSafetyPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BrandSafetyPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BrandSafetyPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( targetingProfilesIds, ",")

		for _, targetingProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TargetingProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TargetingProfile
			// with a matching targetingProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , targetingProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the TargetingProfiles using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TargetingProfiles").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TargetingProfiles", targetingProfilesId )
				return utils.RequestResult{false, msg, "unassignTargetingProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BrandSafetyPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetBrandSafetyPolicy(brandSafetyPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more targetingProfilesIds as a TargetingProfiles from a BrandSafetyPolicy
//----------------------------------------------------------------------------
func RemoveTargetingProfilesFromBrandSafetyPolicy( brandSafetyPolicyId uint64, targetingProfilesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BrandSafetyPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBrandSafetyPolicy(brandSafetyPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BrandSafetyPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BrandSafetyPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( targetingProfilesIds, ",")

		for _, targetingProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TargetingProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TargetingProfile
			// with a matching targetingProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , targetingProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TargetingProfileObj from the TargetingProfiles array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TargetingProfiles").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TargetingProfiles", targetingProfilesId )
				return utils.RequestResult{false, msg, "removeTargetingProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BrandSafetyPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetBrandSafetyPolicy(brandSafetyPolicyId)

	} else {
		return parentRequestResult
	}
}

