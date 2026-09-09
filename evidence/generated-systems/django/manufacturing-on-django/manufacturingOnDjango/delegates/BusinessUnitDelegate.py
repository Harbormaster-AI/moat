from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.BusinessUnit import BusinessUnit
from manufacturingOnDjango.models.Enterprise import Enterprise
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BusinessUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessUnitDelegate Declaration
#======================================================================
class BusinessUnitDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, businessUnitId ):
		try:	
			businessUnit = BusinessUnit.objects.filter(id=businessUnitId)
			return businessUnit.first();
		except BusinessUnit.DoesNotExist:
			raise ProcessingError("BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, businessUnit):
		for model in serializers.deserialize("json", businessUnit):
			model.save()
			return model;

	def create(self, businessUnit):
		businessUnit.save()
		return businessUnit;

	def saveFromJson(self, businessUnit):
		for model in serializers.deserialize("json", businessUnit):
			model.save()
			return businessUnit;
	
	def save(self, businessUnit):
		businessUnit.save()
		return businessUnit;
	
	def delete(self, businessUnitId ):
		errMsg = "Failed to delete BusinessUnit from db using id " + str(businessUnitId)
		
		try:
			businessUnit = BusinessUnit.objects.get(id=businessUnitId)
			businessUnit.delete()
			return True
		except BusinessUnit.DoesNotExist:
			raise ProcessingError("BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BusinessUnit.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BusinessUnit from db")
		except Exception:
			return None;
		
	def assignEnterprise( self, businessUnitId, enterpriseId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

		errMsg = "Failed to assign element " + str(enterpriseId) + " for Enterprise on BusinessUnit"

		try:
			# get the BusinessUnit from db
			businessUnit = self.get( businessUnitId ).first()	
			
			# get the Enterprise from db
			enterprise = EnterpriseDelegate().get(enterpriseId).first();
			
			# assign the Enterprise		
			businessUnit.enterprise = enterprise
			
			#save it
			businessUnit.save()

			# reload and return the appropriate version					
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise with id " + str(enterpriseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEnterprise( self, businessUnitId ):
		errMsg = "Failed to unassign element " + str(enterpriseId) + " for Enterprise on BusinessUnit"

		try:
			# get the BusinessUnit from db
			businessUnit = self.get( businessUnitId ).first()	
			
			# assign to None for unassignment
			businessUnit.enterprise = None			

			#save it
			businessUnit.save()

			# reload and return the appropriate version					
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Exception:
			return None;
		
	def addItems( self, businessUnitId, itemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to add elements " + str(itemsIds) + " for Items on BusinessUnit"

		try:
			# get the BusinessUnit
			businessUnit = self.get( businessUnitId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Item		
				item = ItemDelegate().get(id).first();	
				# add the Item
				businessUnit.items.add(item)
				
			# save it		
			businessUnit.save()
			
			# reload and return the appropriate version
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeItems( self, businessUnitId, itemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to remove elements " + str(itemsIds) + " for Items on BusinessUnit"

		try:
			# get the BusinessUnit
			businessUnit = self.get( businessUnitId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Item		
				item = ItemDelegate().get(id).first();	
				# add the Item
				businessUnit.items.remove(item)
				
			# save it		
			businessUnit.save()
			
			# reload and return the appropriate version
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPlants( self, businessUnitId, plantsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to add elements " + str(plantsIds) + " for Plants on BusinessUnit"

		try:
			# get the BusinessUnit
			businessUnit = self.get( businessUnitId ).first()
				
			# split on a comma with no spaces
			idList = plantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Plant		
				plant = PlantDelegate().get(id).first();	
				# add the Plant
				businessUnit.plants.add(plant)
				
			# save it		
			businessUnit.save()
			
			# reload and return the appropriate version
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlants( self, businessUnitId, plantsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to remove elements " + str(plantsIds) + " for Plants on BusinessUnit"

		try:
			# get the BusinessUnit
			businessUnit = self.get( businessUnitId ).first()
				
			# split on a comma with no spaces
			idList = plantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Plant		
				plant = PlantDelegate().get(id).first();	
				# add the Plant
				businessUnit.plants.remove(plant)
				
			# save it		
			businessUnit.save()
			
			# reload and return the appropriate version
			return self.get( businessUnitId );
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
