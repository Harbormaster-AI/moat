package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AdAccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAdAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateAdAccount(obj model.AdAccount)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AdAccount with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AdAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAdAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAdAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAdAccount(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AdAccount

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AdAccount with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AdAccount using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AdAccount using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAdAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAdAccount - returns all
//----------------------------------------------------------------------------
func GetAllAdAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AdAccount

	//----------------------------------------------------------------------------
	// Request the ORM to find all AdAccount
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AdAccount" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AdAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAdAccount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAdAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAdAccount(obj model.AdAccount)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AdAccount using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AdAccount using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAdAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAdAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAdAccount(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAdAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AdAccount)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AdAccount using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AdAccount using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAdAccount", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Advertiser on a AdAccount
//----------------------------------------------------------------------------
func AssignAdvertiserToAdAccount( adAccountId uint64, advertiserId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Advertiser

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Advertiser with a
		// matching advertiserId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, advertiserId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Advertiser	to the AdAccount
			//----------------------------------------------------------------------------
			parentObj.Advertiser = &childObj

			//----------------------------------------------------------------------------
			// save the AdAccount
			//----------------------------------------------------------------------------
			return UpdateAdAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Advertiser", advertiserId )
			return utils.RequestResult{false, msg, "assignAdvertiser", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Advertiser on a AdAccount
//----------------------------------------------------------------------------
func UnassignAdvertiserFromAdAccount(adAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		//----------------------------------------------------------------------------
		// assign an empty Advertiser to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.Advertiser = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.AdvertiserId = nil;

		//----------------------------------------------------------------------------
		// save the AdAccount
		//----------------------------------------------------------------------------
		return UpdateAdAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a BillingProfile on a AdAccount
//----------------------------------------------------------------------------
func AssignBillingProfileToAdAccount( adAccountId uint64, billingProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BillingProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BillingProfile with a
		// matching billingProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, billingProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BillingProfile	to the AdAccount
			//----------------------------------------------------------------------------
			parentObj.BillingProfile = &childObj

			//----------------------------------------------------------------------------
			// save the AdAccount
			//----------------------------------------------------------------------------
			return UpdateAdAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BillingProfile", billingProfileId )
			return utils.RequestResult{false, msg, "assignBillingProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BillingProfile on a AdAccount
//----------------------------------------------------------------------------
func UnassignBillingProfileFromAdAccount(adAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		//----------------------------------------------------------------------------
		// assign an empty BillingProfile to the BillingProfile
		//----------------------------------------------------------------------------
		parentObj.BillingProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BillingProfile
		//----------------------------------------------------------------------------
		parentObj.BillingProfileId = nil;

		//----------------------------------------------------------------------------
		// save the AdAccount
		//----------------------------------------------------------------------------
		return UpdateAdAccount(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dsp on a AdAccount
//----------------------------------------------------------------------------
func AssignDspToAdAccount( adAccountId uint64, dspId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.DSP

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a DSP with a
		// matching dspId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, dspId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Dsp	to the AdAccount
			//----------------------------------------------------------------------------
			parentObj.Dsp = &childObj

			//----------------------------------------------------------------------------
			// save the AdAccount
			//----------------------------------------------------------------------------
			return UpdateAdAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dsp", dspId )
			return utils.RequestResult{false, msg, "assignDsp", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dsp on a AdAccount
//----------------------------------------------------------------------------
func UnassignDspFromAdAccount(adAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		//----------------------------------------------------------------------------
		// assign an empty DSP to the Dsp
		//----------------------------------------------------------------------------
		parentObj.Dsp = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dsp
		//----------------------------------------------------------------------------
		parentObj.DspId = nil;

		//----------------------------------------------------------------------------
		// save the AdAccount
		//----------------------------------------------------------------------------
		return UpdateAdAccount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more usersIds as a Users to a AdAccount
//----------------------------------------------------------------------------
func AddUsersToAdAccount ( adAccountId uint64, usersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

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
		// retrieve the modified AdAccount from the gorm
		//----------------------------------------------------------------------------
		return GetAdAccount(adAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more usersIds as a Users from a AdAccount
//----------------------------------------------------------------------------
func RemoveUsersFromAdAccount( adAccountId uint64, usersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

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
		// retrieve the modified AdAccount from the gorm
		//----------------------------------------------------------------------------
		return GetAdAccount(adAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a AdAccount
//----------------------------------------------------------------------------
func AddCampaignsToAdAccount ( adAccountId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

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
		// retrieve the modified AdAccount from the gorm
		//----------------------------------------------------------------------------
		return GetAdAccount(adAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a AdAccount
//----------------------------------------------------------------------------
func RemoveCampaignsFromAdAccount( adAccountId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

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
		// retrieve the modified AdAccount from the gorm
		//----------------------------------------------------------------------------
		return GetAdAccount(adAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more performanceMetricsIds as a PerformanceMetrics to a AdAccount
//----------------------------------------------------------------------------
func AddPerformanceMetricsToAdAccount ( adAccountId uint64, performanceMetricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( performanceMetricsIds, ",")

		for _, performanceMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceMetric
			// with a matching performanceMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , performanceMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PerformanceMetrics using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PerformanceMetrics").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PerformanceMetrics", performanceMetricsId )
				return utils.RequestResult{false, msg, "unassignPerformanceMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AdAccount from the gorm
		//----------------------------------------------------------------------------
		return GetAdAccount(adAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more performanceMetricsIds as a PerformanceMetrics from a AdAccount
//----------------------------------------------------------------------------
func RemovePerformanceMetricsFromAdAccount( adAccountId uint64, performanceMetricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AdAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdAccount(adAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( performanceMetricsIds, ",")

		for _, performanceMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PerformanceMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PerformanceMetric
			// with a matching performanceMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , performanceMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PerformanceMetricObj from the PerformanceMetrics array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PerformanceMetrics").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PerformanceMetrics", performanceMetricsId )
				return utils.RequestResult{false, msg, "removePerformanceMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AdAccount from the gorm
		//----------------------------------------------------------------------------
		return GetAdAccount(adAccountId)

	} else {
		return parentRequestResult
	}
}

