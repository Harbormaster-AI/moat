from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.InspectionLot import InspectionLot
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.WorkOrder import WorkOrder
from manufacturingOnDjango.models.GoodsReceipt import GoodsReceipt
from manufacturingOnDjango.models.InspectionResult import InspectionResult
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InspectionLot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionLotDelegate Declaration
#======================================================================
class InspectionLotDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inspectionLotId ):
		try:	
			inspectionLot = InspectionLot.objects.filter(id=inspectionLotId)
			return inspectionLot.first();
		except InspectionLot.DoesNotExist:
			raise ProcessingError("InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inspectionLot):
		for model in serializers.deserialize("json", inspectionLot):
			model.save()
			return model;

	def create(self, inspectionLot):
		inspectionLot.save()
		return inspectionLot;

	def saveFromJson(self, inspectionLot):
		for model in serializers.deserialize("json", inspectionLot):
			model.save()
			return inspectionLot;
	
	def save(self, inspectionLot):
		inspectionLot.save()
		return inspectionLot;
	
	def delete(self, inspectionLotId ):
		errMsg = "Failed to delete InspectionLot from db using id " + str(inspectionLotId)
		
		try:
			inspectionLot = InspectionLot.objects.get(id=inspectionLotId)
			inspectionLot.delete()
			return True
		except InspectionLot.DoesNotExist:
			raise ProcessingError("InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InspectionLot.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InspectionLot from db")
		except Exception:
			return None;
		
	def assignItem( self, inspectionLotId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on InspectionLot"

		try:
			# get the InspectionLot from db
			inspectionLot = self.get( inspectionLotId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			inspectionLot.item = item
			
			#save it
			inspectionLot.save()

			# reload and return the appropriate version					
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, inspectionLotId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on InspectionLot"

		try:
			# get the InspectionLot from db
			inspectionLot = self.get( inspectionLotId ).first()	
			
			# assign to None for unassignment
			inspectionLot.item = None			

			#save it
			inspectionLot.save()

			# reload and return the appropriate version					
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkOrder( self, inspectionLotId, workOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

		errMsg = "Failed to assign element " + str(workOrderId) + " for WorkOrder on InspectionLot"

		try:
			# get the InspectionLot from db
			inspectionLot = self.get( inspectionLotId ).first()	
			
			# get the WorkOrder from db
			workOrder = WorkOrderDelegate().get(workOrderId).first();
			
			# assign the WorkOrder		
			inspectionLot.workOrder = workOrder
			
			#save it
			inspectionLot.save()

			# reload and return the appropriate version					
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkOrder( self, inspectionLotId ):
		errMsg = "Failed to unassign element " + str(workOrderId) + " for WorkOrder on InspectionLot"

		try:
			# get the InspectionLot from db
			inspectionLot = self.get( inspectionLotId ).first()	
			
			# assign to None for unassignment
			inspectionLot.workOrder = None			

			#save it
			inspectionLot.save()

			# reload and return the appropriate version					
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except Exception:
			return None;
		
	def assignGoodsReceipt( self, inspectionLotId, goodsReceiptId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.GoodsReceiptDelegate import GoodsReceiptDelegate

		errMsg = "Failed to assign element " + str(goodsReceiptId) + " for GoodsReceipt on InspectionLot"

		try:
			# get the InspectionLot from db
			inspectionLot = self.get( inspectionLotId ).first()	
			
			# get the GoodsReceipt from db
			goodsReceipt = GoodsReceiptDelegate().get(goodsReceiptId).first();
			
			# assign the GoodsReceipt		
			inspectionLot.goodsReceipt = goodsReceipt
			
			#save it
			inspectionLot.save()

			# reload and return the appropriate version					
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except GoodsReceipt.DoesNotExist:
			raise ProcessingError(errMsg + " : GoodsReceipt with id " + str(goodsReceiptId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignGoodsReceipt( self, inspectionLotId ):
		errMsg = "Failed to unassign element " + str(goodsReceiptId) + " for GoodsReceipt on InspectionLot"

		try:
			# get the InspectionLot from db
			inspectionLot = self.get( inspectionLotId ).first()	
			
			# assign to None for unassignment
			inspectionLot.goodsReceipt = None			

			#save it
			inspectionLot.save()

			# reload and return the appropriate version					
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except Exception:
			return None;
		
	def addResults( self, inspectionLotId, resultsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionResultDelegate import InspectionResultDelegate

		errMsg = "Failed to add elements " + str(resultsIds) + " for Results on InspectionLot"

		try:
			# get the InspectionLot
			inspectionLot = self.get( inspectionLotId ).first()
				
			# split on a comma with no spaces
			idList = resultsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InspectionResult		
				inspectionResult = InspectionResultDelegate().get(id).first();	
				# add the InspectionResult
				inspectionLot.results.add(inspectionResult)
				
			# save it		
			inspectionLot.save()
			
			# reload and return the appropriate version
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except InspectionResult.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionResult does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeResults( self, inspectionLotId, resultsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionResultDelegate import InspectionResultDelegate

		errMsg = "Failed to remove elements " + str(resultsIds) + " for Results on InspectionLot"

		try:
			# get the InspectionLot
			inspectionLot = self.get( inspectionLotId ).first()
				
			# split on a comma with no spaces
			idList = resultsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InspectionResult		
				inspectionResult = InspectionResultDelegate().get(id).first();	
				# add the InspectionResult
				inspectionLot.results.remove(inspectionResult)
				
			# save it		
			inspectionLot.save()
			
			# reload and return the appropriate version
			return self.get( inspectionLotId );
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except InspectionResult.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionResult does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
