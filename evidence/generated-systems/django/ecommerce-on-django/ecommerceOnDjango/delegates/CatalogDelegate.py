from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Catalog import Catalog
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.Category import Category
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Catalog
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CatalogDelegate Declaration
#======================================================================
class CatalogDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, catalogId ):
		try:	
			catalog = Catalog.objects.filter(id=catalogId)
			return catalog.first();
		except Catalog.DoesNotExist:
			raise ProcessingError("Catalog with id " + str(catalogId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, catalog):
		for model in serializers.deserialize("json", catalog):
			model.save()
			return model;

	def create(self, catalog):
		catalog.save()
		return catalog;

	def saveFromJson(self, catalog):
		for model in serializers.deserialize("json", catalog):
			model.save()
			return catalog;
	
	def save(self, catalog):
		catalog.save()
		return catalog;
	
	def delete(self, catalogId ):
		errMsg = "Failed to delete Catalog from db using id " + str(catalogId)
		
		try:
			catalog = Catalog.objects.get(id=catalogId)
			catalog.delete()
			return True
		except Catalog.DoesNotExist:
			raise ProcessingError("Catalog with id " + str(catalogId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Catalog.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Catalog from db")
		except Exception:
			return None;
		
	def assignChannel( self, catalogId, channelId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to assign element " + str(channelId) + " for Channel on Catalog"

		try:
			# get the Catalog from db
			catalog = self.get( catalogId ).first()	
			
			# get the Channel from db
			channel = ChannelDelegate().get(channelId).first();
			
			# assign the Channel		
			catalog.channel = channel
			
			#save it
			catalog.save()

			# reload and return the appropriate version					
			return self.get( catalogId );
		except Catalog.DoesNotExist:
			raise ProcessingError(errMsg + " : Catalog with id " + str(catalogId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignChannel( self, catalogId ):
		errMsg = "Failed to unassign element " + str(channelId) + " for Channel on Catalog"

		try:
			# get the Catalog from db
			catalog = self.get( catalogId ).first()	
			
			# assign to None for unassignment
			catalog.channel = None			

			#save it
			catalog.save()

			# reload and return the appropriate version					
			return self.get( catalogId );
		except Catalog.DoesNotExist:
			raise ProcessingError(errMsg + " : Catalog with id " + str(catalogId) + " does not exist.")
		except Exception:
			return None;
		
	def addCategories( self, catalogId, categoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to add elements " + str(categoriesIds) + " for Categories on Catalog"

		try:
			# get the Catalog
			catalog = self.get( catalogId ).first()
				
			# split on a comma with no spaces
			idList = categoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				catalog.categories.add(category)
				
			# save it		
			catalog.save()
			
			# reload and return the appropriate version
			return self.get( catalogId );
		except Catalog.DoesNotExist:
			raise ProcessingError(errMsg + " : Catalog with id " + str(catalogId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCategories( self, catalogId, categoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to remove elements " + str(categoriesIds) + " for Categories on Catalog"

		try:
			# get the Catalog
			catalog = self.get( catalogId ).first()
				
			# split on a comma with no spaces
			idList = categoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				catalog.categories.remove(category)
				
			# save it		
			catalog.save()
			
			# reload and return the appropriate version
			return self.get( catalogId );
		except Catalog.DoesNotExist:
			raise ProcessingError(errMsg + " : Catalog with id " + str(catalogId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
