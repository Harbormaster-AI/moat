"""mainsite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
urlpatterns = [
    path('FinancialInstitution/', include('fintechOnDjango.urls.FinancialInstitutionUrls')),
    path('Branch/', include('fintechOnDjango.urls.BranchUrls')),
    path('ProductOffering/', include('fintechOnDjango.urls.ProductOfferingUrls')),
    path('PricingPlan/', include('fintechOnDjango.urls.PricingPlanUrls')),
    path('FeeSchedule/', include('fintechOnDjango.urls.FeeScheduleUrls')),
    path('UsageLimit/', include('fintechOnDjango.urls.UsageLimitUrls')),
    path('Customer/', include('fintechOnDjango.urls.CustomerUrls')),
    path('KYCProfile/', include('fintechOnDjango.urls.KYCProfileUrls')),
    path('KYCDocument/', include('fintechOnDjango.urls.KYCDocumentUrls')),
    path('Screening/', include('fintechOnDjango.urls.ScreeningUrls')),
    path('VerifiedAddress/', include('fintechOnDjango.urls.VerifiedAddressUrls')),
    path('CompliancePolicy/', include('fintechOnDjango.urls.CompliancePolicyUrls')),
    path('ComplianceAlert/', include('fintechOnDjango.urls.ComplianceAlertUrls')),
    path('Consent/', include('fintechOnDjango.urls.ConsentUrls')),
    path('APIClient/', include('fintechOnDjango.urls.APIClientUrls')),
    path('Agreement/', include('fintechOnDjango.urls.AgreementUrls')),
    path('Account/', include('fintechOnDjango.urls.AccountUrls')),
    path('Wallet/', include('fintechOnDjango.urls.WalletUrls')),
    path('PaymentCard/', include('fintechOnDjango.urls.PaymentCardUrls')),
    path('CardTokenization/', include('fintechOnDjango.urls.CardTokenizationUrls')),
    path('Merchant/', include('fintechOnDjango.urls.MerchantUrls')),
    path('Terminal/', include('fintechOnDjango.urls.TerminalUrls')),
    path('PaymentContract/', include('fintechOnDjango.urls.PaymentContractUrls')),
    path('PaymentProcessor/', include('fintechOnDjango.urls.PaymentProcessorUrls')),
    path('Transaction/', include('fintechOnDjango.urls.TransactionUrls')),
    path('PaymentOrder/', include('fintechOnDjango.urls.PaymentOrderUrls')),
    path('Beneficiary/', include('fintechOnDjango.urls.BeneficiaryUrls')),
    path('AppliedFee/', include('fintechOnDjango.urls.AppliedFeeUrls')),
    path('FXQuote/', include('fintechOnDjango.urls.FXQuoteUrls')),
    path('FXDeal/', include('fintechOnDjango.urls.FXDealUrls')),
    path('SettlementBatch/', include('fintechOnDjango.urls.SettlementBatchUrls')),
    path('Payout/', include('fintechOnDjango.urls.PayoutUrls')),
    path('Dispute/', include('fintechOnDjango.urls.DisputeUrls')),
    path('Chargeback/', include('fintechOnDjango.urls.ChargebackUrls')),
    path('Invoice/', include('fintechOnDjango.urls.InvoiceUrls')),
    path('AccountStatement/', include('fintechOnDjango.urls.AccountStatementUrls')),
    path('DirectDebitMandate/', include('fintechOnDjango.urls.DirectDebitMandateUrls')),
    path('Creditor/', include('fintechOnDjango.urls.CreditorUrls')),
    path('LoanApplication/', include('fintechOnDjango.urls.LoanApplicationUrls')),
    path('RiskAssessment/', include('fintechOnDjango.urls.RiskAssessmentUrls')),
    path('Loan/', include('fintechOnDjango.urls.LoanUrls')),
    path('RepaymentSchedule/', include('fintechOnDjango.urls.RepaymentScheduleUrls')),
    path('Collateral/', include('fintechOnDjango.urls.CollateralUrls')),
    path('LoanTransaction/', include('fintechOnDjango.urls.LoanTransactionUrls')),
    path('InvestmentPortfolio/', include('fintechOnDjango.urls.InvestmentPortfolioUrls')),
    path('InvestmentAccount/', include('fintechOnDjango.urls.InvestmentAccountUrls')),
    path('Security/', include('fintechOnDjango.urls.SecurityUrls')),
    path('Position/', include('fintechOnDjango.urls.PositionUrls')),
    path('TradeOrder/', include('fintechOnDjango.urls.TradeOrderUrls')),
    path('Trade/', include('fintechOnDjango.urls.TradeUrls')),
    path('ExchangeRate/', include('fintechOnDjango.urls.ExchangeRateUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]