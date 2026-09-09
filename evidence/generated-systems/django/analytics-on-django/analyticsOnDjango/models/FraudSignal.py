from django.db import models
from analyticsOnDjango.models.FraudSignalType import FraudSignalType

#======================================================================
# 
# Encapsulates data for model FraudSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudSignal Declaration
#======================================================================
class FraudSignal (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	ruleLogic = models.CharField(max_length=200, null=True)
	scenario = models.ForeignKey('FraudScenario', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataset = models.ForeignKey('DataSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	modelVersion = models.ForeignKey('ModelVersion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	signalType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FraudSignalType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.ruleLogic
		str = str + self.signalType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FraudSignal";
    
	def objectType(self):
		return "FraudSignal";
