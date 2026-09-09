from django.db import models
from aerospaceOnDjango.models.PackageType import PackageType

#======================================================================
# 
# Encapsulates data for model AircraftPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftPackage Declaration
#======================================================================
class AircraftPackage (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	options = models.ManyToManyField('AircraftOption',  blank=True, related_name='+')
	variants = models.ManyToManyField('AircraftVariant',  blank=True, related_name='+')
	packageType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PackageType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.packageType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AircraftPackage";
    
	def objectType(self):
		return "AircraftPackage";
