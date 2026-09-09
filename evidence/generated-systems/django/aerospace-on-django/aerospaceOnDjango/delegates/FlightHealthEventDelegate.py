from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.FlightHealthEvent import FlightHealthEvent
from aerospaceOnDjango.models.ConnectedAircraft import ConnectedAircraft
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FlightHealthEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FlightHealthEventDelegate Declaration
#======================================================================
class FlightHealthEventDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, flightHealthEventId ):
		try:	
			flightHealthEvent = FlightHealthEvent.objects.filter(id=flightHealthEventId)
			return flightHealthEvent.first();
		except FlightHealthEvent.DoesNotExist:
			raise ProcessingError("FlightHealthEvent with id " + str(flightHealthEventId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, flightHealthEvent):
		for model in serializers.deserialize("json", flightHealthEvent):
			model.save()
			return model;

	def create(self, flightHealthEvent):
		flightHealthEvent.save()
		return flightHealthEvent;

	def saveFromJson(self, flightHealthEvent):
		for model in serializers.deserialize("json", flightHealthEvent):
			model.save()
			return flightHealthEvent;
	
	def save(self, flightHealthEvent):
		flightHealthEvent.save()
		return flightHealthEvent;
	
	def delete(self, flightHealthEventId ):
		errMsg = "Failed to delete FlightHealthEvent from db using id " + str(flightHealthEventId)
		
		try:
			flightHealthEvent = FlightHealthEvent.objects.get(id=flightHealthEventId)
			flightHealthEvent.delete()
			return True
		except FlightHealthEvent.DoesNotExist:
			raise ProcessingError("FlightHealthEvent with id " + str(flightHealthEventId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FlightHealthEvent.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FlightHealthEvent from db")
		except Exception:
			return None;
		
	def assignConnectedAircraft( self, flightHealthEventId, connectedAircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ConnectedAircraftDelegate import ConnectedAircraftDelegate

		errMsg = "Failed to assign element " + str(connectedAircraftId) + " for ConnectedAircraft on FlightHealthEvent"

		try:
			# get the FlightHealthEvent from db
			flightHealthEvent = self.get( flightHealthEventId ).first()	
			
			# get the ConnectedAircraft from db
			connectedAircraft = ConnectedAircraftDelegate().get(connectedAircraftId).first();
			
			# assign the ConnectedAircraft		
			flightHealthEvent.connectedAircraft = connectedAircraft
			
			#save it
			flightHealthEvent.save()

			# reload and return the appropriate version					
			return self.get( flightHealthEventId );
		except FlightHealthEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : FlightHealthEvent with id " + str(flightHealthEventId) + " does not exist.")
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignConnectedAircraft( self, flightHealthEventId ):
		errMsg = "Failed to unassign element " + str(connectedAircraftId) + " for ConnectedAircraft on FlightHealthEvent"

		try:
			# get the FlightHealthEvent from db
			flightHealthEvent = self.get( flightHealthEventId ).first()	
			
			# assign to None for unassignment
			flightHealthEvent.connectedAircraft = None			

			#save it
			flightHealthEvent.save()

			# reload and return the appropriate version					
			return self.get( flightHealthEventId );
		except FlightHealthEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : FlightHealthEvent with id " + str(flightHealthEventId) + " does not exist.")
		except Exception:
			return None;
		
