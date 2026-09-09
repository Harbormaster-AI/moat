from django.db import models
from fintechOnDjango.models.ProductCategory import ProductCategory

#======================================================================
# 
# Encapsulates data for model ProductOffering
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductOffering Declaration
#======================================================================
class ProductOffering (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	productCode = models.CharField(max_length=200, null=True)
	institution = models.ForeignKey('FinancialInstitution', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	pricingPlans = models.ManyToManyField('PricingPlan',  blank=True, related_name='+')
	category = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProductCategory])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.productCode
		str = str + self.category
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProductOffering";
    
	def objectType(self):
		return "ProductOffering";
