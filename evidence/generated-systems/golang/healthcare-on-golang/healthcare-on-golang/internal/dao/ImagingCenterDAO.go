package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ImagingCenterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateImagingCenter - creates a new db entry
//----------------------------------------------------------------------------
func CreateImagingCenter(obj model.ImagingCenter)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ImagingCenter with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ImagingCenter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateImagingCenter", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetImagingCenter - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetImagingCenter(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ImagingCenter

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ImagingCenter with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ImagingCenter using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ImagingCenter using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetImagingCenter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllImagingCenter - returns all
//----------------------------------------------------------------------------
func GetAllImagingCenter()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ImagingCenter

	//----------------------------------------------------------------------------
	// Request the ORM to find all ImagingCenter
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ImagingCenter" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ImagingCenter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllImagingCenter", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateImagingCenter - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateImagingCenter(obj model.ImagingCenter)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ImagingCenter using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ImagingCenter using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateImagingCenter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteImagingCenter - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteImagingCenter(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ImagingCenter with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetImagingCenter(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ImagingCenter)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ImagingCenter using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ImagingCenter using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteImagingCenter", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Facility on a ImagingCenter
//----------------------------------------------------------------------------
func AssignFacilityToImagingCenter( imagingCenterId uint64, facilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ImagingCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingCenter(imagingCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingCenter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Facility

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Facility with a
		// matching facilityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, facilityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Facility	to the ImagingCenter
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the ImagingCenter
			//----------------------------------------------------------------------------
			return UpdateImagingCenter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a ImagingCenter
//----------------------------------------------------------------------------
func UnassignFacilityFromImagingCenter(imagingCenterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingCenter(imagingCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingCenter)

		//----------------------------------------------------------------------------
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the ImagingCenter
		//----------------------------------------------------------------------------
		return UpdateImagingCenter(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more imagingOrdersIds as a ImagingOrders to a ImagingCenter
//----------------------------------------------------------------------------
func AddImagingOrdersToImagingCenter ( imagingCenterId uint64, imagingOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingCenter(imagingCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingOrdersIds, ",")

		for _, imagingOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingOrder
			// with a matching imagingOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ImagingOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingOrders", imagingOrdersId )
				return utils.RequestResult{false, msg, "unassignImagingOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ImagingCenter from the gorm
		//----------------------------------------------------------------------------
		return GetImagingCenter(imagingCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more imagingOrdersIds as a ImagingOrders from a ImagingCenter
//----------------------------------------------------------------------------
func RemoveImagingOrdersFromImagingCenter( imagingCenterId uint64, imagingOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ImagingCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingCenter(imagingCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingOrdersIds, ",")

		for _, imagingOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingOrder
			// with a matching imagingOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ImagingOrderObj from the ImagingOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingOrders", imagingOrdersId )
				return utils.RequestResult{false, msg, "removeImagingOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ImagingCenter from the gorm
		//----------------------------------------------------------------------------
		return GetImagingCenter(imagingCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more imagingReportsIds as a ImagingReports to a ImagingCenter
//----------------------------------------------------------------------------
func AddImagingReportsToImagingCenter ( imagingCenterId uint64, imagingReportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingCenter(imagingCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingReportsIds, ",")

		for _, imagingReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingReport

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingReport
			// with a matching imagingReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ImagingReports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingReports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingReports", imagingReportsId )
				return utils.RequestResult{false, msg, "unassignImagingReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ImagingCenter from the gorm
		//----------------------------------------------------------------------------
		return GetImagingCenter(imagingCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more imagingReportsIds as a ImagingReports from a ImagingCenter
//----------------------------------------------------------------------------
func RemoveImagingReportsFromImagingCenter( imagingCenterId uint64, imagingReportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ImagingCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingCenter(imagingCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingReportsIds, ",")

		for _, imagingReportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingReport

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingReport
			// with a matching imagingReportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingReportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ImagingReportObj from the ImagingReports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingReports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingReports", imagingReportsId )
				return utils.RequestResult{false, msg, "removeImagingReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ImagingCenter from the gorm
		//----------------------------------------------------------------------------
		return GetImagingCenter(imagingCenterId)

	} else {
		return parentRequestResult
	}
}

