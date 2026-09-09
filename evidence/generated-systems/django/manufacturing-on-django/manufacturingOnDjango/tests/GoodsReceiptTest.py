import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.GoodsReceipt import GoodsReceipt
from manufacturingOnDjango.delegates.GoodsReceiptDelegate import GoodsReceiptDelegate

 #======================================================================
# 
# Encapsulates data for model GoodsReceipt
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceiptTest Declaration
#======================================================================
class GoodsReceiptTest (TestCase) :
	def test_crud(self) :
		goodsReceipt = GoodsReceipt()
		goodsReceipt.receiptNumber = "default receiptNumber field value"
		goodsReceipt.receiptDate = datetime.datetime.now()
		goodsReceipt.status = "default status field value"
		
		delegate = GoodsReceiptDelegate()
		responseObj = delegate.create(goodsReceipt)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


