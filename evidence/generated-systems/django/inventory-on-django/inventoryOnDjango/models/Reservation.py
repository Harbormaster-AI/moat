from django.db import models
from inventoryOnDjango.models.ReservationStatus import ReservationStatus
from inventoryOnDjango.models.ReservationType import ReservationType

#======================================================================
# 
# Encapsulates data for model Reservation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Reservation Declaration
#======================================================================
class Reservation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	referenceNumber = models.CharField(max_length=200, null=True)
	reservedQuantity = models.CharField(max_length=64, null=True)
	promisedDate = models.DateField(null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inventoryItem = models.ForeignKey('InventoryItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	demandSignal = models.ForeignKey('DemandSignal', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reservationStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReservationStatus])
	reservationType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReservationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.referenceNumber
		str = str + self.reservedQuantity
		str = str + self.promisedDate
		str = str + self.reservationStatus
		str = str + self.reservationType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Reservation";
    
	def objectType(self):
		return "Reservation";
