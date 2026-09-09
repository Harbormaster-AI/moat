import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.ExpirationPolicy import ExpirationPolicy
from inventoryOnDjango.delegates.ExpirationPolicyDelegate import ExpirationPolicyDelegate

 #======================================================================
# 
# Encapsulates data for model ExpirationPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExpirationPolicyTest Declaration
#======================================================================
class ExpirationPolicyTest (TestCase) :
	def test_crud(self) :
		expirationPolicy = ExpirationPolicy()
		expirationPolicy.rejectIfDaysToExpireLessThan = 22
		expirationPolicy.autoQuarantineDaysToExpire = 22
		expirationPolicy.rotationMethod = "default rotationMethod field value"
		
		delegate = ExpirationPolicyDelegate()
		responseObj = delegate.create(expirationPolicy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


