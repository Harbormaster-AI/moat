from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AvionicsSuite import AvionicsSuite
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.SoftwareLoad import SoftwareLoad
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AvionicsSuite
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AvionicsSuiteDelegate Declaration
#======================================================================
class AvionicsSuiteDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, avionicsSuiteId ):
		try:	
			avionicsSuite = AvionicsSuite.objects.filter(id=avionicsSuiteId)
			return avionicsSuite.first();
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError("AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, avionicsSuite):
		for model in serializers.deserialize("json", avionicsSuite):
			model.save()
			return model;

	def create(self, avionicsSuite):
		avionicsSuite.save()
		return avionicsSuite;

	def saveFromJson(self, avionicsSuite):
		for model in serializers.deserialize("json", avionicsSuite):
			model.save()
			return avionicsSuite;
	
	def save(self, avionicsSuite):
		avionicsSuite.save()
		return avionicsSuite;
	
	def delete(self, avionicsSuiteId ):
		errMsg = "Failed to delete AvionicsSuite from db using id " + str(avionicsSuiteId)
		
		try:
			avionicsSuite = AvionicsSuite.objects.get(id=avionicsSuiteId)
			avionicsSuite.delete()
			return True
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError("AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AvionicsSuite.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AvionicsSuite from db")
		except Exception:
			return None;
		
	def assignSupplier( self, avionicsSuiteId, supplierId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to assign element " + str(supplierId) + " for Supplier on AvionicsSuite"

		try:
			# get the AvionicsSuite from db
			avionicsSuite = self.get( avionicsSuiteId ).first()	
			
			# get the Supplier from db
			supplier = SupplierDelegate().get(supplierId).first();
			
			# assign the Supplier		
			avionicsSuite.supplier = supplier
			
			#save it
			avionicsSuite.save()

			# reload and return the appropriate version					
			return self.get( avionicsSuiteId );
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupplier( self, avionicsSuiteId ):
		errMsg = "Failed to unassign element " + str(supplierId) + " for Supplier on AvionicsSuite"

		try:
			# get the AvionicsSuite from db
			avionicsSuite = self.get( avionicsSuiteId ).first()	
			
			# assign to None for unassignment
			avionicsSuite.supplier = None			

			#save it
			avionicsSuite.save()

			# reload and return the appropriate version					
			return self.get( avionicsSuiteId );
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except Exception:
			return None;
		
	def addVariants( self, avionicsSuiteId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on AvionicsSuite"

		try:
			# get the AvionicsSuite
			avionicsSuite = self.get( avionicsSuiteId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				avionicsSuite.variants.add(aircraftVariant)
				
			# save it		
			avionicsSuite.save()
			
			# reload and return the appropriate version
			return self.get( avionicsSuiteId );
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, avionicsSuiteId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on AvionicsSuite"

		try:
			# get the AvionicsSuite
			avionicsSuite = self.get( avionicsSuiteId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				avionicsSuite.variants.remove(aircraftVariant)
				
			# save it		
			avionicsSuite.save()
			
			# reload and return the appropriate version
			return self.get( avionicsSuiteId );
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSoftwareLoads( self, avionicsSuiteId, softwareLoadsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SoftwareLoadDelegate import SoftwareLoadDelegate

		errMsg = "Failed to add elements " + str(softwareLoadsIds) + " for SoftwareLoads on AvionicsSuite"

		try:
			# get the AvionicsSuite
			avionicsSuite = self.get( avionicsSuiteId ).first()
				
			# split on a comma with no spaces
			idList = softwareLoadsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SoftwareLoad		
				softwareLoad = SoftwareLoadDelegate().get(id).first();	
				# add the SoftwareLoad
				avionicsSuite.softwareLoads.add(softwareLoad)
				
			# save it		
			avionicsSuite.save()
			
			# reload and return the appropriate version
			return self.get( avionicsSuiteId );
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSoftwareLoads( self, avionicsSuiteId, softwareLoadsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SoftwareLoadDelegate import SoftwareLoadDelegate

		errMsg = "Failed to remove elements " + str(softwareLoadsIds) + " for SoftwareLoads on AvionicsSuite"

		try:
			# get the AvionicsSuite
			avionicsSuite = self.get( avionicsSuiteId ).first()
				
			# split on a comma with no spaces
			idList = softwareLoadsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SoftwareLoad		
				softwareLoad = SoftwareLoadDelegate().get(id).first();	
				# add the SoftwareLoad
				avionicsSuite.softwareLoads.remove(softwareLoad)
				
			# save it		
			avionicsSuite.save()
			
			# reload and return the appropriate version
			return self.get( avionicsSuiteId );
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except SoftwareLoad.DoesNotExist:
			raise ProcessingError(errMsg + " : SoftwareLoad does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
