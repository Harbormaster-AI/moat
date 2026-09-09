from django.db import models
from insuranceOnDjango.models.PaymentMethod import PaymentMethod
from insuranceOnDjango.models.PaymentStatus import PaymentStatus

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
	paymentReference = models.CharField(max_length=200, null=True)
	amount = Money
	paymentDate = models.DateField(null=True)
	invoice = models.ForeignKey('Invoice', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	billingAccount = models.ForeignKey('BillingAccount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	method = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentMethod])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.paymentReference
		str = str + self.amount
		str = str + self.paymentDate
		str = str + self.method
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Payment";
    
	def objectType(self):
		return "Payment";
