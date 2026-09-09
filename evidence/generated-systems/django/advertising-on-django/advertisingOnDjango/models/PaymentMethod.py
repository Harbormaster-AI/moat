from django.db import models
from advertisingOnDjango.models.PaymentMethodType import PaymentMethodType

#======================================================================
# 
# Encapsulates data for model PaymentMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethod Declaration
#======================================================================
class PaymentMethod (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	last4 = models.CharField(max_length=200, null=True)
	cardholderName = models.CharField(max_length=200, null=True)
	billingAddress = Address
	billingProfile = models.ForeignKey('BillingProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	methodType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentMethodType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.last4
		str = str + self.cardholderName
		str = str + self.billingAddress
		str = str + self.methodType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PaymentMethod";
    
	def objectType(self):
		return "PaymentMethod";
