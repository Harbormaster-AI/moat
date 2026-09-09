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
	quantityOnHand = Quantity
	quantityReserved = Quantity
	lotNumber = LotId
	serialNumber = SerialId
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('Location', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantityOnHand
		str = str + self.quantityReserved
		str = str + self.lotNumber
		str = str + self.serialNumber
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryItem";
    
	def objectType(self):
		return "InventoryItem";
