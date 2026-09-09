from django.db import models

#======================================================================
# 
# Encapsulates data for model ProductPricing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductPricing Declaration
#======================================================================
class ProductPricing (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	listPrice = Money
	salePrice = Money
	validFrom = models.DateField(null=True)
	validTo = models.DateField(null=True)
	variant = models.ForeignKey('ProductVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	channel = models.ForeignKey('Channel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.listPrice
		str = str + self.salePrice
		str = str + self.validFrom
		str = str + self.validTo
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProductPricing";
    
	def objectType(self):
		return "ProductPricing";
