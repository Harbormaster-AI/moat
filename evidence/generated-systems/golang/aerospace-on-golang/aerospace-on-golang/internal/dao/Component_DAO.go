package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing Component_DAO..." ) )
}

//----------------------------------------------------------------------------
// CreateComponent_ - creates a new db entry
//----------------------------------------------------------------------------
func CreateComponent_(obj model.Component_)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Component_ with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Component_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateComponent_", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetComponent_ - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetComponent_(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Component_

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Component_ with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Component_ using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Component_ using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetComponent_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllComponent_ - returns all
//----------------------------------------------------------------------------
func GetAllComponent_()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Component_

	//----------------------------------------------------------------------------
	// Request the ORM to find all Component_
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Component_" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Component_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllComponent_", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateComponent_ - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateComponent_(obj model.Component_)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Component_ using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Component_ using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateComponent_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteComponent_ - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteComponent_(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Component_ with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetComponent_(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Component_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Component_)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Component_ using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Component_ using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteComponent_", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Supplier on a Component_
//----------------------------------------------------------------------------
func AssignSupplierToComponent_( component_Id uint64, supplierId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Component_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComponent_(component_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Component_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Component_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Supplier

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Supplier with a
		// matching supplierId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, supplierId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Supplier	to the Component_
			//----------------------------------------------------------------------------
			parentObj.Supplier = &childObj

			//----------------------------------------------------------------------------
			// save the Component_
			//----------------------------------------------------------------------------
			return UpdateComponent_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Supplier", supplierId )
			return utils.RequestResult{false, msg, "assignSupplier", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Supplier on a Component_
//----------------------------------------------------------------------------
func UnassignSupplierFromComponent_(component_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Component_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComponent_(component_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Component_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Component_)

		//----------------------------------------------------------------------------
		// assign an empty Supplier to the Supplier
		//----------------------------------------------------------------------------
		parentObj.Supplier = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Supplier
		//----------------------------------------------------------------------------
		parentObj.SupplierId = nil;

		//----------------------------------------------------------------------------
		// save the Component_
		//----------------------------------------------------------------------------
		return UpdateComponent_(parentObj)

	} else {
		return parentRequestResult
	}

}


