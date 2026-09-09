from django.db import models
from inventoryOnDjango.models.InboundShipmentStatus import InboundShipmentStatus

#======================================================================
# 
# Encapsulates data for model InboundShipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipment Declaration
#======================================================================
class InboundShipment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	shipmentNumber = models.CharField(max_length=200, null=True)
	expectedArrivalDate = models.DateField(null=True)
	arrivalDate = models.DateField(null=True)
	carrierName = models.CharField(max_length=200, null=True)
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lines = models.ManyToManyField('InboundShipmentLine',  blank=True, related_name='+')
	transactions = models.ManyToManyField('InventoryTransaction',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InboundShipmentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.shipmentNumber
		str = str + self.expectedArrivalDate
		str = str + self.arrivalDate
		str = str + self.carrierName
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InboundShipment";
    
	def objectType(self):
		return "InboundShipment";
