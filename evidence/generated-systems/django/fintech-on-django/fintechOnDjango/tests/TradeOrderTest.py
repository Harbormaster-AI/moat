import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.TradeOrder import TradeOrder
from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

 #======================================================================
# 
# Encapsulates data for model TradeOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TradeOrderTest Declaration
#======================================================================
class TradeOrderTest (TestCase) :
	def test_crud(self) :
		tradeOrder = TradeOrder()
		tradeOrder.orderId = "default orderId field value"
		tradeOrder.quantity = "default quantity field value"
		tradeOrder.limitPrice = "default limitPrice field value"
		tradeOrder.placedAt = "default placedAt field value"
		tradeOrder.side = "default side field value"
		tradeOrder.type = "default type field value"
		tradeOrder.status = "default status field value"
		tradeOrder.timeInForce = "default timeInForce field value"
		
		delegate = TradeOrderDelegate()
		responseObj = delegate.create(tradeOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


