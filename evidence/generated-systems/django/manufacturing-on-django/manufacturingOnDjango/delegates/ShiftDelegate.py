from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Shift import Shift
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.ShiftAssignment import ShiftAssignment
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Shift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftDelegate Declaration
#======================================================================
class ShiftDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, shiftId ):
		try:	
			shift = Shift.objects.filter(id=shiftId)
			return shift.first();
		except Shift.DoesNotExist:
			raise ProcessingError("Shift with id " + str(shiftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, shift):
		for model in serializers.deserialize("json", shift):
			model.save()
			return model;

	def create(self, shift):
		shift.save()
		return shift;

	def saveFromJson(self, shift):
		for model in serializers.deserialize("json", shift):
			model.save()
			return shift;
	
	def save(self, shift):
		shift.save()
		return shift;
	
	def delete(self, shiftId ):
		errMsg = "Failed to delete Shift from db using id " + str(shiftId)
		
		try:
			shift = Shift.objects.get(id=shiftId)
			shift.delete()
			return True
		except Shift.DoesNotExist:
			raise ProcessingError("Shift with id " + str(shiftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Shift.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Shift from db")
		except Exception:
			return None;
		
	def assignPlant( self, shiftId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on Shift"

		try:
			# get the Shift from db
			shift = self.get( shiftId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			shift.plant = plant
			
			#save it
			shift.save()

			# reload and return the appropriate version					
			return self.get( shiftId );
		except Shift.DoesNotExist:
			raise ProcessingError(errMsg + " : Shift with id " + str(shiftId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, shiftId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on Shift"

		try:
			# get the Shift from db
			shift = self.get( shiftId ).first()	
			
			# assign to None for unassignment
			shift.plant = None			

			#save it
			shift.save()

			# reload and return the appropriate version					
			return self.get( shiftId );
		except Shift.DoesNotExist:
			raise ProcessingError(errMsg + " : Shift with id " + str(shiftId) + " does not exist.")
		except Exception:
			return None;
		
	def addAssignments( self, shiftId, assignmentsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ShiftAssignmentDelegate import ShiftAssignmentDelegate

		errMsg = "Failed to add elements " + str(assignmentsIds) + " for Assignments on Shift"

		try:
			# get the Shift
			shift = self.get( shiftId ).first()
				
			# split on a comma with no spaces
			idList = assignmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ShiftAssignment		
				shiftAssignment = ShiftAssignmentDelegate().get(id).first();	
				# add the ShiftAssignment
				shift.assignments.add(shiftAssignment)
				
			# save it		
			shift.save()
			
			# reload and return the appropriate version
			return self.get( shiftId );
		except Shift.DoesNotExist:
			raise ProcessingError(errMsg + " : Shift with id " + str(shiftId) + " does not exist.")
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAssignments( self, shiftId, assignmentsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ShiftAssignmentDelegate import ShiftAssignmentDelegate

		errMsg = "Failed to remove elements " + str(assignmentsIds) + " for Assignments on Shift"

		try:
			# get the Shift
			shift = self.get( shiftId ).first()
				
			# split on a comma with no spaces
			idList = assignmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ShiftAssignment		
				shiftAssignment = ShiftAssignmentDelegate().get(id).first();	
				# add the ShiftAssignment
				shift.assignments.remove(shiftAssignment)
				
			# save it		
			shift.save()
			
			# reload and return the appropriate version
			return self.get( shiftId );
		except Shift.DoesNotExist:
			raise ProcessingError(errMsg + " : Shift with id " + str(shiftId) + " does not exist.")
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
