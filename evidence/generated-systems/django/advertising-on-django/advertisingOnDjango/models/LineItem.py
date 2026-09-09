from django.db import models
from advertisingOnDjango.models.LineItemStatus import LineItemStatus
from advertisingOnDjango.models.PricingModel import PricingModel
from advertisingOnDjango.models.BidStrategyType import BidStrategyType
from advertisingOnDjango.models.PacingType import PacingType

#======================================================================
# 
# Encapsulates data for model LineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineItem Declaration
#======================================================================
class LineItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	bidAmount = Money
	dailyBudget = Money
	frequencyCap = FrequencyCap
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	placements = models.ManyToManyField('Placement',  blank=True, related_name='+')
	targetingProfile = models.OneToOneField('TargetingProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	deal = models.ForeignKey('Deal', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	creatives = models.ManyToManyField('CreativeAsset',  blank=True, related_name='+')
	performanceMetrics = models.ManyToManyField('PerformanceMetric',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LineItemStatus])
	pricingModel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PricingModel])
	bidStrategy = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BidStrategyType])
	pacing = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PacingType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.bidAmount
		str = str + self.dailyBudget
		str = str + self.frequencyCap
		str = str + self.status
		str = str + self.pricingModel
		str = str + self.bidStrategy
		str = str + self.pacing
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LineItem";
    
	def objectType(self):
		return "LineItem";
