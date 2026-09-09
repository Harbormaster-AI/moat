from django.db import models
from analyticsOnDjango.models.FeatureStoreType import FeatureStoreType

#======================================================================
# 
# Encapsulates data for model FeatureSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureSet Declaration
#======================================================================
class FeatureSet (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	refreshSchedule = CronSchedule
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	features = models.ManyToManyField('Feature',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	modelVersions = models.ManyToManyField('ModelVersion',  blank=True, related_name='+')
	tags = models.ManyToManyField('Tag',  blank=True, related_name='+')
	storeType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FeatureStoreType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.refreshSchedule
		str = str + self.storeType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FeatureSet";
    
	def objectType(self):
		return "FeatureSet";
