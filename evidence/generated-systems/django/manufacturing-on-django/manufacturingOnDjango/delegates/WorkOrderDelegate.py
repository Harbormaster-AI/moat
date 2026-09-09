from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.WorkOrder import WorkOrder
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.Routing import Routing
from manufacturingOnDjango.models.BOM import BOM
from manufacturingOnDjango.models.ProductionSchedule import ProductionSchedule
from manufacturingOnDjango.models.SalesOrder import SalesOrder
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model WorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkOrderDelegate Declaration
#======================================================================
class WorkOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, workOrderId ):
		try:	
			workOrder = WorkOrder.objects.filter(id=workOrderId)
			return workOrder.first();
		except WorkOrder.DoesNotExist:
			raise ProcessingError("WorkOrder with id " + str(workOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, workOrder):
		for model in serializers.deserialize("json", workOrder):
			model.save()
			return model;

	def create(self, workOrder):
		workOrder.save()
		return workOrder;

	def saveFromJson(self, workOrder):
		for model in serializers.deserialize("json", workOrder):
			model.save()
			return workOrder;
	
	def save(self, workOrder):
		workOrder.save()
		return workOrder;
	
	def delete(self, workOrderId ):
		errMsg = "Failed to delete WorkOrder from db using id " + str(workOrderId)
		
		try:
			workOrder = WorkOrder.objects.get(id=workOrderId)
			workOrder.delete()
			return True
		except WorkOrder.DoesNotExist:
			raise ProcessingError("WorkOrder with id " + str(workOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = WorkOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all WorkOrder from db")
		except Exception:
			return None;
		
	def assignItem( self, workOrderId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			workOrder.item = item
			
			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, workOrderId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# assign to None for unassignment
			workOrder.item = None			

			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlant( self, workOrderId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			workOrder.plant = plant
			
			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, workOrderId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# assign to None for unassignment
			workOrder.plant = None			

			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRouting( self, workOrderId, routingId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.RoutingDelegate import RoutingDelegate

		errMsg = "Failed to assign element " + str(routingId) + " for Routing on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# get the Routing from db
			routing = RoutingDelegate().get(routingId).first();
			
			# assign the Routing		
			workOrder.routing = routing
			
			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing with id " + str(routingId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRouting( self, workOrderId ):
		errMsg = "Failed to unassign element " + str(routingId) + " for Routing on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# assign to None for unassignment
			workOrder.routing = None			

			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignBom( self, workOrderId, bomId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BOMDelegate import BOMDelegate

		errMsg = "Failed to assign element " + str(bomId) + " for Bom on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# get the BOM from db
			bOM = BOMDelegate().get(bomId).first();
			
			# assign the Bom		
			workOrder.bom = bOM
			
			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM with id " + str(bomId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBom( self, workOrderId ):
		errMsg = "Failed to unassign element " + str(bomId) + " for Bom on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# assign to None for unassignment
			workOrder.bOM = None			

			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProductionSchedule( self, workOrderId, productionScheduleId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ProductionScheduleDelegate import ProductionScheduleDelegate

		errMsg = "Failed to assign element " + str(productionScheduleId) + " for ProductionSchedule on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# get the ProductionSchedule from db
			productionSchedule = ProductionScheduleDelegate().get(productionScheduleId).first();
			
			# assign the ProductionSchedule		
			workOrder.productionSchedule = productionSchedule
			
			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionSchedule with id " + str(productionScheduleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProductionSchedule( self, workOrderId ):
		errMsg = "Failed to unassign element " + str(productionScheduleId) + " for ProductionSchedule on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# assign to None for unassignment
			workOrder.productionSchedule = None			

			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSalesOrder( self, workOrderId, salesOrderId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SalesOrderDelegate import SalesOrderDelegate

		errMsg = "Failed to assign element " + str(salesOrderId) + " for SalesOrder on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# get the SalesOrder from db
			salesOrder = SalesOrderDelegate().get(salesOrderId).first();
			
			# assign the SalesOrder		
			workOrder.salesOrder = salesOrder
			
			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSalesOrder( self, workOrderId ):
		errMsg = "Failed to unassign element " + str(salesOrderId) + " for SalesOrder on WorkOrder"

		try:
			# get the WorkOrder from db
			workOrder = self.get( workOrderId ).first()	
			
			# assign to None for unassignment
			workOrder.salesOrder = None			

			#save it
			workOrder.save()

			# reload and return the appropriate version					
			return self.get( workOrderId );
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
		
