from django.db import models
from analyticsOnDjango.models.TagCategory import TagCategory

#======================================================================
# 
# Encapsulates data for model Tag
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Tag Declaration
#======================================================================
class Tag (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	modelVersions = models.ManyToManyField('ModelVersion',  blank=True, related_name='+')
	dashboards = models.ManyToManyField('Dashboard',  blank=True, related_name='+')
	reports = models.ManyToManyField('Report',  blank=True, related_name='+')
	featureSets = models.ManyToManyField('FeatureSet',  blank=True, related_name='+')
	metrics = models.ManyToManyField('Metric',  blank=True, related_name='+')
	category = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TagCategory])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.category
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Tag";
    
	def objectType(self):
		return "Tag";
