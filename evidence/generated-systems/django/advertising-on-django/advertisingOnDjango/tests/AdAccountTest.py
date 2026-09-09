import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

 #======================================================================
# 
# Encapsulates data for model AdAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdAccountTest Declaration
#======================================================================
class AdAccountTest (TestCase) :
	def test_crud(self) :
		adAccount = AdAccount()
		adAccount.name = "default name field value"
		adAccount.accountCode = "default accountCode field value"
		adAccount.defaultCurrency = "default defaultCurrency field value"
		adAccount.defaultTimezone = "default defaultTimezone field value"
		
		delegate = AdAccountDelegate()
		responseObj = delegate.create(adAccount)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


