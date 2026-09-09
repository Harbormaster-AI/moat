from django.db import models
from fintechOnDjango.models.CustomerType import CustomerType

#======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Customer Declaration
#======================================================================
class Customer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	dateOfBirth = models.DateField(null=True)
	email = Email
	phone = PhoneNumber
	address = Address
	taxId = TaxId
	riskScore = RiskScore
	institution = models.ForeignKey('FinancialInstitution', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	accounts = models.ManyToManyField('Account',  blank=True, related_name='+')
	wallets = models.ManyToManyField('Wallet',  blank=True, related_name='+')
	cards = models.ManyToManyField('PaymentCard',  blank=True, related_name='+')
	kycProfiles = models.ManyToManyField('KYCProfile',  blank=True, related_name='+')
	consents = models.ManyToManyField('Consent',  blank=True, related_name='+')
	agreements = models.ManyToManyField('Agreement',  blank=True, related_name='+')
	loanApplications = models.ManyToManyField('LoanApplication',  blank=True, related_name='+')
	loans = models.ManyToManyField('Loan',  blank=True, related_name='+')
	portfolios = models.ManyToManyField('InvestmentPortfolio',  blank=True, related_name='+')
	disputes = models.ManyToManyField('Dispute',  blank=True, related_name='+')
	customerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CustomerType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.dateOfBirth
		str = str + self.email
		str = str + self.phone
		str = str + self.address
		str = str + self.taxId
		str = str + self.riskScore
		str = str + self.customerType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Customer";
    
	def objectType(self):
		return "Customer";
