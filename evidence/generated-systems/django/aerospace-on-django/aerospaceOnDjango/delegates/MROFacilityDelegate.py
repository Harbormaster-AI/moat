from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.MROFacility import MROFacility
from aerospaceOnDjango.models.MaintenanceAppointment import MaintenanceAppointment
from aerospaceOnDjango.models.MaintenanceWorkOrder import MaintenanceWorkOrder
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MROFacility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MROFacilityDelegate Declaration
#======================================================================
class MROFacilityDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, mROFacilityId ):
		try:	
			mROFacility = MROFacility.objects.filter(id=mROFacilityId)
			return mROFacility.first();
		except MROFacility.DoesNotExist:
			raise ProcessingError("MROFacility with id " + str(mROFacilityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, mROFacility):
		for model in serializers.deserialize("json", mROFacility):
			model.save()
			return model;

	def create(self, mROFacility):
		mROFacility.save()
		return mROFacility;

	def saveFromJson(self, mROFacility):
		for model in serializers.deserialize("json", mROFacility):
			model.save()
			return mROFacility;
	
	def save(self, mROFacility):
		mROFacility.save()
		return mROFacility;
	
	def delete(self, mROFacilityId ):
		errMsg = "Failed to delete MROFacility from db using id " + str(mROFacilityId)
		
		try:
			mROFacility = MROFacility.objects.get(id=mROFacilityId)
			mROFacility.delete()
			return True
		except MROFacility.DoesNotExist:
			raise ProcessingError("MROFacility with id " + str(mROFacilityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MROFacility.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MROFacility from db")
		except Exception:
			return None;
		
	def addAppointments( self, mROFacilityId, appointmentsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceAppointmentDelegate import MaintenanceAppointmentDelegate

		errMsg = "Failed to add elements " + str(appointmentsIds) + " for Appointments on MROFacility"

		try:
			# get the MROFacility
			mROFacility = self.get( mROFacilityId ).first()
				
			# split on a comma with no spaces
			idList = appointmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenanceAppointment		
				maintenanceAppointment = MaintenanceAppointmentDelegate().get(id).first();	
				# add the MaintenanceAppointment
				mROFacility.appointments.add(maintenanceAppointment)
				
			# save it		
			mROFacility.save()
			
			# reload and return the appropriate version
			return self.get( mROFacilityId );
		except MROFacility.DoesNotExist:
			raise ProcessingError(errMsg + " : MROFacility with id " + str(mROFacilityId) + " does not exist.")
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAppointments( self, mROFacilityId, appointmentsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceAppointmentDelegate import MaintenanceAppointmentDelegate

		errMsg = "Failed to remove elements " + str(appointmentsIds) + " for Appointments on MROFacility"

		try:
			# get the MROFacility
			mROFacility = self.get( mROFacilityId ).first()
				
			# split on a comma with no spaces
			idList = appointmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenanceAppointment		
				maintenanceAppointment = MaintenanceAppointmentDelegate().get(id).first();	
				# add the MaintenanceAppointment
				mROFacility.appointments.remove(maintenanceAppointment)
				
			# save it		
			mROFacility.save()
			
			# reload and return the appropriate version
			return self.get( mROFacilityId );
		except MROFacility.DoesNotExist:
			raise ProcessingError(errMsg + " : MROFacility with id " + str(mROFacilityId) + " does not exist.")
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWorkOrders( self, mROFacilityId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to add elements " + str(workOrdersIds) + " for WorkOrders on MROFacility"

		try:
			# get the MROFacility
			mROFacility = self.get( mROFacilityId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				mROFacility.workOrders.add(maintenanceWorkOrder)
				
			# save it		
			mROFacility.save()
			
			# reload and return the appropriate version
			return self.get( mROFacilityId );
		except MROFacility.DoesNotExist:
			raise ProcessingError(errMsg + " : MROFacility with id " + str(mROFacilityId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkOrders( self, mROFacilityId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to remove elements " + str(workOrdersIds) + " for WorkOrders on MROFacility"

		try:
			# get the MROFacility
			mROFacility = self.get( mROFacilityId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				mROFacility.workOrders.remove(maintenanceWorkOrder)
				
			# save it		
			mROFacility.save()
			
			# reload and return the appropriate version
			return self.get( mROFacilityId );
		except MROFacility.DoesNotExist:
			raise ProcessingError(errMsg + " : MROFacility with id " + str(mROFacilityId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
