from django.db import models
from advertisingOnDjango.models.PublisherType import PublisherType

#======================================================================
# 
# Encapsulates data for model Publisher
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Publisher Declaration
#======================================================================
class Publisher (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	inventorySources = models.ManyToManyField('InventorySource',  blank=True, related_name='+')
	deals = models.ManyToManyField('Deal',  blank=True, related_name='+')
	creativeApprovals = models.ManyToManyField('CreativeApproval',  blank=True, related_name='+')
	insertionOrders = models.ManyToManyField('InsertionOrder',  blank=True, related_name='+')
	rateCards = models.ManyToManyField('RateCard',  blank=True, related_name='+')
	publisherType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PublisherType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.website
		str = str + self.publisherType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Publisher";
    
	def objectType(self):
		return "Publisher";
