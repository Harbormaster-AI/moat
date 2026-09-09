from django.db import models
from advertisingOnDjango.models.ChannelType import ChannelType
from advertisingOnDjango.models.AdFormat import AdFormat

#======================================================================
# 
# Encapsulates data for model InventorySource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventorySource Declaration
#======================================================================
class InventorySource (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	domain = models.CharField(max_length=200, null=True)
	publisher = models.ForeignKey('Publisher', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	adSlots = models.ManyToManyField('AdSlot',  blank=True, related_name='+')
	deals = models.ManyToManyField('Deal',  blank=True, related_name='+')
	channel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ChannelType])
	primaryFormat = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdFormat])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.domain
		str = str + self.channel
		str = str + self.primaryFormat
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventorySource";
    
	def objectType(self):
		return "InventorySource";
