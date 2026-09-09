import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.BillingProfile import BillingProfile
from advertisingOnDjango.delegates.BillingProfileDelegate import BillingProfileDelegate

 #======================================================================
# 
# Encapsulates data for model BillingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingProfileTest Declaration
#======================================================================
class BillingProfileTest (TestCase) :
	def test_crud(self) :
		billingProfile = BillingProfile()
		billingProfile.billingName = "default billingName field value"
		billingProfile.taxId = "default taxId field value"
		billingProfile.billingAddress = "default billingAddress field value"
		billingProfile.paymentTerms = "default paymentTerms field value"
		
		delegate = BillingProfileDelegate()
		responseObj = delegate.create(billingProfile)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


