import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.TrackingPixel import TrackingPixel
from advertisingOnDjango.delegates.TrackingPixelDelegate import TrackingPixelDelegate

 #======================================================================
# 
# Encapsulates data for model TrackingPixel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrackingPixelTest Declaration
#======================================================================
class TrackingPixelTest (TestCase) :
	def test_crud(self) :
		trackingPixel = TrackingPixel()
		trackingPixel.name = "default name field value"
		trackingPixel.url = "default url field value"
		trackingPixel.eventType = "default eventType field value"
		trackingPixel.pixelType = "default pixelType field value"
		
		delegate = TrackingPixelDelegate()
		responseObj = delegate.create(trackingPixel)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


