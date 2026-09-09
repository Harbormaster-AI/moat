import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

 #======================================================================
# 
# Encapsulates data for model Encounter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EncounterTest Declaration
#======================================================================
class EncounterTest (TestCase) :
	def test_crud(self) :
		encounter = Encounter()
		encounter.encounterNumber = "default encounterNumber field value"
		encounter.startDateTime = "default startDateTime field value"
		encounter.endDateTime = "default endDateTime field value"
		encounter.status = "default status field value"
		encounter.encounterType = "default encounterType field value"
		
		delegate = EncounterDelegate()
		responseObj = delegate.create(encounter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


