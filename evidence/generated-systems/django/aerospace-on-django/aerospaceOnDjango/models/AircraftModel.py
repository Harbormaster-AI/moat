from django.db import models
from aerospaceOnDjango.models.AircraftType import AircraftType

#======================================================================
# 
# Encapsulates data for model AircraftModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftModel Declaration
#======================================================================
class AircraftModel (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	modelDesignation = models.CharField(max_length=200, null=True)
	family = models.ForeignKey('AircraftFamily', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variants = models.ManyToManyField('AircraftVariant',  blank=True, related_name='+')
	engineTypes = models.ManyToManyField('EngineType',  blank=True, related_name='+')
	aircraftType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AircraftType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.modelDesignation
		str = str + self.aircraftType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AircraftModel";
    
	def objectType(self):
		return "AircraftModel";
