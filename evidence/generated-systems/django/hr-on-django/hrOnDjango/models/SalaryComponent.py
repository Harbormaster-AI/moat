from django.db import models
from hrOnDjango.models.SalaryComponentType import SalaryComponentType

#======================================================================
# 
# Encapsulates data for model SalaryComponent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalaryComponent Declaration
#======================================================================
class SalaryComponent (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	amount = Money
	recurring = models.BooleanField(null=True)
	compensationPackage = models.ForeignKey('CompensationPackage', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	componentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SalaryComponentType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.amount
		str = str + self.recurring
		str = str + self.componentType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SalaryComponent";
    
	def objectType(self):
		return "SalaryComponent";
