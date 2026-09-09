import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

 #======================================================================
# 
# Encapsulates data for model Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyTest Declaration
#======================================================================
class PolicyTest (TestCase) :
	def test_crud(self) :
		policy = Policy()
		policy.title = "default title field value"
		policy.versionLabel = "default versionLabel field value"
		policy.approvalDate = datetime.datetime.now()
		policy.nextReviewDate = datetime.datetime.now()
		policy.documentUrl = "default documentUrl field value"
		policy.policyType = "default policyType field value"
		policy.status = "default status field value"
		
		delegate = PolicyDelegate()
		responseObj = delegate.create(policy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


