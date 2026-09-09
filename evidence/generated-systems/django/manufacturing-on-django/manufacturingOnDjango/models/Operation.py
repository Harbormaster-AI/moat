from django.db import models
from manufacturingOnDjango.models.OperationType import OperationType

#======================================================================
# 
# Encapsulates data for model Operation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Operation Declaration
#======================================================================
class Operation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	operationNumber = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	setupTime = TimeDuration
	standardCycleTime = TimeDuration
	routing = models.ForeignKey('Routing', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workCenter = models.ForeignKey('WorkCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inspectionPlan = models.ForeignKey('InspectionPlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	operationType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OperationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.operationNumber
		str = str + self.name
		str = str + self.setupTime
		str = str + self.standardCycleTime
		str = str + self.operationType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Operation";
    
	def objectType(self):
		return "Operation";
