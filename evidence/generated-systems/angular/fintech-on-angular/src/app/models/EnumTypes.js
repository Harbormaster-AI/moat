
// enum type ProductCategory
export let ProductCategory = {
	Checking:"Checking",
	Savings:"Savings",
	CreditCard:"CreditCard",
	Loan:"Loan",
	Investment:"Investment",
	Insurance:"Insurance",
	Payments:"Payments",
	FX:"FX",
	Wallet:"Wallet",
}

// enum type PlanStatus
export let PlanStatus = {
	Draft:"Draft",
	Active:"Active",
	Suspended:"Suspended",
	Archived:"Archived",
}

// enum type FeeType
export let FeeType = {
	Fixed:"Fixed",
	Percentage:"Percentage",
	Tiered:"Tiered",
	Interchange:"Interchange",
	Network:"Network",
	Chargeback:"Chargeback",
	ATM:"ATM",
	FX:"FX",
}

// enum type FeeCalculationMethod
export let FeeCalculationMethod = {
	PerTransaction:"PerTransaction",
	PerMonth:"PerMonth",
	PerAnnum:"PerAnnum",
	Slab:"Slab",
	Tiered:"Tiered",
}

// enum type LimitScope
export let LimitScope = {
	PerTransaction:"PerTransaction",
	Daily:"Daily",
	Monthly:"Monthly",
	Yearly:"Yearly",
	Rolling24h:"Rolling24h",
}

// enum type LimitPeriod
export let LimitPeriod = {
	None:"None",
	Day:"Day",
	Week:"Week",
	Month:"Month",
	Year:"Year",
}

// enum type CustomerType
export let CustomerType = {
	Individual:"Individual",
	Business:"Business",
}

// enum type KYCStatus
export let KYCStatus = {
	Pending:"Pending",
	Verified:"Verified",
	Rejected:"Rejected",
	Expired:"Expired",
}

// enum type VerificationLevel
export let VerificationLevel = {
	Basic:"Basic",
	Standard:"Standard",
	Enhanced:"Enhanced",
}

// enum type KYCDocumentType
export let KYCDocumentType = {
	Passport:"Passport",
	NationalID:"NationalID",
	DriverLicense:"DriverLicense",
	BusinessRegistration:"BusinessRegistration",
	ProofOfAddress:"ProofOfAddress",
}

// enum type DocumentStatus
export let DocumentStatus = {
	Submitted:"Submitted",
	Approved:"Approved",
	Rejected:"Rejected",
	Expired:"Expired",
}

// enum type ScreeningType
export let ScreeningType = {
	Sanctions:"Sanctions",
	PEP:"PEP",
	AdverseMedia:"AdverseMedia",
}

// enum type ScreeningStatus
export let ScreeningStatus = {
	Clear:"Clear",
	Review:"Review",
	Match:"Match",
}

// enum type VerificationStatus
export let VerificationStatus = {
	Unverified:"Unverified",
	Verified:"Verified",
	Failed:"Failed",
}

// enum type PolicyStatus
export let PolicyStatus = {
	Draft:"Draft",
	Active:"Active",
	Retired:"Retired",
}

// enum type AlertSeverity
export let AlertSeverity = {
	Low:"Low",
	Medium:"Medium",
	High:"High",
	Critical:"Critical",
}

// enum type AlertStatus
export let AlertStatus = {
	Open:"Open",
	Investigating:"Investigating",
	Resolved:"Resolved",
	Dismissed:"Dismissed",
}

// enum type ConsentType
export let ConsentType = {
	DataAccess:"DataAccess",
	PaymentInitiation:"PaymentInitiation",
}

// enum type ConsentStatus
export let ConsentStatus = {
	Active:"Active",
	Revoked:"Revoked",
	Expired:"Expired",
}

// enum type ClientType
export let ClientType = {
	Confidential:"Confidential",
	Public:"Public",
}

// enum type AgreementType
export let AgreementType = {
	TermsOfService:"TermsOfService",
	PrivacyPolicy:"PrivacyPolicy",
	LoanAgreement:"LoanAgreement",
	AccountAgreement:"AccountAgreement",
}

// enum type AgreementStatus
export let AgreementStatus = {
	Active:"Active",
	Suspended:"Suspended",
	Terminated:"Terminated",
}

// enum type AccountType
export let AccountType = {
	Checking:"Checking",
	Savings:"Savings",
	Current:"Current",
	Brokerage:"Brokerage",
	Settlement:"Settlement",
	Escrow:"Escrow",
}

// enum type AccountStatus
export let AccountStatus = {
	Pending:"Pending",
	Active:"Active",
	Frozen:"Frozen",
	Closed:"Closed",
}

// enum type WalletStatus
export let WalletStatus = {
	Active:"Active",
	Suspended:"Suspended",
	Closed:"Closed",
}

// enum type CardScheme
export let CardScheme = {
	Visa:"Visa",
	Mastercard:"Mastercard",
	Amex:"Amex",
	Discover:"Discover",
	UnionPay:"UnionPay",
}

// enum type CardStatus
export let CardStatus = {
	Active:"Active",
	Blocked:"Blocked",
	Closed:"Closed",
	Expired:"Expired",
}

// enum type WalletProvider
export let WalletProvider = {
	ApplePay:"ApplePay",
	GooglePay:"GooglePay",
	SamsungPay:"SamsungPay",
	Other:"Other",
}

// enum type TokenizationStatus
export let TokenizationStatus = {
	Active:"Active",
	Suspended:"Suspended",
	Deactivated:"Deactivated",
}

// enum type TerminalType
export let TerminalType = {
	POS:"POS",
	mPOS:"mPOS",
	ECommerce:"ECommerce",
}

// enum type TerminalStatus
export let TerminalStatus = {
	Active:"Active",
	Inactive:"Inactive",
	Decommissioned:"Decommissioned",
}

// enum type ContractStatus
export let ContractStatus = {
	Draft:"Draft",
	Active:"Active",
	Suspended:"Suspended",
	Terminated:"Terminated",
}

// enum type TransactionType
export let TransactionType = {
	Deposit:"Deposit",
	Withdrawal:"Withdrawal",
	Transfer:"Transfer",
	Payment:"Payment",
	Refund:"Refund",
	Fee:"Fee",
	Interest:"Interest",
	FXConversion:"FXConversion",
}

// enum type TransactionStatus
export let TransactionStatus = {
	Pending:"Pending",
	Authorized:"Authorized",
	Posted:"Posted",
	Settled:"Settled",
	Reversed:"Reversed",
	Failed:"Failed",
}

// enum type PaymentMethod
export let PaymentMethod = {
	Card:"Card",
	BankTransfer:"BankTransfer",
	DirectDebit:"DirectDebit",
	Wallet:"Wallet",
	Cash:"Cash",
}

// enum type PaymentOrderStatus
export let PaymentOrderStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	Processing:"Processing",
	Completed:"Completed",
	Cancelled:"Cancelled",
	Failed:"Failed",
}

// enum type PaymentPriority
export let PaymentPriority = {
	Normal:"Normal",
	Urgent:"Urgent",
}

// enum type FXPriceType
export let FXPriceType = {
	Indicative:"Indicative",
	Firm:"Firm",
}

// enum type FXDealStatus
export let FXDealStatus = {
	Booked:"Booked",
	Cancelled:"Cancelled",
	Settled:"Settled",
}

// enum type SettlementStatus
export let SettlementStatus = {
	Open:"Open",
	Processing:"Processing",
	Closed:"Closed",
	Reconciled:"Reconciled",
}

// enum type PayoutStatus
export let PayoutStatus = {
	Scheduled:"Scheduled",
	Processing:"Processing",
	Paid:"Paid",
	Failed:"Failed",
}

// enum type DisputeReason
export let DisputeReason = {
	Fraud:"Fraud",
	Duplicate:"Duplicate",
	NotAsDescribed:"NotAsDescribed",
	NotReceived:"NotReceived",
	ProcessingError:"ProcessingError",
}

// enum type DisputeStatus
export let DisputeStatus = {
	Open:"Open",
	Represented:"Represented",
	Won:"Won",
	Lost:"Lost",
	Closed:"Closed",
}

// enum type ChargebackStage
export let ChargebackStage = {
	FirstChargeback:"FirstChargeback",
	SecondChargeback:"SecondChargeback",
	Arbitration:"Arbitration",
}

// enum type ChargebackStatus
export let ChargebackStatus = {
	Pending:"Pending",
	Accepted:"Accepted",
	Reversed:"Reversed",
	Lost:"Lost",
}

// enum type InvoiceStatus
export let InvoiceStatus = {
	Draft:"Draft",
	Issued:"Issued",
	Paid:"Paid",
	Overdue:"Overdue",
	Cancelled:"Cancelled",
}

// enum type DirectDebitScheme
export let DirectDebitScheme = {
	SEPA:"SEPA",
	ACH:"ACH",
	BACS:"BACS",
	BECS:"BECS",
}

// enum type MandateStatus
export let MandateStatus = {
	Active:"Active",
	Suspended:"Suspended",
	Cancelled:"Cancelled",
	Expired:"Expired",
}

// enum type LoanProductType
export let LoanProductType = {
	PersonalLoan:"PersonalLoan",
	Mortgage:"Mortgage",
	InstallmentLoan:"InstallmentLoan",
	CreditLine:"CreditLine",
	SME:"SME",
}

// enum type LoanPurpose
export let LoanPurpose = {
	HomeImprovement:"HomeImprovement",
	Education:"Education",
	DebtConsolidation:"DebtConsolidation",
	Business:"Business",
	Other:"Other",
}

// enum type ApplicationStatus
export let ApplicationStatus = {
	Draft:"Draft",
	Submitted:"Submitted",
	Underwriting:"Underwriting",
	Approved:"Approved",
	Declined:"Declined",
	Withdrawn:"Withdrawn",
}

// enum type DecisionOutcome
export let DecisionOutcome = {
	Approve:"Approve",
	Decline:"Decline",
	Refer:"Refer",
}

// enum type InterestRateType
export let InterestRateType = {
	Fixed:"Fixed",
	Variable:"Variable",
}

// enum type LoanStatus
export let LoanStatus = {
	Active:"Active",
	Delinquent:"Delinquent",
	Closed:"Closed",
	ChargedOff:"ChargedOff",
}

// enum type InstallmentStatus
export let InstallmentStatus = {
	Scheduled:"Scheduled",
	Paid:"Paid",
	Overdue:"Overdue",
	Waived:"Waived",
}

// enum type CollateralType
export let CollateralType = {
	RealEstate:"RealEstate",
	Deposit:"Deposit",
	PersonalGuarantee:"PersonalGuarantee",
	Inventory:"Inventory",
	Equipment:"Equipment",
}

// enum type LoanTransactionType
export let LoanTransactionType = {
	Disbursement:"Disbursement",
	Repayment:"Repayment",
	Interest:"Interest",
	Fee:"Fee",
	Reversal:"Reversal",
}

// enum type PostingStatus
export let PostingStatus = {
	Pending:"Pending",
	Posted:"Posted",
	Reversed:"Reversed",
}

// enum type PortfolioStatus
export let PortfolioStatus = {
	Active:"Active",
	Closed:"Closed",
	Suspended:"Suspended",
}

// enum type InvestmentAccountType
export let InvestmentAccountType = {
	Brokerage:"Brokerage",
	Retirement:"Retirement",
	Custody:"Custody",
	Margin:"Margin",
}

// enum type SecurityType
export let SecurityType = {
	Equity:"Equity",
	Bond:"Bond",
	ETF:"ETF",
	MutualFund:"MutualFund",
	Derivative:"Derivative",
	Crypto:"Crypto",
}

// enum type OrderSide
export let OrderSide = {
	Buy:"Buy",
	Sell:"Sell",
}

// enum type OrderType
export let OrderType = {
	Market:"Market",
	Limit:"Limit",
	Stop:"Stop",
	StopLimit:"StopLimit",
}

// enum type OrderStatus
export let OrderStatus = {
	New:"New",
	PartiallyFilled:"PartiallyFilled",
	Filled:"Filled",
	Cancelled:"Cancelled",
	Rejected:"Rejected",
	Expired:"Expired",
}

// enum type TimeInForce
export let TimeInForce = {
	Day:"Day",
	GTC:"GTC",
	IOC:"IOC",
	FOK:"FOK",
}
