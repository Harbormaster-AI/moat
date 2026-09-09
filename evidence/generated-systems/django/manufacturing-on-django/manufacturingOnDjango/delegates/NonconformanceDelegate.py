from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Nonconformance import Nonconformance
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.WorkOrder import WorkOrder
from manufacturingOnDjango.models.InspectionLot import InspectionLot
from manufacturingOnDjango.models.CorrectiveAction import CorrectiveAction
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Nonconformance
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NonconformanceDelegate Declaration
#======================================================================
class NonconformanceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, nonconformanceId ):
		try:	
			nonconformance = Nonconformance.objects.filter(id=nonconformanceId)
			return nonconformance.first();
		except Nonconformance.DoesNotExist:
			raise ProcessingError("Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, nonconformance):
		for model in serializers.deserialize("json", nonconformance):
			model.save()
			return model;

	def create(self, nonconformance):
		nonconformance.save()
		return nonconformance;

	def saveFromJson(self, nonconformance):
		for model in serializers.deserialize("json", nonconformance):
			model.save()
			return nonconformance;
	
	def save(self, nonconformance):
		nonconformance.save()
		return nonconformance;
	
	def delete(self, nonconformanceId ):
		errMsg = "Failed to delete Nonconformance from db using id " + str(nonconformanceId)
		
		try:
			nonconformance = Nonconformance.objects.get(id=nonconformanceId)
			nonconformance.delete()
			return True
		except Nonconformance.DoesNotExist:
			raise ProcessingError("Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Nonconformance.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Nonconformance from db")
		except Exception:
			return None;
		
	def assignItem( self, nonconformanceId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			nonconformance.item = item
			
			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, nonconformanceId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# assign to None for unassignment
			nonconformance.item = None			

			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkOrder( self, nonconformanceId, workOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

		errMsg = "Failed to assign element " + str(workOrderId) + " for WorkOrder on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# get the WorkOrder from db
			workOrder = WorkOrderDelegate().get(workOrderId).first();
			
			# assign the WorkOrder		
			nonconformance.workOrder = workOrder
			
			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkOrder( self, nonconformanceId ):
		errMsg = "Failed to unassign element " + str(workOrderId) + " for WorkOrder on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# assign to None for unassignment
			nonconformance.workOrder = None			

			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInspectionLot( self, nonconformanceId, inspectionLotId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionLotDelegate import InspectionLotDelegate

		errMsg = "Failed to assign element " + str(inspectionLotId) + " for InspectionLot on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# get the InspectionLot from db
			inspectionLot = InspectionLotDelegate().get(inspectionLotId).first();
			
			# assign the InspectionLot		
			nonconformance.inspectionLot = inspectionLot
			
			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInspectionLot( self, nonconformanceId ):
		errMsg = "Failed to unassign element " + str(inspectionLotId) + " for InspectionLot on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# assign to None for unassignment
			nonconformance.inspectionLot = None			

			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCorrectiveAction( self, nonconformanceId, correctiveActionId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to assign element " + str(correctiveActionId) + " for CorrectiveAction on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# get the CorrectiveAction from db
			correctiveAction = CorrectiveActionDelegate().get(correctiveActionId).first();
			
			# assign the CorrectiveAction		
			nonconformance.correctiveAction = correctiveAction
			
			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCorrectiveAction( self, nonconformanceId ):
		errMsg = "Failed to unassign element " + str(correctiveActionId) + " for CorrectiveAction on Nonconformance"

		try:
			# get the Nonconformance from db
			nonconformance = self.get( nonconformanceId ).first()	
			
			# assign to None for unassignment
			nonconformance.correctiveAction = None			

			#save it
			nonconformance.save()

			# reload and return the appropriate version					
			return self.get( nonconformanceId );
		except Nonconformance.DoesNotExist:
			raise ProcessingError(errMsg + " : Nonconformance with id " + str(nonconformanceId) + " does not exist.")
		except Exception:
			return None;
		
