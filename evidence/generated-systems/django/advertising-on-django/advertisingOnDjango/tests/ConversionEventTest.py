import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.ConversionEvent import ConversionEvent
from advertisingOnDjango.delegates.ConversionEventDelegate import ConversionEventDelegate

 #======================================================================
# 
# Encapsulates data for model ConversionEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConversionEventTest Declaration
#======================================================================
class ConversionEventTest (TestCase) :
	def test_crud(self) :
		conversionEvent = ConversionEvent()
		conversionEvent.timestamp = "default timestamp field value"
		conversionEvent.value = "default value field value"
		conversionEvent.eventType = "default eventType field value"
		conversionEvent.attributionModel = "default attributionModel field value"
		
		delegate = ConversionEventDelegate()
		responseObj = delegate.create(conversionEvent)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


