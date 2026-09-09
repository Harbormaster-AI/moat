import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.FXDeal import FXDeal
from fintechOnDjango.delegates.FXDealDelegate import FXDealDelegate

 #======================================================================
# 
# Encapsulates data for model FXDeal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXDealTest Declaration
#======================================================================
class FXDealTest (TestCase) :
	def test_crud(self) :
		fXDeal = FXDeal()
		fXDeal.dealReference = "default dealReference field value"
		fXDeal.baseCurrency = "default baseCurrency field value"
		fXDeal.quoteCurrency = "default quoteCurrency field value"
		fXDeal.rate = "default rate field value"
		fXDeal.amount = "default amount field value"
		fXDeal.settlementDate = datetime.datetime.now()
		fXDeal.status = "default status field value"
		
		delegate = FXDealDelegate()
		responseObj = delegate.create(fXDeal)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


