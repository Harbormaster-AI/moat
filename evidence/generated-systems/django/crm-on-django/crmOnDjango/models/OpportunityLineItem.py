from django.db import models

#======================================================================
# 
# Encapsulates data for model OpportunityLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityLineItem Declaration
#======================================================================
class OpportunityLineItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantity = models.CharField(max_length=64, null=True)
	unitPrice = Money
	discountPercent = models.CharField(max_length=64, null=True)
	totalPrice = Money
	opportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	product = models.ForeignKey('Product', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	priceBookEntry = models.ForeignKey('PriceBookEntry', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantity
		str = str + self.unitPrice
		str = str + self.discountPercent
		str = str + self.totalPrice
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "OpportunityLineItem";
    
	def objectType(self):
		return "OpportunityLineItem";
