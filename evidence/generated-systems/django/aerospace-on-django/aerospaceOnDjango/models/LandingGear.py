from django.db import models
from aerospaceOnDjango.models.LandingGearType import LandingGearType

#======================================================================
# 
# Encapsulates data for model LandingGear
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LandingGear Declaration
#======================================================================
class LandingGear (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	supplierPartNumber = models.CharField(max_length=200, null=True)
	supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variants = models.ManyToManyField('AircraftVariant',  blank=True, related_name='+')
	gearType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LandingGearType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.supplierPartNumber
		str = str + self.gearType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LandingGear";
    
	def objectType(self):
		return "LandingGear";
