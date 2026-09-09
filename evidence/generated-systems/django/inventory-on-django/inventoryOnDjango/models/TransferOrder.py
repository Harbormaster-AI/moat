from django.db import models
from inventoryOnDjango.models.TransferOrderStatus import TransferOrderStatus

#======================================================================
# 
# Encapsulates data for model TransferOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrder Declaration
#======================================================================
class TransferOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderNumber = models.CharField(max_length=200, null=True)
	requestedShipDate = models.DateField(null=True)
	requestedReceiveDate = models.DateField(null=True)
	shippedDate = models.DateField(null=True)
	receivedDate = models.DateField(null=True)
	originWarehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	destinationWarehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lines = models.ManyToManyField('TransferOrderLine',  blank=True, related_name='+')
	transactions = models.ManyToManyField('InventoryTransaction',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TransferOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.requestedShipDate
		str = str + self.requestedReceiveDate
		str = str + self.shippedDate
		str = str + self.receivedDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TransferOrder";
    
	def objectType(self):
		return "TransferOrder";
