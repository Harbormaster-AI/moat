from django.db import models
from analyticsOnDjango.models.InferenceMode import InferenceMode

#======================================================================
# 
# Encapsulates data for model InferenceEndpoint
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InferenceEndpoint Declaration
#======================================================================
class InferenceEndpoint (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	endpointUrl = models.CharField(max_length=200, null=True)
	trafficShare = Percentage
	modelVersion = models.ForeignKey('ModelVersion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	predictions = models.ManyToManyField('Prediction',  blank=True, related_name='+')
	mode = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InferenceMode])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.endpointUrl
		str = str + self.trafficShare
		str = str + self.mode
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InferenceEndpoint";
    
	def objectType(self):
		return "InferenceEndpoint";
