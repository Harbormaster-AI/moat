from django.db import models
from fintechOnDjango.models.PortfolioStatus import PortfolioStatus

#======================================================================
# 
# Encapsulates data for model InvestmentPortfolio
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentPortfolio Declaration
#======================================================================
class InvestmentPortfolio (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	portfolioCode = models.CharField(max_length=200, null=True)
	baseCurrency = models.CharField(max_length=200, null=True)
	createdAt = DateTime
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	accounts = models.ManyToManyField('InvestmentAccount',  blank=True, related_name='+')
	orders = models.ManyToManyField('TradeOrder',  blank=True, related_name='+')
	holdings = models.ManyToManyField('Position',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PortfolioStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.portfolioCode
		str = str + self.baseCurrency
		str = str + self.createdAt
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InvestmentPortfolio";
    
	def objectType(self):
		return "InvestmentPortfolio";
