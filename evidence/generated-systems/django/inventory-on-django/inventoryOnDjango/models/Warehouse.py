from django.db import models

#======================================================================
# 
# Encapsulates data for model Warehouse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Warehouse Declaration
#======================================================================
class Warehouse (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	code = models.CharField(max_length=200, null=True)
	address = Address
	timeZone = models.CharField(max_length=200, null=True)
	allowsOverAllocation = models.BooleanField(null=True)
	storageLocations = models.ManyToManyField('StorageLocation',  blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	inboundShipments = models.ManyToManyField('InboundShipment',  blank=True, related_name='+')
	outboundAllocations = models.ManyToManyField('OutboundAllocation',  blank=True, related_name='+')
	originTransfers = models.ManyToManyField('TransferOrder',  blank=True, related_name='+')
	destinationTransfers = models.ManyToManyField('TransferOrder',  blank=True, related_name='+')
	cycleCounts = models.ManyToManyField('CycleCount',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.code
		str = str + self.address
		str = str + self.timeZone
		str = str + self.allowsOverAllocation
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Warehouse";
    
	def objectType(self):
		return "Warehouse";
