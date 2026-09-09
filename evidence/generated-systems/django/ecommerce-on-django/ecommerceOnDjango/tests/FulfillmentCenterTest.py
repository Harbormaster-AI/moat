import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.FulfillmentCenter import FulfillmentCenter
from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

 #======================================================================
# 
# Encapsulates data for model FulfillmentCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FulfillmentCenterTest Declaration
#======================================================================
class FulfillmentCenterTest (TestCase) :
	def test_crud(self) :
		fulfillmentCenter = FulfillmentCenter()
		fulfillmentCenter.name = "default name field value"
		fulfillmentCenter.centerCode = "default centerCode field value"
		fulfillmentCenter.address = "default address field value"
		fulfillmentCenter.timezone = "default timezone field value"
		fulfillmentCenter.asActive = False
		
		delegate = FulfillmentCenterDelegate()
		responseObj = delegate.create(fulfillmentCenter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


