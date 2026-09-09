import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Chargeback import Chargeback
from fintechOnDjango.delegates.ChargebackDelegate import ChargebackDelegate

 #======================================================================
# 
# Encapsulates data for model Chargeback
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChargebackTest Declaration
#======================================================================
class ChargebackTest (TestCase) :
	def test_crud(self) :
		chargeback = Chargeback()
		chargeback.chargebackReference = "default chargebackReference field value"
		chargeback.amount = "default amount field value"
		chargeback.postedAt = "default postedAt field value"
		chargeback.stage = "default stage field value"
		chargeback.status = "default status field value"
		
		delegate = ChargebackDelegate()
		responseObj = delegate.create(chargeback)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


