from django.db import models
from inventoryOnDjango.models.AllocationStatus import AllocationStatus

#======================================================================
# 
# Encapsulates data for model OutboundAllocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OutboundAllocation Declaration
#======================================================================
class OutboundAllocation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	allocationNumber = models.CharField(max_length=200, null=True)
	allocatedQuantity = models.CharField(max_length=64, null=True)
	allocationDate = models.DateField(null=True)
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inventoryItem = models.ForeignKey('InventoryItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reservation = models.ForeignKey('Reservation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	sourceLocation = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AllocationStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.allocationNumber
		str = str + self.allocatedQuantity
		str = str + self.allocationDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "OutboundAllocation";
    
	def objectType(self):
		return "OutboundAllocation";
