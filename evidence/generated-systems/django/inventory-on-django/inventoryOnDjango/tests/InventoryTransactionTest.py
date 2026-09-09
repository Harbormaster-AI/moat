import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.InventoryTransaction import InventoryTransaction
from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

 #======================================================================
# 
# Encapsulates data for model InventoryTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryTransactionTest Declaration
#======================================================================
class InventoryTransactionTest (TestCase) :
	def test_crud(self) :
		inventoryTransaction = InventoryTransaction()
		inventoryTransaction.transactionNumber = "default transactionNumber field value"
		inventoryTransaction.quantity = "default quantity field value"
		inventoryTransaction.unitCost = "default unitCost field value"
		inventoryTransaction.transactionDate = datetime.datetime.now()
		inventoryTransaction.reasonCode = "default reasonCode field value"
		inventoryTransaction.transactionType = "default transactionType field value"
		inventoryTransaction.unitOfMeasure = "default unitOfMeasure field value"
		inventoryTransaction.status = "default status field value"
		
		delegate = InventoryTransactionDelegate()
		responseObj = delegate.create(inventoryTransaction)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


