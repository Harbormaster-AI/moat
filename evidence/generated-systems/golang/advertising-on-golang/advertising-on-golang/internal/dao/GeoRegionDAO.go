package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing GeoRegionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateGeoRegion - creates a new db entry
//----------------------------------------------------------------------------
func CreateGeoRegion(obj model.GeoRegion)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a GeoRegion with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a GeoRegion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateGeoRegion", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetGeoRegion - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetGeoRegion(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.GeoRegion

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a GeoRegion with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a GeoRegion using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a GeoRegion using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetGeoRegion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllGeoRegion - returns all
//----------------------------------------------------------------------------
func GetAllGeoRegion()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.GeoRegion

	//----------------------------------------------------------------------------
	// Request the ORM to find all GeoRegion
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all GeoRegion" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all GeoRegion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllGeoRegion", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateGeoRegion - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateGeoRegion(obj model.GeoRegion)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a GeoRegion using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a GeoRegion using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateGeoRegion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteGeoRegion - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteGeoRegion(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the GeoRegion with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetGeoRegion(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GeoRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.GeoRegion)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a GeoRegion using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a GeoRegion using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteGeoRegion", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Parent on a GeoRegion
//----------------------------------------------------------------------------
func AssignParentToGeoRegion( geoRegionId uint64, parentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the GeoRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGeoRegion(geoRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GeoRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GeoRegion)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.GeoRegion

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a GeoRegion with a
		// matching parentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, parentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Parent	to the GeoRegion
			//----------------------------------------------------------------------------
			parentObj.Parent = &childObj

			//----------------------------------------------------------------------------
			// save the GeoRegion
			//----------------------------------------------------------------------------
			return UpdateGeoRegion(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Parent", parentId )
			return utils.RequestResult{false, msg, "assignParent", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Parent on a GeoRegion
//----------------------------------------------------------------------------
func UnassignParentFromGeoRegion(geoRegionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GeoRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGeoRegion(geoRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GeoRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GeoRegion)

		//----------------------------------------------------------------------------
		// assign an empty GeoRegion to the Parent
		//----------------------------------------------------------------------------
		parentObj.Parent = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Parent
		//----------------------------------------------------------------------------
		parentObj.ParentId = nil;

		//----------------------------------------------------------------------------
		// save the GeoRegion
		//----------------------------------------------------------------------------
		return UpdateGeoRegion(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more childrenIds as a Children to a GeoRegion
//----------------------------------------------------------------------------
func AddChildrenToGeoRegion ( geoRegionId uint64, childrenIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GeoRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGeoRegion(geoRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GeoRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GeoRegion)

		// slice the ids on comma with no spaces
		ids := strings.Split( childrenIds, ",")

		for _, childrenId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GeoRegion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GeoRegion
			// with a matching childrenId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childrenId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Children using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Children").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Children", childrenId )
				return utils.RequestResult{false, msg, "unassignChildren", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified GeoRegion from the gorm
		//----------------------------------------------------------------------------
		return GetGeoRegion(geoRegionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more childrenIds as a Children from a GeoRegion
//----------------------------------------------------------------------------
func RemoveChildrenFromGeoRegion( geoRegionId uint64, childrenIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the GeoRegion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGeoRegion(geoRegionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GeoRegion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GeoRegion)

		// slice the ids on comma with no spaces
		ids := strings.Split( childrenIds, ",")

		for _, childrenId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GeoRegion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GeoRegion
			// with a matching childrenId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childrenId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GeoRegionObj from the Children array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Children").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Children", childrenId )
				return utils.RequestResult{false, msg, "removeChildren", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified GeoRegion from the gorm
		//----------------------------------------------------------------------------
		return GetGeoRegion(geoRegionId)

	} else {
		return parentRequestResult
	}
}

