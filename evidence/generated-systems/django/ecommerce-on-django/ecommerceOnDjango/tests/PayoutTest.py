import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Payout import Payout
from ecommerceOnDjango.delegates.PayoutDelegate import PayoutDelegate

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
		payout.payoutNumber = "default payoutNumber field value"
		payout.amount = "default amount field value"
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


