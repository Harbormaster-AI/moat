import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {FinancialInstitutionService} from '../services/FinancialInstitution.service';
import {BranchService} from '../services/Branch.service';
import {ProductOfferingService} from '../services/ProductOffering.service';
import {PricingPlanService} from '../services/PricingPlan.service';
import {FeeScheduleService} from '../services/FeeSchedule.service';
import {UsageLimitService} from '../services/UsageLimit.service';
import {CustomerService} from '../services/Customer.service';
import {KYCProfileService} from '../services/KYCProfile.service';
import {KYCDocumentService} from '../services/KYCDocument.service';
import {ScreeningService} from '../services/Screening.service';
import {VerifiedAddressService} from '../services/VerifiedAddress.service';
import {CompliancePolicyService} from '../services/CompliancePolicy.service';
import {ComplianceAlertService} from '../services/ComplianceAlert.service';
import {ConsentService} from '../services/Consent.service';
import {APIClientService} from '../services/APIClient.service';
import {AgreementService} from '../services/Agreement.service';
import {AccountService} from '../services/Account.service';
import {WalletService} from '../services/Wallet.service';
import {PaymentCardService} from '../services/PaymentCard.service';
import {CardTokenizationService} from '../services/CardTokenization.service';
import {MerchantService} from '../services/Merchant.service';
import {TerminalService} from '../services/Terminal.service';
import {PaymentContractService} from '../services/PaymentContract.service';
import {PaymentProcessorService} from '../services/PaymentProcessor.service';
import {TransactionService} from '../services/Transaction.service';
import {PaymentOrderService} from '../services/PaymentOrder.service';
import {BeneficiaryService} from '../services/Beneficiary.service';
import {AppliedFeeService} from '../services/AppliedFee.service';
import {FXQuoteService} from '../services/FXQuote.service';
import {FXDealService} from '../services/FXDeal.service';
import {SettlementBatchService} from '../services/SettlementBatch.service';
import {PayoutService} from '../services/Payout.service';
import {DisputeService} from '../services/Dispute.service';
import {ChargebackService} from '../services/Chargeback.service';
import {InvoiceService} from '../services/Invoice.service';
import {AccountStatementService} from '../services/AccountStatement.service';
import {DirectDebitMandateService} from '../services/DirectDebitMandate.service';
import {CreditorService} from '../services/Creditor.service';
import {LoanApplicationService} from '../services/LoanApplication.service';
import {RiskAssessmentService} from '../services/RiskAssessment.service';
import {LoanService} from '../services/Loan.service';
import {RepaymentScheduleService} from '../services/RepaymentSchedule.service';
import {CollateralService} from '../services/Collateral.service';
import {LoanTransactionService} from '../services/LoanTransaction.service';
import {InvestmentPortfolioService} from '../services/InvestmentPortfolio.service';
import {InvestmentAccountService} from '../services/InvestmentAccount.service';
import {SecurityService} from '../services/Security.service';
import {PositionService} from '../services/Position.service';
import {TradeOrderService} from '../services/TradeOrder.service';
import {TradeService} from '../services/Trade.service';
import {ExchangeRateService} from '../services/ExchangeRate.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    ProductCategorys = Object.keys(enumTypes.ProductCategory);
    PlanStatuss = Object.keys(enumTypes.PlanStatus);
    FeeTypes = Object.keys(enumTypes.FeeType);
    FeeCalculationMethods = Object.keys(enumTypes.FeeCalculationMethod);
    LimitScopes = Object.keys(enumTypes.LimitScope);
    LimitPeriods = Object.keys(enumTypes.LimitPeriod);
    CustomerTypes = Object.keys(enumTypes.CustomerType);
    KYCStatuss = Object.keys(enumTypes.KYCStatus);
    VerificationLevels = Object.keys(enumTypes.VerificationLevel);
    KYCDocumentTypes = Object.keys(enumTypes.KYCDocumentType);
    DocumentStatuss = Object.keys(enumTypes.DocumentStatus);
    ScreeningTypes = Object.keys(enumTypes.ScreeningType);
    ScreeningStatuss = Object.keys(enumTypes.ScreeningStatus);
    VerificationStatuss = Object.keys(enumTypes.VerificationStatus);
    PolicyStatuss = Object.keys(enumTypes.PolicyStatus);
    AlertSeveritys = Object.keys(enumTypes.AlertSeverity);
    AlertStatuss = Object.keys(enumTypes.AlertStatus);
    ConsentTypes = Object.keys(enumTypes.ConsentType);
    ConsentStatuss = Object.keys(enumTypes.ConsentStatus);
    ClientTypes = Object.keys(enumTypes.ClientType);
    AgreementTypes = Object.keys(enumTypes.AgreementType);
    AgreementStatuss = Object.keys(enumTypes.AgreementStatus);
    AccountTypes = Object.keys(enumTypes.AccountType);
    AccountStatuss = Object.keys(enumTypes.AccountStatus);
    WalletStatuss = Object.keys(enumTypes.WalletStatus);
    CardSchemes = Object.keys(enumTypes.CardScheme);
    CardStatuss = Object.keys(enumTypes.CardStatus);
    WalletProviders = Object.keys(enumTypes.WalletProvider);
    TokenizationStatuss = Object.keys(enumTypes.TokenizationStatus);
    TerminalTypes = Object.keys(enumTypes.TerminalType);
    TerminalStatuss = Object.keys(enumTypes.TerminalStatus);
    ContractStatuss = Object.keys(enumTypes.ContractStatus);
    TransactionTypes = Object.keys(enumTypes.TransactionType);
    TransactionStatuss = Object.keys(enumTypes.TransactionStatus);
    PaymentMethods = Object.keys(enumTypes.PaymentMethod);
    PaymentOrderStatuss = Object.keys(enumTypes.PaymentOrderStatus);
    PaymentPrioritys = Object.keys(enumTypes.PaymentPriority);
    FXPriceTypes = Object.keys(enumTypes.FXPriceType);
    FXDealStatuss = Object.keys(enumTypes.FXDealStatus);
    SettlementStatuss = Object.keys(enumTypes.SettlementStatus);
    PayoutStatuss = Object.keys(enumTypes.PayoutStatus);
    DisputeReasons = Object.keys(enumTypes.DisputeReason);
    DisputeStatuss = Object.keys(enumTypes.DisputeStatus);
    ChargebackStages = Object.keys(enumTypes.ChargebackStage);
    ChargebackStatuss = Object.keys(enumTypes.ChargebackStatus);
    InvoiceStatuss = Object.keys(enumTypes.InvoiceStatus);
    DirectDebitSchemes = Object.keys(enumTypes.DirectDebitScheme);
    MandateStatuss = Object.keys(enumTypes.MandateStatus);
    LoanProductTypes = Object.keys(enumTypes.LoanProductType);
    LoanPurposes = Object.keys(enumTypes.LoanPurpose);
    ApplicationStatuss = Object.keys(enumTypes.ApplicationStatus);
    DecisionOutcomes = Object.keys(enumTypes.DecisionOutcome);
    InterestRateTypes = Object.keys(enumTypes.InterestRateType);
    LoanStatuss = Object.keys(enumTypes.LoanStatus);
    InstallmentStatuss = Object.keys(enumTypes.InstallmentStatus);
    CollateralTypes = Object.keys(enumTypes.CollateralType);
    LoanTransactionTypes = Object.keys(enumTypes.LoanTransactionType);
    PostingStatuss = Object.keys(enumTypes.PostingStatus);
    PortfolioStatuss = Object.keys(enumTypes.PortfolioStatus);
    InvestmentAccountTypes = Object.keys(enumTypes.InvestmentAccountType);
    SecurityTypes = Object.keys(enumTypes.SecurityType);
    OrderSides = Object.keys(enumTypes.OrderSide);
    OrderTypes = Object.keys(enumTypes.OrderType);
    OrderStatuss = Object.keys(enumTypes.OrderStatus);
    TimeInForces = Object.keys(enumTypes.TimeInForce);

// all collection instances
    financialInstitutions : any;
    branchs : any;
    productOfferings : any;
    pricingPlans : any;
    feeSchedules : any;
    usageLimits : any;
    customers : any;
    kYCProfiles : any;
    kYCDocuments : any;
    screenings : any;
    verifiedAddresss : any;
    compliancePolicys : any;
    complianceAlerts : any;
    consents : any;
    aPIClients : any;
    agreements : any;
    accounts : any;
    wallets : any;
    paymentCards : any;
    cardTokenizations : any;
    merchants : any;
    terminals : any;
    paymentContracts : any;
    paymentProcessors : any;
    transactions : any;
    paymentOrders : any;
    beneficiarys : any;
    appliedFees : any;
    fXQuotes : any;
    fXDeals : any;
    settlementBatchs : any;
    payouts : any;
    disputes : any;
    chargebacks : any;
    invoices : any;
    accountStatements : any;
    directDebitMandates : any;
    creditors : any;
    loanApplications : any;
    riskAssessments : any;
    loans : any;
    repaymentSchedules : any;
    collaterals : any;
    loanTransactions : any;
    investmentPortfolios : any;
    investmentAccounts : any;
    securitys : any;
    positions : any;
    tradeOrders : any;
    trades : any;
    exchangeRates : any;
  
// initialization  
    ngOnInit() {
    }

    initFinancialInstitutionList() {
        if ( this.financialInstitutions == null ) {
            new FinancialInstitutionService(this.http).getFinancialInstitutions().subscribe(res => {
                this.financialInstitutions = res;
            });
        }
    }
    
    initBranchList() {
        if ( this.branchs == null ) {
            new BranchService(this.http).getBranchs().subscribe(res => {
                this.branchs = res;
            });
        }
    }
    
    initProductOfferingList() {
        if ( this.productOfferings == null ) {
            new ProductOfferingService(this.http).getProductOfferings().subscribe(res => {
                this.productOfferings = res;
            });
        }
    }
    
    initPricingPlanList() {
        if ( this.pricingPlans == null ) {
            new PricingPlanService(this.http).getPricingPlans().subscribe(res => {
                this.pricingPlans = res;
            });
        }
    }
    
    initFeeScheduleList() {
        if ( this.feeSchedules == null ) {
            new FeeScheduleService(this.http).getFeeSchedules().subscribe(res => {
                this.feeSchedules = res;
            });
        }
    }
    
    initUsageLimitList() {
        if ( this.usageLimits == null ) {
            new UsageLimitService(this.http).getUsageLimits().subscribe(res => {
                this.usageLimits = res;
            });
        }
    }
    
    initCustomerList() {
        if ( this.customers == null ) {
            new CustomerService(this.http).getCustomers().subscribe(res => {
                this.customers = res;
            });
        }
    }
    
    initKYCProfileList() {
        if ( this.kYCProfiles == null ) {
            new KYCProfileService(this.http).getKYCProfiles().subscribe(res => {
                this.kYCProfiles = res;
            });
        }
    }
    
    initKYCDocumentList() {
        if ( this.kYCDocuments == null ) {
            new KYCDocumentService(this.http).getKYCDocuments().subscribe(res => {
                this.kYCDocuments = res;
            });
        }
    }
    
    initScreeningList() {
        if ( this.screenings == null ) {
            new ScreeningService(this.http).getScreenings().subscribe(res => {
                this.screenings = res;
            });
        }
    }
    
    initVerifiedAddressList() {
        if ( this.verifiedAddresss == null ) {
            new VerifiedAddressService(this.http).getVerifiedAddresss().subscribe(res => {
                this.verifiedAddresss = res;
            });
        }
    }
    
    initCompliancePolicyList() {
        if ( this.compliancePolicys == null ) {
            new CompliancePolicyService(this.http).getCompliancePolicys().subscribe(res => {
                this.compliancePolicys = res;
            });
        }
    }
    
    initComplianceAlertList() {
        if ( this.complianceAlerts == null ) {
            new ComplianceAlertService(this.http).getComplianceAlerts().subscribe(res => {
                this.complianceAlerts = res;
            });
        }
    }
    
    initConsentList() {
        if ( this.consents == null ) {
            new ConsentService(this.http).getConsents().subscribe(res => {
                this.consents = res;
            });
        }
    }
    
    initAPIClientList() {
        if ( this.aPIClients == null ) {
            new APIClientService(this.http).getAPIClients().subscribe(res => {
                this.aPIClients = res;
            });
        }
    }
    
    initAgreementList() {
        if ( this.agreements == null ) {
            new AgreementService(this.http).getAgreements().subscribe(res => {
                this.agreements = res;
            });
        }
    }
    
    initAccountList() {
        if ( this.accounts == null ) {
            new AccountService(this.http).getAccounts().subscribe(res => {
                this.accounts = res;
            });
        }
    }
    
    initWalletList() {
        if ( this.wallets == null ) {
            new WalletService(this.http).getWallets().subscribe(res => {
                this.wallets = res;
            });
        }
    }
    
    initPaymentCardList() {
        if ( this.paymentCards == null ) {
            new PaymentCardService(this.http).getPaymentCards().subscribe(res => {
                this.paymentCards = res;
            });
        }
    }
    
    initCardTokenizationList() {
        if ( this.cardTokenizations == null ) {
            new CardTokenizationService(this.http).getCardTokenizations().subscribe(res => {
                this.cardTokenizations = res;
            });
        }
    }
    
    initMerchantList() {
        if ( this.merchants == null ) {
            new MerchantService(this.http).getMerchants().subscribe(res => {
                this.merchants = res;
            });
        }
    }
    
    initTerminalList() {
        if ( this.terminals == null ) {
            new TerminalService(this.http).getTerminals().subscribe(res => {
                this.terminals = res;
            });
        }
    }
    
    initPaymentContractList() {
        if ( this.paymentContracts == null ) {
            new PaymentContractService(this.http).getPaymentContracts().subscribe(res => {
                this.paymentContracts = res;
            });
        }
    }
    
    initPaymentProcessorList() {
        if ( this.paymentProcessors == null ) {
            new PaymentProcessorService(this.http).getPaymentProcessors().subscribe(res => {
                this.paymentProcessors = res;
            });
        }
    }
    
    initTransactionList() {
        if ( this.transactions == null ) {
            new TransactionService(this.http).getTransactions().subscribe(res => {
                this.transactions = res;
            });
        }
    }
    
    initPaymentOrderList() {
        if ( this.paymentOrders == null ) {
            new PaymentOrderService(this.http).getPaymentOrders().subscribe(res => {
                this.paymentOrders = res;
            });
        }
    }
    
    initBeneficiaryList() {
        if ( this.beneficiarys == null ) {
            new BeneficiaryService(this.http).getBeneficiarys().subscribe(res => {
                this.beneficiarys = res;
            });
        }
    }
    
    initAppliedFeeList() {
        if ( this.appliedFees == null ) {
            new AppliedFeeService(this.http).getAppliedFees().subscribe(res => {
                this.appliedFees = res;
            });
        }
    }
    
    initFXQuoteList() {
        if ( this.fXQuotes == null ) {
            new FXQuoteService(this.http).getFXQuotes().subscribe(res => {
                this.fXQuotes = res;
            });
        }
    }
    
    initFXDealList() {
        if ( this.fXDeals == null ) {
            new FXDealService(this.http).getFXDeals().subscribe(res => {
                this.fXDeals = res;
            });
        }
    }
    
    initSettlementBatchList() {
        if ( this.settlementBatchs == null ) {
            new SettlementBatchService(this.http).getSettlementBatchs().subscribe(res => {
                this.settlementBatchs = res;
            });
        }
    }
    
    initPayoutList() {
        if ( this.payouts == null ) {
            new PayoutService(this.http).getPayouts().subscribe(res => {
                this.payouts = res;
            });
        }
    }
    
    initDisputeList() {
        if ( this.disputes == null ) {
            new DisputeService(this.http).getDisputes().subscribe(res => {
                this.disputes = res;
            });
        }
    }
    
    initChargebackList() {
        if ( this.chargebacks == null ) {
            new ChargebackService(this.http).getChargebacks().subscribe(res => {
                this.chargebacks = res;
            });
        }
    }
    
    initInvoiceList() {
        if ( this.invoices == null ) {
            new InvoiceService(this.http).getInvoices().subscribe(res => {
                this.invoices = res;
            });
        }
    }
    
    initAccountStatementList() {
        if ( this.accountStatements == null ) {
            new AccountStatementService(this.http).getAccountStatements().subscribe(res => {
                this.accountStatements = res;
            });
        }
    }
    
    initDirectDebitMandateList() {
        if ( this.directDebitMandates == null ) {
            new DirectDebitMandateService(this.http).getDirectDebitMandates().subscribe(res => {
                this.directDebitMandates = res;
            });
        }
    }
    
    initCreditorList() {
        if ( this.creditors == null ) {
            new CreditorService(this.http).getCreditors().subscribe(res => {
                this.creditors = res;
            });
        }
    }
    
    initLoanApplicationList() {
        if ( this.loanApplications == null ) {
            new LoanApplicationService(this.http).getLoanApplications().subscribe(res => {
                this.loanApplications = res;
            });
        }
    }
    
    initRiskAssessmentList() {
        if ( this.riskAssessments == null ) {
            new RiskAssessmentService(this.http).getRiskAssessments().subscribe(res => {
                this.riskAssessments = res;
            });
        }
    }
    
    initLoanList() {
        if ( this.loans == null ) {
            new LoanService(this.http).getLoans().subscribe(res => {
                this.loans = res;
            });
        }
    }
    
    initRepaymentScheduleList() {
        if ( this.repaymentSchedules == null ) {
            new RepaymentScheduleService(this.http).getRepaymentSchedules().subscribe(res => {
                this.repaymentSchedules = res;
            });
        }
    }
    
    initCollateralList() {
        if ( this.collaterals == null ) {
            new CollateralService(this.http).getCollaterals().subscribe(res => {
                this.collaterals = res;
            });
        }
    }
    
    initLoanTransactionList() {
        if ( this.loanTransactions == null ) {
            new LoanTransactionService(this.http).getLoanTransactions().subscribe(res => {
                this.loanTransactions = res;
            });
        }
    }
    
    initInvestmentPortfolioList() {
        if ( this.investmentPortfolios == null ) {
            new InvestmentPortfolioService(this.http).getInvestmentPortfolios().subscribe(res => {
                this.investmentPortfolios = res;
            });
        }
    }
    
    initInvestmentAccountList() {
        if ( this.investmentAccounts == null ) {
            new InvestmentAccountService(this.http).getInvestmentAccounts().subscribe(res => {
                this.investmentAccounts = res;
            });
        }
    }
    
    initSecurityList() {
        if ( this.securitys == null ) {
            new SecurityService(this.http).getSecuritys().subscribe(res => {
                this.securitys = res;
            });
        }
    }
    
    initPositionList() {
        if ( this.positions == null ) {
            new PositionService(this.http).getPositions().subscribe(res => {
                this.positions = res;
            });
        }
    }
    
    initTradeOrderList() {
        if ( this.tradeOrders == null ) {
            new TradeOrderService(this.http).getTradeOrders().subscribe(res => {
                this.tradeOrders = res;
            });
        }
    }
    
    initTradeList() {
        if ( this.trades == null ) {
            new TradeService(this.http).getTrades().subscribe(res => {
                this.trades = res;
            });
        }
    }
    
    initExchangeRateList() {
        if ( this.exchangeRates == null ) {
            new ExchangeRateService(this.http).getExchangeRates().subscribe(res => {
                this.exchangeRates = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
