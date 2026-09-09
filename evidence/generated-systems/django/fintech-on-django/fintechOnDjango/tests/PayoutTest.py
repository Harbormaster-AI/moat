import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Payout import Payout
from fintechOnDjango.delegates.PayoutDelegate import PayoutDelegate

 #======================================================================
# 
# Encapsulates data for model Payout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayoutTest Declaration
#======================================================================
class PayoutTest (TestCase) :
	def test_crud(self) :
		payout = Payout()
		payout.payoutReference = "default payoutReference field value"
		payout.amount = "default amount field value"
		payout.currency = "default currency field value"
		payout.scheduledDate = datetime.datetime.now()
		payout.paidDate = datetime.datetime.now()
		payout.status = "default status field value"
		
		delegate = PayoutDelegate()
		responseObj = delegate.create(payout)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


