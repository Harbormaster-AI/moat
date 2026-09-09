import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.ImagingOrder import ImagingOrder
from healthcareOnDjango.delegates.ImagingOrderDelegate import ImagingOrderDelegate

 #======================================================================
# 
# Encapsulates data for model ImagingOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingOrderTest Declaration
#======================================================================
class ImagingOrderTest (TestCase) :
	def test_crud(self) :
		imagingOrder = ImagingOrder()
		imagingOrder.bodySite = "default bodySite field value"
		imagingOrder.contrast = False
		imagingOrder.modality = "default modality field value"
		
		delegate = ImagingOrderDelegate()
		responseObj = delegate.create(imagingOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


