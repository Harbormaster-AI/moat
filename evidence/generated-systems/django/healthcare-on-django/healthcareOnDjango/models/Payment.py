from django.db import models
from healthcareOnDjango.models.PaymentMethod import PaymentMethod

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
	paymentDate = models.DateField(null=True)
	invoice = models.ForeignKey('Invoice', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payer = models.ForeignKey('InsurancePayer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	method = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentMethod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.paymentNumber
		str = str + self.amount
		str = str + self.paymentDate
		str = str + self.method
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Payment";
    
	def objectType(self):
		return "Payment";
