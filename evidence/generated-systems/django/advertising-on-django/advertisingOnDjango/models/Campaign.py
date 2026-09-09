from django.db import models
from advertisingOnDjango.models.ObjectiveType import ObjectiveType
from advertisingOnDjango.models.CampaignStatus import CampaignStatus

#======================================================================
# 
# Encapsulates data for model Campaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Campaign Declaration
#======================================================================
class Campaign (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	totalBudget = Money
	flight = DateRange
	adAccount = models.ForeignKey('AdAccount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lineItems = models.ManyToManyField('LineItem',  blank=True, related_name='+')
	kpis = models.ManyToManyField('KPI',  blank=True, related_name='+')
	trackingPixels = models.ManyToManyField('TrackingPixel',  blank=True, related_name='+')
	audiences = models.ManyToManyField('AudienceSegment',  blank=True, related_name='+')
	reports = models.ManyToManyField('Report',  blank=True, related_name='+')
	insertionOrder = models.ForeignKey('InsertionOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	objective = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ObjectiveType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CampaignStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.totalBudget
		str = str + self.flight
		str = str + self.objective
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Campaign";
    
	def objectType(self):
		return "Campaign";
