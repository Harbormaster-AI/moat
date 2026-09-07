// routerConfig.ts

import { Routes } from '@angular/router';
import { CreateFinancialInstitutionComponent } from './components/FinancialInstitution/create/create.component';
import { EditFinancialInstitutionComponent } from './components/FinancialInstitution/edit/edit.component';
import { IndexFinancialInstitutionComponent } from './components/FinancialInstitution/index/index.component';
import { CreateBranchComponent } from './components/Branch/create/create.component';
import { EditBranchComponent } from './components/Branch/edit/edit.component';
import { IndexBranchComponent } from './components/Branch/index/index.component';
import { CreateProductOfferingComponent } from './components/ProductOffering/create/create.component';
import { EditProductOfferingComponent } from './components/ProductOffering/edit/edit.component';
import { IndexProductOfferingComponent } from './components/ProductOffering/index/index.component';
import { CreatePricingPlanComponent } from './components/PricingPlan/create/create.component';
import { EditPricingPlanComponent } from './components/PricingPlan/edit/edit.component';
import { IndexPricingPlanComponent } from './components/PricingPlan/index/index.component';
import { CreateFeeScheduleComponent } from './components/FeeSchedule/create/create.component';
import { EditFeeScheduleComponent } from './components/FeeSchedule/edit/edit.component';
import { IndexFeeScheduleComponent } from './components/FeeSchedule/index/index.component';
import { CreateUsageLimitComponent } from './components/UsageLimit/create/create.component';
import { EditUsageLimitComponent } from './components/UsageLimit/edit/edit.component';
import { IndexUsageLimitComponent } from './components/UsageLimit/index/index.component';
import { CreateCustomerComponent } from './components/Customer/create/create.component';
import { EditCustomerComponent } from './components/Customer/edit/edit.component';
import { IndexCustomerComponent } from './components/Customer/index/index.component';
import { CreateKYCProfileComponent } from './components/KYCProfile/create/create.component';
import { EditKYCProfileComponent } from './components/KYCProfile/edit/edit.component';
import { IndexKYCProfileComponent } from './components/KYCProfile/index/index.component';
import { CreateKYCDocumentComponent } from './components/KYCDocument/create/create.component';
import { EditKYCDocumentComponent } from './components/KYCDocument/edit/edit.component';
import { IndexKYCDocumentComponent } from './components/KYCDocument/index/index.component';
import { CreateScreeningComponent } from './components/Screening/create/create.component';
import { EditScreeningComponent } from './components/Screening/edit/edit.component';
import { IndexScreeningComponent } from './components/Screening/index/index.component';
import { CreateVerifiedAddressComponent } from './components/VerifiedAddress/create/create.component';
import { EditVerifiedAddressComponent } from './components/VerifiedAddress/edit/edit.component';
import { IndexVerifiedAddressComponent } from './components/VerifiedAddress/index/index.component';
import { CreateCompliancePolicyComponent } from './components/CompliancePolicy/create/create.component';
import { EditCompliancePolicyComponent } from './components/CompliancePolicy/edit/edit.component';
import { IndexCompliancePolicyComponent } from './components/CompliancePolicy/index/index.component';
import { CreateComplianceAlertComponent } from './components/ComplianceAlert/create/create.component';
import { EditComplianceAlertComponent } from './components/ComplianceAlert/edit/edit.component';
import { IndexComplianceAlertComponent } from './components/ComplianceAlert/index/index.component';
import { CreateConsentComponent } from './components/Consent/create/create.component';
import { EditConsentComponent } from './components/Consent/edit/edit.component';
import { IndexConsentComponent } from './components/Consent/index/index.component';
import { CreateAPIClientComponent } from './components/APIClient/create/create.component';
import { EditAPIClientComponent } from './components/APIClient/edit/edit.component';
import { IndexAPIClientComponent } from './components/APIClient/index/index.component';
import { CreateAgreementComponent } from './components/Agreement/create/create.component';
import { EditAgreementComponent } from './components/Agreement/edit/edit.component';
import { IndexAgreementComponent } from './components/Agreement/index/index.component';
import { CreateAccountComponent } from './components/Account/create/create.component';
import { EditAccountComponent } from './components/Account/edit/edit.component';
import { IndexAccountComponent } from './components/Account/index/index.component';
import { CreateWalletComponent } from './components/Wallet/create/create.component';
import { EditWalletComponent } from './components/Wallet/edit/edit.component';
import { IndexWalletComponent } from './components/Wallet/index/index.component';
import { CreatePaymentCardComponent } from './components/PaymentCard/create/create.component';
import { EditPaymentCardComponent } from './components/PaymentCard/edit/edit.component';
import { IndexPaymentCardComponent } from './components/PaymentCard/index/index.component';
import { CreateCardTokenizationComponent } from './components/CardTokenization/create/create.component';
import { EditCardTokenizationComponent } from './components/CardTokenization/edit/edit.component';
import { IndexCardTokenizationComponent } from './components/CardTokenization/index/index.component';
import { CreateMerchantComponent } from './components/Merchant/create/create.component';
import { EditMerchantComponent } from './components/Merchant/edit/edit.component';
import { IndexMerchantComponent } from './components/Merchant/index/index.component';
import { CreateTerminalComponent } from './components/Terminal/create/create.component';
import { EditTerminalComponent } from './components/Terminal/edit/edit.component';
import { IndexTerminalComponent } from './components/Terminal/index/index.component';
import { CreatePaymentContractComponent } from './components/PaymentContract/create/create.component';
import { EditPaymentContractComponent } from './components/PaymentContract/edit/edit.component';
import { IndexPaymentContractComponent } from './components/PaymentContract/index/index.component';
import { CreatePaymentProcessorComponent } from './components/PaymentProcessor/create/create.component';
import { EditPaymentProcessorComponent } from './components/PaymentProcessor/edit/edit.component';
import { IndexPaymentProcessorComponent } from './components/PaymentProcessor/index/index.component';
import { CreateTransactionComponent } from './components/Transaction/create/create.component';
import { EditTransactionComponent } from './components/Transaction/edit/edit.component';
import { IndexTransactionComponent } from './components/Transaction/index/index.component';
import { CreatePaymentOrderComponent } from './components/PaymentOrder/create/create.component';
import { EditPaymentOrderComponent } from './components/PaymentOrder/edit/edit.component';
import { IndexPaymentOrderComponent } from './components/PaymentOrder/index/index.component';
import { CreateBeneficiaryComponent } from './components/Beneficiary/create/create.component';
import { EditBeneficiaryComponent } from './components/Beneficiary/edit/edit.component';
import { IndexBeneficiaryComponent } from './components/Beneficiary/index/index.component';
import { CreateAppliedFeeComponent } from './components/AppliedFee/create/create.component';
import { EditAppliedFeeComponent } from './components/AppliedFee/edit/edit.component';
import { IndexAppliedFeeComponent } from './components/AppliedFee/index/index.component';
import { CreateFXQuoteComponent } from './components/FXQuote/create/create.component';
import { EditFXQuoteComponent } from './components/FXQuote/edit/edit.component';
import { IndexFXQuoteComponent } from './components/FXQuote/index/index.component';
import { CreateFXDealComponent } from './components/FXDeal/create/create.component';
import { EditFXDealComponent } from './components/FXDeal/edit/edit.component';
import { IndexFXDealComponent } from './components/FXDeal/index/index.component';
import { CreateSettlementBatchComponent } from './components/SettlementBatch/create/create.component';
import { EditSettlementBatchComponent } from './components/SettlementBatch/edit/edit.component';
import { IndexSettlementBatchComponent } from './components/SettlementBatch/index/index.component';
import { CreatePayoutComponent } from './components/Payout/create/create.component';
import { EditPayoutComponent } from './components/Payout/edit/edit.component';
import { IndexPayoutComponent } from './components/Payout/index/index.component';
import { CreateDisputeComponent } from './components/Dispute/create/create.component';
import { EditDisputeComponent } from './components/Dispute/edit/edit.component';
import { IndexDisputeComponent } from './components/Dispute/index/index.component';
import { CreateChargebackComponent } from './components/Chargeback/create/create.component';
import { EditChargebackComponent } from './components/Chargeback/edit/edit.component';
import { IndexChargebackComponent } from './components/Chargeback/index/index.component';
import { CreateInvoiceComponent } from './components/Invoice/create/create.component';
import { EditInvoiceComponent } from './components/Invoice/edit/edit.component';
import { IndexInvoiceComponent } from './components/Invoice/index/index.component';
import { CreateAccountStatementComponent } from './components/AccountStatement/create/create.component';
import { EditAccountStatementComponent } from './components/AccountStatement/edit/edit.component';
import { IndexAccountStatementComponent } from './components/AccountStatement/index/index.component';
import { CreateDirectDebitMandateComponent } from './components/DirectDebitMandate/create/create.component';
import { EditDirectDebitMandateComponent } from './components/DirectDebitMandate/edit/edit.component';
import { IndexDirectDebitMandateComponent } from './components/DirectDebitMandate/index/index.component';
import { CreateCreditorComponent } from './components/Creditor/create/create.component';
import { EditCreditorComponent } from './components/Creditor/edit/edit.component';
import { IndexCreditorComponent } from './components/Creditor/index/index.component';
import { CreateLoanApplicationComponent } from './components/LoanApplication/create/create.component';
import { EditLoanApplicationComponent } from './components/LoanApplication/edit/edit.component';
import { IndexLoanApplicationComponent } from './components/LoanApplication/index/index.component';
import { CreateRiskAssessmentComponent } from './components/RiskAssessment/create/create.component';
import { EditRiskAssessmentComponent } from './components/RiskAssessment/edit/edit.component';
import { IndexRiskAssessmentComponent } from './components/RiskAssessment/index/index.component';
import { CreateLoanComponent } from './components/Loan/create/create.component';
import { EditLoanComponent } from './components/Loan/edit/edit.component';
import { IndexLoanComponent } from './components/Loan/index/index.component';
import { CreateRepaymentScheduleComponent } from './components/RepaymentSchedule/create/create.component';
import { EditRepaymentScheduleComponent } from './components/RepaymentSchedule/edit/edit.component';
import { IndexRepaymentScheduleComponent } from './components/RepaymentSchedule/index/index.component';
import { CreateCollateralComponent } from './components/Collateral/create/create.component';
import { EditCollateralComponent } from './components/Collateral/edit/edit.component';
import { IndexCollateralComponent } from './components/Collateral/index/index.component';
import { CreateLoanTransactionComponent } from './components/LoanTransaction/create/create.component';
import { EditLoanTransactionComponent } from './components/LoanTransaction/edit/edit.component';
import { IndexLoanTransactionComponent } from './components/LoanTransaction/index/index.component';
import { CreateInvestmentPortfolioComponent } from './components/InvestmentPortfolio/create/create.component';
import { EditInvestmentPortfolioComponent } from './components/InvestmentPortfolio/edit/edit.component';
import { IndexInvestmentPortfolioComponent } from './components/InvestmentPortfolio/index/index.component';
import { CreateInvestmentAccountComponent } from './components/InvestmentAccount/create/create.component';
import { EditInvestmentAccountComponent } from './components/InvestmentAccount/edit/edit.component';
import { IndexInvestmentAccountComponent } from './components/InvestmentAccount/index/index.component';
import { CreateSecurityComponent } from './components/Security/create/create.component';
import { EditSecurityComponent } from './components/Security/edit/edit.component';
import { IndexSecurityComponent } from './components/Security/index/index.component';
import { CreatePositionComponent } from './components/Position/create/create.component';
import { EditPositionComponent } from './components/Position/edit/edit.component';
import { IndexPositionComponent } from './components/Position/index/index.component';
import { CreateTradeOrderComponent } from './components/TradeOrder/create/create.component';
import { EditTradeOrderComponent } from './components/TradeOrder/edit/edit.component';
import { IndexTradeOrderComponent } from './components/TradeOrder/index/index.component';
import { CreateTradeComponent } from './components/Trade/create/create.component';
import { EditTradeComponent } from './components/Trade/edit/edit.component';
import { IndexTradeComponent } from './components/Trade/index/index.component';
import { CreateExchangeRateComponent } from './components/ExchangeRate/create/create.component';
import { EditExchangeRateComponent } from './components/ExchangeRate/edit/edit.component';
import { IndexExchangeRateComponent } from './components/ExchangeRate/index/index.component';

export const FinancialInstitutionRoutes: Routes = [
  { path: 'createFinancialInstitution',
    component: CreateFinancialInstitutionComponent
  },
  {
    path: 'editFinancialInstitution/:id',
    component: EditFinancialInstitutionComponent
  },
  { path: 'indexFinancialInstitution',
    component: IndexFinancialInstitutionComponent
  }
];
export const BranchRoutes: Routes = [
  { path: 'createBranch',
    component: CreateBranchComponent
  },
  {
    path: 'editBranch/:id',
    component: EditBranchComponent
  },
  { path: 'indexBranch',
    component: IndexBranchComponent
  }
];
export const ProductOfferingRoutes: Routes = [
  { path: 'createProductOffering',
    component: CreateProductOfferingComponent
  },
  {
    path: 'editProductOffering/:id',
    component: EditProductOfferingComponent
  },
  { path: 'indexProductOffering',
    component: IndexProductOfferingComponent
  }
];
export const PricingPlanRoutes: Routes = [
  { path: 'createPricingPlan',
    component: CreatePricingPlanComponent
  },
  {
    path: 'editPricingPlan/:id',
    component: EditPricingPlanComponent
  },
  { path: 'indexPricingPlan',
    component: IndexPricingPlanComponent
  }
];
export const FeeScheduleRoutes: Routes = [
  { path: 'createFeeSchedule',
    component: CreateFeeScheduleComponent
  },
  {
    path: 'editFeeSchedule/:id',
    component: EditFeeScheduleComponent
  },
  { path: 'indexFeeSchedule',
    component: IndexFeeScheduleComponent
  }
];
export const UsageLimitRoutes: Routes = [
  { path: 'createUsageLimit',
    component: CreateUsageLimitComponent
  },
  {
    path: 'editUsageLimit/:id',
    component: EditUsageLimitComponent
  },
  { path: 'indexUsageLimit',
    component: IndexUsageLimitComponent
  }
];
export const CustomerRoutes: Routes = [
  { path: 'createCustomer',
    component: CreateCustomerComponent
  },
  {
    path: 'editCustomer/:id',
    component: EditCustomerComponent
  },
  { path: 'indexCustomer',
    component: IndexCustomerComponent
  }
];
export const KYCProfileRoutes: Routes = [
  { path: 'createKYCProfile',
    component: CreateKYCProfileComponent
  },
  {
    path: 'editKYCProfile/:id',
    component: EditKYCProfileComponent
  },
  { path: 'indexKYCProfile',
    component: IndexKYCProfileComponent
  }
];
export const KYCDocumentRoutes: Routes = [
  { path: 'createKYCDocument',
    component: CreateKYCDocumentComponent
  },
  {
    path: 'editKYCDocument/:id',
    component: EditKYCDocumentComponent
  },
  { path: 'indexKYCDocument',
    component: IndexKYCDocumentComponent
  }
];
export const ScreeningRoutes: Routes = [
  { path: 'createScreening',
    component: CreateScreeningComponent
  },
  {
    path: 'editScreening/:id',
    component: EditScreeningComponent
  },
  { path: 'indexScreening',
    component: IndexScreeningComponent
  }
];
export const VerifiedAddressRoutes: Routes = [
  { path: 'createVerifiedAddress',
    component: CreateVerifiedAddressComponent
  },
  {
    path: 'editVerifiedAddress/:id',
    component: EditVerifiedAddressComponent
  },
  { path: 'indexVerifiedAddress',
    component: IndexVerifiedAddressComponent
  }
];
export const CompliancePolicyRoutes: Routes = [
  { path: 'createCompliancePolicy',
    component: CreateCompliancePolicyComponent
  },
  {
    path: 'editCompliancePolicy/:id',
    component: EditCompliancePolicyComponent
  },
  { path: 'indexCompliancePolicy',
    component: IndexCompliancePolicyComponent
  }
];
export const ComplianceAlertRoutes: Routes = [
  { path: 'createComplianceAlert',
    component: CreateComplianceAlertComponent
  },
  {
    path: 'editComplianceAlert/:id',
    component: EditComplianceAlertComponent
  },
  { path: 'indexComplianceAlert',
    component: IndexComplianceAlertComponent
  }
];
export const ConsentRoutes: Routes = [
  { path: 'createConsent',
    component: CreateConsentComponent
  },
  {
    path: 'editConsent/:id',
    component: EditConsentComponent
  },
  { path: 'indexConsent',
    component: IndexConsentComponent
  }
];
export const APIClientRoutes: Routes = [
  { path: 'createAPIClient',
    component: CreateAPIClientComponent
  },
  {
    path: 'editAPIClient/:id',
    component: EditAPIClientComponent
  },
  { path: 'indexAPIClient',
    component: IndexAPIClientComponent
  }
];
export const AgreementRoutes: Routes = [
  { path: 'createAgreement',
    component: CreateAgreementComponent
  },
  {
    path: 'editAgreement/:id',
    component: EditAgreementComponent
  },
  { path: 'indexAgreement',
    component: IndexAgreementComponent
  }
];
export const AccountRoutes: Routes = [
  { path: 'createAccount',
    component: CreateAccountComponent
  },
  {
    path: 'editAccount/:id',
    component: EditAccountComponent
  },
  { path: 'indexAccount',
    component: IndexAccountComponent
  }
];
export const WalletRoutes: Routes = [
  { path: 'createWallet',
    component: CreateWalletComponent
  },
  {
    path: 'editWallet/:id',
    component: EditWalletComponent
  },
  { path: 'indexWallet',
    component: IndexWalletComponent
  }
];
export const PaymentCardRoutes: Routes = [
  { path: 'createPaymentCard',
    component: CreatePaymentCardComponent
  },
  {
    path: 'editPaymentCard/:id',
    component: EditPaymentCardComponent
  },
  { path: 'indexPaymentCard',
    component: IndexPaymentCardComponent
  }
];
export const CardTokenizationRoutes: Routes = [
  { path: 'createCardTokenization',
    component: CreateCardTokenizationComponent
  },
  {
    path: 'editCardTokenization/:id',
    component: EditCardTokenizationComponent
  },
  { path: 'indexCardTokenization',
    component: IndexCardTokenizationComponent
  }
];
export const MerchantRoutes: Routes = [
  { path: 'createMerchant',
    component: CreateMerchantComponent
  },
  {
    path: 'editMerchant/:id',
    component: EditMerchantComponent
  },
  { path: 'indexMerchant',
    component: IndexMerchantComponent
  }
];
export const TerminalRoutes: Routes = [
  { path: 'createTerminal',
    component: CreateTerminalComponent
  },
  {
    path: 'editTerminal/:id',
    component: EditTerminalComponent
  },
  { path: 'indexTerminal',
    component: IndexTerminalComponent
  }
];
export const PaymentContractRoutes: Routes = [
  { path: 'createPaymentContract',
    component: CreatePaymentContractComponent
  },
  {
    path: 'editPaymentContract/:id',
    component: EditPaymentContractComponent
  },
  { path: 'indexPaymentContract',
    component: IndexPaymentContractComponent
  }
];
export const PaymentProcessorRoutes: Routes = [
  { path: 'createPaymentProcessor',
    component: CreatePaymentProcessorComponent
  },
  {
    path: 'editPaymentProcessor/:id',
    component: EditPaymentProcessorComponent
  },
  { path: 'indexPaymentProcessor',
    component: IndexPaymentProcessorComponent
  }
];
export const TransactionRoutes: Routes = [
  { path: 'createTransaction',
    component: CreateTransactionComponent
  },
  {
    path: 'editTransaction/:id',
    component: EditTransactionComponent
  },
  { path: 'indexTransaction',
    component: IndexTransactionComponent
  }
];
export const PaymentOrderRoutes: Routes = [
  { path: 'createPaymentOrder',
    component: CreatePaymentOrderComponent
  },
  {
    path: 'editPaymentOrder/:id',
    component: EditPaymentOrderComponent
  },
  { path: 'indexPaymentOrder',
    component: IndexPaymentOrderComponent
  }
];
export const BeneficiaryRoutes: Routes = [
  { path: 'createBeneficiary',
    component: CreateBeneficiaryComponent
  },
  {
    path: 'editBeneficiary/:id',
    component: EditBeneficiaryComponent
  },
  { path: 'indexBeneficiary',
    component: IndexBeneficiaryComponent
  }
];
export const AppliedFeeRoutes: Routes = [
  { path: 'createAppliedFee',
    component: CreateAppliedFeeComponent
  },
  {
    path: 'editAppliedFee/:id',
    component: EditAppliedFeeComponent
  },
  { path: 'indexAppliedFee',
    component: IndexAppliedFeeComponent
  }
];
export const FXQuoteRoutes: Routes = [
  { path: 'createFXQuote',
    component: CreateFXQuoteComponent
  },
  {
    path: 'editFXQuote/:id',
    component: EditFXQuoteComponent
  },
  { path: 'indexFXQuote',
    component: IndexFXQuoteComponent
  }
];
export const FXDealRoutes: Routes = [
  { path: 'createFXDeal',
    component: CreateFXDealComponent
  },
  {
    path: 'editFXDeal/:id',
    component: EditFXDealComponent
  },
  { path: 'indexFXDeal',
    component: IndexFXDealComponent
  }
];
export const SettlementBatchRoutes: Routes = [
  { path: 'createSettlementBatch',
    component: CreateSettlementBatchComponent
  },
  {
    path: 'editSettlementBatch/:id',
    component: EditSettlementBatchComponent
  },
  { path: 'indexSettlementBatch',
    component: IndexSettlementBatchComponent
  }
];
export const PayoutRoutes: Routes = [
  { path: 'createPayout',
    component: CreatePayoutComponent
  },
  {
    path: 'editPayout/:id',
    component: EditPayoutComponent
  },
  { path: 'indexPayout',
    component: IndexPayoutComponent
  }
];
export const DisputeRoutes: Routes = [
  { path: 'createDispute',
    component: CreateDisputeComponent
  },
  {
    path: 'editDispute/:id',
    component: EditDisputeComponent
  },
  { path: 'indexDispute',
    component: IndexDisputeComponent
  }
];
export const ChargebackRoutes: Routes = [
  { path: 'createChargeback',
    component: CreateChargebackComponent
  },
  {
    path: 'editChargeback/:id',
    component: EditChargebackComponent
  },
  { path: 'indexChargeback',
    component: IndexChargebackComponent
  }
];
export const InvoiceRoutes: Routes = [
  { path: 'createInvoice',
    component: CreateInvoiceComponent
  },
  {
    path: 'editInvoice/:id',
    component: EditInvoiceComponent
  },
  { path: 'indexInvoice',
    component: IndexInvoiceComponent
  }
];
export const AccountStatementRoutes: Routes = [
  { path: 'createAccountStatement',
    component: CreateAccountStatementComponent
  },
  {
    path: 'editAccountStatement/:id',
    component: EditAccountStatementComponent
  },
  { path: 'indexAccountStatement',
    component: IndexAccountStatementComponent
  }
];
export const DirectDebitMandateRoutes: Routes = [
  { path: 'createDirectDebitMandate',
    component: CreateDirectDebitMandateComponent
  },
  {
    path: 'editDirectDebitMandate/:id',
    component: EditDirectDebitMandateComponent
  },
  { path: 'indexDirectDebitMandate',
    component: IndexDirectDebitMandateComponent
  }
];
export const CreditorRoutes: Routes = [
  { path: 'createCreditor',
    component: CreateCreditorComponent
  },
  {
    path: 'editCreditor/:id',
    component: EditCreditorComponent
  },
  { path: 'indexCreditor',
    component: IndexCreditorComponent
  }
];
export const LoanApplicationRoutes: Routes = [
  { path: 'createLoanApplication',
    component: CreateLoanApplicationComponent
  },
  {
    path: 'editLoanApplication/:id',
    component: EditLoanApplicationComponent
  },
  { path: 'indexLoanApplication',
    component: IndexLoanApplicationComponent
  }
];
export const RiskAssessmentRoutes: Routes = [
  { path: 'createRiskAssessment',
    component: CreateRiskAssessmentComponent
  },
  {
    path: 'editRiskAssessment/:id',
    component: EditRiskAssessmentComponent
  },
  { path: 'indexRiskAssessment',
    component: IndexRiskAssessmentComponent
  }
];
export const LoanRoutes: Routes = [
  { path: 'createLoan',
    component: CreateLoanComponent
  },
  {
    path: 'editLoan/:id',
    component: EditLoanComponent
  },
  { path: 'indexLoan',
    component: IndexLoanComponent
  }
];
export const RepaymentScheduleRoutes: Routes = [
  { path: 'createRepaymentSchedule',
    component: CreateRepaymentScheduleComponent
  },
  {
    path: 'editRepaymentSchedule/:id',
    component: EditRepaymentScheduleComponent
  },
  { path: 'indexRepaymentSchedule',
    component: IndexRepaymentScheduleComponent
  }
];
export const CollateralRoutes: Routes = [
  { path: 'createCollateral',
    component: CreateCollateralComponent
  },
  {
    path: 'editCollateral/:id',
    component: EditCollateralComponent
  },
  { path: 'indexCollateral',
    component: IndexCollateralComponent
  }
];
export const LoanTransactionRoutes: Routes = [
  { path: 'createLoanTransaction',
    component: CreateLoanTransactionComponent
  },
  {
    path: 'editLoanTransaction/:id',
    component: EditLoanTransactionComponent
  },
  { path: 'indexLoanTransaction',
    component: IndexLoanTransactionComponent
  }
];
export const InvestmentPortfolioRoutes: Routes = [
  { path: 'createInvestmentPortfolio',
    component: CreateInvestmentPortfolioComponent
  },
  {
    path: 'editInvestmentPortfolio/:id',
    component: EditInvestmentPortfolioComponent
  },
  { path: 'indexInvestmentPortfolio',
    component: IndexInvestmentPortfolioComponent
  }
];
export const InvestmentAccountRoutes: Routes = [
  { path: 'createInvestmentAccount',
    component: CreateInvestmentAccountComponent
  },
  {
    path: 'editInvestmentAccount/:id',
    component: EditInvestmentAccountComponent
  },
  { path: 'indexInvestmentAccount',
    component: IndexInvestmentAccountComponent
  }
];
export const SecurityRoutes: Routes = [
  { path: 'createSecurity',
    component: CreateSecurityComponent
  },
  {
    path: 'editSecurity/:id',
    component: EditSecurityComponent
  },
  { path: 'indexSecurity',
    component: IndexSecurityComponent
  }
];
export const PositionRoutes: Routes = [
  { path: 'createPosition',
    component: CreatePositionComponent
  },
  {
    path: 'editPosition/:id',
    component: EditPositionComponent
  },
  { path: 'indexPosition',
    component: IndexPositionComponent
  }
];
export const TradeOrderRoutes: Routes = [
  { path: 'createTradeOrder',
    component: CreateTradeOrderComponent
  },
  {
    path: 'editTradeOrder/:id',
    component: EditTradeOrderComponent
  },
  { path: 'indexTradeOrder',
    component: IndexTradeOrderComponent
  }
];
export const TradeRoutes: Routes = [
  { path: 'createTrade',
    component: CreateTradeComponent
  },
  {
    path: 'editTrade/:id',
    component: EditTradeComponent
  },
  { path: 'indexTrade',
    component: IndexTradeComponent
  }
];
export const ExchangeRateRoutes: Routes = [
  { path: 'createExchangeRate',
    component: CreateExchangeRateComponent
  },
  {
    path: 'editExchangeRate/:id',
    component: EditExchangeRateComponent
  },
  { path: 'indexExchangeRate',
    component: IndexExchangeRateComponent
  }
];
