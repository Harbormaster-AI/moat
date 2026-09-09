from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.SoftwareLoad import SoftwareLoad
from aerospaceOnDjango.models.ConnectedAircraft import ConnectedAircraft
from aerospaceOnDjango.models.AvionicsSuite import AvionicsSuite
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SoftwareLoad
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareLoadDelegate Declaration
#======================================================================
class SoftwareLoadDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, softwareLoadId ):
		try:	
			softwareLoad = SoftwareLoad.objects.filter(id=softwareLoadId)
			return softwareLoad.first();
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError("SoftwareLoad with id " + str(softwareLoadId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, softwareLoad):
		for model in serializers.deserialize("json", softwareLoad):
			model.save()
			return model;

	def create(self, softwareLoad):
		softwareLoad.save()
		return softwareLoad;

	def saveFromJson(self, softwareLoad):
		for model in serializers.deserialize("json", softwareLoad):
			model.save()
			return softwareLoad;
	
	def save(self, softwareLoad):
		softwareLoad.save()
		return softwareLoad;
	
	def delete(self, softwareLoadId ):
		errMsg = "Failed to delete SoftwareLoad from db using id " + str(softwareLoadId)
		
		try:
			softwareLoad = SoftwareLoad.objects.get(id=softwareLoadId)
			softwareLoad.delete()
			return True
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError("SoftwareLoad with id " + str(softwareLoadId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SoftwareLoad.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SoftwareLoad from db")
		except Exception:
			return None;
		
	def assignConnectedAircraft( self, softwareLoadId, connectedAircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ConnectedAircraftDelegate import ConnectedAircraftDelegate

		errMsg = "Failed to assign element " + str(connectedAircraftId) + " for ConnectedAircraft on SoftwareLoad"

		try:
			# get the SoftwareLoad from db
			softwareLoad = self.get( softwareLoadId ).first()	
			
			# get the ConnectedAircraft from db
			connectedAircraft = ConnectedAircraftDelegate().get(connectedAircraftId).first();
			
			# assign the ConnectedAircraft		
			softwareLoad.connectedAircraft = connectedAircraft
			
			#save it
			softwareLoad.save()

			# reload and return the appropriate version					
			return self.get( softwareLoadId );
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad with id " + str(softwareLoadId) + " does not exist.")
		except ConnectedAircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : ConnectedAircraft with id " + str(connectedAircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignConnectedAircraft( self, softwareLoadId ):
		errMsg = "Failed to unassign element " + str(connectedAircraftId) + " for ConnectedAircraft on SoftwareLoad"

		try:
			# get the SoftwareLoad from db
			softwareLoad = self.get( softwareLoadId ).first()	
			
			# assign to None for unassignment
			softwareLoad.connectedAircraft = None			

			#save it
			softwareLoad.save()

			# reload and return the appropriate version					
			return self.get( softwareLoadId );
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad with id " + str(softwareLoadId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAvionicsSuite( self, softwareLoadId, avionicsSuiteId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AvionicsSuiteDelegate import AvionicsSuiteDelegate

		errMsg = "Failed to assign element " + str(avionicsSuiteId) + " for AvionicsSuite on SoftwareLoad"

		try:
			# get the SoftwareLoad from db
			softwareLoad = self.get( softwareLoadId ).first()	
			
			# get the AvionicsSuite from db
			avionicsSuite = AvionicsSuiteDelegate().get(avionicsSuiteId).first();
			
			# assign the AvionicsSuite		
			softwareLoad.avionicsSuite = avionicsSuite
			
			#save it
			softwareLoad.save()

			# reload and return the appropriate version					
			return self.get( softwareLoadId );
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad with id " + str(softwareLoadId) + " does not exist.")
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAvionicsSuite( self, softwareLoadId ):
		errMsg = "Failed to unassign element " + str(avionicsSuiteId) + " for AvionicsSuite on SoftwareLoad"

		try:
			# get the SoftwareLoad from db
			softwareLoad = self.get( softwareLoadId ).first()	
			
			# assign to None for unassignment
			softwareLoad.avionicsSuite = None			

			#save it
			softwareLoad.save()

			# reload and return the appropriate version					
			return self.get( softwareLoadId );
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad with id " + str(softwareLoadId) + " does not exist.")
		except Exception:
			return None;
		
