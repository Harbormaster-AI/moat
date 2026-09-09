from django.db import models
from governanceOnDjango.models.DataClassificationLevel import DataClassificationLevel

#======================================================================
# 
# Encapsulates data for model DataCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataCategory Declaration
#======================================================================
class DataCategory (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	processingActivities = models.ManyToManyField('DataProcessingActivity',  blank=True, related_name='+')
	records = models.ManyToManyField('Record_',  blank=True, related_name='+')
	dataBreaches = models.ManyToManyField('DataBreach',  blank=True, related_name='+')
	classification = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataClassificationLevel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.description
		str = str + self.classification
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataCategory";
    
	def objectType(self):
		return "DataCategory";
