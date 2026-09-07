package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ContentCategoryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateContentCategory - creates a new db entry
//----------------------------------------------------------------------------
func CreateContentCategory(obj model.ContentCategory)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ContentCategory with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ContentCategory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateContentCategory", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetContentCategory - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetContentCategory(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ContentCategory

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ContentCategory with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ContentCategory using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ContentCategory using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetContentCategory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllContentCategory - returns all
//----------------------------------------------------------------------------
func GetAllContentCategory()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ContentCategory

	//----------------------------------------------------------------------------
	// Request the ORM to find all ContentCategory
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ContentCategory" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ContentCategory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllContentCategory", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateContentCategory - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateContentCategory(obj model.ContentCategory)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ContentCategory using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ContentCategory using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateContentCategory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteContentCategory - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteContentCategory(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ContentCategory with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetContentCategory(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ContentCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ContentCategory)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ContentCategory using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ContentCategory using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteContentCategory", requestResult.Data}

	}

	return requestResult
}



