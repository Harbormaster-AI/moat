from django.db import models
from hrOnDjango.models.PositionStatus import PositionStatus
from hrOnDjango.models.WorkLocationType import WorkLocationType

#======================================================================
# 
# Encapsulates data for model Position
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Position Declaration
#======================================================================
class Position (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	positionCode = models.CharField(max_length=200, null=True)
	fte = models.CharField(max_length=64, null=True)
	department = models.ForeignKey('Department', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	jobProfile = models.ForeignKey('JobProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	costCenter = models.ForeignKey('CostCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('Location', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	managerPosition = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	directReports = models.ManyToManyField('Position',  blank=True, related_name='+')
	assignments = models.ManyToManyField('EmploymentAssignment',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PositionStatus])
	workLocationType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WorkLocationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.positionCode
		str = str + self.fte
		str = str + self.status
		str = str + self.workLocationType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Position";
    
	def objectType(self):
		return "Position";
