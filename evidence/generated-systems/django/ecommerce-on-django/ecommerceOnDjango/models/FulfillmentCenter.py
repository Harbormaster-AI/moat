from django.db import models

#======================================================================
# 
# Encapsulates data for model FulfillmentCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FulfillmentCenter Declaration
#======================================================================
class FulfillmentCenter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	centerCode = models.CharField(max_length=200, null=True)
	address = Address
	timezone = models.CharField(max_length=200, null=True)
	asActive = models.BooleanField(null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	shipments = models.ManyToManyField('Shipment',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.centerCode
		str = str + self.address
		str = str + self.timezone
		str = str + self.asActive
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FulfillmentCenter";
    
	def objectType(self):
		return "FulfillmentCenter";
