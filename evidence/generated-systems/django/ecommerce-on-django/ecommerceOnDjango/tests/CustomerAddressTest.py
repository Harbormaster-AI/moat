import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.CustomerAddress import CustomerAddress
from ecommerceOnDjango.delegates.CustomerAddressDelegate import CustomerAddressDelegate

 #======================================================================
# 
# Encapsulates data for model CustomerAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerAddressTest Declaration
#======================================================================
class CustomerAddressTest (TestCase) :
	def test_crud(self) :
		customerAddress = CustomerAddress()
		customerAddress.label = "default label field value"
		customerAddress.address = "default address field value"
		customerAddress.asDefaultShipping = False
		customerAddress.asDefaultBilling = False
		
		delegate = CustomerAddressDelegate()
		responseObj = delegate.create(customerAddress)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


