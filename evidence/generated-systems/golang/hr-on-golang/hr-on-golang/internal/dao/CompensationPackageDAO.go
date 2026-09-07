package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CompensationPackageDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCompensationPackage - creates a new db entry
//----------------------------------------------------------------------------
func CreateCompensationPackage(obj model.CompensationPackage)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CompensationPackage with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CompensationPackage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCompensationPackage", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCompensationPackage - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCompensationPackage(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CompensationPackage

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CompensationPackage with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CompensationPackage using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CompensationPackage using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCompensationPackage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCompensationPackage - returns all
//----------------------------------------------------------------------------
func GetAllCompensationPackage()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CompensationPackage

	//----------------------------------------------------------------------------
	// Request the ORM to find all CompensationPackage
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CompensationPackage" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CompensationPackage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCompensationPackage", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCompensationPackage - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCompensationPackage(obj model.CompensationPackage)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CompensationPackage using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CompensationPackage using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCompensationPackage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCompensationPackage - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCompensationPackage(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCompensationPackage(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CompensationPackage)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CompensationPackage using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CompensationPackage using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCompensationPackage", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Contract on a CompensationPackage
//----------------------------------------------------------------------------
func AssignContractToCompensationPackage( compensationPackageId uint64, contractId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.EmploymentContract

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a EmploymentContract with a
		// matching contractId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, contractId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Contract	to the CompensationPackage
			//----------------------------------------------------------------------------
			parentObj.Contract = &childObj

			//----------------------------------------------------------------------------
			// save the CompensationPackage
			//----------------------------------------------------------------------------
			return UpdateCompensationPackage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contract", contractId )
			return utils.RequestResult{false, msg, "assignContract", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Contract on a CompensationPackage
//----------------------------------------------------------------------------
func UnassignContractFromCompensationPackage(compensationPackageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		//----------------------------------------------------------------------------
		// assign an empty EmploymentContract to the Contract
		//----------------------------------------------------------------------------
		parentObj.Contract = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Contract
		//----------------------------------------------------------------------------
		parentObj.ContractId = nil;

		//----------------------------------------------------------------------------
		// save the CompensationPackage
		//----------------------------------------------------------------------------
		return UpdateCompensationPackage(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more salaryComponentsIds as a SalaryComponents to a CompensationPackage
//----------------------------------------------------------------------------
func AddSalaryComponentsToCompensationPackage ( compensationPackageId uint64, salaryComponentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		// slice the ids on comma with no spaces
		ids := strings.Split( salaryComponentsIds, ",")

		for _, salaryComponentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalaryComponent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalaryComponent
			// with a matching salaryComponentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , salaryComponentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SalaryComponents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SalaryComponents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalaryComponents", salaryComponentsId )
				return utils.RequestResult{false, msg, "unassignSalaryComponents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CompensationPackage from the gorm
		//----------------------------------------------------------------------------
		return GetCompensationPackage(compensationPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more salaryComponentsIds as a SalaryComponents from a CompensationPackage
//----------------------------------------------------------------------------
func RemoveSalaryComponentsFromCompensationPackage( compensationPackageId uint64, salaryComponentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		// slice the ids on comma with no spaces
		ids := strings.Split( salaryComponentsIds, ",")

		for _, salaryComponentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalaryComponent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalaryComponent
			// with a matching salaryComponentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , salaryComponentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SalaryComponentObj from the SalaryComponents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SalaryComponents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalaryComponents", salaryComponentsId )
				return utils.RequestResult{false, msg, "removeSalaryComponents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CompensationPackage from the gorm
		//----------------------------------------------------------------------------
		return GetCompensationPackage(compensationPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more bonusPlansIds as a BonusPlans to a CompensationPackage
//----------------------------------------------------------------------------
func AddBonusPlansToCompensationPackage ( compensationPackageId uint64, bonusPlansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		// slice the ids on comma with no spaces
		ids := strings.Split( bonusPlansIds, ",")

		for _, bonusPlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BonusPlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BonusPlan
			// with a matching bonusPlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , bonusPlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the BonusPlans using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BonusPlans").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BonusPlans", bonusPlansId )
				return utils.RequestResult{false, msg, "unassignBonusPlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CompensationPackage from the gorm
		//----------------------------------------------------------------------------
		return GetCompensationPackage(compensationPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more bonusPlansIds as a BonusPlans from a CompensationPackage
//----------------------------------------------------------------------------
func RemoveBonusPlansFromCompensationPackage( compensationPackageId uint64, bonusPlansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		// slice the ids on comma with no spaces
		ids := strings.Split( bonusPlansIds, ",")

		for _, bonusPlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BonusPlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BonusPlan
			// with a matching bonusPlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , bonusPlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BonusPlanObj from the BonusPlans array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BonusPlans").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BonusPlans", bonusPlansId )
				return utils.RequestResult{false, msg, "removeBonusPlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CompensationPackage from the gorm
		//----------------------------------------------------------------------------
		return GetCompensationPackage(compensationPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more equityGrantsIds as a EquityGrants to a CompensationPackage
//----------------------------------------------------------------------------
func AddEquityGrantsToCompensationPackage ( compensationPackageId uint64, equityGrantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		// slice the ids on comma with no spaces
		ids := strings.Split( equityGrantsIds, ",")

		for _, equityGrantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EquityGrant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EquityGrant
			// with a matching equityGrantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , equityGrantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EquityGrants using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EquityGrants").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EquityGrants", equityGrantsId )
				return utils.RequestResult{false, msg, "unassignEquityGrants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CompensationPackage from the gorm
		//----------------------------------------------------------------------------
		return GetCompensationPackage(compensationPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more equityGrantsIds as a EquityGrants from a CompensationPackage
//----------------------------------------------------------------------------
func RemoveEquityGrantsFromCompensationPackage( compensationPackageId uint64, equityGrantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CompensationPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompensationPackage(compensationPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompensationPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompensationPackage)

		// slice the ids on comma with no spaces
		ids := strings.Split( equityGrantsIds, ",")

		for _, equityGrantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EquityGrant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EquityGrant
			// with a matching equityGrantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , equityGrantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EquityGrantObj from the EquityGrants array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EquityGrants").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EquityGrants", equityGrantsId )
				return utils.RequestResult{false, msg, "removeEquityGrants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CompensationPackage from the gorm
		//----------------------------------------------------------------------------
		return GetCompensationPackage(compensationPackageId)

	} else {
		return parentRequestResult
	}
}

