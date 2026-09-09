from django.db import models
from inventoryOnDjango.models.StockStatus import StockStatus

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
	quantityOnHand = models.CharField(max_length=64, null=True)
	quantityAvailable = models.CharField(max_length=64, null=True)
	quantityReserved = models.CharField(max_length=64, null=True)
	unitCost = Money
	lastUpdated = models.DateField(null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	transactions = models.ManyToManyField('InventoryTransaction',  blank=True, related_name='+')
	reservations = models.ManyToManyField('Reservation',  blank=True, related_name='+')
	stockStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in StockStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantityOnHand
		str = str + self.quantityAvailable
		str = str + self.quantityReserved
		str = str + self.unitCost
		str = str + self.lastUpdated
		str = str + self.stockStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryItem";
    
	def objectType(self):
		return "InventoryItem";
