from django.db import models
from ecommerceOnDjango.models.SupplierStatus import SupplierStatus

#======================================================================
# 
# Encapsulates data for model Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Supplier Declaration
#======================================================================
class Supplier (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	contactEmail = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	products = models.ManyToManyField('Product',  blank=True, related_name='+')
	fulfillmentCenters = models.ManyToManyField('FulfillmentCenter',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SupplierStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.contactEmail
		str = str + self.website
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Supplier";
    
	def objectType(self):
		return "Supplier";
