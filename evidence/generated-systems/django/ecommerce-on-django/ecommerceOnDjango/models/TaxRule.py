from django.db import models
from ecommerceOnDjango.models.TaxClass import TaxClass

#======================================================================
# 
# Encapsulates data for model TaxRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxRule Declaration
#======================================================================
class TaxRule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	country = models.CharField(max_length=200, null=True)
	region = models.CharField(max_length=200, null=True)
	rate = Percentage
	taxInclusive = models.BooleanField(null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	channels = models.ManyToManyField('Channel',  blank=True, related_name='+')
	taxClass = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TaxClass])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.country
		str = str + self.region
		str = str + self.rate
		str = str + self.taxInclusive
		str = str + self.taxClass
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TaxRule";
    
	def objectType(self):
		return "TaxRule";
