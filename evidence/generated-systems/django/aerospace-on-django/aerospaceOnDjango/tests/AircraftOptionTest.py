import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AircraftOption import AircraftOption
from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

 #======================================================================
# 
# Encapsulates data for model AircraftOption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOptionTest Declaration
#======================================================================
class AircraftOptionTest (TestCase) :
	def test_crud(self) :
		aircraftOption = AircraftOption()
		aircraftOption.code = "default code field value"
		aircraftOption.name = "default name field value"
		aircraftOption.optionCategory = "default optionCategory field value"
		
		delegate = AircraftOptionDelegate()
		responseObj = delegate.create(aircraftOption)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


