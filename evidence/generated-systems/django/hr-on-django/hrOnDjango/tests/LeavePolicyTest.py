import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.LeavePolicy import LeavePolicy
from hrOnDjango.delegates.LeavePolicyDelegate import LeavePolicyDelegate

 #======================================================================
# 
# Encapsulates data for model LeavePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeavePolicyTest Declaration
#======================================================================
class LeavePolicyTest (TestCase) :
	def test_crud(self) :
		leavePolicy = LeavePolicy()
		leavePolicy.name = "default name field value"
		leavePolicy.accrualRate = "default accrualRate field value"
		leavePolicy.carryoverAllowed = False
		leavePolicy.maxBalance = "default maxBalance field value"
		leavePolicy.leaveCategory = "default leaveCategory field value"
		leavePolicy.accrualUnit = "default accrualUnit field value"
		
		delegate = LeavePolicyDelegate()
		responseObj = delegate.create(leavePolicy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


