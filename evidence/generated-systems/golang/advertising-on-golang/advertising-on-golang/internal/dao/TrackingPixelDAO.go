package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TrackingPixelDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTrackingPixel - creates a new db entry
//----------------------------------------------------------------------------
func CreateTrackingPixel(obj model.TrackingPixel)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TrackingPixel with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TrackingPixel", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTrackingPixel", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTrackingPixel - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTrackingPixel(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TrackingPixel

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TrackingPixel with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TrackingPixel using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TrackingPixel using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTrackingPixel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTrackingPixel - returns all
//----------------------------------------------------------------------------
func GetAllTrackingPixel()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TrackingPixel

	//----------------------------------------------------------------------------
	// Request the ORM to find all TrackingPixel
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TrackingPixel" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TrackingPixel", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTrackingPixel", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTrackingPixel - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTrackingPixel(obj model.TrackingPixel)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TrackingPixel using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TrackingPixel using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTrackingPixel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTrackingPixel - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTrackingPixel(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TrackingPixel with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTrackingPixel(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrackingPixel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TrackingPixel)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TrackingPixel using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TrackingPixel using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTrackingPixel", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Campaign on a TrackingPixel
//----------------------------------------------------------------------------
func AssignCampaignToTrackingPixel( trackingPixelId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TrackingPixel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrackingPixel(trackingPixelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrackingPixel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrackingPixel)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Campaign

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Campaign with a
		// matching campaignId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, campaignId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Campaign	to the TrackingPixel
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the TrackingPixel
			//----------------------------------------------------------------------------
			return UpdateTrackingPixel(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a TrackingPixel
//----------------------------------------------------------------------------
func UnassignCampaignFromTrackingPixel(trackingPixelId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrackingPixel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrackingPixel(trackingPixelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrackingPixel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrackingPixel)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the TrackingPixel
		//----------------------------------------------------------------------------
		return UpdateTrackingPixel(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Advertiser on a TrackingPixel
//----------------------------------------------------------------------------
func AssignAdvertiserToTrackingPixel( trackingPixelId uint64, advertiserId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TrackingPixel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrackingPixel(trackingPixelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrackingPixel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrackingPixel)

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
			// assign the Advertiser	to the TrackingPixel
			//----------------------------------------------------------------------------
			parentObj.Advertiser = &childObj

			//----------------------------------------------------------------------------
			// save the TrackingPixel
			//----------------------------------------------------------------------------
			return UpdateTrackingPixel(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Advertiser", advertiserId )
			return utils.RequestResult{false, msg, "assignAdvertiser", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Advertiser on a TrackingPixel
//----------------------------------------------------------------------------
func UnassignAdvertiserFromTrackingPixel(trackingPixelId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrackingPixel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrackingPixel(trackingPixelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrackingPixel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrackingPixel)

		//----------------------------------------------------------------------------
		// assign an empty Advertiser to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.Advertiser = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.AdvertiserId = nil;

		//----------------------------------------------------------------------------
		// save the TrackingPixel
		//----------------------------------------------------------------------------
		return UpdateTrackingPixel(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more conversionEventsIds as a ConversionEvents to a TrackingPixel
//----------------------------------------------------------------------------
func AddConversionEventsToTrackingPixel ( trackingPixelId uint64, conversionEventsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrackingPixel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrackingPixel(trackingPixelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrackingPixel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrackingPixel)

		// slice the ids on comma with no spaces
		ids := strings.Split( conversionEventsIds, ",")

		for _, conversionEventsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ConversionEvent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ConversionEvent
			// with a matching conversionEventsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , conversionEventsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ConversionEvents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ConversionEvents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConversionEvents", conversionEventsId )
				return utils.RequestResult{false, msg, "unassignConversionEvents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrackingPixel from the gorm
		//----------------------------------------------------------------------------
		return GetTrackingPixel(trackingPixelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more conversionEventsIds as a ConversionEvents from a TrackingPixel
//----------------------------------------------------------------------------
func RemoveConversionEventsFromTrackingPixel( trackingPixelId uint64, conversionEventsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrackingPixel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrackingPixel(trackingPixelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrackingPixel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrackingPixel)

		// slice the ids on comma with no spaces
		ids := strings.Split( conversionEventsIds, ",")

		for _, conversionEventsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ConversionEvent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ConversionEvent
			// with a matching conversionEventsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , conversionEventsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConversionEventObj from the ConversionEvents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ConversionEvents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConversionEvents", conversionEventsId )
				return utils.RequestResult{false, msg, "removeConversionEvents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrackingPixel from the gorm
		//----------------------------------------------------------------------------
		return GetTrackingPixel(trackingPixelId)

	} else {
		return parentRequestResult
	}
}

