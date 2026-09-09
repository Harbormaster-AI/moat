import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Allergy import Allergy
from healthcareOnDjango.delegates.AllergyDelegate import AllergyDelegate

 #======================================================================
# 
# Encapsulates data for model Allergy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AllergyTest Declaration
#======================================================================
class AllergyTest (TestCase) :
	def test_crud(self) :
		allergy = Allergy()
		allergy.substance = "default substance field value"
		allergy.reaction = "default reaction field value"
		allergy.severity = "default severity field value"
		allergy.status = "default status field value"
		
		delegate = AllergyDelegate()
		responseObj = delegate.create(allergy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


