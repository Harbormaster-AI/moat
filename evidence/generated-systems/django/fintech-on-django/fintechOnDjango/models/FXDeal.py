from django.db import models
from fintechOnDjango.models.FXDealStatus import FXDealStatus

#======================================================================
# 
# Encapsulates data for model FXDeal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXDeal Declaration
#======================================================================
class FXDeal (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	dealReference = models.CharField(max_length=200, null=True)
	baseCurrency = models.CharField(max_length=200, null=True)
	quoteCurrency = models.CharField(max_length=200, null=True)
	rate = models.CharField(max_length=64, null=True)
	amount = Money
	settlementDate = models.DateField(null=True)
	quote = models.ForeignKey('FXQuote', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	paymentOrders = models.ManyToManyField('PaymentOrder',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FXDealStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.dealReference
		str = str + self.baseCurrency
		str = str + self.quoteCurrency
		str = str + self.rate
		str = str + self.amount
		str = str + self.settlementDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FXDeal";
    
	def objectType(self):
		return "FXDeal";
