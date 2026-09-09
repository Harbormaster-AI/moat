import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.HealthSystem import HealthSystem
from healthcareOnDjango.delegates.HealthSystemDelegate import HealthSystemDelegate

 #======================================================================
# 
# Encapsulates data for model HealthSystem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class HealthSystemTest Declaration
#======================================================================
class HealthSystemTest (TestCase) :
	def test_crud(self) :
		healthSystem = HealthSystem()
		healthSystem.name = "default name field value"
		healthSystem.legalName = "default legalName field value"
		healthSystem.headquartersCountry = "default headquartersCountry field value"
		healthSystem.website = "default website field value"
		
		delegate = HealthSystemDelegate()
		responseObj = delegate.create(healthSystem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


