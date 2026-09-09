from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AerospaceManufacturer import AerospaceManufacturer
from aerospaceOnDjango.models.AircraftProgram import AircraftProgram
from aerospaceOnDjango.models.Plant import Plant
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.models.ProductionCertificate import ProductionCertificate
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AerospaceManufacturer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AerospaceManufacturerDelegate Declaration
#======================================================================
class AerospaceManufacturerDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aerospaceManufacturerId ):
		try:	
			aerospaceManufacturer = AerospaceManufacturer.objects.filter(id=aerospaceManufacturerId)
			return aerospaceManufacturer.first();
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError("AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aerospaceManufacturer):
		for model in serializers.deserialize("json", aerospaceManufacturer):
			model.save()
			return model;

	def create(self, aerospaceManufacturer):
		aerospaceManufacturer.save()
		return aerospaceManufacturer;

	def saveFromJson(self, aerospaceManufacturer):
		for model in serializers.deserialize("json", aerospaceManufacturer):
			model.save()
			return aerospaceManufacturer;
	
	def save(self, aerospaceManufacturer):
		aerospaceManufacturer.save()
		return aerospaceManufacturer;
	
	def delete(self, aerospaceManufacturerId ):
		errMsg = "Failed to delete AerospaceManufacturer from db using id " + str(aerospaceManufacturerId)
		
		try:
			aerospaceManufacturer = AerospaceManufacturer.objects.get(id=aerospaceManufacturerId)
			aerospaceManufacturer.delete()
			return True
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError("AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AerospaceManufacturer.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AerospaceManufacturer from db")
		except Exception:
			return None;
		
	def addPrograms( self, aerospaceManufacturerId, programsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftProgramDelegate import AircraftProgramDelegate

		errMsg = "Failed to add elements " + str(programsIds) + " for Programs on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = programsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftProgram		
				aircraftProgram = AircraftProgramDelegate().get(id).first();	
				# add the AircraftProgram
				aerospaceManufacturer.programs.add(aircraftProgram)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePrograms( self, aerospaceManufacturerId, programsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftProgramDelegate import AircraftProgramDelegate

		errMsg = "Failed to remove elements " + str(programsIds) + " for Programs on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = programsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftProgram		
				aircraftProgram = AircraftProgramDelegate().get(id).first();	
				# add the AircraftProgram
				aerospaceManufacturer.programs.remove(aircraftProgram)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPlants( self, aerospaceManufacturerId, plantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to add elements " + str(plantsIds) + " for Plants on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = plantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Plant		
				plant = PlantDelegate().get(id).first();	
				# add the Plant
				aerospaceManufacturer.plants.add(plant)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlants( self, aerospaceManufacturerId, plantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to remove elements " + str(plantsIds) + " for Plants on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = plantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Plant		
				plant = PlantDelegate().get(id).first();	
				# add the Plant
				aerospaceManufacturer.plants.remove(plant)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSuppliers( self, aerospaceManufacturerId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to add elements " + str(suppliersIds) + " for Suppliers on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				aerospaceManufacturer.suppliers.add(supplier)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSuppliers( self, aerospaceManufacturerId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to remove elements " + str(suppliersIds) + " for Suppliers on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				aerospaceManufacturer.suppliers.remove(supplier)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProductionCertificates( self, aerospaceManufacturerId, productionCertificatesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ProductionCertificateDelegate import ProductionCertificateDelegate

		errMsg = "Failed to add elements " + str(productionCertificatesIds) + " for ProductionCertificates on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = productionCertificatesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProductionCertificate		
				productionCertificate = ProductionCertificateDelegate().get(id).first();	
				# add the ProductionCertificate
				aerospaceManufacturer.productionCertificates.add(productionCertificate)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except ProductionCertificate.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionCertificate does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProductionCertificates( self, aerospaceManufacturerId, productionCertificatesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ProductionCertificateDelegate import ProductionCertificateDelegate

		errMsg = "Failed to remove elements " + str(productionCertificatesIds) + " for ProductionCertificates on AerospaceManufacturer"

		try:
			# get the AerospaceManufacturer
			aerospaceManufacturer = self.get( aerospaceManufacturerId ).first()
				
			# split on a comma with no spaces
			idList = productionCertificatesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProductionCertificate		
				productionCertificate = ProductionCertificateDelegate().get(id).first();	
				# add the ProductionCertificate
				aerospaceManufacturer.productionCertificates.remove(productionCertificate)
				
			# save it		
			aerospaceManufacturer.save()
			
			# reload and return the appropriate version
			return self.get( aerospaceManufacturerId );
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(aerospaceManufacturerId) + " does not exist.")
		except ProductionCertificate.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionCertificate does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
