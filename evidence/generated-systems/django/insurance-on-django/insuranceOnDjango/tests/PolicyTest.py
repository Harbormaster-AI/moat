import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

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
		policy.policyNumber = "default policyNumber field value"
		policy.effectivePeriod = "default effectivePeriod field value"
		policy.totalPremium = "default totalPremium field value"
		policy.status = "default status field value"
		policy.paymentPlan = "default paymentPlan field value"
		
		delegate = PolicyDelegate()
		responseObj = delegate.create(policy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


