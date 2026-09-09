import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.InvestmentAccount import InvestmentAccount
from fintechOnDjango.delegates.InvestmentAccountDelegate import InvestmentAccountDelegate

 #======================================================================
# 
# Encapsulates data for model InvestmentAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentAccountTest Declaration
#======================================================================
class InvestmentAccountTest (TestCase) :
	def test_crud(self) :
		investmentAccount = InvestmentAccount()
		investmentAccount.accountNumber = "default accountNumber field value"
		investmentAccount.baseCurrency = "default baseCurrency field value"
		investmentAccount.balance = "default balance field value"
		investmentAccount.accountType = "default accountType field value"
		investmentAccount.status = "default status field value"
		
		delegate = InvestmentAccountDelegate()
		responseObj = delegate.create(investmentAccount)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


