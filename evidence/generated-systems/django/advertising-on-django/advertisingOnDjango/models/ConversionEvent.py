from django.db import models
from advertisingOnDjango.models.ConversionEventType import ConversionEventType
from advertisingOnDjango.models.AttributionModel import AttributionModel

#======================================================================
# 
# Encapsulates data for model ConversionEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConversionEvent Declaration
#======================================================================
class ConversionEvent (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	timestamp = models.CharField(max_length=64, null=True)
	value = Money
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lineItem = models.ForeignKey('LineItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	trackingPixel = models.ForeignKey('TrackingPixel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	eventType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConversionEventType])
	attributionModel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AttributionModel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.timestamp
		str = str + self.value
		str = str + self.eventType
		str = str + self.attributionModel
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ConversionEvent";
    
	def objectType(self):
		return "ConversionEvent";
