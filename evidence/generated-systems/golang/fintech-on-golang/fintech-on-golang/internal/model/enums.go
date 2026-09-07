package model


//==============================================================
// ProductCategory Declaration
//==============================================================
type ProductCategory int
const (
    ProductCategoryChecking ProductCategory = iota
	ProductCategorySavings
	ProductCategoryCreditCard
	ProductCategoryLoan
	ProductCategoryInvestment
	ProductCategoryInsurance
	ProductCategoryPayments
	ProductCategoryFX
	ProductCategoryWallet
)


//==============================================================
// PlanStatus Declaration
//==============================================================
type PlanStatus int
const (
    PlanStatusDraft PlanStatus = iota
	PlanStatusActive
	PlanStatusSuspended
	PlanStatusArchived
)


//==============================================================
// FeeType Declaration
//==============================================================
type FeeType int
const (
    FeeTypeFixed FeeType = iota
	FeeTypePercentage
	FeeTypeTiered
	FeeTypeInterchange
	FeeTypeNetwork
	FeeTypeChargeback
	FeeTypeATM
	FeeTypeFX
)


//==============================================================
// FeeCalculationMethod Declaration
//==============================================================
type FeeCalculationMethod int
const (
    FeeCalculationMethodPerTransaction FeeCalculationMethod = iota
	FeeCalculationMethodPerMonth
	FeeCalculationMethodPerAnnum
	FeeCalculationMethodSlab
	FeeCalculationMethodTiered
)


//==============================================================
// LimitScope Declaration
//==============================================================
type LimitScope int
const (
    LimitScopePerTransaction LimitScope = iota
	LimitScopeDaily
	LimitScopeMonthly
	LimitScopeYearly
	LimitScopeRolling24h
)


//==============================================================
// LimitPeriod Declaration
//==============================================================
type LimitPeriod int
const (
    LimitPeriodNone LimitPeriod = iota
	LimitPeriodDay
	LimitPeriodWeek
	LimitPeriodMonth
	LimitPeriodYear
)


//==============================================================
// CustomerType Declaration
//==============================================================
type CustomerType int
const (
    CustomerTypeIndividual CustomerType = iota
	CustomerTypeBusiness
)


//==============================================================
// KYCStatus Declaration
//==============================================================
type KYCStatus int
const (
    KYCStatusPending KYCStatus = iota
	KYCStatusVerified
	KYCStatusRejected
	KYCStatusExpired
)


//==============================================================
// VerificationLevel Declaration
//==============================================================
type VerificationLevel int
const (
    VerificationLevelBasic VerificationLevel = iota
	VerificationLevelStandard
	VerificationLevelEnhanced
)


//==============================================================
// KYCDocumentType Declaration
//==============================================================
type KYCDocumentType int
const (
    KYCDocumentTypePassport KYCDocumentType = iota
	KYCDocumentTypeNationalID
	KYCDocumentTypeDriverLicense
	KYCDocumentTypeBusinessRegistration
	KYCDocumentTypeProofOfAddress
)


//==============================================================
// DocumentStatus Declaration
//==============================================================
type DocumentStatus int
const (
    DocumentStatusSubmitted DocumentStatus = iota
	DocumentStatusApproved
	DocumentStatusRejected
	DocumentStatusExpired
)


//==============================================================
// ScreeningType Declaration
//==============================================================
type ScreeningType int
const (
    ScreeningTypeSanctions ScreeningType = iota
	ScreeningTypePEP
	ScreeningTypeAdverseMedia
)


//==============================================================
// ScreeningStatus Declaration
//==============================================================
type ScreeningStatus int
const (
    ScreeningStatusClear ScreeningStatus = iota
	ScreeningStatusReview
	ScreeningStatusMatch
)


//==============================================================
// VerificationStatus Declaration
//==============================================================
type VerificationStatus int
const (
    VerificationStatusUnverified VerificationStatus = iota
	VerificationStatusVerified
	VerificationStatusFailed
)


//==============================================================
// PolicyStatus Declaration
//==============================================================
type PolicyStatus int
const (
    PolicyStatusDraft PolicyStatus = iota
	PolicyStatusActive
	PolicyStatusRetired
)


//==============================================================
// AlertSeverity Declaration
//==============================================================
type AlertSeverity int
const (
    AlertSeverityLow AlertSeverity = iota
	AlertSeverityMedium
	AlertSeverityHigh
	AlertSeverityCritical
)


//==============================================================
// AlertStatus Declaration
//==============================================================
type AlertStatus int
const (
    AlertStatusOpen AlertStatus = iota
	AlertStatusInvestigating
	AlertStatusResolved
	AlertStatusDismissed
)


//==============================================================
// ConsentType Declaration
//==============================================================
type ConsentType int
const (
    ConsentTypeDataAccess ConsentType = iota
	ConsentTypePaymentInitiation
)


//==============================================================
// ConsentStatus Declaration
//==============================================================
type ConsentStatus int
const (
    ConsentStatusActive ConsentStatus = iota
	ConsentStatusRevoked
	ConsentStatusExpired
)


//==============================================================
// ClientType Declaration
//==============================================================
type ClientType int
const (
    ClientTypeConfidential ClientType = iota
	ClientTypePublic
)


//==============================================================
// AgreementType Declaration
//==============================================================
type AgreementType int
const (
    AgreementTypeTermsOfService AgreementType = iota
	AgreementTypePrivacyPolicy
	AgreementTypeLoanAgreement
	AgreementTypeAccountAgreement
)


//==============================================================
// AgreementStatus Declaration
//==============================================================
type AgreementStatus int
const (
    AgreementStatusActive AgreementStatus = iota
	AgreementStatusSuspended
	AgreementStatusTerminated
)


//==============================================================
// AccountType Declaration
//==============================================================
type AccountType int
const (
    AccountTypeChecking AccountType = iota
	AccountTypeSavings
	AccountTypeCurrent
	AccountTypeBrokerage
	AccountTypeSettlement
	AccountTypeEscrow
)


//==============================================================
// AccountStatus Declaration
//==============================================================
type AccountStatus int
const (
    AccountStatusPending AccountStatus = iota
	AccountStatusActive
	AccountStatusFrozen
	AccountStatusClosed
)


//==============================================================
// WalletStatus Declaration
//==============================================================
type WalletStatus int
const (
    WalletStatusActive WalletStatus = iota
	WalletStatusSuspended
	WalletStatusClosed
)


//==============================================================
// CardScheme Declaration
//==============================================================
type CardScheme int
const (
    CardSchemeVisa CardScheme = iota
	CardSchemeMastercard
	CardSchemeAmex
	CardSchemeDiscover
	CardSchemeUnionPay
)


//==============================================================
// CardStatus Declaration
//==============================================================
type CardStatus int
const (
    CardStatusActive CardStatus = iota
	CardStatusBlocked
	CardStatusClosed
	CardStatusExpired
)


//==============================================================
// WalletProvider Declaration
//==============================================================
type WalletProvider int
const (
    WalletProviderApplePay WalletProvider = iota
	WalletProviderGooglePay
	WalletProviderSamsungPay
	WalletProviderOther
)


//==============================================================
// TokenizationStatus Declaration
//==============================================================
type TokenizationStatus int
const (
    TokenizationStatusActive TokenizationStatus = iota
	TokenizationStatusSuspended
	TokenizationStatusDeactivated
)


//==============================================================
// TerminalType Declaration
//==============================================================
type TerminalType int
const (
    TerminalTypePOS TerminalType = iota
	TerminalTypemPOS
	TerminalTypeECommerce
)


//==============================================================
// TerminalStatus Declaration
//==============================================================
type TerminalStatus int
const (
    TerminalStatusActive TerminalStatus = iota
	TerminalStatusInactive
	TerminalStatusDecommissioned
)


//==============================================================
// ContractStatus Declaration
//==============================================================
type ContractStatus int
const (
    ContractStatusDraft ContractStatus = iota
	ContractStatusActive
	ContractStatusSuspended
	ContractStatusTerminated
)


//==============================================================
// TransactionType Declaration
//==============================================================
type TransactionType int
const (
    TransactionTypeDeposit TransactionType = iota
	TransactionTypeWithdrawal
	TransactionTypeTransfer
	TransactionTypePayment
	TransactionTypeRefund
	TransactionTypeFee
	TransactionTypeInterest
	TransactionTypeFXConversion
)


//==============================================================
// TransactionStatus Declaration
//==============================================================
type TransactionStatus int
const (
    TransactionStatusPending TransactionStatus = iota
	TransactionStatusAuthorized
	TransactionStatusPosted
	TransactionStatusSettled
	TransactionStatusReversed
	TransactionStatusFailed
)


//==============================================================
// PaymentMethod Declaration
//==============================================================
type PaymentMethod int
const (
    PaymentMethodCard PaymentMethod = iota
	PaymentMethodBankTransfer
	PaymentMethodDirectDebit
	PaymentMethodWallet
	PaymentMethodCash
)


//==============================================================
// PaymentOrderStatus Declaration
//==============================================================
type PaymentOrderStatus int
const (
    PaymentOrderStatusDraft PaymentOrderStatus = iota
	PaymentOrderStatusSubmitted
	PaymentOrderStatusProcessing
	PaymentOrderStatusCompleted
	PaymentOrderStatusCancelled
	PaymentOrderStatusFailed
)


//==============================================================
// PaymentPriority Declaration
//==============================================================
type PaymentPriority int
const (
    PaymentPriorityNormal PaymentPriority = iota
	PaymentPriorityUrgent
)


//==============================================================
// FXPriceType Declaration
//==============================================================
type FXPriceType int
const (
    FXPriceTypeIndicative FXPriceType = iota
	FXPriceTypeFirm
)


//==============================================================
// FXDealStatus Declaration
//==============================================================
type FXDealStatus int
const (
    FXDealStatusBooked FXDealStatus = iota
	FXDealStatusCancelled
	FXDealStatusSettled
)


//==============================================================
// SettlementStatus Declaration
//==============================================================
type SettlementStatus int
const (
    SettlementStatusOpen SettlementStatus = iota
	SettlementStatusProcessing
	SettlementStatusClosed
	SettlementStatusReconciled
)


//==============================================================
// PayoutStatus Declaration
//==============================================================
type PayoutStatus int
const (
    PayoutStatusScheduled PayoutStatus = iota
	PayoutStatusProcessing
	PayoutStatusPaid
	PayoutStatusFailed
)


//==============================================================
// DisputeReason Declaration
//==============================================================
type DisputeReason int
const (
    DisputeReasonFraud DisputeReason = iota
	DisputeReasonDuplicate
	DisputeReasonNotAsDescribed
	DisputeReasonNotReceived
	DisputeReasonProcessingError
)


//==============================================================
// DisputeStatus Declaration
//==============================================================
type DisputeStatus int
const (
    DisputeStatusOpen DisputeStatus = iota
	DisputeStatusRepresented
	DisputeStatusWon
	DisputeStatusLost
	DisputeStatusClosed
)


//==============================================================
// ChargebackStage Declaration
//==============================================================
type ChargebackStage int
const (
    ChargebackStageFirstChargeback ChargebackStage = iota
	ChargebackStageSecondChargeback
	ChargebackStageArbitration
)


//==============================================================
// ChargebackStatus Declaration
//==============================================================
type ChargebackStatus int
const (
    ChargebackStatusPending ChargebackStatus = iota
	ChargebackStatusAccepted
	ChargebackStatusReversed
	ChargebackStatusLost
)


//==============================================================
// InvoiceStatus Declaration
//==============================================================
type InvoiceStatus int
const (
    InvoiceStatusDraft InvoiceStatus = iota
	InvoiceStatusIssued
	InvoiceStatusPaid
	InvoiceStatusOverdue
	InvoiceStatusCancelled
)


//==============================================================
// DirectDebitScheme Declaration
//==============================================================
type DirectDebitScheme int
const (
    DirectDebitSchemeSEPA DirectDebitScheme = iota
	DirectDebitSchemeACH
	DirectDebitSchemeBACS
	DirectDebitSchemeBECS
)


//==============================================================
// MandateStatus Declaration
//==============================================================
type MandateStatus int
const (
    MandateStatusActive MandateStatus = iota
	MandateStatusSuspended
	MandateStatusCancelled
	MandateStatusExpired
)


//==============================================================
// LoanProductType Declaration
//==============================================================
type LoanProductType int
const (
    LoanProductTypePersonalLoan LoanProductType = iota
	LoanProductTypeMortgage
	LoanProductTypeInstallmentLoan
	LoanProductTypeCreditLine
	LoanProductTypeSME
)


//==============================================================
// LoanPurpose Declaration
//==============================================================
type LoanPurpose int
const (
    LoanPurposeHomeImprovement LoanPurpose = iota
	LoanPurposeEducation
	LoanPurposeDebtConsolidation
	LoanPurposeBusiness
	LoanPurposeOther
)


//==============================================================
// ApplicationStatus Declaration
//==============================================================
type ApplicationStatus int
const (
    ApplicationStatusDraft ApplicationStatus = iota
	ApplicationStatusSubmitted
	ApplicationStatusUnderwriting
	ApplicationStatusApproved
	ApplicationStatusDeclined
	ApplicationStatusWithdrawn
)


//==============================================================
// DecisionOutcome Declaration
//==============================================================
type DecisionOutcome int
const (
    DecisionOutcomeApprove DecisionOutcome = iota
	DecisionOutcomeDecline
	DecisionOutcomeRefer
)


//==============================================================
// InterestRateType Declaration
//==============================================================
type InterestRateType int
const (
    InterestRateTypeFixed InterestRateType = iota
	InterestRateTypeVariable
)


//==============================================================
// LoanStatus Declaration
//==============================================================
type LoanStatus int
const (
    LoanStatusActive LoanStatus = iota
	LoanStatusDelinquent
	LoanStatusClosed
	LoanStatusChargedOff
)


//==============================================================
// InstallmentStatus Declaration
//==============================================================
type InstallmentStatus int
const (
    InstallmentStatusScheduled InstallmentStatus = iota
	InstallmentStatusPaid
	InstallmentStatusOverdue
	InstallmentStatusWaived
)


//==============================================================
// CollateralType Declaration
//==============================================================
type CollateralType int
const (
    CollateralTypeRealEstate CollateralType = iota
	CollateralTypeDeposit
	CollateralTypePersonalGuarantee
	CollateralTypeInventory
	CollateralTypeEquipment
)


//==============================================================
// LoanTransactionType Declaration
//==============================================================
type LoanTransactionType int
const (
    LoanTransactionTypeDisbursement LoanTransactionType = iota
	LoanTransactionTypeRepayment
	LoanTransactionTypeInterest
	LoanTransactionTypeFee
	LoanTransactionTypeReversal
)


//==============================================================
// PostingStatus Declaration
//==============================================================
type PostingStatus int
const (
    PostingStatusPending PostingStatus = iota
	PostingStatusPosted
	PostingStatusReversed
)


//==============================================================
// PortfolioStatus Declaration
//==============================================================
type PortfolioStatus int
const (
    PortfolioStatusActive PortfolioStatus = iota
	PortfolioStatusClosed
	PortfolioStatusSuspended
)


//==============================================================
// InvestmentAccountType Declaration
//==============================================================
type InvestmentAccountType int
const (
    InvestmentAccountTypeBrokerage InvestmentAccountType = iota
	InvestmentAccountTypeRetirement
	InvestmentAccountTypeCustody
	InvestmentAccountTypeMargin
)


//==============================================================
// SecurityType Declaration
//==============================================================
type SecurityType int
const (
    SecurityTypeEquity SecurityType = iota
	SecurityTypeBond
	SecurityTypeETF
	SecurityTypeMutualFund
	SecurityTypeDerivative
	SecurityTypeCrypto
)


//==============================================================
// OrderSide Declaration
//==============================================================
type OrderSide int
const (
    OrderSideBuy OrderSide = iota
	OrderSideSell
)


//==============================================================
// OrderType Declaration
//==============================================================
type OrderType int
const (
    OrderTypeMarket OrderType = iota
	OrderTypeLimit
	OrderTypeStop
	OrderTypeStopLimit
)


//==============================================================
// OrderStatus Declaration
//==============================================================
type OrderStatus int
const (
    OrderStatusNew OrderStatus = iota
	OrderStatusPartiallyFilled
	OrderStatusFilled
	OrderStatusCancelled
	OrderStatusRejected
	OrderStatusExpired
)


//==============================================================
// TimeInForce Declaration
//==============================================================
type TimeInForce int
const (
    TimeInForceDay TimeInForce = iota
	TimeInForceGTC
	TimeInForceIOC
	TimeInForceFOK
)

