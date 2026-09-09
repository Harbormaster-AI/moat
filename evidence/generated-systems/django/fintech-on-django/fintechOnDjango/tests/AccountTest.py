import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Account import Account
from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

 #======================================================================
# 
# Encapsulates data for model Account
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountTest Declaration
#======================================================================
class AccountTest (TestCase) :
	def test_crud(self) :
		account = Account()
		account.accountNumber = "default accountNumber field value"
		account.iban = "default iban field value"
		account.bic = "default bic field value"
		account.openedDate = datetime.datetime.now()
		account.currency = "default currency field value"
		account.balance = "default balance field value"
		account.availableBalance = "default availableBalance field value"
		account.accountType = "default accountType field value"
		account.status = "default status field value"
		
		delegate = AccountDelegate()
		responseObj = delegate.create(account)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


