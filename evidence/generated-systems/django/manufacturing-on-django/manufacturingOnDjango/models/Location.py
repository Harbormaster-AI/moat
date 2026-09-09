from django.db import models
from manufacturingOnDjango.models.LocationType import LocationType

#======================================================================
# 
# Encapsulates data for model Location
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Location Declaration
#======================================================================
class Location (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	locationCode = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	locationType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LocationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.locationCode
		str = str + self.description
		str = str + self.locationType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Location";
    
	def objectType(self):
		return "Location";
