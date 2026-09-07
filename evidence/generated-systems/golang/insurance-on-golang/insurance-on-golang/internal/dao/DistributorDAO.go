package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DistributorDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDistributor - creates a new db entry
//----------------------------------------------------------------------------
func CreateDistributor(obj model.Distributor)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Distributor with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Distributor", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDistributor", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDistributor - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDistributor(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Distributor

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Distributor with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Distributor using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Distributor using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDistributor", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDistributor - returns all
//----------------------------------------------------------------------------
func GetAllDistributor()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Distributor

	//----------------------------------------------------------------------------
	// Request the ORM to find all Distributor
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Distributor" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Distributor", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDistributor", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDistributor - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDistributor(obj model.Distributor)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Distributor using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Distributor using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDistributor", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDistributor - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDistributor(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Distributor with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDistributor(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Distributor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Distributor)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Distributor using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Distributor using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDistributor", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more insurersIds as a Insurers to a Distributor
//----------------------------------------------------------------------------
func AddInsurersToDistributor ( distributorId uint64, insurersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Distributor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDistributor(distributorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Distributor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Distributor)

		// slice the ids on comma with no spaces
		ids := strings.Split( insurersIds, ",")

		for _, insurersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Insurer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Insurer
			// with a matching insurersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insurersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Insurers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Insurers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Insurers", insurersId )
				return utils.RequestResult{false, msg, "unassignInsurers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Distributor from the gorm
		//----------------------------------------------------------------------------
		return GetDistributor(distributorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more insurersIds as a Insurers from a Distributor
//----------------------------------------------------------------------------
func RemoveInsurersFromDistributor( distributorId uint64, insurersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Distributor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDistributor(distributorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Distributor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Distributor)

		// slice the ids on comma with no spaces
		ids := strings.Split( insurersIds, ",")

		for _, insurersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Insurer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Insurer
			// with a matching insurersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insurersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsurerObj from the Insurers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Insurers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Insurers", insurersId )
				return utils.RequestResult{false, msg, "removeInsurers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Distributor from the gorm
		//----------------------------------------------------------------------------
		return GetDistributor(distributorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more agentsIds as a Agents to a Distributor
//----------------------------------------------------------------------------
func AddAgentsToDistributor ( distributorId uint64, agentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Distributor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDistributor(distributorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Distributor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Distributor)

		// slice the ids on comma with no spaces
		ids := strings.Split( agentsIds, ",")

		for _, agentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Agent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Agent
			// with a matching agentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , agentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Agents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Agents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agents", agentsId )
				return utils.RequestResult{false, msg, "unassignAgents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Distributor from the gorm
		//----------------------------------------------------------------------------
		return GetDistributor(distributorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more agentsIds as a Agents from a Distributor
//----------------------------------------------------------------------------
func RemoveAgentsFromDistributor( distributorId uint64, agentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Distributor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDistributor(distributorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Distributor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Distributor)

		// slice the ids on comma with no spaces
		ids := strings.Split( agentsIds, ",")

		for _, agentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Agent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Agent
			// with a matching agentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , agentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AgentObj from the Agents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Agents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agents", agentsId )
				return utils.RequestResult{false, msg, "removeAgents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Distributor from the gorm
		//----------------------------------------------------------------------------
		return GetDistributor(distributorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a Distributor
//----------------------------------------------------------------------------
func AddPoliciesToDistributor ( distributorId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Distributor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDistributor(distributorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Distributor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Distributor)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Policies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "unassignPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Distributor from the gorm
		//----------------------------------------------------------------------------
		return GetDistributor(distributorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a Distributor
//----------------------------------------------------------------------------
func RemovePoliciesFromDistributor( distributorId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Distributor with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDistributor(distributorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Distributor so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Distributor)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyObj from the Policies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "removePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Distributor from the gorm
		//----------------------------------------------------------------------------
		return GetDistributor(distributorId)

	} else {
		return parentRequestResult
	}
}

