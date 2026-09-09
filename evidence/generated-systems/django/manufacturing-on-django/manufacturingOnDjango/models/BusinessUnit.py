from django.db import models
from manufacturingOnDjango.models.BusinessUnitCategory import BusinessUnitCategory

#======================================================================
# 
# Encapsulates data for model BusinessUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessUnit Declaration
#======================================================================
class BusinessUnit (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	code = models.CharField(max_length=200, null=True)
	enterprise = models.ForeignKey('Enterprise', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	items = models.ManyToManyField('Item',  blank=True, related_name='+')
	plants = models.ManyToManyField('Plant',  blank=True, related_name='+')
	category = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BusinessUnitCategory])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.code
		str = str + self.category
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BusinessUnit";
    
	def objectType(self):
		return "BusinessUnit";
