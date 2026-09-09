from django.db import models
from analyticsOnDjango.models.RecommendationType import RecommendationType

#======================================================================
# 
# Encapsulates data for model RecommendationScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecommendationScenario Declaration
#======================================================================
class RecommendationScenario (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	objective = models.CharField(max_length=200, null=True)
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	experiments = models.ManyToManyField('Experiment',  blank=True, related_name='+')
	alerts = models.ManyToManyField('Alert',  blank=True, related_name='+')
	recommendationType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RecommendationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.objective
		str = str + self.recommendationType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RecommendationScenario";
    
	def objectType(self):
		return "RecommendationScenario";
