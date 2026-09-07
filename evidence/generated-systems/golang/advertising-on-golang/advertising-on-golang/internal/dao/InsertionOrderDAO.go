package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InsertionOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInsertionOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateInsertionOrder(obj model.InsertionOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InsertionOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InsertionOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInsertionOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInsertionOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInsertionOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InsertionOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InsertionOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InsertionOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InsertionOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInsertionOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInsertionOrder - returns all
//----------------------------------------------------------------------------
func GetAllInsertionOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InsertionOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all InsertionOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InsertionOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InsertionOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInsertionOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInsertionOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInsertionOrder(obj model.InsertionOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InsertionOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InsertionOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInsertionOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInsertionOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInsertionOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInsertionOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InsertionOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InsertionOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InsertionOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInsertionOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Advertiser on a InsertionOrder
//----------------------------------------------------------------------------
func AssignAdvertiserToInsertionOrder( insertionOrderId uint64, advertiserId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

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
			// assign the Advertiser	to the InsertionOrder
			//----------------------------------------------------------------------------
			parentObj.Advertiser = &childObj

			//----------------------------------------------------------------------------
			// save the InsertionOrder
			//----------------------------------------------------------------------------
			return UpdateInsertionOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Advertiser", advertiserId )
			return utils.RequestResult{false, msg, "assignAdvertiser", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Advertiser on a InsertionOrder
//----------------------------------------------------------------------------
func UnassignAdvertiserFromInsertionOrder(insertionOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

		//----------------------------------------------------------------------------
		// assign an empty Advertiser to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.Advertiser = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.AdvertiserId = nil;

		//----------------------------------------------------------------------------
		// save the InsertionOrder
		//----------------------------------------------------------------------------
		return UpdateInsertionOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Agency on a InsertionOrder
//----------------------------------------------------------------------------
func AssignAgencyToInsertionOrder( insertionOrderId uint64, agencyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

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
			// assign the Agency	to the InsertionOrder
			//----------------------------------------------------------------------------
			parentObj.Agency = &childObj

			//----------------------------------------------------------------------------
			// save the InsertionOrder
			//----------------------------------------------------------------------------
			return UpdateInsertionOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agency", agencyId )
			return utils.RequestResult{false, msg, "assignAgency", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Agency on a InsertionOrder
//----------------------------------------------------------------------------
func UnassignAgencyFromInsertionOrder(insertionOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

		//----------------------------------------------------------------------------
		// assign an empty Agency to the Agency
		//----------------------------------------------------------------------------
		parentObj.Agency = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Agency
		//----------------------------------------------------------------------------
		parentObj.AgencyId = nil;

		//----------------------------------------------------------------------------
		// save the InsertionOrder
		//----------------------------------------------------------------------------
		return UpdateInsertionOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Publisher on a InsertionOrder
//----------------------------------------------------------------------------
func AssignPublisherToInsertionOrder( insertionOrderId uint64, publisherId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Publisher

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Publisher with a
		// matching publisherId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, publisherId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Publisher	to the InsertionOrder
			//----------------------------------------------------------------------------
			parentObj.Publisher = &childObj

			//----------------------------------------------------------------------------
			// save the InsertionOrder
			//----------------------------------------------------------------------------
			return UpdateInsertionOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Publisher", publisherId )
			return utils.RequestResult{false, msg, "assignPublisher", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Publisher on a InsertionOrder
//----------------------------------------------------------------------------
func UnassignPublisherFromInsertionOrder(insertionOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

		//----------------------------------------------------------------------------
		// assign an empty Publisher to the Publisher
		//----------------------------------------------------------------------------
		parentObj.Publisher = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Publisher
		//----------------------------------------------------------------------------
		parentObj.PublisherId = nil;

		//----------------------------------------------------------------------------
		// save the InsertionOrder
		//----------------------------------------------------------------------------
		return UpdateInsertionOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a InsertionOrder
//----------------------------------------------------------------------------
func AddCampaignsToInsertionOrder ( insertionOrderId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

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
		// retrieve the modified InsertionOrder from the gorm
		//----------------------------------------------------------------------------
		return GetInsertionOrder(insertionOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a InsertionOrder
//----------------------------------------------------------------------------
func RemoveCampaignsFromInsertionOrder( insertionOrderId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InsertionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsertionOrder(insertionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsertionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsertionOrder)

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
		// retrieve the modified InsertionOrder from the gorm
		//----------------------------------------------------------------------------
		return GetInsertionOrder(insertionOrderId)

	} else {
		return parentRequestResult
	}
}

