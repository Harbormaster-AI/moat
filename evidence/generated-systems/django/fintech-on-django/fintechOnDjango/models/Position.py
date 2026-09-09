from django.db import models

#======================================================================
# 
# Encapsulates data for model Position
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Position Declaration
#======================================================================
class Position (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantity = models.CharField(max_length=64, null=True)
	averageCost = Money
	marketValue = Money
	portfolio = models.ForeignKey('InvestmentPortfolio', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	security = models.ForeignKey('Security', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantity
		str = str + self.averageCost
		str = str + self.marketValue
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Position";
    
	def objectType(self):
		return "Position";
