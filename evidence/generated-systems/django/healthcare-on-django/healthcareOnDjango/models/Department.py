from django.db import models
from healthcareOnDjango.models.DepartmentType import DepartmentType

#======================================================================
# 
# Encapsulates data for model Department
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Department Declaration
#======================================================================
class Department (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	careTeams = models.ManyToManyField('CareTeam',  blank=True, related_name='+')
	departmentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DepartmentType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.departmentType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Department";
    
	def objectType(self):
		return "Department";
