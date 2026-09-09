import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.CreativeApproval import CreativeApproval
from advertisingOnDjango.delegates.CreativeApprovalDelegate import CreativeApprovalDelegate

 #======================================================================
# 
# Encapsulates data for model CreativeApproval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeApprovalTest Declaration
#======================================================================
class CreativeApprovalTest (TestCase) :
	def test_crud(self) :
		creativeApproval = CreativeApproval()
		creativeApproval.reviewer = "default reviewer field value"
		creativeApproval.reviewedAt = datetime.datetime.now()
		creativeApproval.status = "default status field value"
		
		delegate = CreativeApprovalDelegate()
		responseObj = delegate.create(creativeApproval)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


