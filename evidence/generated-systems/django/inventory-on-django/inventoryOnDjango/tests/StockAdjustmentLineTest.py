import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.StockAdjustmentLine import StockAdjustmentLine
from inventoryOnDjango.delegates.StockAdjustmentLineDelegate import StockAdjustmentLineDelegate

 #======================================================================
# 
# Encapsulates data for model StockAdjustmentLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockAdjustmentLineTest Declaration
#======================================================================
class StockAdjustmentLineTest (TestCase) :
	def test_crud(self) :
		stockAdjustmentLine = StockAdjustmentLine()
		stockAdjustmentLine.lineNumber = 22
		stockAdjustmentLine.quantity = "default quantity field value"
		stockAdjustmentLine.unitOfMeasure = "default unitOfMeasure field value"
		stockAdjustmentLine.stockStatus = "default stockStatus field value"
		
		delegate = StockAdjustmentLineDelegate()
		responseObj = delegate.create(stockAdjustmentLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


