from django.db import models

#======================================================================
# 
# Encapsulates data for model InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItem Declaration
#======================================================================
class InventoryItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	sku = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	quantityOnHand = models.IntegerField(null=True)
	quantityReserved = models.IntegerField(null=True)
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	supplier = models.ForeignKey('MedicalSupplier', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.sku
		str = str + self.name
		str = str + self.quantityOnHand
		str = str + self.quantityReserved
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryItem";
    
	def objectType(self):
		return "InventoryItem";
