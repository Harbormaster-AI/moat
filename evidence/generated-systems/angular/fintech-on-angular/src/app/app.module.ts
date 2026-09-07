import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatInputModule} from '@angular/material/input';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatMomentDateModule} from "@angular/material-moment-adapter";
import {NgModule} from '@angular/core';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {RouterModule} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';
import {ReactiveFormsModule} from '@angular/forms';
import {AppComponent} from './app.component';
import {MatMenuModule} from '@angular/material/menu';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav'

import {IndexFinancialInstitutionComponent} from './components/FinancialInstitution/index/index.component';
import {CreateFinancialInstitutionComponent} from './components/FinancialInstitution/create/create.component';
import {EditFinancialInstitutionComponent} from './components/FinancialInstitution/edit/edit.component';
import {IndexBranchComponent} from './components/Branch/index/index.component';
import {CreateBranchComponent} from './components/Branch/create/create.component';
import {EditBranchComponent} from './components/Branch/edit/edit.component';
import {IndexProductOfferingComponent} from './components/ProductOffering/index/index.component';
import {CreateProductOfferingComponent} from './components/ProductOffering/create/create.component';
import {EditProductOfferingComponent} from './components/ProductOffering/edit/edit.component';
import {IndexPricingPlanComponent} from './components/PricingPlan/index/index.component';
import {CreatePricingPlanComponent} from './components/PricingPlan/create/create.component';
import {EditPricingPlanComponent} from './components/PricingPlan/edit/edit.component';
import {IndexFeeScheduleComponent} from './components/FeeSchedule/index/index.component';
import {CreateFeeScheduleComponent} from './components/FeeSchedule/create/create.component';
import {EditFeeScheduleComponent} from './components/FeeSchedule/edit/edit.component';
import {IndexUsageLimitComponent} from './components/UsageLimit/index/index.component';
import {CreateUsageLimitComponent} from './components/UsageLimit/create/create.component';
import {EditUsageLimitComponent} from './components/UsageLimit/edit/edit.component';
import {IndexCustomerComponent} from './components/Customer/index/index.component';
import {CreateCustomerComponent} from './components/Customer/create/create.component';
import {EditCustomerComponent} from './components/Customer/edit/edit.component';
import {IndexKYCProfileComponent} from './components/KYCProfile/index/index.component';
import {CreateKYCProfileComponent} from './components/KYCProfile/create/create.component';
import {EditKYCProfileComponent} from './components/KYCProfile/edit/edit.component';
import {IndexKYCDocumentComponent} from './components/KYCDocument/index/index.component';
import {CreateKYCDocumentComponent} from './components/KYCDocument/create/create.component';
import {EditKYCDocumentComponent} from './components/KYCDocument/edit/edit.component';
import {IndexScreeningComponent} from './components/Screening/index/index.component';
import {CreateScreeningComponent} from './components/Screening/create/create.component';
import {EditScreeningComponent} from './components/Screening/edit/edit.component';
import {IndexVerifiedAddressComponent} from './components/VerifiedAddress/index/index.component';
import {CreateVerifiedAddressComponent} from './components/VerifiedAddress/create/create.component';
import {EditVerifiedAddressComponent} from './components/VerifiedAddress/edit/edit.component';
import {IndexCompliancePolicyComponent} from './components/CompliancePolicy/index/index.component';
import {CreateCompliancePolicyComponent} from './components/CompliancePolicy/create/create.component';
import {EditCompliancePolicyComponent} from './components/CompliancePolicy/edit/edit.component';
import {IndexComplianceAlertComponent} from './components/ComplianceAlert/index/index.component';
import {CreateComplianceAlertComponent} from './components/ComplianceAlert/create/create.component';
import {EditComplianceAlertComponent} from './components/ComplianceAlert/edit/edit.component';
import {IndexConsentComponent} from './components/Consent/index/index.component';
import {CreateConsentComponent} from './components/Consent/create/create.component';
import {EditConsentComponent} from './components/Consent/edit/edit.component';
import {IndexAPIClientComponent} from './components/APIClient/index/index.component';
import {CreateAPIClientComponent} from './components/APIClient/create/create.component';
import {EditAPIClientComponent} from './components/APIClient/edit/edit.component';
import {IndexAgreementComponent} from './components/Agreement/index/index.component';
import {CreateAgreementComponent} from './components/Agreement/create/create.component';
import {EditAgreementComponent} from './components/Agreement/edit/edit.component';
import {IndexAccountComponent} from './components/Account/index/index.component';
import {CreateAccountComponent} from './components/Account/create/create.component';
import {EditAccountComponent} from './components/Account/edit/edit.component';
import {IndexWalletComponent} from './components/Wallet/index/index.component';
import {CreateWalletComponent} from './components/Wallet/create/create.component';
import {EditWalletComponent} from './components/Wallet/edit/edit.component';
import {IndexPaymentCardComponent} from './components/PaymentCard/index/index.component';
import {CreatePaymentCardComponent} from './components/PaymentCard/create/create.component';
import {EditPaymentCardComponent} from './components/PaymentCard/edit/edit.component';
import {IndexCardTokenizationComponent} from './components/CardTokenization/index/index.component';
import {CreateCardTokenizationComponent} from './components/CardTokenization/create/create.component';
import {EditCardTokenizationComponent} from './components/CardTokenization/edit/edit.component';
import {IndexMerchantComponent} from './components/Merchant/index/index.component';
import {CreateMerchantComponent} from './components/Merchant/create/create.component';
import {EditMerchantComponent} from './components/Merchant/edit/edit.component';
import {IndexTerminalComponent} from './components/Terminal/index/index.component';
import {CreateTerminalComponent} from './components/Terminal/create/create.component';
import {EditTerminalComponent} from './components/Terminal/edit/edit.component';
import {IndexPaymentContractComponent} from './components/PaymentContract/index/index.component';
import {CreatePaymentContractComponent} from './components/PaymentContract/create/create.component';
import {EditPaymentContractComponent} from './components/PaymentContract/edit/edit.component';
import {IndexPaymentProcessorComponent} from './components/PaymentProcessor/index/index.component';
import {CreatePaymentProcessorComponent} from './components/PaymentProcessor/create/create.component';
import {EditPaymentProcessorComponent} from './components/PaymentProcessor/edit/edit.component';
import {IndexTransactionComponent} from './components/Transaction/index/index.component';
import {CreateTransactionComponent} from './components/Transaction/create/create.component';
import {EditTransactionComponent} from './components/Transaction/edit/edit.component';
import {IndexPaymentOrderComponent} from './components/PaymentOrder/index/index.component';
import {CreatePaymentOrderComponent} from './components/PaymentOrder/create/create.component';
import {EditPaymentOrderComponent} from './components/PaymentOrder/edit/edit.component';
import {IndexBeneficiaryComponent} from './components/Beneficiary/index/index.component';
import {CreateBeneficiaryComponent} from './components/Beneficiary/create/create.component';
import {EditBeneficiaryComponent} from './components/Beneficiary/edit/edit.component';
import {IndexAppliedFeeComponent} from './components/AppliedFee/index/index.component';
import {CreateAppliedFeeComponent} from './components/AppliedFee/create/create.component';
import {EditAppliedFeeComponent} from './components/AppliedFee/edit/edit.component';
import {IndexFXQuoteComponent} from './components/FXQuote/index/index.component';
import {CreateFXQuoteComponent} from './components/FXQuote/create/create.component';
import {EditFXQuoteComponent} from './components/FXQuote/edit/edit.component';
import {IndexFXDealComponent} from './components/FXDeal/index/index.component';
import {CreateFXDealComponent} from './components/FXDeal/create/create.component';
import {EditFXDealComponent} from './components/FXDeal/edit/edit.component';
import {IndexSettlementBatchComponent} from './components/SettlementBatch/index/index.component';
import {CreateSettlementBatchComponent} from './components/SettlementBatch/create/create.component';
import {EditSettlementBatchComponent} from './components/SettlementBatch/edit/edit.component';
import {IndexPayoutComponent} from './components/Payout/index/index.component';
import {CreatePayoutComponent} from './components/Payout/create/create.component';
import {EditPayoutComponent} from './components/Payout/edit/edit.component';
import {IndexDisputeComponent} from './components/Dispute/index/index.component';
import {CreateDisputeComponent} from './components/Dispute/create/create.component';
import {EditDisputeComponent} from './components/Dispute/edit/edit.component';
import {IndexChargebackComponent} from './components/Chargeback/index/index.component';
import {CreateChargebackComponent} from './components/Chargeback/create/create.component';
import {EditChargebackComponent} from './components/Chargeback/edit/edit.component';
import {IndexInvoiceComponent} from './components/Invoice/index/index.component';
import {CreateInvoiceComponent} from './components/Invoice/create/create.component';
import {EditInvoiceComponent} from './components/Invoice/edit/edit.component';
import {IndexAccountStatementComponent} from './components/AccountStatement/index/index.component';
import {CreateAccountStatementComponent} from './components/AccountStatement/create/create.component';
import {EditAccountStatementComponent} from './components/AccountStatement/edit/edit.component';
import {IndexDirectDebitMandateComponent} from './components/DirectDebitMandate/index/index.component';
import {CreateDirectDebitMandateComponent} from './components/DirectDebitMandate/create/create.component';
import {EditDirectDebitMandateComponent} from './components/DirectDebitMandate/edit/edit.component';
import {IndexCreditorComponent} from './components/Creditor/index/index.component';
import {CreateCreditorComponent} from './components/Creditor/create/create.component';
import {EditCreditorComponent} from './components/Creditor/edit/edit.component';
import {IndexLoanApplicationComponent} from './components/LoanApplication/index/index.component';
import {CreateLoanApplicationComponent} from './components/LoanApplication/create/create.component';
import {EditLoanApplicationComponent} from './components/LoanApplication/edit/edit.component';
import {IndexRiskAssessmentComponent} from './components/RiskAssessment/index/index.component';
import {CreateRiskAssessmentComponent} from './components/RiskAssessment/create/create.component';
import {EditRiskAssessmentComponent} from './components/RiskAssessment/edit/edit.component';
import {IndexLoanComponent} from './components/Loan/index/index.component';
import {CreateLoanComponent} from './components/Loan/create/create.component';
import {EditLoanComponent} from './components/Loan/edit/edit.component';
import {IndexRepaymentScheduleComponent} from './components/RepaymentSchedule/index/index.component';
import {CreateRepaymentScheduleComponent} from './components/RepaymentSchedule/create/create.component';
import {EditRepaymentScheduleComponent} from './components/RepaymentSchedule/edit/edit.component';
import {IndexCollateralComponent} from './components/Collateral/index/index.component';
import {CreateCollateralComponent} from './components/Collateral/create/create.component';
import {EditCollateralComponent} from './components/Collateral/edit/edit.component';
import {IndexLoanTransactionComponent} from './components/LoanTransaction/index/index.component';
import {CreateLoanTransactionComponent} from './components/LoanTransaction/create/create.component';
import {EditLoanTransactionComponent} from './components/LoanTransaction/edit/edit.component';
import {IndexInvestmentPortfolioComponent} from './components/InvestmentPortfolio/index/index.component';
import {CreateInvestmentPortfolioComponent} from './components/InvestmentPortfolio/create/create.component';
import {EditInvestmentPortfolioComponent} from './components/InvestmentPortfolio/edit/edit.component';
import {IndexInvestmentAccountComponent} from './components/InvestmentAccount/index/index.component';
import {CreateInvestmentAccountComponent} from './components/InvestmentAccount/create/create.component';
import {EditInvestmentAccountComponent} from './components/InvestmentAccount/edit/edit.component';
import {IndexSecurityComponent} from './components/Security/index/index.component';
import {CreateSecurityComponent} from './components/Security/create/create.component';
import {EditSecurityComponent} from './components/Security/edit/edit.component';
import {IndexPositionComponent} from './components/Position/index/index.component';
import {CreatePositionComponent} from './components/Position/create/create.component';
import {EditPositionComponent} from './components/Position/edit/edit.component';
import {IndexTradeOrderComponent} from './components/TradeOrder/index/index.component';
import {CreateTradeOrderComponent} from './components/TradeOrder/create/create.component';
import {EditTradeOrderComponent} from './components/TradeOrder/edit/edit.component';
import {IndexTradeComponent} from './components/Trade/index/index.component';
import {CreateTradeComponent} from './components/Trade/create/create.component';
import {EditTradeComponent} from './components/Trade/edit/edit.component';
import {IndexExchangeRateComponent} from './components/ExchangeRate/index/index.component';
import {CreateExchangeRateComponent} from './components/ExchangeRate/create/create.component';
import {EditExchangeRateComponent} from './components/ExchangeRate/edit/edit.component';

import * as appRoutes from './routerConfig';

import {FinancialInstitutionService} from './services/FinancialInstitution.service';
import {BranchService} from './services/Branch.service';
import {ProductOfferingService} from './services/ProductOffering.service';
import {PricingPlanService} from './services/PricingPlan.service';
import {FeeScheduleService} from './services/FeeSchedule.service';
import {UsageLimitService} from './services/UsageLimit.service';
import {CustomerService} from './services/Customer.service';
import {KYCProfileService} from './services/KYCProfile.service';
import {KYCDocumentService} from './services/KYCDocument.service';
import {ScreeningService} from './services/Screening.service';
import {VerifiedAddressService} from './services/VerifiedAddress.service';
import {CompliancePolicyService} from './services/CompliancePolicy.service';
import {ComplianceAlertService} from './services/ComplianceAlert.service';
import {ConsentService} from './services/Consent.service';
import {APIClientService} from './services/APIClient.service';
import {AgreementService} from './services/Agreement.service';
import {AccountService} from './services/Account.service';
import {WalletService} from './services/Wallet.service';
import {PaymentCardService} from './services/PaymentCard.service';
import {CardTokenizationService} from './services/CardTokenization.service';
import {MerchantService} from './services/Merchant.service';
import {TerminalService} from './services/Terminal.service';
import {PaymentContractService} from './services/PaymentContract.service';
import {PaymentProcessorService} from './services/PaymentProcessor.service';
import {TransactionService} from './services/Transaction.service';
import {PaymentOrderService} from './services/PaymentOrder.service';
import {BeneficiaryService} from './services/Beneficiary.service';
import {AppliedFeeService} from './services/AppliedFee.service';
import {FXQuoteService} from './services/FXQuote.service';
import {FXDealService} from './services/FXDeal.service';
import {SettlementBatchService} from './services/SettlementBatch.service';
import {PayoutService} from './services/Payout.service';
import {DisputeService} from './services/Dispute.service';
import {ChargebackService} from './services/Chargeback.service';
import {InvoiceService} from './services/Invoice.service';
import {AccountStatementService} from './services/AccountStatement.service';
import {DirectDebitMandateService} from './services/DirectDebitMandate.service';
import {CreditorService} from './services/Creditor.service';
import {LoanApplicationService} from './services/LoanApplication.service';
import {RiskAssessmentService} from './services/RiskAssessment.service';
import {LoanService} from './services/Loan.service';
import {RepaymentScheduleService} from './services/RepaymentSchedule.service';
import {CollateralService} from './services/Collateral.service';
import {LoanTransactionService} from './services/LoanTransaction.service';
import {InvestmentPortfolioService} from './services/InvestmentPortfolio.service';
import {InvestmentAccountService} from './services/InvestmentAccount.service';
import {SecurityService} from './services/Security.service';
import {PositionService} from './services/Position.service';
import {TradeOrderService} from './services/TradeOrder.service';
import {TradeService} from './services/Trade.service';
import {ExchangeRateService} from './services/ExchangeRate.service';

@NgModule({
  declarations: [
    IndexFinancialInstitutionComponent,
    CreateFinancialInstitutionComponent,
    EditFinancialInstitutionComponent,
    IndexBranchComponent,
    CreateBranchComponent,
    EditBranchComponent,
    IndexProductOfferingComponent,
    CreateProductOfferingComponent,
    EditProductOfferingComponent,
    IndexPricingPlanComponent,
    CreatePricingPlanComponent,
    EditPricingPlanComponent,
    IndexFeeScheduleComponent,
    CreateFeeScheduleComponent,
    EditFeeScheduleComponent,
    IndexUsageLimitComponent,
    CreateUsageLimitComponent,
    EditUsageLimitComponent,
    IndexCustomerComponent,
    CreateCustomerComponent,
    EditCustomerComponent,
    IndexKYCProfileComponent,
    CreateKYCProfileComponent,
    EditKYCProfileComponent,
    IndexKYCDocumentComponent,
    CreateKYCDocumentComponent,
    EditKYCDocumentComponent,
    IndexScreeningComponent,
    CreateScreeningComponent,
    EditScreeningComponent,
    IndexVerifiedAddressComponent,
    CreateVerifiedAddressComponent,
    EditVerifiedAddressComponent,
    IndexCompliancePolicyComponent,
    CreateCompliancePolicyComponent,
    EditCompliancePolicyComponent,
    IndexComplianceAlertComponent,
    CreateComplianceAlertComponent,
    EditComplianceAlertComponent,
    IndexConsentComponent,
    CreateConsentComponent,
    EditConsentComponent,
    IndexAPIClientComponent,
    CreateAPIClientComponent,
    EditAPIClientComponent,
    IndexAgreementComponent,
    CreateAgreementComponent,
    EditAgreementComponent,
    IndexAccountComponent,
    CreateAccountComponent,
    EditAccountComponent,
    IndexWalletComponent,
    CreateWalletComponent,
    EditWalletComponent,
    IndexPaymentCardComponent,
    CreatePaymentCardComponent,
    EditPaymentCardComponent,
    IndexCardTokenizationComponent,
    CreateCardTokenizationComponent,
    EditCardTokenizationComponent,
    IndexMerchantComponent,
    CreateMerchantComponent,
    EditMerchantComponent,
    IndexTerminalComponent,
    CreateTerminalComponent,
    EditTerminalComponent,
    IndexPaymentContractComponent,
    CreatePaymentContractComponent,
    EditPaymentContractComponent,
    IndexPaymentProcessorComponent,
    CreatePaymentProcessorComponent,
    EditPaymentProcessorComponent,
    IndexTransactionComponent,
    CreateTransactionComponent,
    EditTransactionComponent,
    IndexPaymentOrderComponent,
    CreatePaymentOrderComponent,
    EditPaymentOrderComponent,
    IndexBeneficiaryComponent,
    CreateBeneficiaryComponent,
    EditBeneficiaryComponent,
    IndexAppliedFeeComponent,
    CreateAppliedFeeComponent,
    EditAppliedFeeComponent,
    IndexFXQuoteComponent,
    CreateFXQuoteComponent,
    EditFXQuoteComponent,
    IndexFXDealComponent,
    CreateFXDealComponent,
    EditFXDealComponent,
    IndexSettlementBatchComponent,
    CreateSettlementBatchComponent,
    EditSettlementBatchComponent,
    IndexPayoutComponent,
    CreatePayoutComponent,
    EditPayoutComponent,
    IndexDisputeComponent,
    CreateDisputeComponent,
    EditDisputeComponent,
    IndexChargebackComponent,
    CreateChargebackComponent,
    EditChargebackComponent,
    IndexInvoiceComponent,
    CreateInvoiceComponent,
    EditInvoiceComponent,
    IndexAccountStatementComponent,
    CreateAccountStatementComponent,
    EditAccountStatementComponent,
    IndexDirectDebitMandateComponent,
    CreateDirectDebitMandateComponent,
    EditDirectDebitMandateComponent,
    IndexCreditorComponent,
    CreateCreditorComponent,
    EditCreditorComponent,
    IndexLoanApplicationComponent,
    CreateLoanApplicationComponent,
    EditLoanApplicationComponent,
    IndexRiskAssessmentComponent,
    CreateRiskAssessmentComponent,
    EditRiskAssessmentComponent,
    IndexLoanComponent,
    CreateLoanComponent,
    EditLoanComponent,
    IndexRepaymentScheduleComponent,
    CreateRepaymentScheduleComponent,
    EditRepaymentScheduleComponent,
    IndexCollateralComponent,
    CreateCollateralComponent,
    EditCollateralComponent,
    IndexLoanTransactionComponent,
    CreateLoanTransactionComponent,
    EditLoanTransactionComponent,
    IndexInvestmentPortfolioComponent,
    CreateInvestmentPortfolioComponent,
    EditInvestmentPortfolioComponent,
    IndexInvestmentAccountComponent,
    CreateInvestmentAccountComponent,
    EditInvestmentAccountComponent,
    IndexSecurityComponent,
    CreateSecurityComponent,
    EditSecurityComponent,
    IndexPositionComponent,
    CreatePositionComponent,
    EditPositionComponent,
    IndexTradeOrderComponent,
    CreateTradeOrderComponent,
    EditTradeOrderComponent,
    IndexTradeComponent,
    CreateTradeComponent,
    EditTradeComponent,
    IndexExchangeRateComponent,
    CreateExchangeRateComponent,
    EditExchangeRateComponent,
    AppComponent
  ],
  imports: [

    BrowserModule, 
    NgbModule,
    MatMenuModule,
    MatToolbarModule,
    MatCheckboxModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDatepickerModule,
	MatMomentDateModule,
    BrowserAnimationsModule,
	HttpClientModule, 
    ReactiveFormsModule,
    FormsModule,
    MatSidenavModule,    
    RouterModule.forRoot(appRoutes.FinancialInstitutionRoutes), 
    RouterModule.forRoot(appRoutes.BranchRoutes), 
    RouterModule.forRoot(appRoutes.ProductOfferingRoutes), 
    RouterModule.forRoot(appRoutes.PricingPlanRoutes), 
    RouterModule.forRoot(appRoutes.FeeScheduleRoutes), 
    RouterModule.forRoot(appRoutes.UsageLimitRoutes), 
    RouterModule.forRoot(appRoutes.CustomerRoutes), 
    RouterModule.forRoot(appRoutes.KYCProfileRoutes), 
    RouterModule.forRoot(appRoutes.KYCDocumentRoutes), 
    RouterModule.forRoot(appRoutes.ScreeningRoutes), 
    RouterModule.forRoot(appRoutes.VerifiedAddressRoutes), 
    RouterModule.forRoot(appRoutes.CompliancePolicyRoutes), 
    RouterModule.forRoot(appRoutes.ComplianceAlertRoutes), 
    RouterModule.forRoot(appRoutes.ConsentRoutes), 
    RouterModule.forRoot(appRoutes.APIClientRoutes), 
    RouterModule.forRoot(appRoutes.AgreementRoutes), 
    RouterModule.forRoot(appRoutes.AccountRoutes), 
    RouterModule.forRoot(appRoutes.WalletRoutes), 
    RouterModule.forRoot(appRoutes.PaymentCardRoutes), 
    RouterModule.forRoot(appRoutes.CardTokenizationRoutes), 
    RouterModule.forRoot(appRoutes.MerchantRoutes), 
    RouterModule.forRoot(appRoutes.TerminalRoutes), 
    RouterModule.forRoot(appRoutes.PaymentContractRoutes), 
    RouterModule.forRoot(appRoutes.PaymentProcessorRoutes), 
    RouterModule.forRoot(appRoutes.TransactionRoutes), 
    RouterModule.forRoot(appRoutes.PaymentOrderRoutes), 
    RouterModule.forRoot(appRoutes.BeneficiaryRoutes), 
    RouterModule.forRoot(appRoutes.AppliedFeeRoutes), 
    RouterModule.forRoot(appRoutes.FXQuoteRoutes), 
    RouterModule.forRoot(appRoutes.FXDealRoutes), 
    RouterModule.forRoot(appRoutes.SettlementBatchRoutes), 
    RouterModule.forRoot(appRoutes.PayoutRoutes), 
    RouterModule.forRoot(appRoutes.DisputeRoutes), 
    RouterModule.forRoot(appRoutes.ChargebackRoutes), 
    RouterModule.forRoot(appRoutes.InvoiceRoutes), 
    RouterModule.forRoot(appRoutes.AccountStatementRoutes), 
    RouterModule.forRoot(appRoutes.DirectDebitMandateRoutes), 
    RouterModule.forRoot(appRoutes.CreditorRoutes), 
    RouterModule.forRoot(appRoutes.LoanApplicationRoutes), 
    RouterModule.forRoot(appRoutes.RiskAssessmentRoutes), 
    RouterModule.forRoot(appRoutes.LoanRoutes), 
    RouterModule.forRoot(appRoutes.RepaymentScheduleRoutes), 
    RouterModule.forRoot(appRoutes.CollateralRoutes), 
    RouterModule.forRoot(appRoutes.LoanTransactionRoutes), 
    RouterModule.forRoot(appRoutes.InvestmentPortfolioRoutes), 
    RouterModule.forRoot(appRoutes.InvestmentAccountRoutes), 
    RouterModule.forRoot(appRoutes.SecurityRoutes), 
    RouterModule.forRoot(appRoutes.PositionRoutes), 
    RouterModule.forRoot(appRoutes.TradeOrderRoutes), 
    RouterModule.forRoot(appRoutes.TradeRoutes), 
    RouterModule.forRoot(appRoutes.ExchangeRateRoutes), 
  ],
  providers: [FinancialInstitutionService,BranchService,ProductOfferingService,PricingPlanService,FeeScheduleService,UsageLimitService,CustomerService,KYCProfileService,KYCDocumentService,ScreeningService,VerifiedAddressService,CompliancePolicyService,ComplianceAlertService,ConsentService,APIClientService,AgreementService,AccountService,WalletService,PaymentCardService,CardTokenizationService,MerchantService,TerminalService,PaymentContractService,PaymentProcessorService,TransactionService,PaymentOrderService,BeneficiaryService,AppliedFeeService,FXQuoteService,FXDealService,SettlementBatchService,PayoutService,DisputeService,ChargebackService,InvoiceService,AccountStatementService,DirectDebitMandateService,CreditorService,LoanApplicationService,RiskAssessmentService,LoanService,RepaymentScheduleService,CollateralService,LoanTransactionService,InvestmentPortfolioService,InvestmentAccountService,SecurityService,PositionService,TradeOrderService,TradeService,ExchangeRateService],
  bootstrap: [AppComponent]
})
export class AppModule { }
