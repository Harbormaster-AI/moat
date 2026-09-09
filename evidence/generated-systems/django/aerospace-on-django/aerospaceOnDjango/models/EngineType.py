from django.db import models
from aerospaceOnDjango.models.EngineCategory import EngineCategory

#======================================================================
# 
# Encapsulates data for model EngineType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EngineType Declaration
#======================================================================
class EngineType (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	engineModelCode = models.CharField(max_length=200, null=True)
	maxThrustKn = models.CharField(max_length=64, null=True)
	supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	compatibleModels = models.ManyToManyField('AircraftModel',  blank=True, related_name='+')
	category = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EngineCategory])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.engineModelCode
		str = str + self.maxThrustKn
		str = str + self.category
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "EngineType";
    
	def objectType(self):
		return "EngineType";
