import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Publisher import Publisher
from advertisingOnDjango.delegates.PublisherDelegate import PublisherDelegate

 #======================================================================
# 
# Encapsulates data for model Publisher
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PublisherTest Declaration
#======================================================================
class PublisherTest (TestCase) :
	def test_crud(self) :
		publisher = Publisher()
		publisher.name = "default name field value"
		publisher.website = "default website field value"
		publisher.publisherType = "default publisherType field value"
		
		delegate = PublisherDelegate()
		responseObj = delegate.create(publisher)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


