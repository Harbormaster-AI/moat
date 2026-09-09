from django.db import models
from inventoryOnDjango.models.LotStatus import LotStatus

#======================================================================
# 
# Encapsulates data for model Lot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Lot Declaration
#======================================================================
class Lot (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	batchNumber = BatchNumber
	manufactureDate = models.DateField(null=True)
	expirationDate = models.DateField(null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	lotStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LotStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.batchNumber
		str = str + self.manufactureDate
		str = str + self.expirationDate
		str = str + self.lotStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Lot";
    
	def objectType(self):
		return "Lot";
