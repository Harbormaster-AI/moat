import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Trade import Trade
from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

 #======================================================================
# 
# Encapsulates data for model Trade
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TradeTest Declaration
#======================================================================
class TradeTest (TestCase) :
	def test_crud(self) :
		trade = Trade()
		trade.executedAt = "default executedAt field value"
		trade.quantity = "default quantity field value"
		trade.price = "default price field value"
		trade.fees = "default fees field value"
		trade.settlementDate = datetime.datetime.now()
		
		delegate = TradeDelegate()
		responseObj = delegate.create(trade)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


