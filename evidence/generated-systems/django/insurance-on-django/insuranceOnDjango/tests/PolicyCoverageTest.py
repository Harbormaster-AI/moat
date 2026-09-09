import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.PolicyCoverage import PolicyCoverage
from insuranceOnDjango.delegates.PolicyCoverageDelegate import PolicyCoverageDelegate

 #======================================================================
# 
# Encapsulates data for model PolicyCoverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyCoverageTest Declaration
#======================================================================
class PolicyCoverageTest (TestCase) :
	def test_crud(self) :
		policyCoverage = PolicyCoverage()
		policyCoverage.limit = "default limit field value"
		policyCoverage.deductible = "default deductible field value"
		policyCoverage.premium = "default premium field value"
		policyCoverage.coverageType = "default coverageType field value"
		
		delegate = PolicyCoverageDelegate()
		responseObj = delegate.create(policyCoverage)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


