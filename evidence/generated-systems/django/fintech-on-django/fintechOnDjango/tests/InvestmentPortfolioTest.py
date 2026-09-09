import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.InvestmentPortfolio import InvestmentPortfolio
from fintechOnDjango.delegates.InvestmentPortfolioDelegate import InvestmentPortfolioDelegate

 #======================================================================
# 
# Encapsulates data for model InvestmentPortfolio
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentPortfolioTest Declaration
#======================================================================
class InvestmentPortfolioTest (TestCase) :
	def test_crud(self) :
		investmentPortfolio = InvestmentPortfolio()
		investmentPortfolio.portfolioCode = "default portfolioCode field value"
		investmentPortfolio.baseCurrency = "default baseCurrency field value"
		investmentPortfolio.createdAt = "default createdAt field value"
		investmentPortfolio.status = "default status field value"
		
		delegate = InvestmentPortfolioDelegate()
		responseObj = delegate.create(investmentPortfolio)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


