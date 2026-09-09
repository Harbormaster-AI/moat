import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.ShippingMethod import ShippingMethod
from ecommerceOnDjango.delegates.ShippingMethodDelegate import ShippingMethodDelegate

 #======================================================================
# 
# Encapsulates data for model ShippingMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShippingMethodTest Declaration
#======================================================================
class ShippingMethodTest (TestCase) :
	def test_crud(self) :
		shippingMethod = ShippingMethod()
		shippingMethod.name = "default name field value"
		shippingMethod.flatRate = "default flatRate field value"
		shippingMethod.estimatedDays = 22
		shippingMethod.asActive = False
		shippingMethod.methodType = "default methodType field value"
		
		delegate = ShippingMethodDelegate()
		responseObj = delegate.create(shippingMethod)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


