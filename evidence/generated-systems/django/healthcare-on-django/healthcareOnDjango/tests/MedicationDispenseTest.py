import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.MedicationDispense import MedicationDispense
from healthcareOnDjango.delegates.MedicationDispenseDelegate import MedicationDispenseDelegate

 #======================================================================
# 
# Encapsulates data for model MedicationDispense
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationDispenseTest Declaration
#======================================================================
class MedicationDispenseTest (TestCase) :
	def test_crud(self) :
		medicationDispense = MedicationDispense()
		medicationDispense.dispenseNumber = "default dispenseNumber field value"
		medicationDispense.quantity = "default quantity field value"
		medicationDispense.whenPrepared = "default whenPrepared field value"
		medicationDispense.status = "default status field value"
		
		delegate = MedicationDispenseDelegate()
		responseObj = delegate.create(medicationDispense)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


