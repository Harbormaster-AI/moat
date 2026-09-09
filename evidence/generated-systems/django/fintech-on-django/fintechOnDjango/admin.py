from django.contrib import admin

# Register your models here.
from .models.FinancialInstitution import FinancialInstitution
from .models.Branch import Branch
from .models.ProductOffering import ProductOffering
from .models.PricingPlan import PricingPlan
from .models.FeeSchedule import FeeSchedule
from .models.UsageLimit import UsageLimit
from .models.Customer import Customer
from .models.KYCProfile import KYCProfile
from .models.KYCDocument import KYCDocument
from .models.Screening import Screening
from .models.VerifiedAddress import VerifiedAddress
from .models.CompliancePolicy import CompliancePolicy
from .models.ComplianceAlert import ComplianceAlert
from .models.Consent import Consent
from .models.APIClient import APIClient
from .models.Agreement import Agreement
from .models.Account import Account
from .models.Wallet import Wallet
from .models.PaymentCard import PaymentCard
from .models.CardTokenization import CardTokenization
from .models.Merchant import Merchant
from .models.Terminal import Terminal
from .models.PaymentContract import PaymentContract
from .models.PaymentProcessor import PaymentProcessor
from .models.Transaction import Transaction
from .models.PaymentOrder import PaymentOrder
from .models.Beneficiary import Beneficiary
from .models.AppliedFee import AppliedFee
from .models.FXQuote import FXQuote
from .models.FXDeal import FXDeal
from .models.SettlementBatch import SettlementBatch
from .models.Payout import Payout
from .models.Dispute import Dispute
from .models.Chargeback import Chargeback
from .models.Invoice import Invoice
from .models.AccountStatement import AccountStatement
from .models.DirectDebitMandate import DirectDebitMandate
from .models.Creditor import Creditor
from .models.LoanApplication import LoanApplication
from .models.RiskAssessment import RiskAssessment
from .models.Loan import Loan
from .models.RepaymentSchedule import RepaymentSchedule
from .models.Collateral import Collateral
from .models.LoanTransaction import LoanTransaction
from .models.InvestmentPortfolio import InvestmentPortfolio
from .models.InvestmentAccount import InvestmentAccount
from .models.Security import Security
from .models.Position import Position
from .models.TradeOrder import TradeOrder
from .models.Trade import Trade
from .models.ExchangeRate import ExchangeRate

# Need to add this for each model that requires managing

admin.site.register(FinancialInstitution)
admin.site.register(Branch)
admin.site.register(ProductOffering)
admin.site.register(PricingPlan)
admin.site.register(FeeSchedule)
admin.site.register(UsageLimit)
admin.site.register(Customer)
admin.site.register(KYCProfile)
admin.site.register(KYCDocument)
admin.site.register(Screening)
admin.site.register(VerifiedAddress)
admin.site.register(CompliancePolicy)
admin.site.register(ComplianceAlert)
admin.site.register(Consent)
admin.site.register(APIClient)
admin.site.register(Agreement)
admin.site.register(Account)
admin.site.register(Wallet)
admin.site.register(PaymentCard)
admin.site.register(CardTokenization)
admin.site.register(Merchant)
admin.site.register(Terminal)
admin.site.register(PaymentContract)
admin.site.register(PaymentProcessor)
admin.site.register(Transaction)
admin.site.register(PaymentOrder)
admin.site.register(Beneficiary)
admin.site.register(AppliedFee)
admin.site.register(FXQuote)
admin.site.register(FXDeal)
admin.site.register(SettlementBatch)
admin.site.register(Payout)
admin.site.register(Dispute)
admin.site.register(Chargeback)
admin.site.register(Invoice)
admin.site.register(AccountStatement)
admin.site.register(DirectDebitMandate)
admin.site.register(Creditor)
admin.site.register(LoanApplication)
admin.site.register(RiskAssessment)
admin.site.register(Loan)
admin.site.register(RepaymentSchedule)
admin.site.register(Collateral)
admin.site.register(LoanTransaction)
admin.site.register(InvestmentPortfolio)
admin.site.register(InvestmentAccount)
admin.site.register(Security)
admin.site.register(Position)
admin.site.register(TradeOrder)
admin.site.register(Trade)
admin.site.register(ExchangeRate)
