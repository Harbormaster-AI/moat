from django.db import models
from insuranceOnDjango.models.BillingStatus import BillingStatus

#======================================================================
# 
# Encapsulates data for model BillingAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingAccount Declaration
#======================================================================
class BillingAccount (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	accountNumber = models.CharField(max_length=200, null=True)
	balance = Money
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	invoices = models.ManyToManyField('Invoice',  blank=True, related_name='+')
	payments = models.ManyToManyField('Payment',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BillingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.accountNumber
		str = str + self.balance
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BillingAccount";
    
	def objectType(self):
		return "BillingAccount";
