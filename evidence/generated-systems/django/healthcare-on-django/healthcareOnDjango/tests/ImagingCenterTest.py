import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.ImagingCenter import ImagingCenter
from healthcareOnDjango.delegates.ImagingCenterDelegate import ImagingCenterDelegate

 #======================================================================
# 
# Encapsulates data for model ImagingCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingCenterTest Declaration
#======================================================================
class ImagingCenterTest (TestCase) :
	def test_crud(self) :
		imagingCenter = ImagingCenter()
		imagingCenter.name = "default name field value"
		
		delegate = ImagingCenterDelegate()
		responseObj = delegate.create(imagingCenter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


