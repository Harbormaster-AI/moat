
// Define collection and schema for Customer
export interface Customer {
    firstName:
	type : string
    lastName:
	type : string
    dateOfBirth:
	type : Date
    email:
	type : Email
    phone:
	type : PhoneNumber
    address:
	type : Address
    taxId:
	type : TaxId
    riskScore:
	type : RiskScore
    Institution:
	type : Schema.Types.ObjectId
    Accounts:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Account' }]
    Wallets:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Wallet' }]
    Cards:
 	type : [{ type: Schema.Types.ObjectId, ref: 'PaymentCard' }]
    KycProfiles:
 	type : [{ type: Schema.Types.ObjectId, ref: 'KYCProfile' }]
    Consents:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Consent' }]
    Agreements:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Agreement' }]
    LoanApplications:
 	type : [{ type: Schema.Types.ObjectId, ref: 'LoanApplication' }]
    Loans:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Loan' }]
    Portfolios:
 	type : [{ type: Schema.Types.ObjectId, ref: 'InvestmentPortfolio' }]
    Disputes:
 	type : [{ type: Schema.Types.ObjectId, ref: 'Dispute' }]
    CustomerType:
 	type : String
#
    collection: 'customers'
}
