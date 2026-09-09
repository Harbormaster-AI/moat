import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Coverage import Coverage
from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

 #======================================================================
# 
# Encapsulates data for model Coverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageTest Declaration
#======================================================================
class CoverageTest (TestCase) :
	def test_crud(self) :
		coverage = Coverage()
		coverage.memberId = "default memberId field value"
		coverage.groupNumber = "default groupNumber field value"
		coverage.effectiveDate = datetime.datetime.now()
		coverage.endDate = datetime.datetime.now()
		coverage.coverageType = "default coverageType field value"
		
		delegate = CoverageDelegate()
		responseObj = delegate.create(coverage)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


