from django.db import models
from inventoryOnDjango.models.LocationType import LocationType

#======================================================================
# 
# Encapsulates data for model StorageLocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StorageLocation Declaration
#======================================================================
class StorageLocation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	temperatureControlled = models.BooleanField(null=True)
	capacity = models.CharField(max_length=64, null=True)
	capacityUnit = models.CharField(max_length=200, null=True)
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	parentLocation = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	childLocations = models.ManyToManyField('StorageLocation',  blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	locationType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LocationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.temperatureControlled
		str = str + self.capacity
		str = str + self.capacityUnit
		str = str + self.locationType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "StorageLocation";
    
	def objectType(self):
		return "StorageLocation";
