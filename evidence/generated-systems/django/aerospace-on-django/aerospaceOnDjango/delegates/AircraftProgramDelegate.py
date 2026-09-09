from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AircraftProgram import AircraftProgram
from aerospaceOnDjango.models.AerospaceManufacturer import AerospaceManufacturer
from aerospaceOnDjango.models.AircraftFamily import AircraftFamily
from aerospaceOnDjango.models.TypeCertificate import TypeCertificate
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AircraftProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftProgramDelegate Declaration
#======================================================================
class AircraftProgramDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftProgramId ):
		try:	
			aircraftProgram = AircraftProgram.objects.filter(id=aircraftProgramId)
			return aircraftProgram.first();
		except AircraftProgram.DoesNotExist:
			raise ProcessingError("AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraftProgram):
		for model in serializers.deserialize("json", aircraftProgram):
			model.save()
			return model;

	def create(self, aircraftProgram):
		aircraftProgram.save()
		return aircraftProgram;

	def saveFromJson(self, aircraftProgram):
		for model in serializers.deserialize("json", aircraftProgram):
			model.save()
			return aircraftProgram;
	
	def save(self, aircraftProgram):
		aircraftProgram.save()
		return aircraftProgram;
	
	def delete(self, aircraftProgramId ):
		errMsg = "Failed to delete AircraftProgram from db using id " + str(aircraftProgramId)
		
		try:
			aircraftProgram = AircraftProgram.objects.get(id=aircraftProgramId)
			aircraftProgram.delete()
			return True
		except AircraftProgram.DoesNotExist:
			raise ProcessingError("AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AircraftProgram.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AircraftProgram from db")
		except Exception:
			return None;
		
	def assignManufacturer( self, aircraftProgramId, manufacturerId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AerospaceManufacturerDelegate import AerospaceManufacturerDelegate

		errMsg = "Failed to assign element " + str(manufacturerId) + " for Manufacturer on AircraftProgram"

		try:
			# get the AircraftProgram from db
			aircraftProgram = self.get( aircraftProgramId ).first()	
			
			# get the AerospaceManufacturer from db
			aerospaceManufacturer = AerospaceManufacturerDelegate().get(manufacturerId).first();
			
			# assign the Manufacturer		
			aircraftProgram.manufacturer = aerospaceManufacturer
			
			#save it
			aircraftProgram.save()

			# reload and return the appropriate version					
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(manufacturerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignManufacturer( self, aircraftProgramId ):
		errMsg = "Failed to unassign element " + str(manufacturerId) + " for Manufacturer on AircraftProgram"

		try:
			# get the AircraftProgram from db
			aircraftProgram = self.get( aircraftProgramId ).first()	
			
			# assign to None for unassignment
			aircraftProgram.aerospaceManufacturer = None			

			#save it
			aircraftProgram.save()

			# reload and return the appropriate version					
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTypeCertificate( self, aircraftProgramId, typeCertificateId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.TypeCertificateDelegate import TypeCertificateDelegate

		errMsg = "Failed to assign element " + str(typeCertificateId) + " for TypeCertificate on AircraftProgram"

		try:
			# get the AircraftProgram from db
			aircraftProgram = self.get( aircraftProgramId ).first()	
			
			# get the TypeCertificate from db
			typeCertificate = TypeCertificateDelegate().get(typeCertificateId).first();
			
			# assign the TypeCertificate		
			aircraftProgram.typeCertificate = typeCertificate
			
			#save it
			aircraftProgram.save()

			# reload and return the appropriate version					
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except TypeCertificate.DoesNotExist:
			raise ProcessingError(errMsg + " : TypeCertificate with id " + str(typeCertificateId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTypeCertificate( self, aircraftProgramId ):
		errMsg = "Failed to unassign element " + str(typeCertificateId) + " for TypeCertificate on AircraftProgram"

		try:
			# get the AircraftProgram from db
			aircraftProgram = self.get( aircraftProgramId ).first()	
			
			# assign to None for unassignment
			aircraftProgram.typeCertificate = None			

			#save it
			aircraftProgram.save()

			# reload and return the appropriate version					
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except Exception:
			return None;
		
	def addAircraftFamilies( self, aircraftProgramId, aircraftFamiliesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftFamilyDelegate import AircraftFamilyDelegate

		errMsg = "Failed to add elements " + str(aircraftFamiliesIds) + " for AircraftFamilies on AircraftProgram"

		try:
			# get the AircraftProgram
			aircraftProgram = self.get( aircraftProgramId ).first()
				
			# split on a comma with no spaces
			idList = aircraftFamiliesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftFamily		
				aircraftFamily = AircraftFamilyDelegate().get(id).first();	
				# add the AircraftFamily
				aircraftProgram.aircraftFamilies.add(aircraftFamily)
				
			# save it		
			aircraftProgram.save()
			
			# reload and return the appropriate version
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except AircraftFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftFamily does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAircraftFamilies( self, aircraftProgramId, aircraftFamiliesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftFamilyDelegate import AircraftFamilyDelegate

		errMsg = "Failed to remove elements " + str(aircraftFamiliesIds) + " for AircraftFamilies on AircraftProgram"

		try:
			# get the AircraftProgram
			aircraftProgram = self.get( aircraftProgramId ).first()
				
			# split on a comma with no spaces
			idList = aircraftFamiliesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftFamily		
				aircraftFamily = AircraftFamilyDelegate().get(id).first();	
				# add the AircraftFamily
				aircraftProgram.aircraftFamilies.remove(aircraftFamily)
				
			# save it		
			aircraftProgram.save()
			
			# reload and return the appropriate version
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except AircraftFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftFamily does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addKeySuppliers( self, aircraftProgramId, keySuppliersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to add elements " + str(keySuppliersIds) + " for KeySuppliers on AircraftProgram"

		try:
			# get the AircraftProgram
			aircraftProgram = self.get( aircraftProgramId ).first()
				
			# split on a comma with no spaces
			idList = keySuppliersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				aircraftProgram.keySuppliers.add(supplier)
				
			# save it		
			aircraftProgram.save()
			
			# reload and return the appropriate version
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeKeySuppliers( self, aircraftProgramId, keySuppliersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to remove elements " + str(keySuppliersIds) + " for KeySuppliers on AircraftProgram"

		try:
			# get the AircraftProgram
			aircraftProgram = self.get( aircraftProgramId ).first()
				
			# split on a comma with no spaces
			idList = keySuppliersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				aircraftProgram.keySuppliers.remove(supplier)
				
			# save it		
			aircraftProgram.save()
			
			# reload and return the appropriate version
			return self.get( aircraftProgramId );
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(aircraftProgramId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
