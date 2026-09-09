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
	quantityOnHand = models.IntegerField(null=True)
	quantityReserved = models.IntegerField(null=True)
	lotNumber = models.CharField(max_length=200, null=True)
	component = models.ForeignKey('Component_', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantityOnHand
		str = str + self.quantityReserved
		str = str + self.lotNumber
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryItem";
    
	def objectType(self):
		return "InventoryItem";
