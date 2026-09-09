from django.db import models
from ecommerceOnDjango.models.ShippingMethodType import ShippingMethodType

#======================================================================
# 
# Encapsulates data for model ShippingMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShippingMethod Declaration
#======================================================================
class ShippingMethod (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	flatRate = Money
	estimatedDays = models.IntegerField(null=True)
	asActive = models.BooleanField(null=True)
	channels = models.ManyToManyField('Channel',  blank=True, related_name='+')
	carrierService = models.ForeignKey('CarrierService', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	methodType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ShippingMethodType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.flatRate
		str = str + self.estimatedDays
		str = str + self.asActive
		str = str + self.methodType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ShippingMethod";
    
	def objectType(self):
		return "ShippingMethod";
