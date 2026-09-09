import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Subscriber import Subscriber
from analyticsOnDjango.delegates.SubscriberDelegate import SubscriberDelegate

 #======================================================================
# 
# Encapsulates data for model Subscriber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriberTest Declaration
#======================================================================
class SubscriberTest (TestCase) :
	def test_crud(self) :
		subscriber = Subscriber()
		subscriber.name = "default name field value"
		subscriber.address = "default address field value"
		subscriber.channel = "default channel field value"
		
		delegate = SubscriberDelegate()
		responseObj = delegate.create(subscriber)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


