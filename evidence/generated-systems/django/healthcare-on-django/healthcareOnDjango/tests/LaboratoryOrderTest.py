import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.LaboratoryOrder import LaboratoryOrder
from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

 #======================================================================
# 
# Encapsulates data for model LaboratoryOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LaboratoryOrderTest Declaration
#======================================================================
class LaboratoryOrderTest (TestCase) :
	def test_crud(self) :
		laboratoryOrder = LaboratoryOrder()
		laboratoryOrder.testCode = "default testCode field value"
		laboratoryOrder.fastingRequired = False
		laboratoryOrder.specimenType = "default specimenType field value"
		
		delegate = LaboratoryOrderDelegate()
		responseObj = delegate.create(laboratoryOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


