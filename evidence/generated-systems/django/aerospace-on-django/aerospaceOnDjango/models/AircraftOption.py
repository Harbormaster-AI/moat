from django.db import models
from aerospaceOnDjango.models.OptionCategory import OptionCategory

#======================================================================
# 
# Encapsulates data for model AircraftOption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOption Declaration
#======================================================================
class AircraftOption (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	variants = models.ManyToManyField('AircraftVariant',  blank=True, related_name='+')
	packages = models.ManyToManyField('AircraftPackage',  blank=True, related_name='+')
	optionCategory = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OptionCategory])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.name
		str = str + self.optionCategory
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AircraftOption";
    
	def objectType(self):
		return "AircraftOption";
