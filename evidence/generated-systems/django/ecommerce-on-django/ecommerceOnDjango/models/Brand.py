from django.db import models

#======================================================================
# 
# Encapsulates data for model Brand
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Brand Declaration
#======================================================================
class Brand (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	products = models.ManyToManyField('Product',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.description
		str = str + self.website
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Brand";
    
	def objectType(self):
		return "Brand";
