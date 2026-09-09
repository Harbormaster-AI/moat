from django.db import models
from analyticsOnDjango.models.ExperimentStatus import ExperimentStatus

#======================================================================
# 
# Encapsulates data for model Experiment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Experiment Declaration
#======================================================================
class Experiment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	objective = models.CharField(max_length=200, null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	trainingRuns = models.ManyToManyField('TrainingRun',  blank=True, related_name='+')
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	notebooks = models.ManyToManyField('Notebook',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ExperimentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.objective
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Experiment";
    
	def objectType(self):
		return "Experiment";
