from django.db import models
from manufacturingOnDjango.models.WarehouseType import WarehouseType

#======================================================================
# 
# Encapsulates data for model Warehouse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Warehouse Declaration
#======================================================================
class Warehouse (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	warehouseCode = models.CharField(max_length=200, null=True)
	address = Address
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	locations = models.ManyToManyField('Location',  blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	warehouseType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WarehouseType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.warehouseCode
		str = str + self.address
		str = str + self.warehouseType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Warehouse";
    
	def objectType(self):
		return "Warehouse";
