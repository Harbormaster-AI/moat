from django.db import models
from crmOnDjango.models.QuoteStatus import QuoteStatus

#======================================================================
# 
# Encapsulates data for model Quote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Quote Declaration
#======================================================================
class Quote (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quoteNumber = models.CharField(max_length=200, null=True)
	validityStart = models.DateField(null=True)
	validityEnd = models.DateField(null=True)
	totalAmount = Money
	discountPercent = models.CharField(max_length=64, null=True)
	taxAmount = Money
	shippingAmount = Money
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	opportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lineItems = models.ManyToManyField('QuoteLineItem',  blank=True, related_name='+')
	priceBook = models.ForeignKey('PriceBook', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	order = models.OneToOneField('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in QuoteStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quoteNumber
		str = str + self.validityStart
		str = str + self.validityEnd
		str = str + self.totalAmount
		str = str + self.discountPercent
		str = str + self.taxAmount
		str = str + self.shippingAmount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Quote";
    
	def objectType(self):
		return "Quote";
