from django.db import models

#======================================================================
# 
# Encapsulates data for model Merchant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Merchant Declaration
#======================================================================
class Merchant (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	mcc = models.CharField(max_length=200, null=True)
	url = models.CharField(max_length=200, null=True)
	country = models.CharField(max_length=200, null=True)
	settlementCurrency = models.CharField(max_length=200, null=True)
	terminals = models.ManyToManyField('Terminal',  blank=True, related_name='+')
	paymentContracts = models.ManyToManyField('PaymentContract',  blank=True, related_name='+')
	payouts = models.ManyToManyField('Payout',  blank=True, related_name='+')
	settlements = models.ManyToManyField('SettlementBatch',  blank=True, related_name='+')
	disputes = models.ManyToManyField('Dispute',  blank=True, related_name='+')
	invoices = models.ManyToManyField('Invoice',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.mcc
		str = str + self.url
		str = str + self.country
		str = str + self.settlementCurrency
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Merchant";
    
	def objectType(self):
		return "Merchant";
