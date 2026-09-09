from django.db import models

#======================================================================
# 
# Encapsulates data for model Plant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Plant Declaration
#======================================================================
class Plant (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	plantCode = models.CharField(max_length=200, null=True)
	address = Address
	manufacturer = models.ForeignKey('AerospaceManufacturer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	productionLines = models.ManyToManyField('ProductionLine',  blank=True, related_name='+')
	warehouses = models.ManyToManyField('Warehouse',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.plantCode
		str = str + self.address
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Plant";
    
	def objectType(self):
		return "Plant";
