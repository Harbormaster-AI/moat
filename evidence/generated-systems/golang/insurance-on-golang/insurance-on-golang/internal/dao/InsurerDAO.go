package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InsurerDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInsurer - creates a new db entry
//----------------------------------------------------------------------------
func CreateInsurer(obj model.Insurer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Insurer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Insurer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInsurer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInsurer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInsurer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Insurer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Insurer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Insurer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Insurer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInsurer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInsurer - returns all
//----------------------------------------------------------------------------
func GetAllInsurer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Insurer

	//----------------------------------------------------------------------------
	// Request the ORM to find all Insurer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Insurer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Insurer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInsurer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInsurer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInsurer(obj model.Insurer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Insurer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Insurer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInsurer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInsurer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInsurer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInsurer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Insurer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Insurer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Insurer using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInsurer", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more productsIds as a Products to a Insurer
//----------------------------------------------------------------------------
func AddProductsToInsurer ( insurerId uint64, productsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( productsIds, ",")

		for _, productsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuranceProduct

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuranceProduct
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Products using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Products").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )
				return utils.RequestResult{false, msg, "unassignProducts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productsIds as a Products from a Insurer
//----------------------------------------------------------------------------
func RemoveProductsFromInsurer( insurerId uint64, productsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( productsIds, ",")

		for _, productsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuranceProduct

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuranceProduct
			// with a matching productsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsuranceProductObj from the Products array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Products").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Products", productsId )
				return utils.RequestResult{false, msg, "removeProducts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more distributionPartnersIds as a DistributionPartners to a Insurer
//----------------------------------------------------------------------------
func AddDistributionPartnersToInsurer ( insurerId uint64, distributionPartnersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( distributionPartnersIds, ",")

		for _, distributionPartnersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Distributor

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Distributor
			// with a matching distributionPartnersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , distributionPartnersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DistributionPartners using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DistributionPartners").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DistributionPartners", distributionPartnersId )
				return utils.RequestResult{false, msg, "unassignDistributionPartners", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more distributionPartnersIds as a DistributionPartners from a Insurer
//----------------------------------------------------------------------------
func RemoveDistributionPartnersFromInsurer( insurerId uint64, distributionPartnersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( distributionPartnersIds, ",")

		for _, distributionPartnersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Distributor

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Distributor
			// with a matching distributionPartnersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , distributionPartnersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DistributorObj from the DistributionPartners array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DistributionPartners").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DistributionPartners", distributionPartnersId )
				return utils.RequestResult{false, msg, "removeDistributionPartners", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a Insurer
//----------------------------------------------------------------------------
func AddPoliciesToInsurer ( insurerId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

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
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a Insurer
//----------------------------------------------------------------------------
func RemovePoliciesFromInsurer( insurerId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

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
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more claimsIds as a Claims to a Insurer
//----------------------------------------------------------------------------
func AddClaimsToInsurer ( insurerId uint64, claimsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Claims using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "unassignClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a Insurer
//----------------------------------------------------------------------------
func RemoveClaimsFromInsurer( insurerId uint64, claimsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimObj from the Claims array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "removeClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reinsuranceAgreementsIds as a ReinsuranceAgreements to a Insurer
//----------------------------------------------------------------------------
func AddReinsuranceAgreementsToInsurer ( insurerId uint64, reinsuranceAgreementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( reinsuranceAgreementsIds, ",")

		for _, reinsuranceAgreementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ReinsuranceAgreement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ReinsuranceAgreement
			// with a matching reinsuranceAgreementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reinsuranceAgreementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ReinsuranceAgreements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ReinsuranceAgreements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ReinsuranceAgreements", reinsuranceAgreementsId )
				return utils.RequestResult{false, msg, "unassignReinsuranceAgreements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reinsuranceAgreementsIds as a ReinsuranceAgreements from a Insurer
//----------------------------------------------------------------------------
func RemoveReinsuranceAgreementsFromInsurer( insurerId uint64, reinsuranceAgreementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Insurer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurer(insurerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Insurer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Insurer)

		// slice the ids on comma with no spaces
		ids := strings.Split( reinsuranceAgreementsIds, ",")

		for _, reinsuranceAgreementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ReinsuranceAgreement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ReinsuranceAgreement
			// with a matching reinsuranceAgreementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reinsuranceAgreementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReinsuranceAgreementObj from the ReinsuranceAgreements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ReinsuranceAgreements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ReinsuranceAgreements", reinsuranceAgreementsId )
				return utils.RequestResult{false, msg, "removeReinsuranceAgreements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Insurer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurer(insurerId)

	} else {
		return parentRequestResult
	}
}

