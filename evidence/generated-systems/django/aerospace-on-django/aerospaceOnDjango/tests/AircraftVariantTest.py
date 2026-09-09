import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

 #======================================================================
# 
# Encapsulates data for model AircraftVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftVariantTest Declaration
#======================================================================
class AircraftVariantTest (TestCase) :
	def test_crud(self) :
		aircraftVariant = AircraftVariant()
		aircraftVariant.variantCode = "default variantCode field value"
		aircraftVariant.rangeNm = 22
		aircraftVariant.maxTakeoffWeightKg = "default maxTakeoffWeightKg field value"
		
		delegate = AircraftVariantDelegate()
		responseObj = delegate.create(aircraftVariant)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


