from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Warranty import Warranty
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Warranty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarrantyDelegate Declaration
#======================================================================
class WarrantyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, warrantyId ):
		try:	
			warranty = Warranty.objects.filter(id=warrantyId)
			return warranty.first();
		except Warranty.DoesNotExist:
			raise ProcessingError("Warranty with id " + str(warrantyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, warranty):
		for model in serializers.deserialize("json", warranty):
			model.save()
			return model;

	def create(self, warranty):
		warranty.save()
		return warranty;

	def saveFromJson(self, warranty):
		for model in serializers.deserialize("json", warranty):
			model.save()
			return warranty;
	
	def save(self, warranty):
		warranty.save()
		return warranty;
	
	def delete(self, warrantyId ):
		errMsg = "Failed to delete Warranty from db using id " + str(warrantyId)
		
		try:
			warranty = Warranty.objects.get(id=warrantyId)
			warranty.delete()
			return True
		except Warranty.DoesNotExist:
			raise ProcessingError("Warranty with id " + str(warrantyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Warranty.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Warranty from db")
		except Exception:
			return None;
		
	def assignAircraft( self, warrantyId, aircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to assign element " + str(aircraftId) + " for Aircraft on Warranty"

		try:
			# get the Warranty from db
			warranty = self.get( warrantyId ).first()	
			
			# get the Aircraft from db
			aircraft = AircraftDelegate().get(aircraftId).first();
			
			# assign the Aircraft		
			warranty.aircraft = aircraft
			
			#save it
			warranty.save()

			# reload and return the appropriate version					
			return self.get( warrantyId );
		except Warranty.DoesNotExist:
			raise ProcessingError(errMsg + " : Warranty with id " + str(warrantyId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraft( self, warrantyId ):
		errMsg = "Failed to unassign element " + str(aircraftId) + " for Aircraft on Warranty"

		try:
			# get the Warranty from db
			warranty = self.get( warrantyId ).first()	
			
			# assign to None for unassignment
			warranty.aircraft = None			

			#save it
			warranty.save()

			# reload and return the appropriate version					
			return self.get( warrantyId );
		except Warranty.DoesNotExist:
			raise ProcessingError(errMsg + " : Warranty with id " + str(warrantyId) + " does not exist.")
		except Exception:
			return None;
		
