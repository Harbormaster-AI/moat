from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.InventoryTransaction import InventoryTransaction
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.Location import Location
from manufacturingOnDjango.models.WorkOrder import WorkOrder
from manufacturingOnDjango.models.PurchaseOrder import PurchaseOrder
from manufacturingOnDjango.models.SalesOrder import SalesOrder
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InventoryTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryTransactionDelegate Declaration
#======================================================================
class InventoryTransactionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inventoryTransactionId ):
		try:	
			inventoryTransaction = InventoryTransaction.objects.filter(id=inventoryTransactionId)
			return inventoryTransaction.first();
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError("InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inventoryTransaction):
		for model in serializers.deserialize("json", inventoryTransaction):
			model.save()
			return model;

	def create(self, inventoryTransaction):
		inventoryTransaction.save()
		return inventoryTransaction;

	def saveFromJson(self, inventoryTransaction):
		for model in serializers.deserialize("json", inventoryTransaction):
			model.save()
			return inventoryTransaction;
	
	def save(self, inventoryTransaction):
		inventoryTransaction.save()
		return inventoryTransaction;
	
	def delete(self, inventoryTransactionId ):
		errMsg = "Failed to delete InventoryTransaction from db using id " + str(inventoryTransactionId)
		
		try:
			inventoryTransaction = InventoryTransaction.objects.get(id=inventoryTransactionId)
			inventoryTransaction.delete()
			return True
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError("InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InventoryTransaction.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InventoryTransaction from db")
		except Exception:
			return None;
		
	def assignItem( self, inventoryTransactionId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			inventoryTransaction.item = item
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.item = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, inventoryTransactionId, locationId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the Location from db
			location = LocationDelegate().get(locationId).first();
			
			# assign the Location		
			inventoryTransaction.location = location
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.location = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkOrder( self, inventoryTransactionId, workOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

		errMsg = "Failed to assign element " + str(workOrderId) + " for WorkOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the WorkOrder from db
			workOrder = WorkOrderDelegate().get(workOrderId).first();
			
			# assign the WorkOrder		
			inventoryTransaction.workOrder = workOrder
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkOrder( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(workOrderId) + " for WorkOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.workOrder = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPurchaseOrder( self, inventoryTransactionId, purchaseOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PurchaseOrderDelegate import PurchaseOrderDelegate

		errMsg = "Failed to assign element " + str(purchaseOrderId) + " for PurchaseOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the PurchaseOrder from db
			purchaseOrder = PurchaseOrderDelegate().get(purchaseOrderId).first();
			
			# assign the PurchaseOrder		
			inventoryTransaction.purchaseOrder = purchaseOrder
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPurchaseOrder( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(purchaseOrderId) + " for PurchaseOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.purchaseOrder = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSalesOrder( self, inventoryTransactionId, salesOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SalesOrderDelegate import SalesOrderDelegate

		errMsg = "Failed to assign element " + str(salesOrderId) + " for SalesOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the SalesOrder from db
			salesOrder = SalesOrderDelegate().get(salesOrderId).first();
			
			# assign the SalesOrder		
			inventoryTransaction.salesOrder = salesOrder
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSalesOrder( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(salesOrderId) + " for SalesOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.salesOrder = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
