from django.db import models

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
	code = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	manager = models.OneToOneField('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	positions = models.ManyToManyField('Position',  blank=True, related_name='+')
	employees = models.ManyToManyField('Employee',  blank=True, related_name='+')
	costCenter = models.ForeignKey('CostCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.code
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Department";
    
	def objectType(self):
		return "Department";
