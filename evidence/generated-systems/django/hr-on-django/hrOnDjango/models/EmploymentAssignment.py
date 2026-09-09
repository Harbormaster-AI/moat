from django.db import models
from hrOnDjango.models.AssignmentType import AssignmentType
from hrOnDjango.models.AssignmentStatus import AssignmentStatus

#======================================================================
# 
# Encapsulates data for model EmploymentAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentAssignment Declaration
#======================================================================
class EmploymentAssignment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	primary = models.BooleanField(null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	position = models.ForeignKey('Position', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	supervisor = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	assignmentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AssignmentType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AssignmentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.primary
		str = str + self.assignmentType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "EmploymentAssignment";
    
	def objectType(self):
		return "EmploymentAssignment";
