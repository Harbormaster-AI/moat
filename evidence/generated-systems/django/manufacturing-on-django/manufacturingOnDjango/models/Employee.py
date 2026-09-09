from django.db import models
from manufacturingOnDjango.models.EmployeeRole import EmployeeRole
from manufacturingOnDjango.models.SkillLevel import SkillLevel

#======================================================================
# 
# Encapsulates data for model Employee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Employee Declaration
#======================================================================
class Employee (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	workCenter = models.ForeignKey('WorkCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	shiftAssignments = models.ManyToManyField('ShiftAssignment',  blank=True, related_name='+')
	correctiveActions = models.ManyToManyField('CorrectiveAction',  blank=True, related_name='+')
	role = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EmployeeRole])
	skillLevel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SkillLevel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.role
		str = str + self.skillLevel
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Employee";
    
	def objectType(self):
		return "Employee";
