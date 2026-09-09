from django.db import models
from inventoryOnDjango.models.UnitOfMeasure import UnitOfMeasure
from inventoryOnDjango.models.StockStatus import StockStatus

#======================================================================
# 
# Encapsulates data for model TransferOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderLine Declaration
#======================================================================
class TransferOrderLine (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	lineNumber = models.IntegerField(null=True)
	quantity = models.CharField(max_length=64, null=True)
	transferOrder = models.ForeignKey('TransferOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	fromLocation = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	toLocation = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	unitOfMeasure = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnitOfMeasure])
	stockStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in StockStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.lineNumber
		str = str + self.quantity
		str = str + self.unitOfMeasure
		str = str + self.stockStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TransferOrderLine";
    
	def objectType(self):
		return "TransferOrderLine";
