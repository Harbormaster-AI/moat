from django.db import models
from advertisingOnDjango.models.ExperimentStatus import ExperimentStatus

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
	hypothesis = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variants = models.ManyToManyField('ExperimentVariant',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ExperimentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.hypothesis
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Experiment";
    
	def objectType(self):
		return "Experiment";
