from django.db import models
from crmOnDjango.models.ProductType import ProductType
from crmOnDjango.models.UnitOfMeasure import UnitOfMeasure

#======================================================================
# 
# Encapsulates data for model Product
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Product Declaration
#======================================================================
class Product (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	sku = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	asActive = models.BooleanField(null=True)
	standardPrice = Money
	description = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	priceBookEntries = models.ManyToManyField('PriceBookEntry',  blank=True, related_name='+')
	opportunityLineItems = models.ManyToManyField('OpportunityLineItem',  blank=True, related_name='+')
	quoteLineItems = models.ManyToManyField('QuoteLineItem',  blank=True, related_name='+')
	orderItems = models.ManyToManyField('OrderItem',  blank=True, related_name='+')
	productType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProductType])
	uom = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnitOfMeasure])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.sku
		str = str + self.name
		str = str + self.asActive
		str = str + self.standardPrice
		str = str + self.description
		str = str + self.productType
		str = str + self.uom
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Product";
    
	def objectType(self):
		return "Product";
