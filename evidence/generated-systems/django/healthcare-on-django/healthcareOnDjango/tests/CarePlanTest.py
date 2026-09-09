import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.CarePlan import CarePlan
from healthcareOnDjango.delegates.CarePlanDelegate import CarePlanDelegate

 #======================================================================
# 
# Encapsulates data for model CarePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarePlanTest Declaration
#======================================================================
class CarePlanTest (TestCase) :
	def test_crud(self) :
		carePlan = CarePlan()
		carePlan.planNumber = "default planNumber field value"
		carePlan.goalSummary = "default goalSummary field value"
		carePlan.status = "default status field value"
		
		delegate = CarePlanDelegate()
		responseObj = delegate.create(carePlan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


