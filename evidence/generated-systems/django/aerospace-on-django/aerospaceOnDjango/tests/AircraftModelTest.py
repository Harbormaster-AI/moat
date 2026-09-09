import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AircraftModel import AircraftModel
from aerospaceOnDjango.delegates.AircraftModelDelegate import AircraftModelDelegate

 #======================================================================
# 
# Encapsulates data for model AircraftModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftModelTest Declaration
#======================================================================
class AircraftModelTest (TestCase) :
	def test_crud(self) :
		aircraftModel = AircraftModel()
		aircraftModel.name = "default name field value"
		aircraftModel.modelDesignation = "default modelDesignation field value"
		aircraftModel.aircraftType = "default aircraftType field value"
		
		delegate = AircraftModelDelegate()
		responseObj = delegate.create(aircraftModel)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


