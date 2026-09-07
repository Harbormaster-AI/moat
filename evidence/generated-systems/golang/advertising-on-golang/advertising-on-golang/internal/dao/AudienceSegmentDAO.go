package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AudienceSegmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAudienceSegment - creates a new db entry
//----------------------------------------------------------------------------
func CreateAudienceSegment(obj model.AudienceSegment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AudienceSegment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AudienceSegment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAudienceSegment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAudienceSegment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAudienceSegment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AudienceSegment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AudienceSegment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AudienceSegment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AudienceSegment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAudienceSegment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAudienceSegment - returns all
//----------------------------------------------------------------------------
func GetAllAudienceSegment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AudienceSegment

	//----------------------------------------------------------------------------
	// Request the ORM to find all AudienceSegment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AudienceSegment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AudienceSegment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAudienceSegment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAudienceSegment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAudienceSegment(obj model.AudienceSegment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AudienceSegment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AudienceSegment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAudienceSegment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAudienceSegment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAudienceSegment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AudienceSegment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAudienceSegment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AudienceSegment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AudienceSegment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AudienceSegment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AudienceSegment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAudienceSegment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Provider on a AudienceSegment
//----------------------------------------------------------------------------
func AssignProviderToAudienceSegment( audienceSegmentId uint64, providerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AudienceSegment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAudienceSegment(audienceSegmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AudienceSegment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AudienceSegment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.DataProvider

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a DataProvider with a
		// matching providerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, providerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Provider	to the AudienceSegment
			//----------------------------------------------------------------------------
			parentObj.Provider = &childObj

			//----------------------------------------------------------------------------
			// save the AudienceSegment
			//----------------------------------------------------------------------------
			return UpdateAudienceSegment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Provider", providerId )
			return utils.RequestResult{false, msg, "assignProvider", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Provider on a AudienceSegment
//----------------------------------------------------------------------------
func UnassignProviderFromAudienceSegment(audienceSegmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AudienceSegment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAudienceSegment(audienceSegmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AudienceSegment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AudienceSegment)

		//----------------------------------------------------------------------------
		// assign an empty DataProvider to the Provider
		//----------------------------------------------------------------------------
		parentObj.Provider = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Provider
		//----------------------------------------------------------------------------
		parentObj.ProviderId = nil;

		//----------------------------------------------------------------------------
		// save the AudienceSegment
		//----------------------------------------------------------------------------
		return UpdateAudienceSegment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a AudienceSegment
//----------------------------------------------------------------------------
func AddCampaignsToAudienceSegment ( audienceSegmentId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AudienceSegment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAudienceSegment(audienceSegmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AudienceSegment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AudienceSegment)

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
		// retrieve the modified AudienceSegment from the gorm
		//----------------------------------------------------------------------------
		return GetAudienceSegment(audienceSegmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a AudienceSegment
//----------------------------------------------------------------------------
func RemoveCampaignsFromAudienceSegment( audienceSegmentId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AudienceSegment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAudienceSegment(audienceSegmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AudienceSegment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AudienceSegment)

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
		// retrieve the modified AudienceSegment from the gorm
		//----------------------------------------------------------------------------
		return GetAudienceSegment(audienceSegmentId)

	} else {
		return parentRequestResult
	}
}

