from django.db import models
from analyticsOnDjango.models.ModelType import ModelType

#======================================================================
# 
# Encapsulates data for model Model
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Model Declaration
#======================================================================
class Model (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	taskDescription = models.CharField(max_length=200, null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	versions = models.ManyToManyField('ModelVersion',  blank=True, related_name='+')
	featureSets = models.ManyToManyField('FeatureSet',  blank=True, related_name='+')
	experiments = models.ManyToManyField('Experiment',  blank=True, related_name='+')
	tags = models.ManyToManyField('Tag',  blank=True, related_name='+')
	modelType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ModelType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.taskDescription
		str = str + self.modelType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Model";
    
	def objectType(self):
		return "Model";
