import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

 #======================================================================
# 
# Encapsulates data for model Claim
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimTest Declaration
#======================================================================
class ClaimTest (TestCase) :
	def test_crud(self) :
		claim = Claim()
		claim.claimNumber = "default claimNumber field value"
		claim.noticeDate = datetime.datetime.now()
		claim.lossDate = datetime.datetime.now()
		claim.reportedBy = "default reportedBy field value"
		claim.status = "default status field value"
		claim.lossCause = "default lossCause field value"
		
		delegate = ClaimDelegate()
		responseObj = delegate.create(claim)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


