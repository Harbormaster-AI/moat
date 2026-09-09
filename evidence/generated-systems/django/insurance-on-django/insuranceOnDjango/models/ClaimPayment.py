from django.db import models
from insuranceOnDjango.models.PayeeType import PayeeType
from insuranceOnDjango.models.PaymentMethod import PaymentMethod
from insuranceOnDjango.models.PaymentStatus import PaymentStatus

#======================================================================
# 
# Encapsulates data for model ClaimPayment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimPayment Declaration
#======================================================================
class ClaimPayment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	paymentNumber = models.CharField(max_length=200, null=True)
	amount = Money
	paymentDate = models.DateField(null=True)
	claim = models.ForeignKey('Claim', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	exposure = models.ForeignKey('Exposure', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	beneficiary = models.ForeignKey('Beneficiary', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serviceProvider = models.ForeignKey('ServiceProvider', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	payeeType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PayeeType])
	method = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentMethod])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PaymentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.paymentNumber
		str = str + self.amount
		str = str + self.paymentDate
		str = str + self.payeeType
		str = str + self.method
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ClaimPayment";
    
	def objectType(self):
		return "ClaimPayment";
