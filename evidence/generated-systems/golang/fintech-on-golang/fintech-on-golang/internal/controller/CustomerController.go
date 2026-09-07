package controller

import (
    CustomerDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CustomerDAO for database creation
//----------------------------------------------------------------------------
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Customer model
	//----------------------------------------------------------------------------
	data := model.Customer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Customer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object to create
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.CreateCustomer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CustomerDAO to find the relevant Customer
//----------------------------------------------------------------------------
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]
	
	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.GetCustomer(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CustomerDAO for database read of all Customers
//----------------------------------------------------------------------------
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.GetAllCustomer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CustomerDAO for database save
//----------------------------------------------------------------------------
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Customer model
	//----------------------------------------------------------------------------
	var data = model.Customer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Customer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.UpdateCustomer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CustomerDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]

	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	//----------------------------------------------------------------------------
	// Delegate to the Customer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CustomerDAO.DeleteCustomer(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Institution on a Customer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInstitutionToCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	institutionId,_ := strconv.ParseUint( vars["institutionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AssignInstitutionToCustomer(customerId, institutionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Institution on a Customer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInstitutionFromCustomer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.UnassignInstitutionFromCustomer(customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a Customer
	//----------------------------------------------------------------------------
func AddAccountsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddAccountsToCustomer(customerId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveAccountsFromCustomer(customerId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more walletsIds as a Wallets to a Customer
	//----------------------------------------------------------------------------
func AddWalletsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	walletsIds,_ := vars["walletsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddWalletsToCustomer(customerId, walletsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more walletsIds as a Wallets from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWalletsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	walletsIds,_ := vars["walletsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveWalletsFromCustomer(customerId, walletsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more cardsIds as a Cards to a Customer
	//----------------------------------------------------------------------------
func AddCardsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cardsIds,_ := vars["cardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddCardsToCustomer(customerId, cardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more cardsIds as a Cards from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCardsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cardsIds,_ := vars["cardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveCardsFromCustomer(customerId, cardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more kycProfilesIds as a KycProfiles to a Customer
	//----------------------------------------------------------------------------
func AddKycProfilesToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfilesIds,_ := vars["kycProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddKycProfilesToCustomer(customerId, kycProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more kycProfilesIds as a KycProfiles from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveKycProfilesFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfilesIds,_ := vars["kycProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveKycProfilesFromCustomer(customerId, kycProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a Customer
	//----------------------------------------------------------------------------
func AddConsentsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddConsentsToCustomer(customerId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConsentsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveConsentsFromCustomer(customerId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more agreementsIds as a Agreements to a Customer
	//----------------------------------------------------------------------------
func AddAgreementsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agreementsIds,_ := vars["agreementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddAgreementsToCustomer(customerId, agreementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more agreementsIds as a Agreements from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAgreementsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agreementsIds,_ := vars["agreementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveAgreementsFromCustomer(customerId, agreementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more loanApplicationsIds as a LoanApplications to a Customer
	//----------------------------------------------------------------------------
func AddLoanApplicationsToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanApplicationsIds,_ := vars["loanApplicationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddLoanApplicationsToCustomer(customerId, loanApplicationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more loanApplicationsIds as a LoanApplications from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLoanApplicationsFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanApplicationsIds,_ := vars["loanApplicationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveLoanApplicationsFromCustomer(customerId, loanApplicationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more loansIds as a Loans to a Customer
	//----------------------------------------------------------------------------
func AddLoansToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loansIds,_ := vars["loansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddLoansToCustomer(customerId, loansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more loansIds as a Loans from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLoansFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loansIds,_ := vars["loansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveLoansFromCustomer(customerId, loansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more portfoliosIds as a Portfolios to a Customer
	//----------------------------------------------------------------------------
func AddPortfoliosToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	portfoliosIds,_ := vars["portfoliosIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddPortfoliosToCustomer(customerId, portfoliosIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more portfoliosIds as a Portfolios from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePortfoliosFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	portfoliosIds,_ := vars["portfoliosIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemovePortfoliosFromCustomer(customerId, portfoliosIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more disputesIds as a Disputes to a Customer
	//----------------------------------------------------------------------------
func AddDisputesToCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.AddDisputesToCustomer(customerId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more disputesIds as a Disputes from a Customer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDisputesFromCustomer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	customerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Customer DAO
	//----------------------------------------------------------------------------
	requestResult := CustomerDAO.RemoveDisputesFromCustomer(customerId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
