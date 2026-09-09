from django.db import models
from aerospaceOnDjango.models.ProductionOrderStatus import ProductionOrderStatus

#======================================================================
# 
# Encapsulates data for model ProductionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionOrder Declaration
#======================================================================
class ProductionOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderNumber = models.CharField(max_length=200, null=True)
	variant = models.ForeignKey('AircraftVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	aircraftOrder = models.ForeignKey('AircraftOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProductionOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProductionOrder";
    
	def objectType(self):
		return "ProductionOrder";
