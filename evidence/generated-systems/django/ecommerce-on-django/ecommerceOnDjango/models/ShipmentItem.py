from django.db import models

#======================================================================
# 
# Encapsulates data for model ShipmentItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentItem Declaration
#======================================================================
class ShipmentItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantity = models.IntegerField(null=True)
	shipment = models.ForeignKey('Shipment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	orderLine = models.ForeignKey('OrderLine', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantity
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ShipmentItem";
    
	def objectType(self):
		return "ShipmentItem";
