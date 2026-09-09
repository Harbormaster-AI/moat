import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.InsurancePlan import InsurancePlan
from healthcareOnDjango.delegates.InsurancePlanDelegate import InsurancePlanDelegate

 #======================================================================
# 
# Encapsulates data for model InsurancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePlanTest Declaration
#======================================================================
class InsurancePlanTest (TestCase) :
	def test_crud(self) :
		insurancePlan = InsurancePlan()
		insurancePlan.name = "default name field value"
		insurancePlan.planCode = "default planCode field value"
		insurancePlan.planType = "default planType field value"
		
		delegate = InsurancePlanDelegate()
		responseObj = delegate.create(insurancePlan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


