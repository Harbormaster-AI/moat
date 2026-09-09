from django.db import models
from inventoryOnDjango.models.TransactionType import TransactionType
from inventoryOnDjango.models.UnitOfMeasure import UnitOfMeasure
from inventoryOnDjango.models.TransactionStatus import TransactionStatus

#======================================================================
# 
# Encapsulates data for model InventoryTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryTransaction Declaration
#======================================================================
class InventoryTransaction (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	transactionNumber = models.CharField(max_length=200, null=True)
	quantity = models.CharField(max_length=64, null=True)
	unitCost = Money
	transactionDate = models.DateField(null=True)
	reasonCode = models.CharField(max_length=200, null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	relatedReservation = models.ForeignKey('Reservation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transferOrder = models.ForeignKey('TransferOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	adjustment = models.ForeignKey('StockAdjustment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	cycleCount = models.ForeignKey('CycleCount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transactionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TransactionType])
	unitOfMeasure = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnitOfMeasure])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TransactionStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.transactionNumber
		str = str + self.quantity
		str = str + self.unitCost
		str = str + self.transactionDate
		str = str + self.reasonCode
		str = str + self.transactionType
		str = str + self.unitOfMeasure
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryTransaction";
    
	def objectType(self):
		return "InventoryTransaction";
