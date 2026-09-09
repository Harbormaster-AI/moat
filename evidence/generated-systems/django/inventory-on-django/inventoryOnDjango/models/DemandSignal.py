from django.db import models
from inventoryOnDjango.models.DemandType import DemandType

#======================================================================
# 
# Encapsulates data for model DemandSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DemandSignal Declaration
#======================================================================
class DemandSignal (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	externalReference = models.CharField(max_length=200, null=True)
	requestedDate = models.DateField(null=True)
	quantity = models.CharField(max_length=64, null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reservations = models.ManyToManyField('Reservation',  blank=True, related_name='+')
	demandType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DemandType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.externalReference
		str = str + self.requestedDate
		str = str + self.quantity
		str = str + self.demandType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DemandSignal";
    
	def objectType(self):
		return "DemandSignal";
