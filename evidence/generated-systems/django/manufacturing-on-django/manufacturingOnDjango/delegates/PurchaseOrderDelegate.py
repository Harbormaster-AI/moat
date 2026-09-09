from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.PurchaseOrder import PurchaseOrder
from manufacturingOnDjango.models.Supplier import Supplier
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.PurchaseOrderLine import PurchaseOrderLine
from manufacturingOnDjango.models.GoodsReceipt import GoodsReceipt
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PurchaseOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderDelegate Declaration
#======================================================================
class PurchaseOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, purchaseOrderId ):
		try:	
			purchaseOrder = PurchaseOrder.objects.filter(id=purchaseOrderId)
			return purchaseOrder.first();
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError("PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, purchaseOrder):
		for model in serializers.deserialize("json", purchaseOrder):
			model.save()
			return model;

	def create(self, purchaseOrder):
		purchaseOrder.save()
		return purchaseOrder;

	def saveFromJson(self, purchaseOrder):
		for model in serializers.deserialize("json", purchaseOrder):
			model.save()
			return purchaseOrder;
	
	def save(self, purchaseOrder):
		purchaseOrder.save()
		return purchaseOrder;
	
	def delete(self, purchaseOrderId ):
		errMsg = "Failed to delete PurchaseOrder from db using id " + str(purchaseOrderId)
		
		try:
			purchaseOrder = PurchaseOrder.objects.get(id=purchaseOrderId)
			purchaseOrder.delete()
			return True
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError("PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PurchaseOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PurchaseOrder from db")
		except Exception:
			return None;
		
	def assignSupplier( self, purchaseOrderId, supplierId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to assign element " + str(supplierId) + " for Supplier on PurchaseOrder"

		try:
			# get the PurchaseOrder from db
			purchaseOrder = self.get( purchaseOrderId ).first()	
			
			# get the Supplier from db
			supplier = SupplierDelegate().get(supplierId).first();
			
			# assign the Supplier		
			purchaseOrder.supplier = supplier
			
			#save it
			purchaseOrder.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupplier( self, purchaseOrderId ):
		errMsg = "Failed to unassign element " + str(supplierId) + " for Supplier on PurchaseOrder"

		try:
			# get the PurchaseOrder from db
			purchaseOrder = self.get( purchaseOrderId ).first()	
			
			# assign to None for unassignment
			purchaseOrder.supplier = None			

			#save it
			purchaseOrder.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlant( self, purchaseOrderId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on PurchaseOrder"

		try:
			# get the PurchaseOrder from db
			purchaseOrder = self.get( purchaseOrderId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			purchaseOrder.plant = plant
			
			#save it
			purchaseOrder.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, purchaseOrderId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on PurchaseOrder"

		try:
			# get the PurchaseOrder from db
			purchaseOrder = self.get( purchaseOrderId ).first()	
			
			# assign to None for unassignment
			purchaseOrder.plant = None			

			#save it
			purchaseOrder.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addLines( self, purchaseOrderId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PurchaseOrderLineDelegate import PurchaseOrderLineDelegate

		errMsg = "Failed to add elements " + str(linesIds) + " for Lines on PurchaseOrder"

		try:
			# get the PurchaseOrder
			purchaseOrder = self.get( purchaseOrderId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PurchaseOrderLine		
				purchaseOrderLine = PurchaseOrderLineDelegate().get(id).first();	
				# add the PurchaseOrderLine
				purchaseOrder.lines.add(purchaseOrderLine)
				
			# save it		
			purchaseOrder.save()
			
			# reload and return the appropriate version
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrderLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLines( self, purchaseOrderId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PurchaseOrderLineDelegate import PurchaseOrderLineDelegate

		errMsg = "Failed to remove elements " + str(linesIds) + " for Lines on PurchaseOrder"

		try:
			# get the PurchaseOrder
			purchaseOrder = self.get( purchaseOrderId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PurchaseOrderLine		
				purchaseOrderLine = PurchaseOrderLineDelegate().get(id).first();	
				# add the PurchaseOrderLine
				purchaseOrder.lines.remove(purchaseOrderLine)
				
			# save it		
			purchaseOrder.save()
			
			# reload and return the appropriate version
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrderLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGoodsReceipts( self, purchaseOrderId, goodsReceiptsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.GoodsReceiptDelegate import GoodsReceiptDelegate

		errMsg = "Failed to add elements " + str(goodsReceiptsIds) + " for GoodsReceipts on PurchaseOrder"

		try:
			# get the PurchaseOrder
			purchaseOrder = self.get( purchaseOrderId ).first()
				
			# split on a comma with no spaces
			idList = goodsReceiptsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GoodsReceipt		
				goodsReceipt = GoodsReceiptDelegate().get(id).first();	
				# add the GoodsReceipt
				purchaseOrder.goodsReceipts.add(goodsReceipt)
				
			# save it		
			purchaseOrder.save()
			
			# reload and return the appropriate version
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGoodsReceipts( self, purchaseOrderId, goodsReceiptsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.GoodsReceiptDelegate import GoodsReceiptDelegate

		errMsg = "Failed to remove elements " + str(goodsReceiptsIds) + " for GoodsReceipts on PurchaseOrder"

		try:
			# get the PurchaseOrder
			purchaseOrder = self.get( purchaseOrderId ).first()
				
			# split on a comma with no spaces
			idList = goodsReceiptsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GoodsReceipt		
				goodsReceipt = GoodsReceiptDelegate().get(id).first();	
				# add the GoodsReceipt
				purchaseOrder.goodsReceipts.remove(goodsReceipt)
				
			# save it		
			purchaseOrder.save()
			
			# reload and return the appropriate version
			return self.get( purchaseOrderId );
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
