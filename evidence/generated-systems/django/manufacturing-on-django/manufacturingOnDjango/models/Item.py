from django.db import models
from manufacturingOnDjango.models.ItemType import ItemType
from manufacturingOnDjango.models.ProcurementType import ProcurementType
from manufacturingOnDjango.models.UnitOfMeasure import UnitOfMeasure
from manufacturingOnDjango.models.ProductLifecycleStatus import ProductLifecycleStatus

#======================================================================
# 
# Encapsulates data for model Item
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Item Declaration
#======================================================================
class Item (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	itemNumber = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	standardCost = Money
	weight = Measurement
	asSerialControlled = models.BooleanField(null=True)
	businessUnit = models.ForeignKey('BusinessUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	boms = models.ManyToManyField('BOM',  blank=True, related_name='+')
	routings = models.ManyToManyField('Routing',  blank=True, related_name='+')
	suppliers = models.ManyToManyField('Supplier',  blank=True, related_name='+')
	qualitySpecifications = models.ManyToManyField('QualitySpecification',  blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	itemType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ItemType])
	procurementType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProcurementType])
	unitOfMeasure = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnitOfMeasure])
	lifecycleStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProductLifecycleStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.itemNumber
		str = str + self.name
		str = str + self.standardCost
		str = str + self.weight
		str = str + self.asSerialControlled
		str = str + self.itemType
		str = str + self.procurementType
		str = str + self.unitOfMeasure
		str = str + self.lifecycleStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Item";
    
	def objectType(self):
		return "Item";
