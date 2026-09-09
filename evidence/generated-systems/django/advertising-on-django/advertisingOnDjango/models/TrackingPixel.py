from django.db import models
from advertisingOnDjango.models.ConversionEventType import ConversionEventType
from advertisingOnDjango.models.PixelType import PixelType

#======================================================================
# 
# Encapsulates data for model TrackingPixel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrackingPixel Declaration
#======================================================================
class TrackingPixel (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	url = URL
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	advertiser = models.ForeignKey('Advertiser', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	conversionEvents = models.ManyToManyField('ConversionEvent',  blank=True, related_name='+')
	eventType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ConversionEventType])
	pixelType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PixelType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.url
		str = str + self.eventType
		str = str + self.pixelType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TrackingPixel";
    
	def objectType(self):
		return "TrackingPixel";
