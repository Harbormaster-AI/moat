import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.MedicationOrder import MedicationOrder
from healthcareOnDjango.delegates.MedicationOrderDelegate import MedicationOrderDelegate

 #======================================================================
# 
# Encapsulates data for model MedicationOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicationOrderTest Declaration
#======================================================================
class MedicationOrderTest (TestCase) :
	def test_crud(self) :
		medicationOrder = MedicationOrder()
		medicationOrder.medicationCode = "default medicationCode field value"
		medicationOrder.dose = "default dose field value"
		medicationOrder.frequency = "default frequency field value"
		medicationOrder.duration = "default duration field value"
		medicationOrder.route = "default route field value"
		
		delegate = MedicationOrderDelegate()
		responseObj = delegate.create(medicationOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


