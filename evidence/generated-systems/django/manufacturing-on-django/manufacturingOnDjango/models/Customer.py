from django.db import models
from manufacturingOnDjango.models.CustomerType import CustomerType

#======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Customer Declaration
#======================================================================
class Customer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	customerCode = models.CharField(max_length=200, null=True)
	address = Address
	enterprises = models.ManyToManyField('Enterprise',  blank=True, related_name='+')
	salesOrders = models.ManyToManyField('SalesOrder',  blank=True, related_name='+')
	customerType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CustomerType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.customerCode
		str = str + self.address
		str = str + self.customerType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Customer";
    
	def objectType(self):
		return "Customer";
