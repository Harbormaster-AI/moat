from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.WorkShift import WorkShift
from hrOnDjango.models.WorkSchedule import WorkSchedule
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model WorkShift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkShiftDelegate Declaration
#======================================================================
class WorkShiftDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, workShiftId ):
		try:	
			workShift = WorkShift.objects.filter(id=workShiftId)
			return workShift.first();
		except WorkShift.DoesNotExist:
			raise ProcessingError("WorkShift with id " + str(workShiftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, workShift):
		for model in serializers.deserialize("json", workShift):
			model.save()
			return model;

	def create(self, workShift):
		workShift.save()
		return workShift;

	def saveFromJson(self, workShift):
		for model in serializers.deserialize("json", workShift):
			model.save()
			return workShift;
	
	def save(self, workShift):
		workShift.save()
		return workShift;
	
	def delete(self, workShiftId ):
		errMsg = "Failed to delete WorkShift from db using id " + str(workShiftId)
		
		try:
			workShift = WorkShift.objects.get(id=workShiftId)
			workShift.delete()
			return True
		except WorkShift.DoesNotExist:
			raise ProcessingError("WorkShift with id " + str(workShiftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = WorkShift.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all WorkShift from db")
		except Exception:
			return None;
		
	def assignWorkSchedule( self, workShiftId, workScheduleId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.WorkScheduleDelegate import WorkScheduleDelegate

		errMsg = "Failed to assign element " + str(workScheduleId) + " for WorkSchedule on WorkShift"

		try:
			# get the WorkShift from db
			workShift = self.get( workShiftId ).first()	
			
			# get the WorkSchedule from db
			workSchedule = WorkScheduleDelegate().get(workScheduleId).first();
			
			# assign the WorkSchedule		
			workShift.workSchedule = workSchedule
			
			#save it
			workShift.save()

			# reload and return the appropriate version					
			return self.get( workShiftId );
		except WorkShift.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkShift with id " + str(workShiftId) + " does not exist.")
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkSchedule( self, workShiftId ):
		errMsg = "Failed to unassign element " + str(workScheduleId) + " for WorkSchedule on WorkShift"

		try:
			# get the WorkShift from db
			workShift = self.get( workShiftId ).first()	
			
			# assign to None for unassignment
			workShift.workSchedule = None			

			#save it
			workShift.save()

			# reload and return the appropriate version					
			return self.get( workShiftId );
		except WorkShift.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkShift with id " + str(workShiftId) + " does not exist.")
		except Exception:
			return None;
		
