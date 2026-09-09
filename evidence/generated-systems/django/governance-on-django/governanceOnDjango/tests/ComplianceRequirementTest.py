import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.ComplianceRequirement import ComplianceRequirement
from governanceOnDjango.delegates.ComplianceRequirementDelegate import ComplianceRequirementDelegate

 #======================================================================
# 
# Encapsulates data for model ComplianceRequirement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceRequirementTest Declaration
#======================================================================
class ComplianceRequirementTest (TestCase) :
	def test_crud(self) :
		complianceRequirement = ComplianceRequirement()
		complianceRequirement.name = "default name field value"
		complianceRequirement.source = "default source field value"
		complianceRequirement.citation = "default citation field value"
		complianceRequirement.applicability = "default applicability field value"
		complianceRequirement.status = "default status field value"
		
		delegate = ComplianceRequirementDelegate()
		responseObj = delegate.create(complianceRequirement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


