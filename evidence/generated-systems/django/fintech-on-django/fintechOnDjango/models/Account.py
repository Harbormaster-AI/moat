from django.db import models
from fintechOnDjango.models.AccountType import AccountType
from fintechOnDjango.models.AccountStatus import AccountStatus

#======================================================================
# 
# Encapsulates data for model Account
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Account Declaration
#======================================================================
class Account (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	accountNumber = AccountNumber
	iban = IBAN
	bic = BIC
	openedDate = models.DateField(null=True)
	currency = models.CharField(max_length=200, null=True)
	balance = Money
	availableBalance = Money
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	institution = models.ForeignKey('FinancialInstitution', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transactions = models.ManyToManyField('Transaction',  blank=True, related_name='+')
	cards = models.ManyToManyField('PaymentCard',  blank=True, related_name='+')
	statements = models.ManyToManyField('AccountStatement',  blank=True, related_name='+')
	mandates = models.ManyToManyField('DirectDebitMandate',  blank=True, related_name='+')
	accountType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccountType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccountStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.accountNumber
		str = str + self.iban
		str = str + self.bic
		str = str + self.openedDate
		str = str + self.currency
		str = str + self.balance
		str = str + self.availableBalance
		str = str + self.accountType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Account";
    
	def objectType(self):
		return "Account";
