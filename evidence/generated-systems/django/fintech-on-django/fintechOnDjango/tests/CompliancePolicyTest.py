import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.CompliancePolicy import CompliancePolicy
from fintechOnDjango.delegates.CompliancePolicyDelegate import CompliancePolicyDelegate

 #======================================================================
# 
# Encapsulates data for model CompliancePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompliancePolicyTest Declaration
#======================================================================
class CompliancePolicyTest (TestCase) :
	def test_crud(self) :
		compliancePolicy = CompliancePolicy()
		compliancePolicy.name = "default name field value"
		compliancePolicy.policyCode = "default policyCode field value"
		compliancePolicy.description = "default description field value"
		compliancePolicy.status = "default status field value"
		
		delegate = CompliancePolicyDelegate()
		responseObj = delegate.create(compliancePolicy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


