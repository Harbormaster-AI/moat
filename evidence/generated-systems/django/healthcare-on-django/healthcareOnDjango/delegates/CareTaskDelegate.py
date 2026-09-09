from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.CareTask import CareTask
from healthcareOnDjango.models.CarePlan import CarePlan
from healthcareOnDjango.models.Clinician import Clinician
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CareTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTaskDelegate Declaration
#======================================================================
class CareTaskDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, careTaskId ):
		try:	
			careTask = CareTask.objects.filter(id=careTaskId)
			return careTask.first();
		except CareTask.DoesNotExist:
			raise ProcessingError("CareTask with id " + str(careTaskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, careTask):
		for model in serializers.deserialize("json", careTask):
			model.save()
			return model;

	def create(self, careTask):
		careTask.save()
		return careTask;

	def saveFromJson(self, careTask):
		for model in serializers.deserialize("json", careTask):
			model.save()
			return careTask;
	
	def save(self, careTask):
		careTask.save()
		return careTask;
	
	def delete(self, careTaskId ):
		errMsg = "Failed to delete CareTask from db using id " + str(careTaskId)
		
		try:
			careTask = CareTask.objects.get(id=careTaskId)
			careTask.delete()
			return True
		except CareTask.DoesNotExist:
			raise ProcessingError("CareTask with id " + str(careTaskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CareTask.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CareTask from db")
		except Exception:
			return None;
		
	def assignCarePlan( self, careTaskId, carePlanId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CarePlanDelegate import CarePlanDelegate

		errMsg = "Failed to assign element " + str(carePlanId) + " for CarePlan on CareTask"

		try:
			# get the CareTask from db
			careTask = self.get( careTaskId ).first()	
			
			# get the CarePlan from db
			carePlan = CarePlanDelegate().get(carePlanId).first();
			
			# assign the CarePlan		
			careTask.carePlan = carePlan
			
			#save it
			careTask.save()

			# reload and return the appropriate version					
			return self.get( careTaskId );
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask with id " + str(careTaskId) + " does not exist.")
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCarePlan( self, careTaskId ):
		errMsg = "Failed to unassign element " + str(carePlanId) + " for CarePlan on CareTask"

		try:
			# get the CareTask from db
			careTask = self.get( careTaskId ).first()	
			
			# assign to None for unassignment
			careTask.carePlan = None			

			#save it
			careTask.save()

			# reload and return the appropriate version					
			return self.get( careTaskId );
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask with id " + str(careTaskId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAssignedTo( self, careTaskId, assignedToId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicianDelegate import ClinicianDelegate

		errMsg = "Failed to assign element " + str(assignedToId) + " for AssignedTo on CareTask"

		try:
			# get the CareTask from db
			careTask = self.get( careTaskId ).first()	
			
			# get the Clinician from db
			clinician = ClinicianDelegate().get(assignedToId).first();
			
			# assign the AssignedTo		
			careTask.assignedTo = clinician
			
			#save it
			careTask.save()

			# reload and return the appropriate version					
			return self.get( careTaskId );
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask with id " + str(careTaskId) + " does not exist.")
		except Clinician.DoesNotExist:
			raise ProcessingError(errMsg + " : Clinician with id " + str(assignedToId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAssignedTo( self, careTaskId ):
		errMsg = "Failed to unassign element " + str(assignedToId) + " for AssignedTo on CareTask"

		try:
			# get the CareTask from db
			careTask = self.get( careTaskId ).first()	
			
			# assign to None for unassignment
			careTask.clinician = None			

			#save it
			careTask.save()

			# reload and return the appropriate version					
			return self.get( careTaskId );
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask with id " + str(careTaskId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEncounter( self, careTaskId, encounterId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to assign element " + str(encounterId) + " for Encounter on CareTask"

		try:
			# get the CareTask from db
			careTask = self.get( careTaskId ).first()	
			
			# get the Encounter from db
			encounter = EncounterDelegate().get(encounterId).first();
			
			# assign the Encounter		
			careTask.encounter = encounter
			
			#save it
			careTask.save()

			# reload and return the appropriate version					
			return self.get( careTaskId );
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask with id " + str(careTaskId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter with id " + str(encounterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEncounter( self, careTaskId ):
		errMsg = "Failed to unassign element " + str(encounterId) + " for Encounter on CareTask"

		try:
			# get the CareTask from db
			careTask = self.get( careTaskId ).first()	
			
			# assign to None for unassignment
			careTask.encounter = None			

			#save it
			careTask.save()

			# reload and return the appropriate version					
			return self.get( careTaskId );
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask with id " + str(careTaskId) + " does not exist.")
		except Exception:
			return None;
		
