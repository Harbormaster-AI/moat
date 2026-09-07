package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProductionCertificateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProductionCertificate - creates a new db entry
//----------------------------------------------------------------------------
func CreateProductionCertificate(obj model.ProductionCertificate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ProductionCertificate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ProductionCertificate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProductionCertificate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProductionCertificate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProductionCertificate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ProductionCertificate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ProductionCertificate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ProductionCertificate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ProductionCertificate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProductionCertificate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProductionCertificate - returns all
//----------------------------------------------------------------------------
func GetAllProductionCertificate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ProductionCertificate

	//----------------------------------------------------------------------------
	// Request the ORM to find all ProductionCertificate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ProductionCertificate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ProductionCertificate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProductionCertificate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProductionCertificate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProductionCertificate(obj model.ProductionCertificate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ProductionCertificate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ProductionCertificate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProductionCertificate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProductionCertificate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProductionCertificate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ProductionCertificate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProductionCertificate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionCertificate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ProductionCertificate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ProductionCertificate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ProductionCertificate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProductionCertificate", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Manufacturer on a ProductionCertificate
//----------------------------------------------------------------------------
func AssignManufacturerToProductionCertificate( productionCertificateId uint64, manufacturerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProductionCertificate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionCertificate(productionCertificateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionCertificate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionCertificate)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AerospaceManufacturer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AerospaceManufacturer with a
		// matching manufacturerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, manufacturerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Manufacturer	to the ProductionCertificate
			//----------------------------------------------------------------------------
			parentObj.Manufacturer = &childObj

			//----------------------------------------------------------------------------
			// save the ProductionCertificate
			//----------------------------------------------------------------------------
			return UpdateProductionCertificate(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Manufacturer", manufacturerId )
			return utils.RequestResult{false, msg, "assignManufacturer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Manufacturer on a ProductionCertificate
//----------------------------------------------------------------------------
func UnassignManufacturerFromProductionCertificate(productionCertificateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionCertificate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionCertificate(productionCertificateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionCertificate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionCertificate)

		//----------------------------------------------------------------------------
		// assign an empty AerospaceManufacturer to the Manufacturer
		//----------------------------------------------------------------------------
		parentObj.Manufacturer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Manufacturer
		//----------------------------------------------------------------------------
		parentObj.ManufacturerId = nil;

		//----------------------------------------------------------------------------
		// save the ProductionCertificate
		//----------------------------------------------------------------------------
		return UpdateProductionCertificate(parentObj)

	} else {
		return parentRequestResult
	}

}


