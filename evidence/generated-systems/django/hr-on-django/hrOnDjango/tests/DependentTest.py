import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Dependent import Dependent
from hrOnDjango.delegates.DependentDelegate import DependentDelegate

 #======================================================================
# 
# Encapsulates data for model Dependent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DependentTest Declaration
#======================================================================
class DependentTest (TestCase) :
	def test_crud(self) :
		dependent = Dependent()
		dependent.firstName = "default firstName field value"
		dependent.lastName = "default lastName field value"
		dependent.birthDate = datetime.datetime.now()
		dependent.relationship = "default relationship field value"
		
		delegate = DependentDelegate()
		responseObj = delegate.create(dependent)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


