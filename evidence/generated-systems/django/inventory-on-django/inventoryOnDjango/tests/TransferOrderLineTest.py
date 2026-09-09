import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.TransferOrderLine import TransferOrderLine
from inventoryOnDjango.delegates.TransferOrderLineDelegate import TransferOrderLineDelegate

 #======================================================================
# 
# Encapsulates data for model TransferOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderLineTest Declaration
#======================================================================
class TransferOrderLineTest (TestCase) :
	def test_crud(self) :
		transferOrderLine = TransferOrderLine()
		transferOrderLine.lineNumber = 22
		transferOrderLine.quantity = "default quantity field value"
		transferOrderLine.unitOfMeasure = "default unitOfMeasure field value"
		transferOrderLine.stockStatus = "default stockStatus field value"
		
		delegate = TransferOrderLineDelegate()
		responseObj = delegate.create(transferOrderLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


