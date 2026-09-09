import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Claim import Claim
from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

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
		claim.totalAmount = "default totalAmount field value"
		claim.status = "default status field value"
		
		delegate = ClaimDelegate()
		responseObj = delegate.create(claim)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


