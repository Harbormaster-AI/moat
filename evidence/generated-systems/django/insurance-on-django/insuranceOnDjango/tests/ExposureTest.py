import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Exposure import Exposure
from insuranceOnDjango.delegates.ExposureDelegate import ExposureDelegate

 #======================================================================
# 
# Encapsulates data for model Exposure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExposureTest Declaration
#======================================================================
class ExposureTest (TestCase) :
	def test_crud(self) :
		exposure = Exposure()
		exposure.exposureType = "default exposureType field value"
		exposure.status = "default status field value"
		
		delegate = ExposureDelegate()
		responseObj = delegate.create(exposure)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


