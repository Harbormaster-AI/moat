import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Customer import Customer
from manufacturingOnDjango.delegates.CustomerDelegate import CustomerDelegate

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
		customer.name = "default name field value"
		customer.customerCode = "default customerCode field value"
		customer.address = "default address field value"
		customer.customerType = "default customerType field value"
		
		delegate = CustomerDelegate()
		responseObj = delegate.create(customer)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


