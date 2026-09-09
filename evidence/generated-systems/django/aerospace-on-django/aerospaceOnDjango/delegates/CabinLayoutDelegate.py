from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.CabinLayout import CabinLayout
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.models.AircraftOption import AircraftOption
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CabinLayout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CabinLayoutDelegate Declaration
#======================================================================
class CabinLayoutDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, cabinLayoutId ):
		try:	
			cabinLayout = CabinLayout.objects.filter(id=cabinLayoutId)
			return cabinLayout.first();
		except CabinLayout.DoesNotExist:
			raise ProcessingError("CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, cabinLayout):
		for model in serializers.deserialize("json", cabinLayout):
			model.save()
			return model;

	def create(self, cabinLayout):
		cabinLayout.save()
		return cabinLayout;

	def saveFromJson(self, cabinLayout):
		for model in serializers.deserialize("json", cabinLayout):
			model.save()
			return cabinLayout;
	
	def save(self, cabinLayout):
		cabinLayout.save()
		return cabinLayout;
	
	def delete(self, cabinLayoutId ):
		errMsg = "Failed to delete CabinLayout from db using id " + str(cabinLayoutId)
		
		try:
			cabinLayout = CabinLayout.objects.get(id=cabinLayoutId)
			cabinLayout.delete()
			return True
		except CabinLayout.DoesNotExist:
			raise ProcessingError("CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CabinLayout.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CabinLayout from db")
		except Exception:
			return None;
		
	def assignVariant( self, cabinLayoutId, variantId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on CabinLayout"

		try:
			# get the CabinLayout from db
			cabinLayout = self.get( cabinLayoutId ).first()	
			
			# get the AircraftVariant from db
			aircraftVariant = AircraftVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			cabinLayout.variant = aircraftVariant
			
			#save it
			cabinLayout.save()

			# reload and return the appropriate version					
			return self.get( cabinLayoutId );
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, cabinLayoutId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on CabinLayout"

		try:
			# get the CabinLayout from db
			cabinLayout = self.get( cabinLayoutId ).first()	
			
			# assign to None for unassignment
			cabinLayout.aircraftVariant = None			

			#save it
			cabinLayout.save()

			# reload and return the appropriate version					
			return self.get( cabinLayoutId );
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except Exception:
			return None;
		
	def addAircraft( self, cabinLayoutId, aircraftIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to add elements " + str(aircraftIds) + " for Aircraft on CabinLayout"

		try:
			# get the CabinLayout
			cabinLayout = self.get( cabinLayoutId ).first()
				
			# split on a comma with no spaces
			idList = aircraftIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Aircraft		
				aircraft = AircraftDelegate().get(id).first();	
				# add the Aircraft
				cabinLayout.aircraft.add(aircraft)
				
			# save it		
			cabinLayout.save()
			
			# reload and return the appropriate version
			return self.get( cabinLayoutId );
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAircraft( self, cabinLayoutId, aircraftIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to remove elements " + str(aircraftIds) + " for Aircraft on CabinLayout"

		try:
			# get the CabinLayout
			cabinLayout = self.get( cabinLayoutId ).first()
				
			# split on a comma with no spaces
			idList = aircraftIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Aircraft		
				aircraft = AircraftDelegate().get(id).first();	
				# add the Aircraft
				cabinLayout.aircraft.remove(aircraft)
				
			# save it		
			cabinLayout.save()
			
			# reload and return the appropriate version
			return self.get( cabinLayoutId );
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOptions( self, cabinLayoutId, optionsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

		errMsg = "Failed to add elements " + str(optionsIds) + " for Options on CabinLayout"

		try:
			# get the CabinLayout
			cabinLayout = self.get( cabinLayoutId ).first()
				
			# split on a comma with no spaces
			idList = optionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftOption		
				aircraftOption = AircraftOptionDelegate().get(id).first();	
				# add the AircraftOption
				cabinLayout.options.add(aircraftOption)
				
			# save it		
			cabinLayout.save()
			
			# reload and return the appropriate version
			return self.get( cabinLayoutId );
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOptions( self, cabinLayoutId, optionsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

		errMsg = "Failed to remove elements " + str(optionsIds) + " for Options on CabinLayout"

		try:
			# get the CabinLayout
			cabinLayout = self.get( cabinLayoutId ).first()
				
			# split on a comma with no spaces
			idList = optionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftOption		
				aircraftOption = AircraftOptionDelegate().get(id).first();	
				# add the AircraftOption
				cabinLayout.options.remove(aircraftOption)
				
			# save it		
			cabinLayout.save()
			
			# reload and return the appropriate version
			return self.get( cabinLayoutId );
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout with id " + str(cabinLayoutId) + " does not exist.")
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
