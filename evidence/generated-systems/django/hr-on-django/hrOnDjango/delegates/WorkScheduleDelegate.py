from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.WorkSchedule import WorkSchedule
from hrOnDjango.models.EmploymentContract import EmploymentContract
from hrOnDjango.models.WorkShift import WorkShift
from hrOnDjango.models.ScheduleException import ScheduleException
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model WorkSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkScheduleDelegate Declaration
#======================================================================
class WorkScheduleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, workScheduleId ):
		try:	
			workSchedule = WorkSchedule.objects.filter(id=workScheduleId)
			return workSchedule.first();
		except WorkSchedule.DoesNotExist:
			raise ProcessingError("WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, workSchedule):
		for model in serializers.deserialize("json", workSchedule):
			model.save()
			return model;

	def create(self, workSchedule):
		workSchedule.save()
		return workSchedule;

	def saveFromJson(self, workSchedule):
		for model in serializers.deserialize("json", workSchedule):
			model.save()
			return workSchedule;
	
	def save(self, workSchedule):
		workSchedule.save()
		return workSchedule;
	
	def delete(self, workScheduleId ):
		errMsg = "Failed to delete WorkSchedule from db using id " + str(workScheduleId)
		
		try:
			workSchedule = WorkSchedule.objects.get(id=workScheduleId)
			workSchedule.delete()
			return True
		except WorkSchedule.DoesNotExist:
			raise ProcessingError("WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = WorkSchedule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all WorkSchedule from db")
		except Exception:
			return None;
		
	def addContracts( self, workScheduleId, contractsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on WorkSchedule"

		try:
			# get the WorkSchedule
			workSchedule = self.get( workScheduleId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmploymentContract		
				employmentContract = EmploymentContractDelegate().get(id).first();	
				# add the EmploymentContract
				workSchedule.contracts.add(employmentContract)
				
			# save it		
			workSchedule.save()
			
			# reload and return the appropriate version
			return self.get( workScheduleId );
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, workScheduleId, contractsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on WorkSchedule"

		try:
			# get the WorkSchedule
			workSchedule = self.get( workScheduleId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmploymentContract		
				employmentContract = EmploymentContractDelegate().get(id).first();	
				# add the EmploymentContract
				workSchedule.contracts.remove(employmentContract)
				
			# save it		
			workSchedule.save()
			
			# reload and return the appropriate version
			return self.get( workScheduleId );
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addShifts( self, workScheduleId, shiftsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.WorkShiftDelegate import WorkShiftDelegate

		errMsg = "Failed to add elements " + str(shiftsIds) + " for Shifts on WorkSchedule"

		try:
			# get the WorkSchedule
			workSchedule = self.get( workScheduleId ).first()
				
			# split on a comma with no spaces
			idList = shiftsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WorkShift		
				workShift = WorkShiftDelegate().get(id).first();	
				# add the WorkShift
				workSchedule.shifts.add(workShift)
				
			# save it		
			workSchedule.save()
			
			# reload and return the appropriate version
			return self.get( workScheduleId );
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except WorkShift.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkShift does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeShifts( self, workScheduleId, shiftsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.WorkShiftDelegate import WorkShiftDelegate

		errMsg = "Failed to remove elements " + str(shiftsIds) + " for Shifts on WorkSchedule"

		try:
			# get the WorkSchedule
			workSchedule = self.get( workScheduleId ).first()
				
			# split on a comma with no spaces
			idList = shiftsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WorkShift		
				workShift = WorkShiftDelegate().get(id).first();	
				# add the WorkShift
				workSchedule.shifts.remove(workShift)
				
			# save it		
			workSchedule.save()
			
			# reload and return the appropriate version
			return self.get( workScheduleId );
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except WorkShift.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkShift does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addExceptions( self, workScheduleId, exceptionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ScheduleExceptionDelegate import ScheduleExceptionDelegate

		errMsg = "Failed to add elements " + str(exceptionsIds) + " for Exceptions on WorkSchedule"

		try:
			# get the WorkSchedule
			workSchedule = self.get( workScheduleId ).first()
				
			# split on a comma with no spaces
			idList = exceptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ScheduleException		
				scheduleException = ScheduleExceptionDelegate().get(id).first();	
				# add the ScheduleException
				workSchedule.exceptions.add(scheduleException)
				
			# save it		
			workSchedule.save()
			
			# reload and return the appropriate version
			return self.get( workScheduleId );
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except ScheduleException.DoesNotExist:
			raise ProcessingError(errMsg + " : ScheduleException does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeExceptions( self, workScheduleId, exceptionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ScheduleExceptionDelegate import ScheduleExceptionDelegate

		errMsg = "Failed to remove elements " + str(exceptionsIds) + " for Exceptions on WorkSchedule"

		try:
			# get the WorkSchedule
			workSchedule = self.get( workScheduleId ).first()
				
			# split on a comma with no spaces
			idList = exceptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ScheduleException		
				scheduleException = ScheduleExceptionDelegate().get(id).first();	
				# add the ScheduleException
				workSchedule.exceptions.remove(scheduleException)
				
			# save it		
			workSchedule.save()
			
			# reload and return the appropriate version
			return self.get( workScheduleId );
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except ScheduleException.DoesNotExist:
			raise ProcessingError(errMsg + " : ScheduleException does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
