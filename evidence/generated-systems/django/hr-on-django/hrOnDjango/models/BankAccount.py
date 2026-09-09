from django.db import models

#======================================================================
# 
# Encapsulates data for model BankAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BankAccount Declaration
#======================================================================
class BankAccount (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	accountHolder = models.CharField(max_length=200, null=True)
	bankName = models.CharField(max_length=200, null=True)
	iban = models.CharField(max_length=200, null=True)
	bic = models.CharField(max_length=200, null=True)
	accountNumber = models.CharField(max_length=200, null=True)
	routingNumber = models.CharField(max_length=200, null=True)

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.accountHolder
		str = str + self.bankName
		str = str + self.iban
		str = str + self.bic
		str = str + self.accountNumber
		str = str + self.routingNumber
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BankAccount";
    
	def objectType(self):
		return "BankAccount";
