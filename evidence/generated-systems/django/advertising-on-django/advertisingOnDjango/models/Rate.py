from django.db import models
from advertisingOnDjango.models.AdFormat import AdFormat
from advertisingOnDjango.models.PricingModel import PricingModel

#======================================================================
# 
# Encapsulates data for model Rate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Rate Declaration
#======================================================================
class Rate (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	unitPrice = Money
	rateCard = models.ForeignKey('RateCard', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	adSlot = models.ForeignKey('AdSlot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	adFormat = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdFormat])
	pricingModel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PricingModel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.unitPrice
		str = str + self.adFormat
		str = str + self.pricingModel
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Rate";
    
	def objectType(self):
		return "Rate";
