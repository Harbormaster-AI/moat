from django.db import models
from fintechOnDjango.models.OrderSide import OrderSide
from fintechOnDjango.models.OrderType import OrderType
from fintechOnDjango.models.OrderStatus import OrderStatus
from fintechOnDjango.models.TimeInForce import TimeInForce

#======================================================================
# 
# Encapsulates data for model TradeOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TradeOrder Declaration
#======================================================================
class TradeOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderId = models.CharField(max_length=200, null=True)
	quantity = models.CharField(max_length=64, null=True)
	limitPrice = Money
	placedAt = DateTime
	portfolio = models.ForeignKey('InvestmentPortfolio', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	security = models.ForeignKey('Security', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	trades = models.ManyToManyField('Trade',  blank=True, related_name='+')
	side = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OrderSide])
	type = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OrderType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OrderStatus])
	timeInForce = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TimeInForce])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderId
		str = str + self.quantity
		str = str + self.limitPrice
		str = str + self.placedAt
		str = str + self.side
		str = str + self.type
		str = str + self.status
		str = str + self.timeInForce
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TradeOrder";
    
	def objectType(self):
		return "TradeOrder";
