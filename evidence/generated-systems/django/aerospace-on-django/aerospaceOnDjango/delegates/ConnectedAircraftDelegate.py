from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.ConnectedAircraft import ConnectedAircraft
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.models.FlightHealthEvent import FlightHealthEvent
from aerospaceOnDjango.models.SoftwareLoad import SoftwareLoad
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ConnectedAircraft
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConnectedAircraftDelegate Declaration
#======================================================================
class ConnectedAircraftDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, connectedAircraftId ):
		try:	
			connectedAircraft = ConnectedAircraft.objects.filter(id=connectedAircraftId)
			return connectedAircraft.first();
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError("ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, connectedAircraft):
		for model in serializers.deserialize("json", connectedAircraft):
			model.save()
			return model;

	def create(self, connectedAircraft):
		connectedAircraft.save()
		return connectedAircraft;

	def saveFromJson(self, connectedAircraft):
		for model in serializers.deserialize("json", connectedAircraft):
			model.save()
			return connectedAircraft;
	
	def save(self, connectedAircraft):
		connectedAircraft.save()
		return connectedAircraft;
	
	def delete(self, connectedAircraftId ):
		errMsg = "Failed to delete ConnectedAircraft from db using id " + str(connectedAircraftId)
		
		try:
			connectedAircraft = ConnectedAircraft.objects.get(id=connectedAircraftId)
			connectedAircraft.delete()
			return True
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError("ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ConnectedAircraft.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ConnectedAircraft from db")
		except Exception:
			return None;
		
	def assignAircraft( self, connectedAircraftId, aircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to assign element " + str(aircraftId) + " for Aircraft on ConnectedAircraft"

		try:
			# get the ConnectedAircraft from db
			connectedAircraft = self.get( connectedAircraftId ).first()	
			
			# get the Aircraft from db
			aircraft = AircraftDelegate().get(aircraftId).first();
			
			# assign the Aircraft		
			connectedAircraft.aircraft = aircraft
			
			#save it
			connectedAircraft.save()

			# reload and return the appropriate version					
			return self.get( connectedAircraftId );
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraft( self, connectedAircraftId ):
		errMsg = "Failed to unassign element " + str(aircraftId) + " for Aircraft on ConnectedAircraft"

		try:
			# get the ConnectedAircraft from db
			connectedAircraft = self.get( connectedAircraftId ).first()	
			
			# assign to None for unassignment
			connectedAircraft.aircraft = None			

			#save it
			connectedAircraft.save()

			# reload and return the appropriate version					
			return self.get( connectedAircraftId );
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except Exception:
			return None;
		
	def addFlightHealthEvents( self, connectedAircraftId, flightHealthEventsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.FlightHealthEventDelegate import FlightHealthEventDelegate

		errMsg = "Failed to add elements " + str(flightHealthEventsIds) + " for FlightHealthEvents on ConnectedAircraft"

		try:
			# get the ConnectedAircraft
			connectedAircraft = self.get( connectedAircraftId ).first()
				
			# split on a comma with no spaces
			idList = flightHealthEventsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FlightHealthEvent		
				flightHealthEvent = FlightHealthEventDelegate().get(id).first();	
				# add the FlightHealthEvent
				connectedAircraft.flightHealthEvents.add(flightHealthEvent)
				
			# save it		
			connectedAircraft.save()
			
			# reload and return the appropriate version
			return self.get( connectedAircraftId );
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except FlightHealthEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : FlightHealthEvent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFlightHealthEvents( self, connectedAircraftId, flightHealthEventsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.FlightHealthEventDelegate import FlightHealthEventDelegate

		errMsg = "Failed to remove elements " + str(flightHealthEventsIds) + " for FlightHealthEvents on ConnectedAircraft"

		try:
			# get the ConnectedAircraft
			connectedAircraft = self.get( connectedAircraftId ).first()
				
			# split on a comma with no spaces
			idList = flightHealthEventsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FlightHealthEvent		
				flightHealthEvent = FlightHealthEventDelegate().get(id).first();	
				# add the FlightHealthEvent
				connectedAircraft.flightHealthEvents.remove(flightHealthEvent)
				
			# save it		
			connectedAircraft.save()
			
			# reload and return the appropriate version
			return self.get( connectedAircraftId );
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except FlightHealthEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : FlightHealthEvent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSoftwareLoads( self, connectedAircraftId, softwareLoadsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SoftwareLoadDelegate import SoftwareLoadDelegate

		errMsg = "Failed to add elements " + str(softwareLoadsIds) + " for SoftwareLoads on ConnectedAircraft"

		try:
			# get the ConnectedAircraft
			connectedAircraft = self.get( connectedAircraftId ).first()
				
			# split on a comma with no spaces
			idList = softwareLoadsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SoftwareLoad		
				softwareLoad = SoftwareLoadDelegate().get(id).first();	
				# add the SoftwareLoad
				connectedAircraft.softwareLoads.add(softwareLoad)
				
			# save it		
			connectedAircraft.save()
			
			# reload and return the appropriate version
			return self.get( connectedAircraftId );
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSoftwareLoads( self, connectedAircraftId, softwareLoadsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SoftwareLoadDelegate import SoftwareLoadDelegate

		errMsg = "Failed to remove elements " + str(softwareLoadsIds) + " for SoftwareLoads on ConnectedAircraft"

		try:
			# get the ConnectedAircraft
			connectedAircraft = self.get( connectedAircraftId ).first()
				
			# split on a comma with no spaces
			idList = softwareLoadsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SoftwareLoad		
				softwareLoad = SoftwareLoadDelegate().get(id).first();	
				# add the SoftwareLoad
				connectedAircraft.softwareLoads.remove(softwareLoad)
				
			# save it		
			connectedAircraft.save()
			
			# reload and return the appropriate version
			return self.get( connectedAircraftId );
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
