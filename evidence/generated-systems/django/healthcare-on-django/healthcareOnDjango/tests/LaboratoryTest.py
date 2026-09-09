import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Laboratory import Laboratory
from healthcareOnDjango.delegates.LaboratoryDelegate import LaboratoryDelegate

 #======================================================================
# 
# Encapsulates data for model Laboratory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LaboratoryTest Declaration
#======================================================================
class LaboratoryTest (TestCase) :
	def test_crud(self) :
		laboratory = Laboratory()
		laboratory.name = "default name field value"
		laboratory.cliaNumber = "default cliaNumber field value"
		
		delegate = LaboratoryDelegate()
		responseObj = delegate.create(laboratory)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


