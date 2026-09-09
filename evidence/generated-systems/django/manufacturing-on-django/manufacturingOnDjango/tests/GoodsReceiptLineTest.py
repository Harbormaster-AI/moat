import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.GoodsReceiptLine import GoodsReceiptLine
from manufacturingOnDjango.delegates.GoodsReceiptLineDelegate import GoodsReceiptLineDelegate

 #======================================================================
# 
# Encapsulates data for model GoodsReceiptLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceiptLineTest Declaration
#======================================================================
class GoodsReceiptLineTest (TestCase) :
	def test_crud(self) :
		goodsReceiptLine = GoodsReceiptLine()
		goodsReceiptLine.lineNumber = 22
		goodsReceiptLine.receivedQuantity = "default receivedQuantity field value"
		goodsReceiptLine.acceptedQuantity = "default acceptedQuantity field value"
		goodsReceiptLine.rejectedQuantity = "default rejectedQuantity field value"
		goodsReceiptLine.lot = "default lot field value"
		
		delegate = GoodsReceiptLineDelegate()
		responseObj = delegate.create(goodsReceiptLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


