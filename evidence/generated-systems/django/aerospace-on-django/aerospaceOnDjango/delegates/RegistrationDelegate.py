from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Registration import Registration
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Registration
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RegistrationDelegate Declaration
#======================================================================
class RegistrationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, registrationId ):
		try:	
			registration = Registration.objects.filter(id=registrationId)
			return registration.first();
		except Registration.DoesNotExist:
			raise ProcessingError("Registration with id " + str(registrationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, registration):
		for model in serializers.deserialize("json", registration):
			model.save()
			return model;

	def create(self, registration):
		registration.save()
		return registration;

	def saveFromJson(self, registration):
		for model in serializers.deserialize("json", registration):
			model.save()
			return registration;
	
	def save(self, registration):
		registration.save()
		return registration;
	
	def delete(self, registrationId ):
		errMsg = "Failed to delete Registration from db using id " + str(registrationId)
		
		try:
			registration = Registration.objects.get(id=registrationId)
			registration.delete()
			return True
		except Registration.DoesNotExist:
			raise ProcessingError("Registration with id " + str(registrationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Registration.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Registration from db")
		except Exception:
			return None;
		
	def assignAircraft( self, registrationId, aircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to assign element " + str(aircraftId) + " for Aircraft on Registration"

		try:
			# get the Registration from db
			registration = self.get( registrationId ).first()	
			
			# get the Aircraft from db
			aircraft = AircraftDelegate().get(aircraftId).first();
			
			# assign the Aircraft		
			registration.aircraft = aircraft
			
			#save it
			registration.save()

			# reload and return the appropriate version					
			return self.get( registrationId );
		except Registration.DoesNotExist:
			raise ProcessingError(errMsg + " : Registration with id " + str(registrationId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraft( self, registrationId ):
		errMsg = "Failed to unassign element " + str(aircraftId) + " for Aircraft on Registration"

		try:
			# get the Registration from db
			registration = self.get( registrationId ).first()	
			
			# assign to None for unassignment
			registration.aircraft = None			

			#save it
			registration.save()

			# reload and return the appropriate version					
			return self.get( registrationId );
		except Registration.DoesNotExist:
			raise ProcessingError(errMsg + " : Registration with id " + str(registrationId) + " does not exist.")
		except Exception:
			return None;
		
