from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.APU import APU
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model APU
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APUDelegate Declaration
#======================================================================
class APUDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aPUId ):
		try:	
			aPU = APU.objects.filter(id=aPUId)
			return aPU.first();
		except APU.DoesNotExist:
			raise ProcessingError("APU with id " + str(aPUId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aPU):
		for model in serializers.deserialize("json", aPU):
			model.save()
			return model;

	def create(self, aPU):
		aPU.save()
		return aPU;

	def saveFromJson(self, aPU):
		for model in serializers.deserialize("json", aPU):
			model.save()
			return aPU;
	
	def save(self, aPU):
		aPU.save()
		return aPU;
	
	def delete(self, aPUId ):
		errMsg = "Failed to delete APU from db using id " + str(aPUId)
		
		try:
			aPU = APU.objects.get(id=aPUId)
			aPU.delete()
			return True
		except APU.DoesNotExist:
			raise ProcessingError("APU with id " + str(aPUId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = APU.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all APU from db")
		except Exception:
			return None;
		
	def assignSupplier( self, aPUId, supplierId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to assign element " + str(supplierId) + " for Supplier on APU"

		try:
			# get the APU from db
			aPU = self.get( aPUId ).first()	
			
			# get the Supplier from db
			supplier = SupplierDelegate().get(supplierId).first();
			
			# assign the Supplier		
			aPU.supplier = supplier
			
			#save it
			aPU.save()

			# reload and return the appropriate version					
			return self.get( aPUId );
		except APU.DoesNotExist:
			raise ProcessingError(errMsg + " : APU with id " + str(aPUId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupplier( self, aPUId ):
		errMsg = "Failed to unassign element " + str(supplierId) + " for Supplier on APU"

		try:
			# get the APU from db
			aPU = self.get( aPUId ).first()	
			
			# assign to None for unassignment
			aPU.supplier = None			

			#save it
			aPU.save()

			# reload and return the appropriate version					
			return self.get( aPUId );
		except APU.DoesNotExist:
			raise ProcessingError(errMsg + " : APU with id " + str(aPUId) + " does not exist.")
		except Exception:
			return None;
		
	def addVariants( self, aPUId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on APU"

		try:
			# get the APU
			aPU = self.get( aPUId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aPU.variants.add(aircraftVariant)
				
			# save it		
			aPU.save()
			
			# reload and return the appropriate version
			return self.get( aPUId );
		except APU.DoesNotExist:
			raise ProcessingError(errMsg + " : APU with id " + str(aPUId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, aPUId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on APU"

		try:
			# get the APU
			aPU = self.get( aPUId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aPU.variants.remove(aircraftVariant)
				
			# save it		
			aPU.save()
			
			# reload and return the appropriate version
			return self.get( aPUId );
		except APU.DoesNotExist:
			raise ProcessingError(errMsg + " : APU with id " + str(aPUId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
