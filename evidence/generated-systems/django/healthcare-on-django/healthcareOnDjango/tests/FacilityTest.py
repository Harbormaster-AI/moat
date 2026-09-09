import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

 #======================================================================
# 
# Encapsulates data for model Facility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FacilityTest Declaration
#======================================================================
class FacilityTest (TestCase) :
	def test_crud(self) :
		facility = Facility()
		facility.name = "default name field value"
		facility.facilityCode = "default facilityCode field value"
		facility.address = "default address field value"
		facility.facilityType = "default facilityType field value"
		
		delegate = FacilityDelegate()
		responseObj = delegate.create(facility)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


