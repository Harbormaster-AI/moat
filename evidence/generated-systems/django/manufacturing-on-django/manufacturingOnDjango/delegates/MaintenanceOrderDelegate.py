from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.MaintenanceOrder import MaintenanceOrder
from manufacturingOnDjango.models.Asset import Asset
from manufacturingOnDjango.models.MaintenancePlan import MaintenancePlan
from manufacturingOnDjango.models.WorkCenter import WorkCenter
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MaintenanceOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceOrderDelegate Declaration
#======================================================================
class MaintenanceOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, maintenanceOrderId ):
		try:	
			maintenanceOrder = MaintenanceOrder.objects.filter(id=maintenanceOrderId)
			return maintenanceOrder.first();
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError("MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, maintenanceOrder):
		for model in serializers.deserialize("json", maintenanceOrder):
			model.save()
			return model;

	def create(self, maintenanceOrder):
		maintenanceOrder.save()
		return maintenanceOrder;

	def saveFromJson(self, maintenanceOrder):
		for model in serializers.deserialize("json", maintenanceOrder):
			model.save()
			return maintenanceOrder;
	
	def save(self, maintenanceOrder):
		maintenanceOrder.save()
		return maintenanceOrder;
	
	def delete(self, maintenanceOrderId ):
		errMsg = "Failed to delete MaintenanceOrder from db using id " + str(maintenanceOrderId)
		
		try:
			maintenanceOrder = MaintenanceOrder.objects.get(id=maintenanceOrderId)
			maintenanceOrder.delete()
			return True
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError("MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MaintenanceOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MaintenanceOrder from db")
		except Exception:
			return None;
		
	def assignAsset( self, maintenanceOrderId, assetId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.AssetDelegate import AssetDelegate

		errMsg = "Failed to assign element " + str(assetId) + " for Asset on MaintenanceOrder"

		try:
			# get the MaintenanceOrder from db
			maintenanceOrder = self.get( maintenanceOrderId ).first()	
			
			# get the Asset from db
			asset = AssetDelegate().get(assetId).first();
			
			# assign the Asset		
			maintenanceOrder.asset = asset
			
			#save it
			maintenanceOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceOrderId );
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except Asset.DoesNotExist:
			raise ProcessingError(errMsg + " : Asset with id " + str(assetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAsset( self, maintenanceOrderId ):
		errMsg = "Failed to unassign element " + str(assetId) + " for Asset on MaintenanceOrder"

		try:
			# get the MaintenanceOrder from db
			maintenanceOrder = self.get( maintenanceOrderId ).first()	
			
			# assign to None for unassignment
			maintenanceOrder.asset = None			

			#save it
			maintenanceOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceOrderId );
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlan( self, maintenanceOrderId, planId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MaintenancePlanDelegate import MaintenancePlanDelegate

		errMsg = "Failed to assign element " + str(planId) + " for Plan on MaintenanceOrder"

		try:
			# get the MaintenanceOrder from db
			maintenanceOrder = self.get( maintenanceOrderId ).first()	
			
			# get the MaintenancePlan from db
			maintenancePlan = MaintenancePlanDelegate().get(planId).first();
			
			# assign the Plan		
			maintenanceOrder.plan = maintenancePlan
			
			#save it
			maintenanceOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceOrderId );
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except MaintenancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenancePlan with id " + str(planId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlan( self, maintenanceOrderId ):
		errMsg = "Failed to unassign element " + str(planId) + " for Plan on MaintenanceOrder"

		try:
			# get the MaintenanceOrder from db
			maintenanceOrder = self.get( maintenanceOrderId ).first()	
			
			# assign to None for unassignment
			maintenanceOrder.maintenancePlan = None			

			#save it
			maintenanceOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceOrderId );
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkCenter( self, maintenanceOrderId, workCenterId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to assign element " + str(workCenterId) + " for WorkCenter on MaintenanceOrder"

		try:
			# get the MaintenanceOrder from db
			maintenanceOrder = self.get( maintenanceOrderId ).first()	
			
			# get the WorkCenter from db
			workCenter = WorkCenterDelegate().get(workCenterId).first();
			
			# assign the WorkCenter		
			maintenanceOrder.workCenter = workCenter
			
			#save it
			maintenanceOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceOrderId );
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter with id " + str(workCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkCenter( self, maintenanceOrderId ):
		errMsg = "Failed to unassign element " + str(workCenterId) + " for WorkCenter on MaintenanceOrder"

		try:
			# get the MaintenanceOrder from db
			maintenanceOrder = self.get( maintenanceOrderId ).first()	
			
			# assign to None for unassignment
			maintenanceOrder.workCenter = None			

			#save it
			maintenanceOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceOrderId );
		except MaintenanceOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceOrder with id " + str(maintenanceOrderId) + " does not exist.")
		except Exception:
			return None;
		
