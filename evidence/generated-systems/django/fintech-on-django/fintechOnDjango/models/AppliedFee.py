from django.db import models
from fintechOnDjango.models.FeeType import FeeType

#======================================================================
# 
# Encapsulates data for model AppliedFee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppliedFee Declaration
#======================================================================
class AppliedFee (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	amount = Money
	description = models.CharField(max_length=200, null=True)
	paymentOrder = models.ForeignKey('PaymentOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transaction = models.ForeignKey('Transaction', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	feeType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FeeType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.amount
		str = str + self.description
		str = str + self.feeType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AppliedFee";
    
	def objectType(self):
		return "AppliedFee";
