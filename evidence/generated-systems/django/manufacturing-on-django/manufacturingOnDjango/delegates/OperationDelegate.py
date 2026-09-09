from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Operation import Operation
from manufacturingOnDjango.models.Routing import Routing
from manufacturingOnDjango.models.WorkCenter import WorkCenter
from manufacturingOnDjango.models.InspectionPlan import InspectionPlan
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Operation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OperationDelegate Declaration
#======================================================================
class OperationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, operationId ):
		try:	
			operation = Operation.objects.filter(id=operationId)
			return operation.first();
		except Operation.DoesNotExist:
			raise ProcessingError("Operation with id " + str(operationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, operation):
		for model in serializers.deserialize("json", operation):
			model.save()
			return model;

	def create(self, operation):
		operation.save()
		return operation;

	def saveFromJson(self, operation):
		for model in serializers.deserialize("json", operation):
			model.save()
			return operation;
	
	def save(self, operation):
		operation.save()
		return operation;
	
	def delete(self, operationId ):
		errMsg = "Failed to delete Operation from db using id " + str(operationId)
		
		try:
			operation = Operation.objects.get(id=operationId)
			operation.delete()
			return True
		except Operation.DoesNotExist:
			raise ProcessingError("Operation with id " + str(operationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Operation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Operation from db")
		except Exception:
			return None;
		
	def assignRouting( self, operationId, routingId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.RoutingDelegate import RoutingDelegate

		errMsg = "Failed to assign element " + str(routingId) + " for Routing on Operation"

		try:
			# get the Operation from db
			operation = self.get( operationId ).first()	
			
			# get the Routing from db
			routing = RoutingDelegate().get(routingId).first();
			
			# assign the Routing		
			operation.routing = routing
			
			#save it
			operation.save()

			# reload and return the appropriate version					
			return self.get( operationId );
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation with id " + str(operationId) + " does not exist.")
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing with id " + str(routingId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRouting( self, operationId ):
		errMsg = "Failed to unassign element " + str(routingId) + " for Routing on Operation"

		try:
			# get the Operation from db
			operation = self.get( operationId ).first()	
			
			# assign to None for unassignment
			operation.routing = None			

			#save it
			operation.save()

			# reload and return the appropriate version					
			return self.get( operationId );
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation with id " + str(operationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkCenter( self, operationId, workCenterId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to assign element " + str(workCenterId) + " for WorkCenter on Operation"

		try:
			# get the Operation from db
			operation = self.get( operationId ).first()	
			
			# get the WorkCenter from db
			workCenter = WorkCenterDelegate().get(workCenterId).first();
			
			# assign the WorkCenter		
			operation.workCenter = workCenter
			
			#save it
			operation.save()

			# reload and return the appropriate version					
			return self.get( operationId );
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation with id " + str(operationId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter with id " + str(workCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkCenter( self, operationId ):
		errMsg = "Failed to unassign element " + str(workCenterId) + " for WorkCenter on Operation"

		try:
			# get the Operation from db
			operation = self.get( operationId ).first()	
			
			# assign to None for unassignment
			operation.workCenter = None			

			#save it
			operation.save()

			# reload and return the appropriate version					
			return self.get( operationId );
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation with id " + str(operationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInspectionPlan( self, operationId, inspectionPlanId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionPlanDelegate import InspectionPlanDelegate

		errMsg = "Failed to assign element " + str(inspectionPlanId) + " for InspectionPlan on Operation"

		try:
			# get the Operation from db
			operation = self.get( operationId ).first()	
			
			# get the InspectionPlan from db
			inspectionPlan = InspectionPlanDelegate().get(inspectionPlanId).first();
			
			# assign the InspectionPlan		
			operation.inspectionPlan = inspectionPlan
			
			#save it
			operation.save()

			# reload and return the appropriate version					
			return self.get( operationId );
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation with id " + str(operationId) + " does not exist.")
		except InspectionPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInspectionPlan( self, operationId ):
		errMsg = "Failed to unassign element " + str(inspectionPlanId) + " for InspectionPlan on Operation"

		try:
			# get the Operation from db
			operation = self.get( operationId ).first()	
			
			# assign to None for unassignment
			operation.inspectionPlan = None			

			#save it
			operation.save()

			# reload and return the appropriate version					
			return self.get( operationId );
		except Operation.DoesNotExist:
			raise ProcessingError(errMsg + " : Operation with id " + str(operationId) + " does not exist.")
		except Exception:
			return None;
		
