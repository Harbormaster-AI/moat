from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.PurchaseOrderLine import PurchaseOrderLine
from manufacturingOnDjango.models.PurchaseOrder import PurchaseOrder
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PurchaseOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderLineDelegate Declaration
#======================================================================
class PurchaseOrderLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, purchaseOrderLineId ):
		try:	
			purchaseOrderLine = PurchaseOrderLine.objects.filter(id=purchaseOrderLineId)
			return purchaseOrderLine.first();
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError("PurchaseOrderLine with id " + str(purchaseOrderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, purchaseOrderLine):
		for model in serializers.deserialize("json", purchaseOrderLine):
			model.save()
			return model;

	def create(self, purchaseOrderLine):
		purchaseOrderLine.save()
		return purchaseOrderLine;

	def saveFromJson(self, purchaseOrderLine):
		for model in serializers.deserialize("json", purchaseOrderLine):
			model.save()
			return purchaseOrderLine;
	
	def save(self, purchaseOrderLine):
		purchaseOrderLine.save()
		return purchaseOrderLine;
	
	def delete(self, purchaseOrderLineId ):
		errMsg = "Failed to delete PurchaseOrderLine from db using id " + str(purchaseOrderLineId)
		
		try:
			purchaseOrderLine = PurchaseOrderLine.objects.get(id=purchaseOrderLineId)
			purchaseOrderLine.delete()
			return True
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError("PurchaseOrderLine with id " + str(purchaseOrderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PurchaseOrderLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PurchaseOrderLine from db")
		except Exception:
			return None;
		
	def assignPurchaseOrder( self, purchaseOrderLineId, purchaseOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PurchaseOrderDelegate import PurchaseOrderDelegate

		errMsg = "Failed to assign element " + str(purchaseOrderId) + " for PurchaseOrder on PurchaseOrderLine"

		try:
			# get the PurchaseOrderLine from db
			purchaseOrderLine = self.get( purchaseOrderLineId ).first()	
			
			# get the PurchaseOrder from db
			purchaseOrder = PurchaseOrderDelegate().get(purchaseOrderId).first();
			
			# assign the PurchaseOrder		
			purchaseOrderLine.purchaseOrder = purchaseOrder
			
			#save it
			purchaseOrderLine.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderLineId );
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrderLine with id " + str(purchaseOrderLineId) + " does not exist.")
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder with id " + str(purchaseOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPurchaseOrder( self, purchaseOrderLineId ):
		errMsg = "Failed to unassign element " + str(purchaseOrderId) + " for PurchaseOrder on PurchaseOrderLine"

		try:
			# get the PurchaseOrderLine from db
			purchaseOrderLine = self.get( purchaseOrderLineId ).first()	
			
			# assign to None for unassignment
			purchaseOrderLine.purchaseOrder = None			

			#save it
			purchaseOrderLine.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderLineId );
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrderLine with id " + str(purchaseOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignItem( self, purchaseOrderLineId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on PurchaseOrderLine"

		try:
			# get the PurchaseOrderLine from db
			purchaseOrderLine = self.get( purchaseOrderLineId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			purchaseOrderLine.item = item
			
			#save it
			purchaseOrderLine.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderLineId );
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrderLine with id " + str(purchaseOrderLineId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, purchaseOrderLineId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on PurchaseOrderLine"

		try:
			# get the PurchaseOrderLine from db
			purchaseOrderLine = self.get( purchaseOrderLineId ).first()	
			
			# assign to None for unassignment
			purchaseOrderLine.item = None			

			#save it
			purchaseOrderLine.save()

			# reload and return the appropriate version					
			return self.get( purchaseOrderLineId );
		except PurchaseOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrderLine with id " + str(purchaseOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
