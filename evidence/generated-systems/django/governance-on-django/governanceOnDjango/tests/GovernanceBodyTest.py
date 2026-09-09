import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.GovernanceBody import GovernanceBody
from governanceOnDjango.delegates.GovernanceBodyDelegate import GovernanceBodyDelegate

 #======================================================================
# 
# Encapsulates data for model GovernanceBody
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GovernanceBodyTest Declaration
#======================================================================
class GovernanceBodyTest (TestCase) :
	def test_crud(self) :
		governanceBody = GovernanceBody()
		governanceBody.name = "default name field value"
		governanceBody.charterUrl = "default charterUrl field value"
		governanceBody.chair = "default chair field value"
		governanceBody.bodyType = "default bodyType field value"
		
		delegate = GovernanceBodyDelegate()
		responseObj = delegate.create(governanceBody)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


