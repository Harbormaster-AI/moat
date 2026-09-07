package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ImagingOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateImagingOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateImagingOrder(obj model.ImagingOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ImagingOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ImagingOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateImagingOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetImagingOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetImagingOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ImagingOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ImagingOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ImagingOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ImagingOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetImagingOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllImagingOrder - returns all
//----------------------------------------------------------------------------
func GetAllImagingOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ImagingOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all ImagingOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ImagingOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ImagingOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllImagingOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateImagingOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateImagingOrder(obj model.ImagingOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ImagingOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ImagingOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateImagingOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteImagingOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteImagingOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ImagingOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetImagingOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ImagingOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ImagingOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ImagingOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteImagingOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Order on a ImagingOrder
//----------------------------------------------------------------------------
func AssignOrderToImagingOrder( imagingOrderId uint64, orderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ImagingOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingOrder(imagingOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ClinicalOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ClinicalOrder with a
		// matching orderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, orderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Order	to the ImagingOrder
			//----------------------------------------------------------------------------
			parentObj.Order = &childObj

			//----------------------------------------------------------------------------
			// save the ImagingOrder
			//----------------------------------------------------------------------------
			return UpdateImagingOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Order", orderId )
			return utils.RequestResult{false, msg, "assignOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Order on a ImagingOrder
//----------------------------------------------------------------------------
func UnassignOrderFromImagingOrder(imagingOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingOrder(imagingOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingOrder)

		//----------------------------------------------------------------------------
		// assign an empty ClinicalOrder to the Order
		//----------------------------------------------------------------------------
		parentObj.Order = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Order
		//----------------------------------------------------------------------------
		parentObj.OrderId = nil;

		//----------------------------------------------------------------------------
		// save the ImagingOrder
		//----------------------------------------------------------------------------
		return UpdateImagingOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ImagingCenter on a ImagingOrder
//----------------------------------------------------------------------------
func AssignImagingCenterToImagingOrder( imagingOrderId uint64, imagingCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ImagingOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingOrder(imagingOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ImagingCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ImagingCenter with a
		// matching imagingCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, imagingCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ImagingCenter	to the ImagingOrder
			//----------------------------------------------------------------------------
			parentObj.ImagingCenter = &childObj

			//----------------------------------------------------------------------------
			// save the ImagingOrder
			//----------------------------------------------------------------------------
			return UpdateImagingOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingCenter", imagingCenterId )
			return utils.RequestResult{false, msg, "assignImagingCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ImagingCenter on a ImagingOrder
//----------------------------------------------------------------------------
func UnassignImagingCenterFromImagingOrder(imagingOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingOrder(imagingOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingOrder)

		//----------------------------------------------------------------------------
		// assign an empty ImagingCenter to the ImagingCenter
		//----------------------------------------------------------------------------
		parentObj.ImagingCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ImagingCenter
		//----------------------------------------------------------------------------
		parentObj.ImagingCenterId = nil;

		//----------------------------------------------------------------------------
		// save the ImagingOrder
		//----------------------------------------------------------------------------
		return UpdateImagingOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a ImagingOrder
//----------------------------------------------------------------------------
func AddReportsToImagingOrder ( imagingOrderId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ImagingOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingOrder(imagingOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingReport

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingReport
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "unassignReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ImagingOrder from the gorm
		//----------------------------------------------------------------------------
		return GetImagingOrder(imagingOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a ImagingOrder
//----------------------------------------------------------------------------
func RemoveReportsFromImagingOrder( imagingOrderId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ImagingOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetImagingOrder(imagingOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ImagingOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ImagingOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingReport

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingReport
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ImagingReportObj from the Reports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "removeReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ImagingOrder from the gorm
		//----------------------------------------------------------------------------
		return GetImagingOrder(imagingOrderId)

	} else {
		return parentRequestResult
	}
}

