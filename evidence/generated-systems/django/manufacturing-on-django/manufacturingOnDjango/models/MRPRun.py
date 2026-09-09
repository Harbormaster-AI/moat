from django.db import models
from manufacturingOnDjango.models.MRPRunStatus import MRPRunStatus

#======================================================================
# 
# Encapsulates data for model MRPRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MRPRun Declaration
#======================================================================
class MRPRun (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	runNumber = models.CharField(max_length=200, null=True)
	runDateTime = models.CharField(max_length=64, null=True)
	planningHorizonDays = models.IntegerField(null=True)
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plannedOrders = models.ManyToManyField('PlannedOrder',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MRPRunStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.runNumber
		str = str + self.runDateTime
		str = str + self.planningHorizonDays
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MRPRun";
    
	def objectType(self):
		return "MRPRun";
