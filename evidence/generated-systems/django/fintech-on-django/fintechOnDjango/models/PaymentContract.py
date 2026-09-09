from django.db import models
from fintechOnDjango.models.ContractStatus import ContractStatus

#======================================================================
# 
# Encapsulates data for model PaymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentContract Declaration
#======================================================================
class PaymentContract (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	contractNumber = models.CharField(max_length=200, null=True)
	pricingPlanCode = models.CharField(max_length=200, null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	acquirer = models.ForeignKey('PaymentProcessor', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ContractStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.contractNumber
		str = str + self.pricingPlanCode
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PaymentContract";
    
	def objectType(self):
		return "PaymentContract";
