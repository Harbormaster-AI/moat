import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.AccessPolicy import AccessPolicy
from analyticsOnDjango.delegates.AccessPolicyDelegate import AccessPolicyDelegate

 #======================================================================
# 
# Encapsulates data for model AccessPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccessPolicyTest Declaration
#======================================================================
class AccessPolicyTest (TestCase) :
	def test_crud(self) :
		accessPolicy = AccessPolicy()
		accessPolicy.name = "default name field value"
		accessPolicy.subjectName = "default subjectName field value"
		accessPolicy.accessLevel = "default accessLevel field value"
		accessPolicy.subjectType = "default subjectType field value"
		
		delegate = AccessPolicyDelegate()
		responseObj = delegate.create(accessPolicy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


