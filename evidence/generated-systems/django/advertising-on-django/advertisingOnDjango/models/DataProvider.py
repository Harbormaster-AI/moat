from django.db import models
from advertisingOnDjango.models.DataProviderType import DataProviderType

#======================================================================
# 
# Encapsulates data for model DataProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProvider Declaration
#======================================================================
class DataProvider (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	audienceSegments = models.ManyToManyField('AudienceSegment',  blank=True, related_name='+')
	providerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataProviderType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.website
		str = str + self.providerType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataProvider";
    
	def objectType(self):
		return "DataProvider";
