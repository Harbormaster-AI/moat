from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.ProductionCertificate import ProductionCertificate
from aerospaceOnDjango.models.AerospaceManufacturer import AerospaceManufacturer
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProductionCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionCertificateDelegate Declaration
#======================================================================
class ProductionCertificateDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productionCertificateId ):
		try:	
			productionCertificate = ProductionCertificate.objects.filter(id=productionCertificateId)
			return productionCertificate.first();
		except ProductionCertificate.DoesNotExist:
			raise ProcessingError("ProductionCertificate with id " + str(productionCertificateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, productionCertificate):
		for model in serializers.deserialize("json", productionCertificate):
			model.save()
			return model;

	def create(self, productionCertificate):
		productionCertificate.save()
		return productionCertificate;

	def saveFromJson(self, productionCertificate):
		for model in serializers.deserialize("json", productionCertificate):
			model.save()
			return productionCertificate;
	
	def save(self, productionCertificate):
		productionCertificate.save()
		return productionCertificate;
	
	def delete(self, productionCertificateId ):
		errMsg = "Failed to delete ProductionCertificate from db using id " + str(productionCertificateId)
		
		try:
			productionCertificate = ProductionCertificate.objects.get(id=productionCertificateId)
			productionCertificate.delete()
			return True
		except ProductionCertificate.DoesNotExist:
			raise ProcessingError("ProductionCertificate with id " + str(productionCertificateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProductionCertificate.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProductionCertificate from db")
		except Exception:
			return None;
		
	def assignManufacturer( self, productionCertificateId, manufacturerId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AerospaceManufacturerDelegate import AerospaceManufacturerDelegate

		errMsg = "Failed to assign element " + str(manufacturerId) + " for Manufacturer on ProductionCertificate"

		try:
			# get the ProductionCertificate from db
			productionCertificate = self.get( productionCertificateId ).first()	
			
			# get the AerospaceManufacturer from db
			aerospaceManufacturer = AerospaceManufacturerDelegate().get(manufacturerId).first();
			
			# assign the Manufacturer		
			productionCertificate.manufacturer = aerospaceManufacturer
			
			#save it
			productionCertificate.save()

			# reload and return the appropriate version					
			return self.get( productionCertificateId );
		except ProductionCertificate.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionCertificate with id " + str(productionCertificateId) + " does not exist.")
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer with id " + str(manufacturerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignManufacturer( self, productionCertificateId ):
		errMsg = "Failed to unassign element " + str(manufacturerId) + " for Manufacturer on ProductionCertificate"

		try:
			# get the ProductionCertificate from db
			productionCertificate = self.get( productionCertificateId ).first()	
			
			# assign to None for unassignment
			productionCertificate.aerospaceManufacturer = None			

			#save it
			productionCertificate.save()

			# reload and return the appropriate version					
			return self.get( productionCertificateId );
		except ProductionCertificate.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionCertificate with id " + str(productionCertificateId) + " does not exist.")
		except Exception:
			return None;
		
