from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.GoodsReceiptLine import GoodsReceiptLine
from manufacturingOnDjango.models.GoodsReceipt import GoodsReceipt
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.InventoryTransaction import InventoryTransaction
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model GoodsReceiptLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceiptLineDelegate Declaration
#======================================================================
class GoodsReceiptLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, goodsReceiptLineId ):
		try:	
			goodsReceiptLine = GoodsReceiptLine.objects.filter(id=goodsReceiptLineId)
			return goodsReceiptLine.first();
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError("GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, goodsReceiptLine):
		for model in serializers.deserialize("json", goodsReceiptLine):
			model.save()
			return model;

	def create(self, goodsReceiptLine):
		goodsReceiptLine.save()
		return goodsReceiptLine;

	def saveFromJson(self, goodsReceiptLine):
		for model in serializers.deserialize("json", goodsReceiptLine):
			model.save()
			return goodsReceiptLine;
	
	def save(self, goodsReceiptLine):
		goodsReceiptLine.save()
		return goodsReceiptLine;
	
	def delete(self, goodsReceiptLineId ):
		errMsg = "Failed to delete GoodsReceiptLine from db using id " + str(goodsReceiptLineId)
		
		try:
			goodsReceiptLine = GoodsReceiptLine.objects.get(id=goodsReceiptLineId)
			goodsReceiptLine.delete()
			return True
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError("GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = GoodsReceiptLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all GoodsReceiptLine from db")
		except Exception:
			return None;
		
	def assignGoodsReceipt( self, goodsReceiptLineId, goodsReceiptId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.GoodsReceiptDelegate import GoodsReceiptDelegate

		errMsg = "Failed to assign element " + str(goodsReceiptId) + " for GoodsReceipt on GoodsReceiptLine"

		try:
			# get the GoodsReceiptLine from db
			goodsReceiptLine = self.get( goodsReceiptLineId ).first()	
			
			# get the GoodsReceipt from db
			goodsReceipt = GoodsReceiptDelegate().get(goodsReceiptId).first();
			
			# assign the GoodsReceipt		
			goodsReceiptLine.goodsReceipt = goodsReceipt
			
			#save it
			goodsReceiptLine.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptLineId );
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignGoodsReceipt( self, goodsReceiptLineId ):
		errMsg = "Failed to unassign element " + str(goodsReceiptId) + " for GoodsReceipt on GoodsReceiptLine"

		try:
			# get the GoodsReceiptLine from db
			goodsReceiptLine = self.get( goodsReceiptLineId ).first()	
			
			# assign to None for unassignment
			goodsReceiptLine.goodsReceipt = None			

			#save it
			goodsReceiptLine.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptLineId );
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignItem( self, goodsReceiptLineId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on GoodsReceiptLine"

		try:
			# get the GoodsReceiptLine from db
			goodsReceiptLine = self.get( goodsReceiptLineId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			goodsReceiptLine.item = item
			
			#save it
			goodsReceiptLine.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptLineId );
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, goodsReceiptLineId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on GoodsReceiptLine"

		try:
			# get the GoodsReceiptLine from db
			goodsReceiptLine = self.get( goodsReceiptLineId ).first()	
			
			# assign to None for unassignment
			goodsReceiptLine.item = None			

			#save it
			goodsReceiptLine.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptLineId );
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInventoryTransaction( self, goodsReceiptLineId, inventoryTransactionId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to assign element " + str(inventoryTransactionId) + " for InventoryTransaction on GoodsReceiptLine"

		try:
			# get the GoodsReceiptLine from db
			goodsReceiptLine = self.get( goodsReceiptLineId ).first()	
			
			# get the InventoryTransaction from db
			inventoryTransaction = InventoryTransactionDelegate().get(inventoryTransactionId).first();
			
			# assign the InventoryTransaction		
			goodsReceiptLine.inventoryTransaction = inventoryTransaction
			
			#save it
			goodsReceiptLine.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptLineId );
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInventoryTransaction( self, goodsReceiptLineId ):
		errMsg = "Failed to unassign element " + str(inventoryTransactionId) + " for InventoryTransaction on GoodsReceiptLine"

		try:
			# get the GoodsReceiptLine from db
			goodsReceiptLine = self.get( goodsReceiptLineId ).first()	
			
			# assign to None for unassignment
			goodsReceiptLine.inventoryTransaction = None			

			#save it
			goodsReceiptLine.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptLineId );
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine with id " + str(goodsReceiptLineId) + " does not exist.")
		except Exception:
			return None;
		
