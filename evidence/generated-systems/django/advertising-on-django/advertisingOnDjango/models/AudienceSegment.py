from django.db import models
from advertisingOnDjango.models.DataProviderType import DataProviderType

#======================================================================
# 
# Encapsulates data for model AudienceSegment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AudienceSegment Declaration
#======================================================================
class AudienceSegment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	estimatedReach = models.IntegerField(null=True)
	description = models.CharField(max_length=200, null=True)
	provider = models.ForeignKey('DataProvider', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	providerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataProviderType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.estimatedReach
		str = str + self.description
		str = str + self.providerType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AudienceSegment";
    
	def objectType(self):
		return "AudienceSegment";
