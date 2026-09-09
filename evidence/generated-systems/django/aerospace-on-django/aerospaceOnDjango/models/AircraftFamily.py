from django.db import models

#======================================================================
# 
# Encapsulates data for model AircraftFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftFamily Declaration
#======================================================================
class AircraftFamily (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	familyCode = models.CharField(max_length=200, null=True)
	program = models.ForeignKey('AircraftProgram', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	aircraftModels = models.ManyToManyField('AircraftModel',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.familyCode
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AircraftFamily";
    
	def objectType(self):
		return "AircraftFamily";
