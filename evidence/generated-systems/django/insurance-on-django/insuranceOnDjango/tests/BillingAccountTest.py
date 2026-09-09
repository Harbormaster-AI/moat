import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.BillingAccount import BillingAccount
from insuranceOnDjango.delegates.BillingAccountDelegate import BillingAccountDelegate

 #======================================================================
# 
# Encapsulates data for model BillingAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingAccountTest Declaration
#======================================================================
class BillingAccountTest (TestCase) :
	def test_crud(self) :
		billingAccount = BillingAccount()
		billingAccount.accountNumber = "default accountNumber field value"
		billingAccount.balance = "default balance field value"
		billingAccount.status = "default status field value"
		
		delegate = BillingAccountDelegate()
		responseObj = delegate.create(billingAccount)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


