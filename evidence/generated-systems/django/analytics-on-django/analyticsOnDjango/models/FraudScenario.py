from django.db import models
from analyticsOnDjango.models.FraudDetectionType import FraudDetectionType

#======================================================================
# 
# Encapsulates data for model FraudScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudScenario Declaration
#======================================================================
class FraudScenario (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	riskAppetite = models.CharField(max_length=200, null=True)
	models = models.ManyToManyField('Model',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	alerts = models.ManyToManyField('Alert',  blank=True, related_name='+')
	signals = models.ManyToManyField('FraudSignal',  blank=True, related_name='+')
	detectionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FraudDetectionType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.riskAppetite
		str = str + self.detectionType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FraudScenario";
    
	def objectType(self):
		return "FraudScenario";
