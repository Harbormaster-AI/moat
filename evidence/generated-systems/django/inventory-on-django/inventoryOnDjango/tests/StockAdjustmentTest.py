import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.StockAdjustment import StockAdjustment
from inventoryOnDjango.delegates.StockAdjustmentDelegate import StockAdjustmentDelegate

 #======================================================================
# 
# Encapsulates data for model StockAdjustment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockAdjustmentTest Declaration
#======================================================================
class StockAdjustmentTest (TestCase) :
	def test_crud(self) :
		stockAdjustment = StockAdjustment()
		stockAdjustment.adjustmentNumber = "default adjustmentNumber field value"
		stockAdjustment.reason = "default reason field value"
		stockAdjustment.adjustmentDate = datetime.datetime.now()
		stockAdjustment.adjustmentType = "default adjustmentType field value"
		stockAdjustment.status = "default status field value"
		
		delegate = StockAdjustmentDelegate()
		responseObj = delegate.create(stockAdjustment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


