import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Account import Account
from crmOnDjango.delegates.AccountDelegate import AccountDelegate

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
		account.name = "default name field value"
		account.accountNumber = "default accountNumber field value"
		account.industry = "default industry field value"
		account.billingAddress = "default billingAddress field value"
		account.shippingAddress = "default shippingAddress field value"
		account.website = "default website field value"
		account.phone = "default phone field value"
		account.asActive = False
		account.accountType = "default accountType field value"
		account.lifecycleStage = "default lifecycleStage field value"
		
		delegate = AccountDelegate()
		responseObj = delegate.create(account)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


