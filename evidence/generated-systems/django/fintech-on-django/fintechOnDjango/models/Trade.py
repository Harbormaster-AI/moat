from django.db import models

#======================================================================
# 
# Encapsulates data for model Trade
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Trade Declaration
#======================================================================
class Trade (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	executedAt = DateTime
	quantity = models.CharField(max_length=64, null=True)
	price = Money
	fees = Money
	settlementDate = models.DateField(null=True)
	order = models.ForeignKey('TradeOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	security = models.ForeignKey('Security', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	investmentAccount = models.ForeignKey('InvestmentAccount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.executedAt
		str = str + self.quantity
		str = str + self.price
		str = str + self.fees
		str = str + self.settlementDate
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Trade";
    
	def objectType(self):
		return "Trade";
