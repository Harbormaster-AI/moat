package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AdvertiserDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAdvertiser - creates a new db entry
//----------------------------------------------------------------------------
func CreateAdvertiser(obj model.Advertiser)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Advertiser with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Advertiser", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAdvertiser", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAdvertiser - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAdvertiser(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Advertiser

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Advertiser with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Advertiser using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Advertiser using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAdvertiser", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAdvertiser - returns all
//----------------------------------------------------------------------------
func GetAllAdvertiser()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Advertiser

	//----------------------------------------------------------------------------
	// Request the ORM to find all Advertiser
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Advertiser" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Advertiser", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAdvertiser", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAdvertiser - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAdvertiser(obj model.Advertiser)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Advertiser using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Advertiser using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAdvertiser", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAdvertiser - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAdvertiser(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAdvertiser(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Advertiser)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Advertiser using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Advertiser using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAdvertiser", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Agency on a Advertiser
//----------------------------------------------------------------------------
func AssignAgencyToAdvertiser( advertiserId uint64, agencyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Agency

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Agency with a
		// matching agencyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, agencyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Agency	to the Advertiser
			//----------------------------------------------------------------------------
			parentObj.Agency = &childObj

			//----------------------------------------------------------------------------
			// save the Advertiser
			//----------------------------------------------------------------------------
			return UpdateAdvertiser(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agency", agencyId )
			return utils.RequestResult{false, msg, "assignAgency", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Agency on a Advertiser
//----------------------------------------------------------------------------
func UnassignAgencyFromAdvertiser(advertiserId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

		//----------------------------------------------------------------------------
		// assign an empty Agency to the Agency
		//----------------------------------------------------------------------------
		parentObj.Agency = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Agency
		//----------------------------------------------------------------------------
		parentObj.AgencyId = nil;

		//----------------------------------------------------------------------------
		// save the Advertiser
		//----------------------------------------------------------------------------
		return UpdateAdvertiser(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more adAccountsIds as a AdAccounts to a Advertiser
//----------------------------------------------------------------------------
func AddAdAccountsToAdvertiser ( advertiserId uint64, adAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

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
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more adAccountsIds as a AdAccounts from a Advertiser
//----------------------------------------------------------------------------
func RemoveAdAccountsFromAdvertiser( advertiserId uint64, adAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

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
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more billingProfilesIds as a BillingProfiles to a Advertiser
//----------------------------------------------------------------------------
func AddBillingProfilesToAdvertiser ( advertiserId uint64, billingProfilesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

		// slice the ids on comma with no spaces
		ids := strings.Split( billingProfilesIds, ",")

		for _, billingProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BillingProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BillingProfile
			// with a matching billingProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , billingProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the BillingProfiles using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BillingProfiles").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BillingProfiles", billingProfilesId )
				return utils.RequestResult{false, msg, "unassignBillingProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more billingProfilesIds as a BillingProfiles from a Advertiser
//----------------------------------------------------------------------------
func RemoveBillingProfilesFromAdvertiser( advertiserId uint64, billingProfilesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

		// slice the ids on comma with no spaces
		ids := strings.Split( billingProfilesIds, ",")

		for _, billingProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BillingProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BillingProfile
			// with a matching billingProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , billingProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BillingProfileObj from the BillingProfiles array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BillingProfiles").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BillingProfiles", billingProfilesId )
				return utils.RequestResult{false, msg, "removeBillingProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a Advertiser
//----------------------------------------------------------------------------
func AddCampaignsToAdvertiser ( advertiserId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

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
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a Advertiser
//----------------------------------------------------------------------------
func RemoveCampaignsFromAdvertiser( advertiserId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

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
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more trackingPixelsIds as a TrackingPixels to a Advertiser
//----------------------------------------------------------------------------
func AddTrackingPixelsToAdvertiser ( advertiserId uint64, trackingPixelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

		// slice the ids on comma with no spaces
		ids := strings.Split( trackingPixelsIds, ",")

		for _, trackingPixelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrackingPixel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrackingPixel
			// with a matching trackingPixelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trackingPixelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the TrackingPixels using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrackingPixels").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrackingPixels", trackingPixelsId )
				return utils.RequestResult{false, msg, "unassignTrackingPixels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more trackingPixelsIds as a TrackingPixels from a Advertiser
//----------------------------------------------------------------------------
func RemoveTrackingPixelsFromAdvertiser( advertiserId uint64, trackingPixelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Advertiser with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdvertiser(advertiserId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Advertiser so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Advertiser)

		// slice the ids on comma with no spaces
		ids := strings.Split( trackingPixelsIds, ",")

		for _, trackingPixelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrackingPixel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrackingPixel
			// with a matching trackingPixelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trackingPixelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TrackingPixelObj from the TrackingPixels array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrackingPixels").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrackingPixels", trackingPixelsId )
				return utils.RequestResult{false, msg, "removeTrackingPixels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Advertiser from the gorm
		//----------------------------------------------------------------------------
		return GetAdvertiser(advertiserId)

	} else {
		return parentRequestResult
	}
}

