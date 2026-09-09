import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AircraftProgram import AircraftProgram
from aerospaceOnDjango.delegates.AircraftProgramDelegate import AircraftProgramDelegate

 #======================================================================
# 
# Encapsulates data for model AircraftProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftProgramTest Declaration
#======================================================================
class AircraftProgramTest (TestCase) :
	def test_crud(self) :
		aircraftProgram = AircraftProgram()
		aircraftProgram.name = "default name field value"
		aircraftProgram.programCode = "default programCode field value"
		aircraftProgram.entryIntoServiceYear = 22
		aircraftProgram.status = "default status field value"
		
		delegate = AircraftProgramDelegate()
		responseObj = delegate.create(aircraftProgram)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


