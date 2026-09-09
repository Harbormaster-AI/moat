from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.MRPRun import MRPRun
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.PlannedOrder import PlannedOrder
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MRPRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MRPRunDelegate Declaration
#======================================================================
class MRPRunDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, mRPRunId ):
		try:	
			mRPRun = MRPRun.objects.filter(id=mRPRunId)
			return mRPRun.first();
		except MRPRun.DoesNotExist:
			raise ProcessingError("MRPRun with id " + str(mRPRunId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, mRPRun):
		for model in serializers.deserialize("json", mRPRun):
			model.save()
			return model;

	def create(self, mRPRun):
		mRPRun.save()
		return mRPRun;

	def saveFromJson(self, mRPRun):
		for model in serializers.deserialize("json", mRPRun):
			model.save()
			return mRPRun;
	
	def save(self, mRPRun):
		mRPRun.save()
		return mRPRun;
	
	def delete(self, mRPRunId ):
		errMsg = "Failed to delete MRPRun from db using id " + str(mRPRunId)
		
		try:
			mRPRun = MRPRun.objects.get(id=mRPRunId)
			mRPRun.delete()
			return True
		except MRPRun.DoesNotExist:
			raise ProcessingError("MRPRun with id " + str(mRPRunId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MRPRun.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MRPRun from db")
		except Exception:
			return None;
		
	def assignPlant( self, mRPRunId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on MRPRun"

		try:
			# get the MRPRun from db
			mRPRun = self.get( mRPRunId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			mRPRun.plant = plant
			
			#save it
			mRPRun.save()

			# reload and return the appropriate version					
			return self.get( mRPRunId );
		except MRPRun.DoesNotExist:
			raise ProcessingError(errMsg + " : MRPRun with id " + str(mRPRunId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, mRPRunId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on MRPRun"

		try:
			# get the MRPRun from db
			mRPRun = self.get( mRPRunId ).first()	
			
			# assign to None for unassignment
			mRPRun.plant = None			

			#save it
			mRPRun.save()

			# reload and return the appropriate version					
			return self.get( mRPRunId );
		except MRPRun.DoesNotExist:
			raise ProcessingError(errMsg + " : MRPRun with id " + str(mRPRunId) + " does not exist.")
		except Exception:
			return None;
		
	def addPlannedOrders( self, mRPRunId, plannedOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlannedOrderDelegate import PlannedOrderDelegate

		errMsg = "Failed to add elements " + str(plannedOrdersIds) + " for PlannedOrders on MRPRun"

		try:
			# get the MRPRun
			mRPRun = self.get( mRPRunId ).first()
				
			# split on a comma with no spaces
			idList = plannedOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PlannedOrder		
				plannedOrder = PlannedOrderDelegate().get(id).first();	
				# add the PlannedOrder
				mRPRun.plannedOrders.add(plannedOrder)
				
			# save it		
			mRPRun.save()
			
			# reload and return the appropriate version
			return self.get( mRPRunId );
		except MRPRun.DoesNotExist:
			raise ProcessingError(errMsg + " : MRPRun with id " + str(mRPRunId) + " does not exist.")
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlannedOrders( self, mRPRunId, plannedOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlannedOrderDelegate import PlannedOrderDelegate

		errMsg = "Failed to remove elements " + str(plannedOrdersIds) + " for PlannedOrders on MRPRun"

		try:
			# get the MRPRun
			mRPRun = self.get( mRPRunId ).first()
				
			# split on a comma with no spaces
			idList = plannedOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PlannedOrder		
				plannedOrder = PlannedOrderDelegate().get(id).first();	
				# add the PlannedOrder
				mRPRun.plannedOrders.remove(plannedOrder)
				
			# save it		
			mRPRun.save()
			
			# reload and return the appropriate version
			return self.get( mRPRunId );
		except MRPRun.DoesNotExist:
			raise ProcessingError(errMsg + " : MRPRun with id " + str(mRPRunId) + " does not exist.")
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
