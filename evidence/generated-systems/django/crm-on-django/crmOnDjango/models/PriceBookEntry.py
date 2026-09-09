from django.db import models

#======================================================================
# 
# Encapsulates data for model PriceBookEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBookEntry Declaration
#======================================================================
class PriceBookEntry (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	unitPrice = Money
	effectiveDate = models.DateField(null=True)
	expirationDate = models.DateField(null=True)
	asActive = models.BooleanField(null=True)
	priceBook = models.ForeignKey('PriceBook', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	product = models.ForeignKey('Product', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.unitPrice
		str = str + self.effectiveDate
		str = str + self.expirationDate
		str = str + self.asActive
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PriceBookEntry";
    
	def objectType(self):
		return "PriceBookEntry";
