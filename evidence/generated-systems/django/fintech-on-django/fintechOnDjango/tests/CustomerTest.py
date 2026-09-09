import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

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
		customer.dateOfBirth = datetime.datetime.now()
		customer.email = "default email field value"
		customer.phone = "default phone field value"
		customer.address = "default address field value"
		customer.taxId = "default taxId field value"
		customer.riskScore = "default riskScore field value"
		customer.customerType = "default customerType field value"
		
		delegate = CustomerDelegate()
		responseObj = delegate.create(customer)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


