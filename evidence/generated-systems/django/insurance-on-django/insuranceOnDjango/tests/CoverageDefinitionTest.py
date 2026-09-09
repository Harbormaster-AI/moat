import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.CoverageDefinition import CoverageDefinition
from insuranceOnDjango.delegates.CoverageDefinitionDelegate import CoverageDefinitionDelegate

 #======================================================================
# 
# Encapsulates data for model CoverageDefinition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageDefinitionTest Declaration
#======================================================================
class CoverageDefinitionTest (TestCase) :
	def test_crud(self) :
		coverageDefinition = CoverageDefinition()
		coverageDefinition.name = "default name field value"
		coverageDefinition.defaultLimit = "default defaultLimit field value"
		coverageDefinition.defaultDeductible = "default defaultDeductible field value"
		coverageDefinition.asMandatory = False
		coverageDefinition.coverageType = "default coverageType field value"
		
		delegate = CoverageDefinitionDelegate()
		responseObj = delegate.create(coverageDefinition)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


