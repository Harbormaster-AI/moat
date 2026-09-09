from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.CarePlan import CarePlan
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.Encounter import Encounter
from healthcareOnDjango.models.CareTask import CareTask
from healthcareOnDjango.models.CareTeam import CareTeam
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CarePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarePlanDelegate Declaration
#======================================================================
class CarePlanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, carePlanId ):
		try:	
			carePlan = CarePlan.objects.filter(id=carePlanId)
			return carePlan.first();
		except CarePlan.DoesNotExist:
			raise ProcessingError("CarePlan with id " + str(carePlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, carePlan):
		for model in serializers.deserialize("json", carePlan):
			model.save()
			return model;

	def create(self, carePlan):
		carePlan.save()
		return carePlan;

	def saveFromJson(self, carePlan):
		for model in serializers.deserialize("json", carePlan):
			model.save()
			return carePlan;
	
	def save(self, carePlan):
		carePlan.save()
		return carePlan;
	
	def delete(self, carePlanId ):
		errMsg = "Failed to delete CarePlan from db using id " + str(carePlanId)
		
		try:
			carePlan = CarePlan.objects.get(id=carePlanId)
			carePlan.delete()
			return True
		except CarePlan.DoesNotExist:
			raise ProcessingError("CarePlan with id " + str(carePlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CarePlan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CarePlan from db")
		except Exception:
			return None;
		
	def assignPatient( self, carePlanId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on CarePlan"

		try:
			# get the CarePlan from db
			carePlan = self.get( carePlanId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			carePlan.patient = patient
			
			#save it
			carePlan.save()

			# reload and return the appropriate version					
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, carePlanId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on CarePlan"

		try:
			# get the CarePlan from db
			carePlan = self.get( carePlanId ).first()	
			
			# assign to None for unassignment
			carePlan.patient = None			

			#save it
			carePlan.save()

			# reload and return the appropriate version					
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCareTeam( self, carePlanId, careTeamId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

		errMsg = "Failed to assign element " + str(careTeamId) + " for CareTeam on CarePlan"

		try:
			# get the CarePlan from db
			carePlan = self.get( carePlanId ).first()	
			
			# get the CareTeam from db
			careTeam = CareTeamDelegate().get(careTeamId).first();
			
			# assign the CareTeam		
			carePlan.careTeam = careTeam
			
			#save it
			carePlan.save()

			# reload and return the appropriate version					
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam with id " + str(careTeamId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCareTeam( self, carePlanId ):
		errMsg = "Failed to unassign element " + str(careTeamId) + " for CareTeam on CarePlan"

		try:
			# get the CarePlan from db
			carePlan = self.get( carePlanId ).first()	
			
			# assign to None for unassignment
			carePlan.careTeam = None			

			#save it
			carePlan.save()

			# reload and return the appropriate version					
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except Exception:
			return None;
		
	def addEncounters( self, carePlanId, encountersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to add elements " + str(encountersIds) + " for Encounters on CarePlan"

		try:
			# get the CarePlan
			carePlan = self.get( carePlanId ).first()
				
			# split on a comma with no spaces
			idList = encountersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Encounter		
				encounter = EncounterDelegate().get(id).first();	
				# add the Encounter
				carePlan.encounters.add(encounter)
				
			# save it		
			carePlan.save()
			
			# reload and return the appropriate version
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEncounters( self, carePlanId, encountersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.EncounterDelegate import EncounterDelegate

		errMsg = "Failed to remove elements " + str(encountersIds) + " for Encounters on CarePlan"

		try:
			# get the CarePlan
			carePlan = self.get( carePlanId ).first()
				
			# split on a comma with no spaces
			idList = encountersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Encounter		
				encounter = EncounterDelegate().get(id).first();	
				# add the Encounter
				carePlan.encounters.remove(encounter)
				
			# save it		
			carePlan.save()
			
			# reload and return the appropriate version
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except Encounter.DoesNotExist:
			raise ProcessingError(errMsg + " : Encounter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTasks( self, carePlanId, tasksIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTaskDelegate import CareTaskDelegate

		errMsg = "Failed to add elements " + str(tasksIds) + " for Tasks on CarePlan"

		try:
			# get the CarePlan
			carePlan = self.get( carePlanId ).first()
				
			# split on a comma with no spaces
			idList = tasksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CareTask		
				careTask = CareTaskDelegate().get(id).first();	
				# add the CareTask
				carePlan.tasks.add(careTask)
				
			# save it		
			carePlan.save()
			
			# reload and return the appropriate version
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTasks( self, carePlanId, tasksIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTaskDelegate import CareTaskDelegate

		errMsg = "Failed to remove elements " + str(tasksIds) + " for Tasks on CarePlan"

		try:
			# get the CarePlan
			carePlan = self.get( carePlanId ).first()
				
			# split on a comma with no spaces
			idList = tasksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CareTask		
				careTask = CareTaskDelegate().get(id).first();	
				# add the CareTask
				carePlan.tasks.remove(careTask)
				
			# save it		
			carePlan.save()
			
			# reload and return the appropriate version
			return self.get( carePlanId );
		except CarePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : CarePlan with id " + str(carePlanId) + " does not exist.")
		except CareTask.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTask does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
