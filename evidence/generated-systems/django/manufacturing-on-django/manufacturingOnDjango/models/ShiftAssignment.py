from django.db import models

#======================================================================
# 
# Encapsulates data for model ShiftAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftAssignment Declaration
#======================================================================
class ShiftAssignment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	assignmentDate = models.DateField(null=True)
	shift = models.ForeignKey('Shift', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workCenter = models.ForeignKey('WorkCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.assignmentDate
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ShiftAssignment";
    
	def objectType(self):
		return "ShiftAssignment";
