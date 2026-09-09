import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Subscription import Subscription
from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

 #======================================================================
# 
# Encapsulates data for model Subscription
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriptionTest Declaration
#======================================================================
class SubscriptionTest (TestCase) :
	def test_crud(self) :
		subscription = Subscription()
		subscription.subscriptionNumber = "default subscriptionNumber field value"
		subscription.nextBillingDate = datetime.datetime.now()
		subscription.startDate = datetime.datetime.now()
		subscription.endDate = datetime.datetime.now()
		subscription.status = "default status field value"
		subscription.interval = "default interval field value"
		
		delegate = SubscriptionDelegate()
		responseObj = delegate.create(subscription)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


