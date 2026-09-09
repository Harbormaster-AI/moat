import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AircraftFamily import AircraftFamily
from aerospaceOnDjango.delegates.AircraftFamilyDelegate import AircraftFamilyDelegate

 #======================================================================
# 
# Encapsulates data for model AircraftFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftFamilyTest Declaration
#======================================================================
class AircraftFamilyTest (TestCase) :
	def test_crud(self) :
		aircraftFamily = AircraftFamily()
		aircraftFamily.name = "default name field value"
		aircraftFamily.familyCode = "default familyCode field value"
		
		delegate = AircraftFamilyDelegate()
		responseObj = delegate.create(aircraftFamily)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


