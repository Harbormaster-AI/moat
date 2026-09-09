from django.db import models
from inventoryOnDjango.models.ItemType import ItemType
from inventoryOnDjango.models.UnitOfMeasure import UnitOfMeasure

#======================================================================
# 
# Encapsulates data for model StockKeepingUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockKeepingUnit Declaration
#======================================================================
class StockKeepingUnit (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	skuCode = SKU
	name = models.CharField(max_length=200, null=True)
	weight = models.CharField(max_length=64, null=True)
	weightUnit = models.CharField(max_length=200, null=True)
	volume = models.CharField(max_length=64, null=True)
	volumeUnit = models.CharField(max_length=200, null=True)
	shelfLifeDays = models.IntegerField(null=True)
	hazardousMaterial = models.BooleanField(null=True)
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	uomConversions = models.ManyToManyField('UoMConversion',  blank=True, related_name='+')
	replenishmentPolicies = models.ManyToManyField('ReplenishmentPolicy',  blank=True, related_name='+')
	lots = models.ManyToManyField('Lot',  blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	itemType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ItemType])
	unitOfMeasure = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnitOfMeasure])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.skuCode
		str = str + self.name
		str = str + self.weight
		str = str + self.weightUnit
		str = str + self.volume
		str = str + self.volumeUnit
		str = str + self.shelfLifeDays
		str = str + self.hazardousMaterial
		str = str + self.itemType
		str = str + self.unitOfMeasure
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "StockKeepingUnit";
    
	def objectType(self):
		return "StockKeepingUnit";
