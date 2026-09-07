package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TargetingProfileDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTargetingProfile - creates a new db entry
//----------------------------------------------------------------------------
func CreateTargetingProfile(obj model.TargetingProfile)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TargetingProfile with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TargetingProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTargetingProfile", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTargetingProfile - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTargetingProfile(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TargetingProfile

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TargetingProfile with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TargetingProfile using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TargetingProfile using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTargetingProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTargetingProfile - returns all
//----------------------------------------------------------------------------
func GetAllTargetingProfile()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TargetingProfile

	//----------------------------------------------------------------------------
	// Request the ORM to find all TargetingProfile
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TargetingProfile" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TargetingProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTargetingProfile", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTargetingProfile - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTargetingProfile(obj model.TargetingProfile)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TargetingProfile using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TargetingProfile using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTargetingProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTargetingProfile - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTargetingProfile(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTargetingProfile(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TargetingProfile)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TargetingProfile using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TargetingProfile using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTargetingProfile", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a BrandSafetyPolicy on a TargetingProfile
//----------------------------------------------------------------------------
func AssignBrandSafetyPolicyToTargetingProfile( targetingProfileId uint64, brandSafetyPolicyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BrandSafetyPolicy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BrandSafetyPolicy with a
		// matching brandSafetyPolicyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, brandSafetyPolicyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BrandSafetyPolicy	to the TargetingProfile
			//----------------------------------------------------------------------------
			parentObj.BrandSafetyPolicy = &childObj

			//----------------------------------------------------------------------------
			// save the TargetingProfile
			//----------------------------------------------------------------------------
			return UpdateTargetingProfile(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BrandSafetyPolicy", brandSafetyPolicyId )
			return utils.RequestResult{false, msg, "assignBrandSafetyPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BrandSafetyPolicy on a TargetingProfile
//----------------------------------------------------------------------------
func UnassignBrandSafetyPolicyFromTargetingProfile(targetingProfileId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		//----------------------------------------------------------------------------
		// assign an empty BrandSafetyPolicy to the BrandSafetyPolicy
		//----------------------------------------------------------------------------
		parentObj.BrandSafetyPolicy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BrandSafetyPolicy
		//----------------------------------------------------------------------------
		parentObj.BrandSafetyPolicyId = nil;

		//----------------------------------------------------------------------------
		// save the TargetingProfile
		//----------------------------------------------------------------------------
		return UpdateTargetingProfile(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more audienceSegmentsIds as a AudienceSegments to a TargetingProfile
//----------------------------------------------------------------------------
func AddAudienceSegmentsToTargetingProfile ( targetingProfileId uint64, audienceSegmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( audienceSegmentsIds, ",")

		for _, audienceSegmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AudienceSegment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AudienceSegment
			// with a matching audienceSegmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , audienceSegmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AudienceSegments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AudienceSegments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AudienceSegments", audienceSegmentsId )
				return utils.RequestResult{false, msg, "unassignAudienceSegments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more audienceSegmentsIds as a AudienceSegments from a TargetingProfile
//----------------------------------------------------------------------------
func RemoveAudienceSegmentsFromTargetingProfile( targetingProfileId uint64, audienceSegmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( audienceSegmentsIds, ",")

		for _, audienceSegmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AudienceSegment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AudienceSegment
			// with a matching audienceSegmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , audienceSegmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AudienceSegmentObj from the AudienceSegments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AudienceSegments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AudienceSegments", audienceSegmentsId )
				return utils.RequestResult{false, msg, "removeAudienceSegments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more geoRegionsIds as a GeoRegions to a TargetingProfile
//----------------------------------------------------------------------------
func AddGeoRegionsToTargetingProfile ( targetingProfileId uint64, geoRegionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( geoRegionsIds, ",")

		for _, geoRegionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GeoRegion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GeoRegion
			// with a matching geoRegionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , geoRegionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the GeoRegions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GeoRegions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GeoRegions", geoRegionsId )
				return utils.RequestResult{false, msg, "unassignGeoRegions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more geoRegionsIds as a GeoRegions from a TargetingProfile
//----------------------------------------------------------------------------
func RemoveGeoRegionsFromTargetingProfile( targetingProfileId uint64, geoRegionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( geoRegionsIds, ",")

		for _, geoRegionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GeoRegion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GeoRegion
			// with a matching geoRegionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , geoRegionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GeoRegionObj from the GeoRegions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GeoRegions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GeoRegions", geoRegionsId )
				return utils.RequestResult{false, msg, "removeGeoRegions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contentCategoriesIds as a ContentCategories to a TargetingProfile
//----------------------------------------------------------------------------
func AddContentCategoriesToTargetingProfile ( targetingProfileId uint64, contentCategoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( contentCategoriesIds, ",")

		for _, contentCategoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ContentCategory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ContentCategory
			// with a matching contentCategoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contentCategoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ContentCategories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ContentCategories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ContentCategories", contentCategoriesId )
				return utils.RequestResult{false, msg, "unassignContentCategories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contentCategoriesIds as a ContentCategories from a TargetingProfile
//----------------------------------------------------------------------------
func RemoveContentCategoriesFromTargetingProfile( targetingProfileId uint64, contentCategoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( contentCategoriesIds, ",")

		for _, contentCategoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ContentCategory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ContentCategory
			// with a matching contentCategoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contentCategoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ContentCategoryObj from the ContentCategories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ContentCategories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ContentCategories", contentCategoriesId )
				return utils.RequestResult{false, msg, "removeContentCategories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more deviceCriteriaIds as a DeviceCriteria to a TargetingProfile
//----------------------------------------------------------------------------
func AddDeviceCriteriaToTargetingProfile ( targetingProfileId uint64, deviceCriteriaIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( deviceCriteriaIds, ",")

		for _, deviceCriteriaId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DeviceCriterion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DeviceCriterion
			// with a matching deviceCriteriaId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , deviceCriteriaId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DeviceCriteria using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DeviceCriteria").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DeviceCriteria", deviceCriteriaId )
				return utils.RequestResult{false, msg, "unassignDeviceCriteria", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more deviceCriteriaIds as a DeviceCriteria from a TargetingProfile
//----------------------------------------------------------------------------
func RemoveDeviceCriteriaFromTargetingProfile( targetingProfileId uint64, deviceCriteriaIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TargetingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTargetingProfile(targetingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TargetingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TargetingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( deviceCriteriaIds, ",")

		for _, deviceCriteriaId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DeviceCriterion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DeviceCriterion
			// with a matching deviceCriteriaId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , deviceCriteriaId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DeviceCriterionObj from the DeviceCriteria array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DeviceCriteria").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DeviceCriteria", deviceCriteriaId )
				return utils.RequestResult{false, msg, "removeDeviceCriteria", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TargetingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetTargetingProfile(targetingProfileId)

	} else {
		return parentRequestResult
	}
}

