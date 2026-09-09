import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.AppliedFee import AppliedFee
from fintechOnDjango.delegates.AppliedFeeDelegate import AppliedFeeDelegate

 #======================================================================
# 
# Encapsulates data for model AppliedFee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppliedFeeTest Declaration
#======================================================================
class AppliedFeeTest (TestCase) :
	def test_crud(self) :
		appliedFee = AppliedFee()
		appliedFee.amount = "default amount field value"
		appliedFee.description = "default description field value"
		appliedFee.feeType = "default feeType field value"
		
		delegate = AppliedFeeDelegate()
		responseObj = delegate.create(appliedFee)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


