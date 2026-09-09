from django.db import models
from ecommerceOnDjango.models.PaymentProviderType import PaymentProviderType

#======================================================================
# 
# Encapsulates data for model PaymentProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProvider Declaration
#======================================================================
class PaymentProvider (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	enabled = models.BooleanField(null=True)
	merchantAccountId = models.CharField(max_length=200, null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	channels = models.ManyToManyField('Channel',  blank=True, related_name='+')
	payments = models.ManyToManyField('Payment',  blank=True, related_name='+')
	subscriptions = models.ManyToManyField('Subscription',  blank=True, related_name='+')
	providerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentProviderType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.enabled
		str = str + self.merchantAccountId
		str = str + self.providerType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PaymentProvider";
    
	def objectType(self):
		return "PaymentProvider";
