from django.db import models
from inventoryOnDjango.models.AdjustmentType import AdjustmentType
from inventoryOnDjango.models.AdjustmentStatus import AdjustmentStatus

#======================================================================
# 
# Encapsulates data for model StockAdjustment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockAdjustment Declaration
#======================================================================
class StockAdjustment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	adjustmentNumber = models.CharField(max_length=200, null=True)
	reason = models.CharField(max_length=200, null=True)
	adjustmentDate = models.DateField(null=True)
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lines = models.ManyToManyField('StockAdjustmentLine',  blank=True, related_name='+')
	transactions = models.ManyToManyField('InventoryTransaction',  blank=True, related_name='+')
	adjustmentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdjustmentType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdjustmentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.adjustmentNumber
		str = str + self.reason
		str = str + self.adjustmentDate
		str = str + self.adjustmentType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "StockAdjustment";
    
	def objectType(self):
		return "StockAdjustment";
