import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Approval import Approval
from hrOnDjango.delegates.ApprovalDelegate import ApprovalDelegate

 #======================================================================
# 
# Encapsulates data for model Approval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApprovalTest Declaration
#======================================================================
class ApprovalTest (TestCase) :
	def test_crud(self) :
		approval = Approval()
		approval.approverComment = "default approverComment field value"
		approval.actionDate = datetime.datetime.now()
		approval.status = "default status field value"
		
		delegate = ApprovalDelegate()
		responseObj = delegate.create(approval)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


