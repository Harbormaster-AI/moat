import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.AudienceSegment import AudienceSegment
from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

 #======================================================================
# 
# Encapsulates data for model AudienceSegment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AudienceSegmentTest Declaration
#======================================================================
class AudienceSegmentTest (TestCase) :
	def test_crud(self) :
		audienceSegment = AudienceSegment()
		audienceSegment.name = "default name field value"
		audienceSegment.estimatedReach = 22
		audienceSegment.description = "default description field value"
		audienceSegment.providerType = "default providerType field value"
		
		delegate = AudienceSegmentDelegate()
		responseObj = delegate.create(audienceSegment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


