from django.db import models

#======================================================================
# 
# Encapsulates data for model PriceBook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBook Declaration
#======================================================================
class PriceBook (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	asActive = models.BooleanField(null=True)
	description = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	entries = models.ManyToManyField('PriceBookEntry',  blank=True, related_name='+')
	quotes = models.ManyToManyField('Quote',  blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.asActive
		str = str + self.description
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PriceBook";
    
	def objectType(self):
		return "PriceBook";
