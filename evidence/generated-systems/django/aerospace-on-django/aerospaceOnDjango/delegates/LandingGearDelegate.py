from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.LandingGear import LandingGear
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LandingGear
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LandingGearDelegate Declaration
#======================================================================
class LandingGearDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, landingGearId ):
		try:	
			landingGear = LandingGear.objects.filter(id=landingGearId)
			return landingGear.first();
		except LandingGear.DoesNotExist:
			raise ProcessingError("LandingGear with id " + str(landingGearId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, landingGear):
		for model in serializers.deserialize("json", landingGear):
			model.save()
			return model;

	def create(self, landingGear):
		landingGear.save()
		return landingGear;

	def saveFromJson(self, landingGear):
		for model in serializers.deserialize("json", landingGear):
			model.save()
			return landingGear;
	
	def save(self, landingGear):
		landingGear.save()
		return landingGear;
	
	def delete(self, landingGearId ):
		errMsg = "Failed to delete LandingGear from db using id " + str(landingGearId)
		
		try:
			landingGear = LandingGear.objects.get(id=landingGearId)
			landingGear.delete()
			return True
		except LandingGear.DoesNotExist:
			raise ProcessingError("LandingGear with id " + str(landingGearId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LandingGear.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LandingGear from db")
		except Exception:
			return None;
		
	def assignSupplier( self, landingGearId, supplierId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to assign element " + str(supplierId) + " for Supplier on LandingGear"

		try:
			# get the LandingGear from db
			landingGear = self.get( landingGearId ).first()	
			
			# get the Supplier from db
			supplier = SupplierDelegate().get(supplierId).first();
			
			# assign the Supplier		
			landingGear.supplier = supplier
			
			#save it
			landingGear.save()

			# reload and return the appropriate version					
			return self.get( landingGearId );
		except LandingGear.DoesNotExist:
			raise ProcessingError(errMsg + " : LandingGear with id " + str(landingGearId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupplier( self, landingGearId ):
		errMsg = "Failed to unassign element " + str(supplierId) + " for Supplier on LandingGear"

		try:
			# get the LandingGear from db
			landingGear = self.get( landingGearId ).first()	
			
			# assign to None for unassignment
			landingGear.supplier = None			

			#save it
			landingGear.save()

			# reload and return the appropriate version					
			return self.get( landingGearId );
		except LandingGear.DoesNotExist:
			raise ProcessingError(errMsg + " : LandingGear with id " + str(landingGearId) + " does not exist.")
		except Exception:
			return None;
		
	def addVariants( self, landingGearId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on LandingGear"

		try:
			# get the LandingGear
			landingGear = self.get( landingGearId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				landingGear.variants.add(aircraftVariant)
				
			# save it		
			landingGear.save()
			
			# reload and return the appropriate version
			return self.get( landingGearId );
		except LandingGear.DoesNotExist:
			raise ProcessingError(errMsg + " : LandingGear with id " + str(landingGearId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, landingGearId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on LandingGear"

		try:
			# get the LandingGear
			landingGear = self.get( landingGearId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				landingGear.variants.remove(aircraftVariant)
				
			# save it		
			landingGear.save()
			
			# reload and return the appropriate version
			return self.get( landingGearId );
		except LandingGear.DoesNotExist:
			raise ProcessingError(errMsg + " : LandingGear with id " + str(landingGearId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
