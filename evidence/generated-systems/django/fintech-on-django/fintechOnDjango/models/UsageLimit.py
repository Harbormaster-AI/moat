from django.db import models
from fintechOnDjango.models.LimitScope import LimitScope
from fintechOnDjango.models.LimitPeriod import LimitPeriod

#======================================================================
# 
# Encapsulates data for model UsageLimit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UsageLimit Declaration
#======================================================================
class UsageLimit (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	amount = Money
	count = models.IntegerField(null=True)
	pricingPlan = models.ForeignKey('PricingPlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	scope = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LimitScope])
	period = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LimitPeriod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.amount
		str = str + self.count
		str = str + self.scope
		str = str + self.period
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "UsageLimit";
    
	def objectType(self):
		return "UsageLimit";
