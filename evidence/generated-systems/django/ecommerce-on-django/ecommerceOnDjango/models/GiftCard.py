from django.db import models
from ecommerceOnDjango.models.GiftCardStatus import GiftCardStatus

#======================================================================
# 
# Encapsulates data for model GiftCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCard Declaration
#======================================================================
class GiftCard (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	balance = Money
	expirationDate = models.DateField(null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	issuedOrder = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	redemptions = models.ManyToManyField('GiftCardRedemption',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in GiftCardStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.balance
		str = str + self.expirationDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "GiftCard";
    
	def objectType(self):
		return "GiftCard";
