from django.db import models
from ecommerceOnDjango.models.Carrier import Carrier
from ecommerceOnDjango.models.ServiceLevel import ServiceLevel

#======================================================================
# 
# Encapsulates data for model CarrierService
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarrierService Declaration
#======================================================================
class CarrierService (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	code = models.CharField(max_length=200, null=True)
	shippingMethods = models.ManyToManyField('ShippingMethod',  blank=True, related_name='+')
	carrier = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Carrier])
	serviceLevel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ServiceLevel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.code
		str = str + self.carrier
		str = str + self.serviceLevel
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CarrierService";
    
	def objectType(self):
		return "CarrierService";
