import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.BusinessGlossaryTerm import BusinessGlossaryTerm
from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

 #======================================================================
# 
# Encapsulates data for model BusinessGlossaryTerm
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessGlossaryTermTest Declaration
#======================================================================
class BusinessGlossaryTermTest (TestCase) :
	def test_crud(self) :
		businessGlossaryTerm = BusinessGlossaryTerm()
		businessGlossaryTerm.term = "default term field value"
		businessGlossaryTerm.definition = "default definition field value"
		businessGlossaryTerm.steward = "default steward field value"
		
		delegate = BusinessGlossaryTermDelegate()
		responseObj = delegate.create(businessGlossaryTerm)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


