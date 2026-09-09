from django.db import models
from manufacturingOnDjango.models.NonconformanceType import NonconformanceType
from manufacturingOnDjango.models.QualitySeverity import QualitySeverity
from manufacturingOnDjango.models.NonconformanceStatus import NonconformanceStatus

#======================================================================
# 
# Encapsulates data for model Nonconformance
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Nonconformance Declaration
#======================================================================
class Nonconformance (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	ncNumber = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	containmentAction = models.CharField(max_length=200, null=True)
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workOrder = models.ForeignKey('WorkOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inspectionLot = models.ForeignKey('InspectionLot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	correctiveAction = models.OneToOneField('CorrectiveAction', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	ncType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in NonconformanceType])
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in QualitySeverity])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in NonconformanceStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.ncNumber
		str = str + self.description
		str = str + self.containmentAction
		str = str + self.ncType
		str = str + self.severity
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Nonconformance";
    
	def objectType(self):
		return "Nonconformance";
