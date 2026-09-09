import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.DeviceCriterion import DeviceCriterion
from advertisingOnDjango.delegates.DeviceCriterionDelegate import DeviceCriterionDelegate

 #======================================================================
# 
# Encapsulates data for model DeviceCriterion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DeviceCriterionTest Declaration
#======================================================================
class DeviceCriterionTest (TestCase) :
	def test_crud(self) :
		deviceCriterion = DeviceCriterion()
		deviceCriterion.deviceType = "default deviceType field value"
		deviceCriterion.platformType = "default platformType field value"
		deviceCriterion.operator = "default operator field value"
		
		delegate = DeviceCriterionDelegate()
		responseObj = delegate.create(deviceCriterion)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


