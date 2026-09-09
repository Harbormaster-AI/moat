from django.db import models

#======================================================================
# 
# Encapsulates data for model CostCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CostCenter Declaration
#======================================================================
class CostCenter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	departments = models.ManyToManyField('Department',  blank=True, related_name='+')
	positions = models.ManyToManyField('Position',  blank=True, related_name='+')
	employees = models.ManyToManyField('Employee',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.name
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CostCenter";
    
	def objectType(self):
		return "CostCenter";
