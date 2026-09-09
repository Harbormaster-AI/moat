from django.db import models
from fintechOnDjango.models.InvestmentAccountType import InvestmentAccountType
from fintechOnDjango.models.AccountStatus import AccountStatus

#======================================================================
# 
# Encapsulates data for model InvestmentAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentAccount Declaration
#======================================================================
class InvestmentAccount (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	accountNumber = AccountNumber
	baseCurrency = models.CharField(max_length=200, null=True)
	balance = Money
	portfolio = models.ForeignKey('InvestmentPortfolio', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	trades = models.ManyToManyField('Trade',  blank=True, related_name='+')
	orders = models.ManyToManyField('TradeOrder',  blank=True, related_name='+')
	accountType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InvestmentAccountType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AccountStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.accountNumber
		str = str + self.baseCurrency
		str = str + self.balance
		str = str + self.accountType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InvestmentAccount";
    
	def objectType(self):
		return "InvestmentAccount";
