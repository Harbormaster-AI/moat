from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.models.AerospaceManufacturer import AerospaceManufacturer
from aerospaceOnDjango.models.Component_ import Component_
from aerospaceOnDjango.models.EngineType import EngineType
from aerospaceOnDjango.models.AvionicsSuite import AvionicsSuite
from aerospaceOnDjango.models.APU import APU
from aerospaceOnDjango.models.LandingGear import LandingGear
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierDelegate Declaration
#======================================================================
class SupplierDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, supplierId ):
		try:	
			supplier = Supplier.objects.filter(id=supplierId)
			return supplier.first();
		except Supplier.DoesNotExist:
			raise ProcessingError("Supplier with id " + str(supplierId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, supplier):
		for model in serializers.deserialize("json", supplier):
			model.save()
			return model;

	def create(self, supplier):
		supplier.save()
		return supplier;

	def saveFromJson(self, supplier):
		for model in serializers.deserialize("json", supplier):
			model.save()
			return supplier;
	
	def save(self, supplier):
		supplier.save()
		return supplier;
	
	def delete(self, supplierId ):
		errMsg = "Failed to delete Supplier from db using id " + str(supplierId)
		
		try:
			supplier = Supplier.objects.get(id=supplierId)
			supplier.delete()
			return True
		except Supplier.DoesNotExist:
			raise ProcessingError("Supplier with id " + str(supplierId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Supplier.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Supplier from db")
		except Exception:
			return None;
		
	def addManufacturers( self, supplierId, manufacturersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AerospaceManufacturerDelegate import AerospaceManufacturerDelegate

		errMsg = "Failed to add elements " + str(manufacturersIds) + " for Manufacturers on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = manufacturersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AerospaceManufacturer		
				aerospaceManufacturer = AerospaceManufacturerDelegate().get(id).first();	
				# add the AerospaceManufacturer
				supplier.manufacturers.add(aerospaceManufacturer)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeManufacturers( self, supplierId, manufacturersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AerospaceManufacturerDelegate import AerospaceManufacturerDelegate

		errMsg = "Failed to remove elements " + str(manufacturersIds) + " for Manufacturers on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = manufacturersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AerospaceManufacturer		
				aerospaceManufacturer = AerospaceManufacturerDelegate().get(id).first();	
				# add the AerospaceManufacturer
				supplier.manufacturers.remove(aerospaceManufacturer)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except AerospaceManufacturer.DoesNotExist:
			raise ProcessingError(errMsg + " : AerospaceManufacturer does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addComponents( self, supplierId, componentsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.Component_Delegate import Component_Delegate

		errMsg = "Failed to add elements " + str(componentsIds) + " for Components on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = componentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Component_		
				component_ = Component_Delegate().get(id).first();	
				# add the Component_
				supplier.components.add(component_)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Component_.DoesNotExist:
			raise ProcessingError(errMsg + " : Component_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeComponents( self, supplierId, componentsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.Component_Delegate import Component_Delegate

		errMsg = "Failed to remove elements " + str(componentsIds) + " for Components on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = componentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Component_		
				component_ = Component_Delegate().get(id).first();	
				# add the Component_
				supplier.components.remove(component_)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Component_.DoesNotExist:
			raise ProcessingError(errMsg + " : Component_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEngineTypes( self, supplierId, engineTypesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.EngineTypeDelegate import EngineTypeDelegate

		errMsg = "Failed to add elements " + str(engineTypesIds) + " for EngineTypes on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = engineTypesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EngineType		
				engineType = EngineTypeDelegate().get(id).first();	
				# add the EngineType
				supplier.engineTypes.add(engineType)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEngineTypes( self, supplierId, engineTypesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.EngineTypeDelegate import EngineTypeDelegate

		errMsg = "Failed to remove elements " + str(engineTypesIds) + " for EngineTypes on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = engineTypesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EngineType		
				engineType = EngineTypeDelegate().get(id).first();	
				# add the EngineType
				supplier.engineTypes.remove(engineType)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAvionicsSuites( self, supplierId, avionicsSuitesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AvionicsSuiteDelegate import AvionicsSuiteDelegate

		errMsg = "Failed to add elements " + str(avionicsSuitesIds) + " for AvionicsSuites on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = avionicsSuitesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AvionicsSuite		
				avionicsSuite = AvionicsSuiteDelegate().get(id).first();	
				# add the AvionicsSuite
				supplier.avionicsSuites.add(avionicsSuite)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAvionicsSuites( self, supplierId, avionicsSuitesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AvionicsSuiteDelegate import AvionicsSuiteDelegate

		errMsg = "Failed to remove elements " + str(avionicsSuitesIds) + " for AvionicsSuites on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = avionicsSuitesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AvionicsSuite		
				avionicsSuite = AvionicsSuiteDelegate().get(id).first();	
				# add the AvionicsSuite
				supplier.avionicsSuites.remove(avionicsSuite)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except AvionicsSuite.DoesNotExist:
			raise ProcessingError(errMsg + " : AvionicsSuite does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addApus( self, supplierId, apusIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.APUDelegate import APUDelegate

		errMsg = "Failed to add elements " + str(apusIds) + " for Apus on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = apusIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the APU		
				aPU = APUDelegate().get(id).first();	
				# add the APU
				supplier.apus.add(aPU)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except APU.DoesNotExist:
			raise ProcessingError(errMsg + " : APU does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApus( self, supplierId, apusIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.APUDelegate import APUDelegate

		errMsg = "Failed to remove elements " + str(apusIds) + " for Apus on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = apusIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the APU		
				aPU = APUDelegate().get(id).first();	
				# add the APU
				supplier.apus.remove(aPU)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except APU.DoesNotExist:
			raise ProcessingError(errMsg + " : APU does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLandingGears( self, supplierId, landingGearsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.LandingGearDelegate import LandingGearDelegate

		errMsg = "Failed to add elements " + str(landingGearsIds) + " for LandingGears on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = landingGearsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LandingGear		
				landingGear = LandingGearDelegate().get(id).first();	
				# add the LandingGear
				supplier.landingGears.add(landingGear)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except LandingGear.DoesNotExist:
			raise ProcessingError(errMsg + " : LandingGear does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLandingGears( self, supplierId, landingGearsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.LandingGearDelegate import LandingGearDelegate

		errMsg = "Failed to remove elements " + str(landingGearsIds) + " for LandingGears on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = landingGearsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LandingGear		
				landingGear = LandingGearDelegate().get(id).first();	
				# add the LandingGear
				supplier.landingGears.remove(landingGear)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except LandingGear.DoesNotExist:
			raise ProcessingError(errMsg + " : LandingGear does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
