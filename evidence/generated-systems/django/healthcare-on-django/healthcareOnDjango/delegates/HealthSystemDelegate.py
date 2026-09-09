from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.HealthSystem import HealthSystem
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.MedicalSupplier import MedicalSupplier
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model HealthSystem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class HealthSystemDelegate Declaration
#======================================================================
class HealthSystemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, healthSystemId ):
		try:	
			healthSystem = HealthSystem.objects.filter(id=healthSystemId)
			return healthSystem.first();
		except HealthSystem.DoesNotExist:
			raise ProcessingError("HealthSystem with id " + str(healthSystemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, healthSystem):
		for model in serializers.deserialize("json", healthSystem):
			model.save()
			return model;

	def create(self, healthSystem):
		healthSystem.save()
		return healthSystem;

	def saveFromJson(self, healthSystem):
		for model in serializers.deserialize("json", healthSystem):
			model.save()
			return healthSystem;
	
	def save(self, healthSystem):
		healthSystem.save()
		return healthSystem;
	
	def delete(self, healthSystemId ):
		errMsg = "Failed to delete HealthSystem from db using id " + str(healthSystemId)
		
		try:
			healthSystem = HealthSystem.objects.get(id=healthSystemId)
			healthSystem.delete()
			return True
		except HealthSystem.DoesNotExist:
			raise ProcessingError("HealthSystem with id " + str(healthSystemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = HealthSystem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all HealthSystem from db")
		except Exception:
			return None;
		
	def addFacilities( self, healthSystemId, facilitiesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to add elements " + str(facilitiesIds) + " for Facilities on HealthSystem"

		try:
			# get the HealthSystem
			healthSystem = self.get( healthSystemId ).first()
				
			# split on a comma with no spaces
			idList = facilitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Facility		
				facility = FacilityDelegate().get(id).first();	
				# add the Facility
				healthSystem.facilities.add(facility)
				
			# save it		
			healthSystem.save()
			
			# reload and return the appropriate version
			return self.get( healthSystemId );
		except HealthSystem.DoesNotExist:
			raise ProcessingError(errMsg + " : HealthSystem with id " + str(healthSystemId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFacilities( self, healthSystemId, facilitiesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to remove elements " + str(facilitiesIds) + " for Facilities on HealthSystem"

		try:
			# get the HealthSystem
			healthSystem = self.get( healthSystemId ).first()
				
			# split on a comma with no spaces
			idList = facilitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Facility		
				facility = FacilityDelegate().get(id).first();	
				# add the Facility
				healthSystem.facilities.remove(facility)
				
			# save it		
			healthSystem.save()
			
			# reload and return the appropriate version
			return self.get( healthSystemId );
		except HealthSystem.DoesNotExist:
			raise ProcessingError(errMsg + " : HealthSystem with id " + str(healthSystemId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSuppliers( self, healthSystemId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicalSupplierDelegate import MedicalSupplierDelegate

		errMsg = "Failed to add elements " + str(suppliersIds) + " for Suppliers on HealthSystem"

		try:
			# get the HealthSystem
			healthSystem = self.get( healthSystemId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MedicalSupplier		
				medicalSupplier = MedicalSupplierDelegate().get(id).first();	
				# add the MedicalSupplier
				healthSystem.suppliers.add(medicalSupplier)
				
			# save it		
			healthSystem.save()
			
			# reload and return the appropriate version
			return self.get( healthSystemId );
		except HealthSystem.DoesNotExist:
			raise ProcessingError(errMsg + " : HealthSystem with id " + str(healthSystemId) + " does not exist.")
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalSupplier does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSuppliers( self, healthSystemId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicalSupplierDelegate import MedicalSupplierDelegate

		errMsg = "Failed to remove elements " + str(suppliersIds) + " for Suppliers on HealthSystem"

		try:
			# get the HealthSystem
			healthSystem = self.get( healthSystemId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MedicalSupplier		
				medicalSupplier = MedicalSupplierDelegate().get(id).first();	
				# add the MedicalSupplier
				healthSystem.suppliers.remove(medicalSupplier)
				
			# save it		
			healthSystem.save()
			
			# reload and return the appropriate version
			return self.get( healthSystemId );
		except HealthSystem.DoesNotExist:
			raise ProcessingError(errMsg + " : HealthSystem with id " + str(healthSystemId) + " does not exist.")
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalSupplier does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
