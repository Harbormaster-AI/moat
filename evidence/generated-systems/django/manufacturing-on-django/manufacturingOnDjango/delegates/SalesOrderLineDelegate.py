from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.SalesOrderLine import SalesOrderLine
from manufacturingOnDjango.models.SalesOrder import SalesOrder
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SalesOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrderLineDelegate Declaration
#======================================================================
class SalesOrderLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, salesOrderLineId ):
		try:	
			salesOrderLine = SalesOrderLine.objects.filter(id=salesOrderLineId)
			return salesOrderLine.first();
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError("SalesOrderLine with id " + str(salesOrderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, salesOrderLine):
		for model in serializers.deserialize("json", salesOrderLine):
			model.save()
			return model;

	def create(self, salesOrderLine):
		salesOrderLine.save()
		return salesOrderLine;

	def saveFromJson(self, salesOrderLine):
		for model in serializers.deserialize("json", salesOrderLine):
			model.save()
			return salesOrderLine;
	
	def save(self, salesOrderLine):
		salesOrderLine.save()
		return salesOrderLine;
	
	def delete(self, salesOrderLineId ):
		errMsg = "Failed to delete SalesOrderLine from db using id " + str(salesOrderLineId)
		
		try:
			salesOrderLine = SalesOrderLine.objects.get(id=salesOrderLineId)
			salesOrderLine.delete()
			return True
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError("SalesOrderLine with id " + str(salesOrderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SalesOrderLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SalesOrderLine from db")
		except Exception:
			return None;
		
	def assignSalesOrder( self, salesOrderLineId, salesOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SalesOrderDelegate import SalesOrderDelegate

		errMsg = "Failed to assign element " + str(salesOrderId) + " for SalesOrder on SalesOrderLine"

		try:
			# get the SalesOrderLine from db
			salesOrderLine = self.get( salesOrderLineId ).first()	
			
			# get the SalesOrder from db
			salesOrder = SalesOrderDelegate().get(salesOrderId).first();
			
			# assign the SalesOrder		
			salesOrderLine.salesOrder = salesOrder
			
			#save it
			salesOrderLine.save()

			# reload and return the appropriate version					
			return self.get( salesOrderLineId );
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrderLine with id " + str(salesOrderLineId) + " does not exist.")
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSalesOrder( self, salesOrderLineId ):
		errMsg = "Failed to unassign element " + str(salesOrderId) + " for SalesOrder on SalesOrderLine"

		try:
			# get the SalesOrderLine from db
			salesOrderLine = self.get( salesOrderLineId ).first()	
			
			# assign to None for unassignment
			salesOrderLine.salesOrder = None			

			#save it
			salesOrderLine.save()

			# reload and return the appropriate version					
			return self.get( salesOrderLineId );
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrderLine with id " + str(salesOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignItem( self, salesOrderLineId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on SalesOrderLine"

		try:
			# get the SalesOrderLine from db
			salesOrderLine = self.get( salesOrderLineId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			salesOrderLine.item = item
			
			#save it
			salesOrderLine.save()

			# reload and return the appropriate version					
			return self.get( salesOrderLineId );
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrderLine with id " + str(salesOrderLineId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, salesOrderLineId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on SalesOrderLine"

		try:
			# get the SalesOrderLine from db
			salesOrderLine = self.get( salesOrderLineId ).first()	
			
			# assign to None for unassignment
			salesOrderLine.item = None			

			#save it
			salesOrderLine.save()

			# reload and return the appropriate version					
			return self.get( salesOrderLineId );
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrderLine with id " + str(salesOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
