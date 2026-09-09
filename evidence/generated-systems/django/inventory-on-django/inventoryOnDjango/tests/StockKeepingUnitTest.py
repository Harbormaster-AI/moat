import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

 #======================================================================
# 
# Encapsulates data for model StockKeepingUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockKeepingUnitTest Declaration
#======================================================================
class StockKeepingUnitTest (TestCase) :
	def test_crud(self) :
		stockKeepingUnit = StockKeepingUnit()
		stockKeepingUnit.skuCode = "default skuCode field value"
		stockKeepingUnit.name = "default name field value"
		stockKeepingUnit.weight = "default weight field value"
		stockKeepingUnit.weightUnit = "default weightUnit field value"
		stockKeepingUnit.volume = "default volume field value"
		stockKeepingUnit.volumeUnit = "default volumeUnit field value"
		stockKeepingUnit.shelfLifeDays = 22
		stockKeepingUnit.hazardousMaterial = False
		stockKeepingUnit.itemType = "default itemType field value"
		stockKeepingUnit.unitOfMeasure = "default unitOfMeasure field value"
		
		delegate = StockKeepingUnitDelegate()
		responseObj = delegate.create(stockKeepingUnit)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


