from django.db import models
from fintechOnDjango.models.FeeType import FeeType
from fintechOnDjango.models.FeeCalculationMethod import FeeCalculationMethod

#======================================================================
# 
# Encapsulates data for model FeeSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeeSchedule Declaration
#======================================================================
class FeeSchedule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	amount = Money
	percentage = models.CharField(max_length=64, null=True)
	minimum = Money
	maximum = Money
	pricingPlan = models.ForeignKey('PricingPlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	feeType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FeeType])
	calculationMethod = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FeeCalculationMethod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.amount
		str = str + self.percentage
		str = str + self.minimum
		str = str + self.maximum
		str = str + self.feeType
		str = str + self.calculationMethod
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FeeSchedule";
    
	def objectType(self):
		return "FeeSchedule";
