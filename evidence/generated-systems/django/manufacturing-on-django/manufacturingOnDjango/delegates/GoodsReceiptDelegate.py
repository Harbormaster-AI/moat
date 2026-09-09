from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.GoodsReceipt import GoodsReceipt
from manufacturingOnDjango.models.PurchaseOrder import PurchaseOrder
from manufacturingOnDjango.models.Warehouse import Warehouse
from manufacturingOnDjango.models.GoodsReceiptLine import GoodsReceiptLine
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model GoodsReceipt
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceiptDelegate Declaration
#======================================================================
class GoodsReceiptDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, goodsReceiptId ):
		try:	
			goodsReceipt = GoodsReceipt.objects.filter(id=goodsReceiptId)
			return goodsReceipt.first();
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError("GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, goodsReceipt):
		for model in serializers.deserialize("json", goodsReceipt):
			model.save()
			return model;

	def create(self, goodsReceipt):
		goodsReceipt.save()
		return goodsReceipt;

	def saveFromJson(self, goodsReceipt):
		for model in serializers.deserialize("json", goodsReceipt):
			model.save()
			return goodsReceipt;
	
	def save(self, goodsReceipt):
		goodsReceipt.save()
		return goodsReceipt;
	
	def delete(self, goodsReceiptId ):
		errMsg = "Failed to delete GoodsReceipt from db using id " + str(goodsReceiptId)
		
		try:
			goodsReceipt = GoodsReceipt.objects.get(id=goodsReceiptId)
			goodsReceipt.delete()
			return True
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError("GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = GoodsReceipt.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all GoodsReceipt from db")
		except Exception:
			return None;
		
	def assignPurchaseOrder( self, goodsReceiptId, purchaseOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PurchaseOrderDelegate import PurchaseOrderDelegate

		errMsg = "Failed to assign element " + str(purchaseOrderId) + " for PurchaseOrder on GoodsReceipt"

		try:
			# get the GoodsReceipt from db
			goodsReceipt = self.get( goodsReceiptId ).first()	
			
			# get the PurchaseOrder from db
			purchaseOrder = PurchaseOrderDelegate().get(purchaseOrderId).first();
			
			# assign the PurchaseOrder		
			goodsReceipt.purchaseOrder = purchaseOrder
			
			#save it
			goodsReceipt.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptId );
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPurchaseOrder( self, goodsReceiptId ):
		errMsg = "Failed to unassign element " + str(purchaseOrderId) + " for PurchaseOrder on GoodsReceipt"

		try:
			# get the GoodsReceipt from db
			goodsReceipt = self.get( goodsReceiptId ).first()	
			
			# assign to None for unassignment
			goodsReceipt.purchaseOrder = None			

			#save it
			goodsReceipt.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptId );
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarehouse( self, goodsReceiptId, warehouseId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on GoodsReceipt"

		try:
			# get the GoodsReceipt from db
			goodsReceipt = self.get( goodsReceiptId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			goodsReceipt.warehouse = warehouse
			
			#save it
			goodsReceipt.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptId );
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, goodsReceiptId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on GoodsReceipt"

		try:
			# get the GoodsReceipt from db
			goodsReceipt = self.get( goodsReceiptId ).first()	
			
			# assign to None for unassignment
			goodsReceipt.warehouse = None			

			#save it
			goodsReceipt.save()

			# reload and return the appropriate version					
			return self.get( goodsReceiptId );
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except Exception:
			return None;
		
	def addLines( self, goodsReceiptId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.GoodsReceiptLineDelegate import GoodsReceiptLineDelegate

		errMsg = "Failed to add elements " + str(linesIds) + " for Lines on GoodsReceipt"

		try:
			# get the GoodsReceipt
			goodsReceipt = self.get( goodsReceiptId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GoodsReceiptLine		
				goodsReceiptLine = GoodsReceiptLineDelegate().get(id).first();	
				# add the GoodsReceiptLine
				goodsReceipt.lines.add(goodsReceiptLine)
				
			# save it		
			goodsReceipt.save()
			
			# reload and return the appropriate version
			return self.get( goodsReceiptId );
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLines( self, goodsReceiptId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.GoodsReceiptLineDelegate import GoodsReceiptLineDelegate

		errMsg = "Failed to remove elements " + str(linesIds) + " for Lines on GoodsReceipt"

		try:
			# get the GoodsReceipt
			goodsReceipt = self.get( goodsReceiptId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GoodsReceiptLine		
				goodsReceiptLine = GoodsReceiptLineDelegate().get(id).first();	
				# add the GoodsReceiptLine
				goodsReceipt.lines.remove(goodsReceiptLine)
				
			# save it		
			goodsReceipt.save()
			
			# reload and return the appropriate version
			return self.get( goodsReceiptId );
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except GoodsReceiptLine.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceiptLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
