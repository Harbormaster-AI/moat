from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.AircraftModel import AircraftModel
from aerospaceOnDjango.models.EngineType import EngineType
from aerospaceOnDjango.models.AvionicsSuite import AvionicsSuite
from aerospaceOnDjango.models.APU import APU
from aerospaceOnDjango.models.LandingGear import LandingGear
from aerospaceOnDjango.models.CabinLayout import CabinLayout
from aerospaceOnDjango.models.AircraftOption import AircraftOption
from aerospaceOnDjango.models.AircraftPackage import AircraftPackage
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AircraftVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftVariantDelegate Declaration
#======================================================================
class AircraftVariantDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftVariantId ):
		try:	
			aircraftVariant = AircraftVariant.objects.filter(id=aircraftVariantId)
			return aircraftVariant.first();
		except AircraftVariant.DoesNotExist:
			raise ProcessingError("AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraftVariant):
		for model in serializers.deserialize("json", aircraftVariant):
			model.save()
			return model;

	def create(self, aircraftVariant):
		aircraftVariant.save()
		return aircraftVariant;

	def saveFromJson(self, aircraftVariant):
		for model in serializers.deserialize("json", aircraftVariant):
			model.save()
			return aircraftVariant;
	
	def save(self, aircraftVariant):
		aircraftVariant.save()
		return aircraftVariant;
	
	def delete(self, aircraftVariantId ):
		errMsg = "Failed to delete AircraftVariant from db using id " + str(aircraftVariantId)
		
		try:
			aircraftVariant = AircraftVariant.objects.get(id=aircraftVariantId)
			aircraftVariant.delete()
			return True
		except AircraftVariant.DoesNotExist:
			raise ProcessingError("AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AircraftVariant.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AircraftVariant from db")
		except Exception:
			return None;
		
	def assignModel( self, aircraftVariantId, modelId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftModelDelegate import AircraftModelDelegate

		errMsg = "Failed to assign element " + str(modelId) + " for Model on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# get the AircraftModel from db
			aircraftModel = AircraftModelDelegate().get(modelId).first();
			
			# assign the Model		
			aircraftVariant.model = aircraftModel
			
			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel with id " + str(modelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModel( self, aircraftVariantId ):
		errMsg = "Failed to unassign element " + str(modelId) + " for Model on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# assign to None for unassignment
			aircraftVariant.aircraftModel = None			

			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEngineType( self, aircraftVariantId, engineTypeId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.EngineTypeDelegate import EngineTypeDelegate

		errMsg = "Failed to assign element " + str(engineTypeId) + " for EngineType on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# get the EngineType from db
			engineType = EngineTypeDelegate().get(engineTypeId).first();
			
			# assign the EngineType		
			aircraftVariant.engineType = engineType
			
			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType with id " + str(engineTypeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEngineType( self, aircraftVariantId ):
		errMsg = "Failed to unassign element " + str(engineTypeId) + " for EngineType on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# assign to None for unassignment
			aircraftVariant.engineType = None			

			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAvionicsSuite( self, aircraftVariantId, avionicsSuiteId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AvionicsSuiteDelegate import AvionicsSuiteDelegate

		errMsg = "Failed to assign element " + str(avionicsSuiteId) + " for AvionicsSuite on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# get the AvionicsSuite from db
			avionicsSuite = AvionicsSuiteDelegate().get(avionicsSuiteId).first();
			
			# assign the AvionicsSuite		
			aircraftVariant.avionicsSuite = avionicsSuite
			
			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite with id " + str(avionicsSuiteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAvionicsSuite( self, aircraftVariantId ):
		errMsg = "Failed to unassign element " + str(avionicsSuiteId) + " for AvionicsSuite on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# assign to None for unassignment
			aircraftVariant.avionicsSuite = None			

			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def assignApu( self, aircraftVariantId, apuId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.APUDelegate import APUDelegate

		errMsg = "Failed to assign element " + str(apuId) + " for Apu on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# get the APU from db
			aPU = APUDelegate().get(apuId).first();
			
			# assign the Apu		
			aircraftVariant.apu = aPU
			
			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except APU.DoesNotExist:
			raise ProcessingError(errMsg + " : APU with id " + str(apuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignApu( self, aircraftVariantId ):
		errMsg = "Failed to unassign element " + str(apuId) + " for Apu on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# assign to None for unassignment
			aircraftVariant.aPU = None			

			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLandingGear( self, aircraftVariantId, landingGearId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.LandingGearDelegate import LandingGearDelegate

		errMsg = "Failed to assign element " + str(landingGearId) + " for LandingGear on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# get the LandingGear from db
			landingGear = LandingGearDelegate().get(landingGearId).first();
			
			# assign the LandingGear		
			aircraftVariant.landingGear = landingGear
			
			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except LandingGear.DoesNotExist:
			raise ProcessingError(errMsg + " : LandingGear with id " + str(landingGearId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLandingGear( self, aircraftVariantId ):
		errMsg = "Failed to unassign element " + str(landingGearId) + " for LandingGear on AircraftVariant"

		try:
			# get the AircraftVariant from db
			aircraftVariant = self.get( aircraftVariantId ).first()	
			
			# assign to None for unassignment
			aircraftVariant.landingGear = None			

			#save it
			aircraftVariant.save()

			# reload and return the appropriate version					
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def addCabinLayouts( self, aircraftVariantId, cabinLayoutsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.CabinLayoutDelegate import CabinLayoutDelegate

		errMsg = "Failed to add elements " + str(cabinLayoutsIds) + " for CabinLayouts on AircraftVariant"

		try:
			# get the AircraftVariant
			aircraftVariant = self.get( aircraftVariantId ).first()
				
			# split on a comma with no spaces
			idList = cabinLayoutsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CabinLayout		
				cabinLayout = CabinLayoutDelegate().get(id).first();	
				# add the CabinLayout
				aircraftVariant.cabinLayouts.add(cabinLayout)
				
			# save it		
			aircraftVariant.save()
			
			# reload and return the appropriate version
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCabinLayouts( self, aircraftVariantId, cabinLayoutsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.CabinLayoutDelegate import CabinLayoutDelegate

		errMsg = "Failed to remove elements " + str(cabinLayoutsIds) + " for CabinLayouts on AircraftVariant"

		try:
			# get the AircraftVariant
			aircraftVariant = self.get( aircraftVariantId ).first()
				
			# split on a comma with no spaces
			idList = cabinLayoutsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CabinLayout		
				cabinLayout = CabinLayoutDelegate().get(id).first();	
				# add the CabinLayout
				aircraftVariant.cabinLayouts.remove(cabinLayout)
				
			# save it		
			aircraftVariant.save()
			
			# reload and return the appropriate version
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except CabinLayout.DoesNotExist:
			raise ProcessingError(errMsg + " : CabinLayout does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOptions( self, aircraftVariantId, optionsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

		errMsg = "Failed to add elements " + str(optionsIds) + " for Options on AircraftVariant"

		try:
			# get the AircraftVariant
			aircraftVariant = self.get( aircraftVariantId ).first()
				
			# split on a comma with no spaces
			idList = optionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftOption		
				aircraftOption = AircraftOptionDelegate().get(id).first();	
				# add the AircraftOption
				aircraftVariant.options.add(aircraftOption)
				
			# save it		
			aircraftVariant.save()
			
			# reload and return the appropriate version
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOptions( self, aircraftVariantId, optionsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

		errMsg = "Failed to remove elements " + str(optionsIds) + " for Options on AircraftVariant"

		try:
			# get the AircraftVariant
			aircraftVariant = self.get( aircraftVariantId ).first()
				
			# split on a comma with no spaces
			idList = optionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftOption		
				aircraftOption = AircraftOptionDelegate().get(id).first();	
				# add the AircraftOption
				aircraftVariant.options.remove(aircraftOption)
				
			# save it		
			aircraftVariant.save()
			
			# reload and return the appropriate version
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPackages( self, aircraftVariantId, packagesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftPackageDelegate import AircraftPackageDelegate

		errMsg = "Failed to add elements " + str(packagesIds) + " for Packages on AircraftVariant"

		try:
			# get the AircraftVariant
			aircraftVariant = self.get( aircraftVariantId ).first()
				
			# split on a comma with no spaces
			idList = packagesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftPackage		
				aircraftPackage = AircraftPackageDelegate().get(id).first();	
				# add the AircraftPackage
				aircraftVariant.packages.add(aircraftPackage)
				
			# save it		
			aircraftVariant.save()
			
			# reload and return the appropriate version
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePackages( self, aircraftVariantId, packagesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftPackageDelegate import AircraftPackageDelegate

		errMsg = "Failed to remove elements " + str(packagesIds) + " for Packages on AircraftVariant"

		try:
			# get the AircraftVariant
			aircraftVariant = self.get( aircraftVariantId ).first()
				
			# split on a comma with no spaces
			idList = packagesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftPackage		
				aircraftPackage = AircraftPackageDelegate().get(id).first();	
				# add the AircraftPackage
				aircraftVariant.packages.remove(aircraftPackage)
				
			# save it		
			aircraftVariant.save()
			
			# reload and return the appropriate version
			return self.get( aircraftVariantId );
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant with id " + str(aircraftVariantId) + " does not exist.")
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
