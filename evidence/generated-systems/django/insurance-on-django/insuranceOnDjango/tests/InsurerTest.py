import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Insurer import Insurer
from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

 #======================================================================
# 
# Encapsulates data for model Insurer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurerTest Declaration
#======================================================================
class InsurerTest (TestCase) :
	def test_crud(self) :
		insurer = Insurer()
		insurer.name = "default name field value"
		insurer.legalName = "default legalName field value"
		insurer.domicileCountry = "default domicileCountry field value"
		insurer.naicNumber = "default naicNumber field value"
		insurer.website = "default website field value"
		
		delegate = InsurerDelegate()
		responseObj = delegate.create(insurer)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


