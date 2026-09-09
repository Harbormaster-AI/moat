from django.db import models
from fintechOnDjango.models.PlanStatus import PlanStatus

#======================================================================
# 
# Encapsulates data for model PricingPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PricingPlan Declaration
#======================================================================
class PricingPlan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	planCode = models.CharField(max_length=200, null=True)
	baseCurrency = models.CharField(max_length=200, null=True)
	productOffering = models.ForeignKey('ProductOffering', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	feeSchedules = models.ManyToManyField('FeeSchedule',  blank=True, related_name='+')
	limits = models.ManyToManyField('UsageLimit',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PlanStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.planCode
		str = str + self.baseCurrency
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PricingPlan";
    
	def objectType(self):
		return "PricingPlan";
