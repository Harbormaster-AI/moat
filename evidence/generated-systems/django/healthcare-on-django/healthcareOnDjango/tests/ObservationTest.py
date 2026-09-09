import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Observation import Observation
from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

 #======================================================================
# 
# Encapsulates data for model Observation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObservationTest Declaration
#======================================================================
class ObservationTest (TestCase) :
	def test_crud(self) :
		observation = Observation()
		observation.code = "default code field value"
		observation.value = "default value field value"
		observation.unit = "default unit field value"
		observation.effectiveDateTime = "default effectiveDateTime field value"
		observation.interpretation = "default interpretation field value"
		
		delegate = ObservationDelegate()
		responseObj = delegate.create(observation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


