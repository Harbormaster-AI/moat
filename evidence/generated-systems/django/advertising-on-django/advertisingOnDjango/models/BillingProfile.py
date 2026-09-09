from django.db import models
from advertisingOnDjango.models.PaymentTerms import PaymentTerms

#======================================================================
# 
# Encapsulates data for model BillingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingProfile Declaration
#======================================================================
class BillingProfile (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	billingName = models.CharField(max_length=200, null=True)
	taxId = models.CharField(max_length=200, null=True)
	billingAddress = Address
	advertiser = models.ForeignKey('Advertiser', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	paymentMethods = models.ManyToManyField('PaymentMethod',  blank=True, related_name='+')
	adAccounts = models.ManyToManyField('AdAccount',  blank=True, related_name='+')
	paymentTerms = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentTerms])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.billingName
		str = str + self.taxId
		str = str + self.billingAddress
		str = str + self.paymentTerms
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BillingProfile";
    
	def objectType(self):
		return "BillingProfile";
