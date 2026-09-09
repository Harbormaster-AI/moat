from django.db import models

#======================================================================
# 
# Encapsulates data for model QuoteLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteLineItem Declaration
#======================================================================
class QuoteLineItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantity = models.CharField(max_length=64, null=True)
	unitPrice = Money
	discountAmount = Money
	taxAmount = Money
	totalAmount = Money
	quote = models.ForeignKey('Quote', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	product = models.ForeignKey('Product', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	priceBookEntry = models.ForeignKey('PriceBookEntry', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	opportunityLineItem = models.ForeignKey('OpportunityLineItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantity
		str = str + self.unitPrice
		str = str + self.discountAmount
		str = str + self.taxAmount
		str = str + self.totalAmount
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "QuoteLineItem";
    
	def objectType(self):
		return "QuoteLineItem";
