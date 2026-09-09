from django.db import models
from manufacturingOnDjango.models.InventoryTransactionType import InventoryTransactionType

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
	quantity = Quantity
	transactionDateTime = models.CharField(max_length=64, null=True)
	referenceDocument = models.CharField(max_length=200, null=True)
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('Location', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workOrder = models.ForeignKey('WorkOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	purchaseOrder = models.ForeignKey('PurchaseOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	salesOrder = models.ForeignKey('SalesOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	transactionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InventoryTransactionType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.transactionNumber
		str = str + self.quantity
		str = str + self.transactionDateTime
		str = str + self.referenceDocument
		str = str + self.transactionType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryTransaction";
    
	def objectType(self):
		return "InventoryTransaction";
