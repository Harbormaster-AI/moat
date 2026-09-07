package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProductOfferingDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProductOffering - creates a new db entry
//----------------------------------------------------------------------------
func CreateProductOffering(obj model.ProductOffering)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ProductOffering with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ProductOffering", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProductOffering", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProductOffering - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProductOffering(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ProductOffering

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ProductOffering with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ProductOffering using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ProductOffering using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProductOffering", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProductOffering - returns all
//----------------------------------------------------------------------------
func GetAllProductOffering()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ProductOffering

	//----------------------------------------------------------------------------
	// Request the ORM to find all ProductOffering
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ProductOffering" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ProductOffering", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProductOffering", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProductOffering - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProductOffering(obj model.ProductOffering)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ProductOffering using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ProductOffering using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProductOffering", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProductOffering - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProductOffering(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ProductOffering with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProductOffering(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductOffering so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ProductOffering)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ProductOffering using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ProductOffering using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProductOffering", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Institution on a ProductOffering
//----------------------------------------------------------------------------
func AssignInstitutionToProductOffering( productOfferingId uint64, institutionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProductOffering with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductOffering(productOfferingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductOffering so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductOffering)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FinancialInstitution

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FinancialInstitution with a
		// matching institutionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, institutionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Institution	to the ProductOffering
			//----------------------------------------------------------------------------
			parentObj.Institution = &childObj

			//----------------------------------------------------------------------------
			// save the ProductOffering
			//----------------------------------------------------------------------------
			return UpdateProductOffering(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Institution", institutionId )
			return utils.RequestResult{false, msg, "assignInstitution", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Institution on a ProductOffering
//----------------------------------------------------------------------------
func UnassignInstitutionFromProductOffering(productOfferingId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductOffering with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductOffering(productOfferingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductOffering so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductOffering)

		//----------------------------------------------------------------------------
		// assign an empty FinancialInstitution to the Institution
		//----------------------------------------------------------------------------
		parentObj.Institution = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Institution
		//----------------------------------------------------------------------------
		parentObj.InstitutionId = nil;

		//----------------------------------------------------------------------------
		// save the ProductOffering
		//----------------------------------------------------------------------------
		return UpdateProductOffering(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more pricingPlansIds as a PricingPlans to a ProductOffering
//----------------------------------------------------------------------------
func AddPricingPlansToProductOffering ( productOfferingId uint64, pricingPlansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductOffering with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductOffering(productOfferingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductOffering so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductOffering)

		// slice the ids on comma with no spaces
		ids := strings.Split( pricingPlansIds, ",")

		for _, pricingPlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PricingPlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PricingPlan
			// with a matching pricingPlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pricingPlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PricingPlans using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PricingPlans").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PricingPlans", pricingPlansId )
				return utils.RequestResult{false, msg, "unassignPricingPlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ProductOffering from the gorm
		//----------------------------------------------------------------------------
		return GetProductOffering(productOfferingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more pricingPlansIds as a PricingPlans from a ProductOffering
//----------------------------------------------------------------------------
func RemovePricingPlansFromProductOffering( productOfferingId uint64, pricingPlansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ProductOffering with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductOffering(productOfferingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductOffering so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductOffering)

		// slice the ids on comma with no spaces
		ids := strings.Split( pricingPlansIds, ",")

		for _, pricingPlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PricingPlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PricingPlan
			// with a matching pricingPlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pricingPlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PricingPlanObj from the PricingPlans array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PricingPlans").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PricingPlans", pricingPlansId )
				return utils.RequestResult{false, msg, "removePricingPlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ProductOffering from the gorm
		//----------------------------------------------------------------------------
		return GetProductOffering(productOfferingId)

	} else {
		return parentRequestResult
	}
}

