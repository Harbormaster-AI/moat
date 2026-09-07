package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataProviderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataProvider - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataProvider(obj model.DataProvider)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataProvider with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataProvider", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataProvider", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataProvider - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataProvider(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataProvider

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataProvider with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataProvider using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataProvider using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataProvider", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataProvider - returns all
//----------------------------------------------------------------------------
func GetAllDataProvider()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataProvider

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataProvider
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataProvider" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataProvider", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataProvider", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataProvider - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataProvider(obj model.DataProvider)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataProvider using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataProvider using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataProvider", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataProvider - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataProvider(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataProvider with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataProvider(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataProvider)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataProvider using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataProvider using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataProvider", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more audienceSegmentsIds as a AudienceSegments to a DataProvider
//----------------------------------------------------------------------------
func AddAudienceSegmentsToDataProvider ( dataProviderId uint64, audienceSegmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProvider(dataProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProvider)

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
		// retrieve the modified DataProvider from the gorm
		//----------------------------------------------------------------------------
		return GetDataProvider(dataProviderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more audienceSegmentsIds as a AudienceSegments from a DataProvider
//----------------------------------------------------------------------------
func RemoveAudienceSegmentsFromDataProvider( dataProviderId uint64, audienceSegmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProvider with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProvider(dataProviderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProvider so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProvider)

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
		// retrieve the modified DataProvider from the gorm
		//----------------------------------------------------------------------------
		return GetDataProvider(dataProviderId)

	} else {
		return parentRequestResult
	}
}

