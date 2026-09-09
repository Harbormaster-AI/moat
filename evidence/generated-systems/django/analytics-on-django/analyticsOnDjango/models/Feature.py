from django.db import models
from analyticsOnDjango.models.DataType import DataType

#======================================================================
# 
# Encapsulates data for model Feature
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Feature Declaration
#======================================================================
class Feature (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	featureSet = models.ForeignKey('FeatureSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	sourceDatasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	trainingRuns = models.ManyToManyField('TrainingRun',  blank=True, related_name='+')
	dataType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.description
		str = str + self.dataType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Feature";
    
	def objectType(self):
		return "Feature";
