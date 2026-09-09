from django.db import models
from ecommerceOnDjango.models.PaymentStatus import PaymentStatus
from ecommerceOnDjango.models.PaymentMethodType import PaymentMethodType

#======================================================================
# 
# Encapsulates data for model Payment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Payment Declaration
#======================================================================
class Payment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	paymentNumber = models.CharField(max_length=200, null=True)
	amount = Money
	transactionId = models.CharField(max_length=200, null=True)
	authorizedAt = models.DateField(null=True)
	capturedAt = models.DateField(null=True)
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	paymentProvider = models.ForeignKey('PaymentProvider', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	refunds = models.ManyToManyField('Refund',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentStatus])
	paymentMethod = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentMethodType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.paymentNumber
		str = str + self.amount
		str = str + self.transactionId
		str = str + self.authorizedAt
		str = str + self.capturedAt
		str = str + self.status
		str = str + self.paymentMethod
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Payment";
    
	def objectType(self):
		return "Payment";
