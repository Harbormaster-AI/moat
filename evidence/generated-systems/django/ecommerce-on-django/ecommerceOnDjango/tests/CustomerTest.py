import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

 #======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerTest Declaration
#======================================================================
class CustomerTest (TestCase) :
	def test_crud(self) :
		customer = Customer()
		customer.firstName = "default firstName field value"
		customer.lastName = "default lastName field value"
		customer.email = "default email field value"
		customer.phone = "default phone field value"
		customer.marketingOptIn = False
		customer.customerGroup = "default customerGroup field value"
		
		delegate = CustomerDelegate()
		responseObj = delegate.create(customer)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


