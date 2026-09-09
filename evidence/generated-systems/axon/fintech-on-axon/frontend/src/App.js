import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListFinancialInstitutionComponent from './components/ListFinancialInstitutionComponent';
import CreateFinancialInstitutionComponent from './components/CreateFinancialInstitutionComponent';
import ViewFinancialInstitutionComponent from './components/ViewFinancialInstitutionComponent';
import ListBranchComponent from './components/ListBranchComponent';
import CreateBranchComponent from './components/CreateBranchComponent';
import ViewBranchComponent from './components/ViewBranchComponent';
import ListProductOfferingComponent from './components/ListProductOfferingComponent';
import CreateProductOfferingComponent from './components/CreateProductOfferingComponent';
import ViewProductOfferingComponent from './components/ViewProductOfferingComponent';
import ListPricingPlanComponent from './components/ListPricingPlanComponent';
import CreatePricingPlanComponent from './components/CreatePricingPlanComponent';
import ViewPricingPlanComponent from './components/ViewPricingPlanComponent';
import ListFeeScheduleComponent from './components/ListFeeScheduleComponent';
import CreateFeeScheduleComponent from './components/CreateFeeScheduleComponent';
import ViewFeeScheduleComponent from './components/ViewFeeScheduleComponent';
import ListUsageLimitComponent from './components/ListUsageLimitComponent';
import CreateUsageLimitComponent from './components/CreateUsageLimitComponent';
import ViewUsageLimitComponent from './components/ViewUsageLimitComponent';
import ListCustomerComponent from './components/ListCustomerComponent';
import CreateCustomerComponent from './components/CreateCustomerComponent';
import ViewCustomerComponent from './components/ViewCustomerComponent';
import ListKYCProfileComponent from './components/ListKYCProfileComponent';
import CreateKYCProfileComponent from './components/CreateKYCProfileComponent';
import ViewKYCProfileComponent from './components/ViewKYCProfileComponent';
import ListKYCDocumentComponent from './components/ListKYCDocumentComponent';
import CreateKYCDocumentComponent from './components/CreateKYCDocumentComponent';
import ViewKYCDocumentComponent from './components/ViewKYCDocumentComponent';
import ListScreeningComponent from './components/ListScreeningComponent';
import CreateScreeningComponent from './components/CreateScreeningComponent';
import ViewScreeningComponent from './components/ViewScreeningComponent';
import ListVerifiedAddressComponent from './components/ListVerifiedAddressComponent';
import CreateVerifiedAddressComponent from './components/CreateVerifiedAddressComponent';
import ViewVerifiedAddressComponent from './components/ViewVerifiedAddressComponent';
import ListCompliancePolicyComponent from './components/ListCompliancePolicyComponent';
import CreateCompliancePolicyComponent from './components/CreateCompliancePolicyComponent';
import ViewCompliancePolicyComponent from './components/ViewCompliancePolicyComponent';
import ListComplianceAlertComponent from './components/ListComplianceAlertComponent';
import CreateComplianceAlertComponent from './components/CreateComplianceAlertComponent';
import ViewComplianceAlertComponent from './components/ViewComplianceAlertComponent';
import ListConsentComponent from './components/ListConsentComponent';
import CreateConsentComponent from './components/CreateConsentComponent';
import ViewConsentComponent from './components/ViewConsentComponent';
import ListAPIClientComponent from './components/ListAPIClientComponent';
import CreateAPIClientComponent from './components/CreateAPIClientComponent';
import ViewAPIClientComponent from './components/ViewAPIClientComponent';
import ListAgreementComponent from './components/ListAgreementComponent';
import CreateAgreementComponent from './components/CreateAgreementComponent';
import ViewAgreementComponent from './components/ViewAgreementComponent';
import ListAccountComponent from './components/ListAccountComponent';
import CreateAccountComponent from './components/CreateAccountComponent';
import ViewAccountComponent from './components/ViewAccountComponent';
import ListWalletComponent from './components/ListWalletComponent';
import CreateWalletComponent from './components/CreateWalletComponent';
import ViewWalletComponent from './components/ViewWalletComponent';
import ListPaymentCardComponent from './components/ListPaymentCardComponent';
import CreatePaymentCardComponent from './components/CreatePaymentCardComponent';
import ViewPaymentCardComponent from './components/ViewPaymentCardComponent';
import ListCardTokenizationComponent from './components/ListCardTokenizationComponent';
import CreateCardTokenizationComponent from './components/CreateCardTokenizationComponent';
import ViewCardTokenizationComponent from './components/ViewCardTokenizationComponent';
import ListMerchantComponent from './components/ListMerchantComponent';
import CreateMerchantComponent from './components/CreateMerchantComponent';
import ViewMerchantComponent from './components/ViewMerchantComponent';
import ListTerminalComponent from './components/ListTerminalComponent';
import CreateTerminalComponent from './components/CreateTerminalComponent';
import ViewTerminalComponent from './components/ViewTerminalComponent';
import ListPaymentContractComponent from './components/ListPaymentContractComponent';
import CreatePaymentContractComponent from './components/CreatePaymentContractComponent';
import ViewPaymentContractComponent from './components/ViewPaymentContractComponent';
import ListPaymentProcessorComponent from './components/ListPaymentProcessorComponent';
import CreatePaymentProcessorComponent from './components/CreatePaymentProcessorComponent';
import ViewPaymentProcessorComponent from './components/ViewPaymentProcessorComponent';
import ListTransactionComponent from './components/ListTransactionComponent';
import CreateTransactionComponent from './components/CreateTransactionComponent';
import ViewTransactionComponent from './components/ViewTransactionComponent';
import ListPaymentOrderComponent from './components/ListPaymentOrderComponent';
import CreatePaymentOrderComponent from './components/CreatePaymentOrderComponent';
import ViewPaymentOrderComponent from './components/ViewPaymentOrderComponent';
import ListBeneficiaryComponent from './components/ListBeneficiaryComponent';
import CreateBeneficiaryComponent from './components/CreateBeneficiaryComponent';
import ViewBeneficiaryComponent from './components/ViewBeneficiaryComponent';
import ListAppliedFeeComponent from './components/ListAppliedFeeComponent';
import CreateAppliedFeeComponent from './components/CreateAppliedFeeComponent';
import ViewAppliedFeeComponent from './components/ViewAppliedFeeComponent';
import ListFXQuoteComponent from './components/ListFXQuoteComponent';
import CreateFXQuoteComponent from './components/CreateFXQuoteComponent';
import ViewFXQuoteComponent from './components/ViewFXQuoteComponent';
import ListFXDealComponent from './components/ListFXDealComponent';
import CreateFXDealComponent from './components/CreateFXDealComponent';
import ViewFXDealComponent from './components/ViewFXDealComponent';
import ListSettlementBatchComponent from './components/ListSettlementBatchComponent';
import CreateSettlementBatchComponent from './components/CreateSettlementBatchComponent';
import ViewSettlementBatchComponent from './components/ViewSettlementBatchComponent';
import ListPayoutComponent from './components/ListPayoutComponent';
import CreatePayoutComponent from './components/CreatePayoutComponent';
import ViewPayoutComponent from './components/ViewPayoutComponent';
import ListDisputeComponent from './components/ListDisputeComponent';
import CreateDisputeComponent from './components/CreateDisputeComponent';
import ViewDisputeComponent from './components/ViewDisputeComponent';
import ListChargebackComponent from './components/ListChargebackComponent';
import CreateChargebackComponent from './components/CreateChargebackComponent';
import ViewChargebackComponent from './components/ViewChargebackComponent';
import ListInvoiceComponent from './components/ListInvoiceComponent';
import CreateInvoiceComponent from './components/CreateInvoiceComponent';
import ViewInvoiceComponent from './components/ViewInvoiceComponent';
import ListAccountStatementComponent from './components/ListAccountStatementComponent';
import CreateAccountStatementComponent from './components/CreateAccountStatementComponent';
import ViewAccountStatementComponent from './components/ViewAccountStatementComponent';
import ListDirectDebitMandateComponent from './components/ListDirectDebitMandateComponent';
import CreateDirectDebitMandateComponent from './components/CreateDirectDebitMandateComponent';
import ViewDirectDebitMandateComponent from './components/ViewDirectDebitMandateComponent';
import ListCreditorComponent from './components/ListCreditorComponent';
import CreateCreditorComponent from './components/CreateCreditorComponent';
import ViewCreditorComponent from './components/ViewCreditorComponent';
import ListLoanApplicationComponent from './components/ListLoanApplicationComponent';
import CreateLoanApplicationComponent from './components/CreateLoanApplicationComponent';
import ViewLoanApplicationComponent from './components/ViewLoanApplicationComponent';
import ListRiskAssessmentComponent from './components/ListRiskAssessmentComponent';
import CreateRiskAssessmentComponent from './components/CreateRiskAssessmentComponent';
import ViewRiskAssessmentComponent from './components/ViewRiskAssessmentComponent';
import ListLoanComponent from './components/ListLoanComponent';
import CreateLoanComponent from './components/CreateLoanComponent';
import ViewLoanComponent from './components/ViewLoanComponent';
import ListRepaymentScheduleComponent from './components/ListRepaymentScheduleComponent';
import CreateRepaymentScheduleComponent from './components/CreateRepaymentScheduleComponent';
import ViewRepaymentScheduleComponent from './components/ViewRepaymentScheduleComponent';
import ListCollateralComponent from './components/ListCollateralComponent';
import CreateCollateralComponent from './components/CreateCollateralComponent';
import ViewCollateralComponent from './components/ViewCollateralComponent';
import ListLoanTransactionComponent from './components/ListLoanTransactionComponent';
import CreateLoanTransactionComponent from './components/CreateLoanTransactionComponent';
import ViewLoanTransactionComponent from './components/ViewLoanTransactionComponent';
import ListInvestmentPortfolioComponent from './components/ListInvestmentPortfolioComponent';
import CreateInvestmentPortfolioComponent from './components/CreateInvestmentPortfolioComponent';
import ViewInvestmentPortfolioComponent from './components/ViewInvestmentPortfolioComponent';
import ListInvestmentAccountComponent from './components/ListInvestmentAccountComponent';
import CreateInvestmentAccountComponent from './components/CreateInvestmentAccountComponent';
import ViewInvestmentAccountComponent from './components/ViewInvestmentAccountComponent';
import ListSecurityComponent from './components/ListSecurityComponent';
import CreateSecurityComponent from './components/CreateSecurityComponent';
import ViewSecurityComponent from './components/ViewSecurityComponent';
import ListPositionComponent from './components/ListPositionComponent';
import CreatePositionComponent from './components/CreatePositionComponent';
import ViewPositionComponent from './components/ViewPositionComponent';
import ListTradeOrderComponent from './components/ListTradeOrderComponent';
import CreateTradeOrderComponent from './components/CreateTradeOrderComponent';
import ViewTradeOrderComponent from './components/ViewTradeOrderComponent';
import ListTradeComponent from './components/ListTradeComponent';
import CreateTradeComponent from './components/CreateTradeComponent';
import ViewTradeComponent from './components/ViewTradeComponent';
import ListExchangeRateComponent from './components/ListExchangeRateComponent';
import CreateExchangeRateComponent from './components/CreateExchangeRateComponent';
import ViewExchangeRateComponent from './components/ViewExchangeRateComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/financialInstitutions" component = {ListFinancialInstitutionComponent}></Route>
                            <Route path = "/add-financialInstitution/:id" component = {CreateFinancialInstitutionComponent}></Route>
                            <Route path = "/view-financialInstitution/:id" component = {ViewFinancialInstitutionComponent}></Route>
                          {/* <Route path = "/update-financialInstitution/:id" component = {UpdateFinancialInstitutionComponent}></Route> */}
                            <Route path = "/branchs" component = {ListBranchComponent}></Route>
                            <Route path = "/add-branch/:id" component = {CreateBranchComponent}></Route>
                            <Route path = "/view-branch/:id" component = {ViewBranchComponent}></Route>
                          {/* <Route path = "/update-branch/:id" component = {UpdateBranchComponent}></Route> */}
                            <Route path = "/productOfferings" component = {ListProductOfferingComponent}></Route>
                            <Route path = "/add-productOffering/:id" component = {CreateProductOfferingComponent}></Route>
                            <Route path = "/view-productOffering/:id" component = {ViewProductOfferingComponent}></Route>
                          {/* <Route path = "/update-productOffering/:id" component = {UpdateProductOfferingComponent}></Route> */}
                            <Route path = "/pricingPlans" component = {ListPricingPlanComponent}></Route>
                            <Route path = "/add-pricingPlan/:id" component = {CreatePricingPlanComponent}></Route>
                            <Route path = "/view-pricingPlan/:id" component = {ViewPricingPlanComponent}></Route>
                          {/* <Route path = "/update-pricingPlan/:id" component = {UpdatePricingPlanComponent}></Route> */}
                            <Route path = "/feeSchedules" component = {ListFeeScheduleComponent}></Route>
                            <Route path = "/add-feeSchedule/:id" component = {CreateFeeScheduleComponent}></Route>
                            <Route path = "/view-feeSchedule/:id" component = {ViewFeeScheduleComponent}></Route>
                          {/* <Route path = "/update-feeSchedule/:id" component = {UpdateFeeScheduleComponent}></Route> */}
                            <Route path = "/usageLimits" component = {ListUsageLimitComponent}></Route>
                            <Route path = "/add-usageLimit/:id" component = {CreateUsageLimitComponent}></Route>
                            <Route path = "/view-usageLimit/:id" component = {ViewUsageLimitComponent}></Route>
                          {/* <Route path = "/update-usageLimit/:id" component = {UpdateUsageLimitComponent}></Route> */}
                            <Route path = "/customers" component = {ListCustomerComponent}></Route>
                            <Route path = "/add-customer/:id" component = {CreateCustomerComponent}></Route>
                            <Route path = "/view-customer/:id" component = {ViewCustomerComponent}></Route>
                          {/* <Route path = "/update-customer/:id" component = {UpdateCustomerComponent}></Route> */}
                            <Route path = "/kYCProfiles" component = {ListKYCProfileComponent}></Route>
                            <Route path = "/add-kYCProfile/:id" component = {CreateKYCProfileComponent}></Route>
                            <Route path = "/view-kYCProfile/:id" component = {ViewKYCProfileComponent}></Route>
                          {/* <Route path = "/update-kYCProfile/:id" component = {UpdateKYCProfileComponent}></Route> */}
                            <Route path = "/kYCDocuments" component = {ListKYCDocumentComponent}></Route>
                            <Route path = "/add-kYCDocument/:id" component = {CreateKYCDocumentComponent}></Route>
                            <Route path = "/view-kYCDocument/:id" component = {ViewKYCDocumentComponent}></Route>
                          {/* <Route path = "/update-kYCDocument/:id" component = {UpdateKYCDocumentComponent}></Route> */}
                            <Route path = "/screenings" component = {ListScreeningComponent}></Route>
                            <Route path = "/add-screening/:id" component = {CreateScreeningComponent}></Route>
                            <Route path = "/view-screening/:id" component = {ViewScreeningComponent}></Route>
                          {/* <Route path = "/update-screening/:id" component = {UpdateScreeningComponent}></Route> */}
                            <Route path = "/verifiedAddresss" component = {ListVerifiedAddressComponent}></Route>
                            <Route path = "/add-verifiedAddress/:id" component = {CreateVerifiedAddressComponent}></Route>
                            <Route path = "/view-verifiedAddress/:id" component = {ViewVerifiedAddressComponent}></Route>
                          {/* <Route path = "/update-verifiedAddress/:id" component = {UpdateVerifiedAddressComponent}></Route> */}
                            <Route path = "/compliancePolicys" component = {ListCompliancePolicyComponent}></Route>
                            <Route path = "/add-compliancePolicy/:id" component = {CreateCompliancePolicyComponent}></Route>
                            <Route path = "/view-compliancePolicy/:id" component = {ViewCompliancePolicyComponent}></Route>
                          {/* <Route path = "/update-compliancePolicy/:id" component = {UpdateCompliancePolicyComponent}></Route> */}
                            <Route path = "/complianceAlerts" component = {ListComplianceAlertComponent}></Route>
                            <Route path = "/add-complianceAlert/:id" component = {CreateComplianceAlertComponent}></Route>
                            <Route path = "/view-complianceAlert/:id" component = {ViewComplianceAlertComponent}></Route>
                          {/* <Route path = "/update-complianceAlert/:id" component = {UpdateComplianceAlertComponent}></Route> */}
                            <Route path = "/consents" component = {ListConsentComponent}></Route>
                            <Route path = "/add-consent/:id" component = {CreateConsentComponent}></Route>
                            <Route path = "/view-consent/:id" component = {ViewConsentComponent}></Route>
                          {/* <Route path = "/update-consent/:id" component = {UpdateConsentComponent}></Route> */}
                            <Route path = "/aPIClients" component = {ListAPIClientComponent}></Route>
                            <Route path = "/add-aPIClient/:id" component = {CreateAPIClientComponent}></Route>
                            <Route path = "/view-aPIClient/:id" component = {ViewAPIClientComponent}></Route>
                          {/* <Route path = "/update-aPIClient/:id" component = {UpdateAPIClientComponent}></Route> */}
                            <Route path = "/agreements" component = {ListAgreementComponent}></Route>
                            <Route path = "/add-agreement/:id" component = {CreateAgreementComponent}></Route>
                            <Route path = "/view-agreement/:id" component = {ViewAgreementComponent}></Route>
                          {/* <Route path = "/update-agreement/:id" component = {UpdateAgreementComponent}></Route> */}
                            <Route path = "/accounts" component = {ListAccountComponent}></Route>
                            <Route path = "/add-account/:id" component = {CreateAccountComponent}></Route>
                            <Route path = "/view-account/:id" component = {ViewAccountComponent}></Route>
                          {/* <Route path = "/update-account/:id" component = {UpdateAccountComponent}></Route> */}
                            <Route path = "/wallets" component = {ListWalletComponent}></Route>
                            <Route path = "/add-wallet/:id" component = {CreateWalletComponent}></Route>
                            <Route path = "/view-wallet/:id" component = {ViewWalletComponent}></Route>
                          {/* <Route path = "/update-wallet/:id" component = {UpdateWalletComponent}></Route> */}
                            <Route path = "/paymentCards" component = {ListPaymentCardComponent}></Route>
                            <Route path = "/add-paymentCard/:id" component = {CreatePaymentCardComponent}></Route>
                            <Route path = "/view-paymentCard/:id" component = {ViewPaymentCardComponent}></Route>
                          {/* <Route path = "/update-paymentCard/:id" component = {UpdatePaymentCardComponent}></Route> */}
                            <Route path = "/cardTokenizations" component = {ListCardTokenizationComponent}></Route>
                            <Route path = "/add-cardTokenization/:id" component = {CreateCardTokenizationComponent}></Route>
                            <Route path = "/view-cardTokenization/:id" component = {ViewCardTokenizationComponent}></Route>
                          {/* <Route path = "/update-cardTokenization/:id" component = {UpdateCardTokenizationComponent}></Route> */}
                            <Route path = "/merchants" component = {ListMerchantComponent}></Route>
                            <Route path = "/add-merchant/:id" component = {CreateMerchantComponent}></Route>
                            <Route path = "/view-merchant/:id" component = {ViewMerchantComponent}></Route>
                          {/* <Route path = "/update-merchant/:id" component = {UpdateMerchantComponent}></Route> */}
                            <Route path = "/terminals" component = {ListTerminalComponent}></Route>
                            <Route path = "/add-terminal/:id" component = {CreateTerminalComponent}></Route>
                            <Route path = "/view-terminal/:id" component = {ViewTerminalComponent}></Route>
                          {/* <Route path = "/update-terminal/:id" component = {UpdateTerminalComponent}></Route> */}
                            <Route path = "/paymentContracts" component = {ListPaymentContractComponent}></Route>
                            <Route path = "/add-paymentContract/:id" component = {CreatePaymentContractComponent}></Route>
                            <Route path = "/view-paymentContract/:id" component = {ViewPaymentContractComponent}></Route>
                          {/* <Route path = "/update-paymentContract/:id" component = {UpdatePaymentContractComponent}></Route> */}
                            <Route path = "/paymentProcessors" component = {ListPaymentProcessorComponent}></Route>
                            <Route path = "/add-paymentProcessor/:id" component = {CreatePaymentProcessorComponent}></Route>
                            <Route path = "/view-paymentProcessor/:id" component = {ViewPaymentProcessorComponent}></Route>
                          {/* <Route path = "/update-paymentProcessor/:id" component = {UpdatePaymentProcessorComponent}></Route> */}
                            <Route path = "/transactions" component = {ListTransactionComponent}></Route>
                            <Route path = "/add-transaction/:id" component = {CreateTransactionComponent}></Route>
                            <Route path = "/view-transaction/:id" component = {ViewTransactionComponent}></Route>
                          {/* <Route path = "/update-transaction/:id" component = {UpdateTransactionComponent}></Route> */}
                            <Route path = "/paymentOrders" component = {ListPaymentOrderComponent}></Route>
                            <Route path = "/add-paymentOrder/:id" component = {CreatePaymentOrderComponent}></Route>
                            <Route path = "/view-paymentOrder/:id" component = {ViewPaymentOrderComponent}></Route>
                          {/* <Route path = "/update-paymentOrder/:id" component = {UpdatePaymentOrderComponent}></Route> */}
                            <Route path = "/beneficiarys" component = {ListBeneficiaryComponent}></Route>
                            <Route path = "/add-beneficiary/:id" component = {CreateBeneficiaryComponent}></Route>
                            <Route path = "/view-beneficiary/:id" component = {ViewBeneficiaryComponent}></Route>
                          {/* <Route path = "/update-beneficiary/:id" component = {UpdateBeneficiaryComponent}></Route> */}
                            <Route path = "/appliedFees" component = {ListAppliedFeeComponent}></Route>
                            <Route path = "/add-appliedFee/:id" component = {CreateAppliedFeeComponent}></Route>
                            <Route path = "/view-appliedFee/:id" component = {ViewAppliedFeeComponent}></Route>
                          {/* <Route path = "/update-appliedFee/:id" component = {UpdateAppliedFeeComponent}></Route> */}
                            <Route path = "/fXQuotes" component = {ListFXQuoteComponent}></Route>
                            <Route path = "/add-fXQuote/:id" component = {CreateFXQuoteComponent}></Route>
                            <Route path = "/view-fXQuote/:id" component = {ViewFXQuoteComponent}></Route>
                          {/* <Route path = "/update-fXQuote/:id" component = {UpdateFXQuoteComponent}></Route> */}
                            <Route path = "/fXDeals" component = {ListFXDealComponent}></Route>
                            <Route path = "/add-fXDeal/:id" component = {CreateFXDealComponent}></Route>
                            <Route path = "/view-fXDeal/:id" component = {ViewFXDealComponent}></Route>
                          {/* <Route path = "/update-fXDeal/:id" component = {UpdateFXDealComponent}></Route> */}
                            <Route path = "/settlementBatchs" component = {ListSettlementBatchComponent}></Route>
                            <Route path = "/add-settlementBatch/:id" component = {CreateSettlementBatchComponent}></Route>
                            <Route path = "/view-settlementBatch/:id" component = {ViewSettlementBatchComponent}></Route>
                          {/* <Route path = "/update-settlementBatch/:id" component = {UpdateSettlementBatchComponent}></Route> */}
                            <Route path = "/payouts" component = {ListPayoutComponent}></Route>
                            <Route path = "/add-payout/:id" component = {CreatePayoutComponent}></Route>
                            <Route path = "/view-payout/:id" component = {ViewPayoutComponent}></Route>
                          {/* <Route path = "/update-payout/:id" component = {UpdatePayoutComponent}></Route> */}
                            <Route path = "/disputes" component = {ListDisputeComponent}></Route>
                            <Route path = "/add-dispute/:id" component = {CreateDisputeComponent}></Route>
                            <Route path = "/view-dispute/:id" component = {ViewDisputeComponent}></Route>
                          {/* <Route path = "/update-dispute/:id" component = {UpdateDisputeComponent}></Route> */}
                            <Route path = "/chargebacks" component = {ListChargebackComponent}></Route>
                            <Route path = "/add-chargeback/:id" component = {CreateChargebackComponent}></Route>
                            <Route path = "/view-chargeback/:id" component = {ViewChargebackComponent}></Route>
                          {/* <Route path = "/update-chargeback/:id" component = {UpdateChargebackComponent}></Route> */}
                            <Route path = "/invoices" component = {ListInvoiceComponent}></Route>
                            <Route path = "/add-invoice/:id" component = {CreateInvoiceComponent}></Route>
                            <Route path = "/view-invoice/:id" component = {ViewInvoiceComponent}></Route>
                          {/* <Route path = "/update-invoice/:id" component = {UpdateInvoiceComponent}></Route> */}
                            <Route path = "/accountStatements" component = {ListAccountStatementComponent}></Route>
                            <Route path = "/add-accountStatement/:id" component = {CreateAccountStatementComponent}></Route>
                            <Route path = "/view-accountStatement/:id" component = {ViewAccountStatementComponent}></Route>
                          {/* <Route path = "/update-accountStatement/:id" component = {UpdateAccountStatementComponent}></Route> */}
                            <Route path = "/directDebitMandates" component = {ListDirectDebitMandateComponent}></Route>
                            <Route path = "/add-directDebitMandate/:id" component = {CreateDirectDebitMandateComponent}></Route>
                            <Route path = "/view-directDebitMandate/:id" component = {ViewDirectDebitMandateComponent}></Route>
                          {/* <Route path = "/update-directDebitMandate/:id" component = {UpdateDirectDebitMandateComponent}></Route> */}
                            <Route path = "/creditors" component = {ListCreditorComponent}></Route>
                            <Route path = "/add-creditor/:id" component = {CreateCreditorComponent}></Route>
                            <Route path = "/view-creditor/:id" component = {ViewCreditorComponent}></Route>
                          {/* <Route path = "/update-creditor/:id" component = {UpdateCreditorComponent}></Route> */}
                            <Route path = "/loanApplications" component = {ListLoanApplicationComponent}></Route>
                            <Route path = "/add-loanApplication/:id" component = {CreateLoanApplicationComponent}></Route>
                            <Route path = "/view-loanApplication/:id" component = {ViewLoanApplicationComponent}></Route>
                          {/* <Route path = "/update-loanApplication/:id" component = {UpdateLoanApplicationComponent}></Route> */}
                            <Route path = "/riskAssessments" component = {ListRiskAssessmentComponent}></Route>
                            <Route path = "/add-riskAssessment/:id" component = {CreateRiskAssessmentComponent}></Route>
                            <Route path = "/view-riskAssessment/:id" component = {ViewRiskAssessmentComponent}></Route>
                          {/* <Route path = "/update-riskAssessment/:id" component = {UpdateRiskAssessmentComponent}></Route> */}
                            <Route path = "/loans" component = {ListLoanComponent}></Route>
                            <Route path = "/add-loan/:id" component = {CreateLoanComponent}></Route>
                            <Route path = "/view-loan/:id" component = {ViewLoanComponent}></Route>
                          {/* <Route path = "/update-loan/:id" component = {UpdateLoanComponent}></Route> */}
                            <Route path = "/repaymentSchedules" component = {ListRepaymentScheduleComponent}></Route>
                            <Route path = "/add-repaymentSchedule/:id" component = {CreateRepaymentScheduleComponent}></Route>
                            <Route path = "/view-repaymentSchedule/:id" component = {ViewRepaymentScheduleComponent}></Route>
                          {/* <Route path = "/update-repaymentSchedule/:id" component = {UpdateRepaymentScheduleComponent}></Route> */}
                            <Route path = "/collaterals" component = {ListCollateralComponent}></Route>
                            <Route path = "/add-collateral/:id" component = {CreateCollateralComponent}></Route>
                            <Route path = "/view-collateral/:id" component = {ViewCollateralComponent}></Route>
                          {/* <Route path = "/update-collateral/:id" component = {UpdateCollateralComponent}></Route> */}
                            <Route path = "/loanTransactions" component = {ListLoanTransactionComponent}></Route>
                            <Route path = "/add-loanTransaction/:id" component = {CreateLoanTransactionComponent}></Route>
                            <Route path = "/view-loanTransaction/:id" component = {ViewLoanTransactionComponent}></Route>
                          {/* <Route path = "/update-loanTransaction/:id" component = {UpdateLoanTransactionComponent}></Route> */}
                            <Route path = "/investmentPortfolios" component = {ListInvestmentPortfolioComponent}></Route>
                            <Route path = "/add-investmentPortfolio/:id" component = {CreateInvestmentPortfolioComponent}></Route>
                            <Route path = "/view-investmentPortfolio/:id" component = {ViewInvestmentPortfolioComponent}></Route>
                          {/* <Route path = "/update-investmentPortfolio/:id" component = {UpdateInvestmentPortfolioComponent}></Route> */}
                            <Route path = "/investmentAccounts" component = {ListInvestmentAccountComponent}></Route>
                            <Route path = "/add-investmentAccount/:id" component = {CreateInvestmentAccountComponent}></Route>
                            <Route path = "/view-investmentAccount/:id" component = {ViewInvestmentAccountComponent}></Route>
                          {/* <Route path = "/update-investmentAccount/:id" component = {UpdateInvestmentAccountComponent}></Route> */}
                            <Route path = "/securitys" component = {ListSecurityComponent}></Route>
                            <Route path = "/add-security/:id" component = {CreateSecurityComponent}></Route>
                            <Route path = "/view-security/:id" component = {ViewSecurityComponent}></Route>
                          {/* <Route path = "/update-security/:id" component = {UpdateSecurityComponent}></Route> */}
                            <Route path = "/positions" component = {ListPositionComponent}></Route>
                            <Route path = "/add-position/:id" component = {CreatePositionComponent}></Route>
                            <Route path = "/view-position/:id" component = {ViewPositionComponent}></Route>
                          {/* <Route path = "/update-position/:id" component = {UpdatePositionComponent}></Route> */}
                            <Route path = "/tradeOrders" component = {ListTradeOrderComponent}></Route>
                            <Route path = "/add-tradeOrder/:id" component = {CreateTradeOrderComponent}></Route>
                            <Route path = "/view-tradeOrder/:id" component = {ViewTradeOrderComponent}></Route>
                          {/* <Route path = "/update-tradeOrder/:id" component = {UpdateTradeOrderComponent}></Route> */}
                            <Route path = "/trades" component = {ListTradeComponent}></Route>
                            <Route path = "/add-trade/:id" component = {CreateTradeComponent}></Route>
                            <Route path = "/view-trade/:id" component = {ViewTradeComponent}></Route>
                          {/* <Route path = "/update-trade/:id" component = {UpdateTradeComponent}></Route> */}
                            <Route path = "/exchangeRates" component = {ListExchangeRateComponent}></Route>
                            <Route path = "/add-exchangeRate/:id" component = {CreateExchangeRateComponent}></Route>
                            <Route path = "/view-exchangeRate/:id" component = {ViewExchangeRateComponent}></Route>
                          {/* <Route path = "/update-exchangeRate/:id" component = {UpdateExchangeRateComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
