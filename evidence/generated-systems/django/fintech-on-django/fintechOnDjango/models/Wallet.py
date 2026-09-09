from django.db import models
from fintechOnDjango.models.WalletStatus import WalletStatus

#======================================================================
# 
# Encapsulates data for model Wallet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Wallet Declaration
#======================================================================
class Wallet (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	currency = models.CharField(max_length=200, null=True)
	balance = Money
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transactions = models.ManyToManyField('Transaction',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WalletStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.currency
		str = str + self.balance
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Wallet";
    
	def objectType(self):
		return "Wallet";
