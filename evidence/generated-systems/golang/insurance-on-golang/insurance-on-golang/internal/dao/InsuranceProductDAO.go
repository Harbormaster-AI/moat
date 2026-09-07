package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InsuranceProductDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInsuranceProduct - creates a new db entry
//----------------------------------------------------------------------------
func CreateInsuranceProduct(obj model.InsuranceProduct)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InsuranceProduct with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InsuranceProduct", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInsuranceProduct", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInsuranceProduct - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInsuranceProduct(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InsuranceProduct

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InsuranceProduct with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InsuranceProduct using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InsuranceProduct using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInsuranceProduct", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInsuranceProduct - returns all
//----------------------------------------------------------------------------
func GetAllInsuranceProduct()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InsuranceProduct

	//----------------------------------------------------------------------------
	// Request the ORM to find all InsuranceProduct
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InsuranceProduct" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InsuranceProduct", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInsuranceProduct", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInsuranceProduct - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInsuranceProduct(obj model.InsuranceProduct)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InsuranceProduct using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InsuranceProduct using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInsuranceProduct", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInsuranceProduct - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInsuranceProduct(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InsuranceProduct with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInsuranceProduct(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuranceProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InsuranceProduct)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InsuranceProduct using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InsuranceProduct using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInsuranceProduct", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Insurer on a InsuranceProduct
//----------------------------------------------------------------------------
func AssignInsurerToInsuranceProduct( insuranceProductId uint64, insurerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InsuranceProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuranceProduct(insuranceProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuranceProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuranceProduct)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Insurer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Insurer with a
		// matching insurerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, insurerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Insurer	to the InsuranceProduct
			//----------------------------------------------------------------------------
			parentObj.Insurer = &childObj

			//----------------------------------------------------------------------------
			// save the InsuranceProduct
			//----------------------------------------------------------------------------
			return UpdateInsuranceProduct(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Insurer", insurerId )
			return utils.RequestResult{false, msg, "assignInsurer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Insurer on a InsuranceProduct
//----------------------------------------------------------------------------
func UnassignInsurerFromInsuranceProduct(insuranceProductId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsuranceProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuranceProduct(insuranceProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuranceProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuranceProduct)

		//----------------------------------------------------------------------------
		// assign an empty Insurer to the Insurer
		//----------------------------------------------------------------------------
		parentObj.Insurer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Insurer
		//----------------------------------------------------------------------------
		parentObj.InsurerId = nil;

		//----------------------------------------------------------------------------
		// save the InsuranceProduct
		//----------------------------------------------------------------------------
		return UpdateInsuranceProduct(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more coverageDefinitionsIds as a CoverageDefinitions to a InsuranceProduct
//----------------------------------------------------------------------------
func AddCoverageDefinitionsToInsuranceProduct ( insuranceProductId uint64, coverageDefinitionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsuranceProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuranceProduct(insuranceProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuranceProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuranceProduct)

		// slice the ids on comma with no spaces
		ids := strings.Split( coverageDefinitionsIds, ",")

		for _, coverageDefinitionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CoverageDefinition

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CoverageDefinition
			// with a matching coverageDefinitionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coverageDefinitionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CoverageDefinitions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CoverageDefinitions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CoverageDefinitions", coverageDefinitionsId )
				return utils.RequestResult{false, msg, "unassignCoverageDefinitions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsuranceProduct from the gorm
		//----------------------------------------------------------------------------
		return GetInsuranceProduct(insuranceProductId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more coverageDefinitionsIds as a CoverageDefinitions from a InsuranceProduct
//----------------------------------------------------------------------------
func RemoveCoverageDefinitionsFromInsuranceProduct( insuranceProductId uint64, coverageDefinitionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InsuranceProduct with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuranceProduct(insuranceProductId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuranceProduct so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuranceProduct)

		// slice the ids on comma with no spaces
		ids := strings.Split( coverageDefinitionsIds, ",")

		for _, coverageDefinitionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CoverageDefinition

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CoverageDefinition
			// with a matching coverageDefinitionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coverageDefinitionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CoverageDefinitionObj from the CoverageDefinitions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CoverageDefinitions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CoverageDefinitions", coverageDefinitionsId )
				return utils.RequestResult{false, msg, "removeCoverageDefinitions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsuranceProduct from the gorm
		//----------------------------------------------------------------------------
		return GetInsuranceProduct(insuranceProductId)

	} else {
		return parentRequestResult
	}
}

