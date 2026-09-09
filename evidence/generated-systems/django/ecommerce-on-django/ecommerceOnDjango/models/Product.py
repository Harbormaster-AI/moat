from django.db import models
from ecommerceOnDjango.models.ProductType import ProductType
from ecommerceOnDjango.models.TaxClass import TaxClass

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
	name = models.CharField(max_length=200, null=True)
	slug = models.CharField(max_length=200, null=True)
	asActive = models.BooleanField(null=True)
	brand = models.ForeignKey('Brand', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	categories = models.ManyToManyField('Category',  blank=True, related_name='+')
	variants = models.ManyToManyField('ProductVariant',  blank=True, related_name='+')
	mediaAssets = models.ManyToManyField('MediaAsset',  blank=True, related_name='+')
	reviews = models.ManyToManyField('Review',  blank=True, related_name='+')
	seller = models.ForeignKey('Seller', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	productType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProductType])
	defaultTaxClass = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TaxClass])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.slug
		str = str + self.asActive
		str = str + self.productType
		str = str + self.defaultTaxClass
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Product";
    
	def objectType(self):
		return "Product";
