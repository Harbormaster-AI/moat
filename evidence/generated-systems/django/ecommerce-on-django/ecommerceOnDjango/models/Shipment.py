from django.db import models
from ecommerceOnDjango.models.ShipmentStatus import ShipmentStatus
from ecommerceOnDjango.models.Carrier import Carrier

#======================================================================
# 
# Encapsulates data for model Shipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Shipment Declaration
#======================================================================
class Shipment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	shipmentNumber = models.CharField(max_length=200, null=True)
	shippedDate = models.DateField(null=True)
	deliveredDate = models.DateField(null=True)
	trackingNumber = models.CharField(max_length=200, null=True)
	shippingAddress = Address
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	shipmentItems = models.ManyToManyField('ShipmentItem',  blank=True, related_name='+')
	fulfillmentCenter = models.ForeignKey('FulfillmentCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ShipmentStatus])
	carrier = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Carrier])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.shipmentNumber
		str = str + self.shippedDate
		str = str + self.deliveredDate
		str = str + self.trackingNumber
		str = str + self.shippingAddress
		str = str + self.status
		str = str + self.carrier
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Shipment";
    
	def objectType(self):
		return "Shipment";
