from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Category import Category
from ecommerceOnDjango.models.Catalog import Catalog
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Category
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CategoryDelegate Declaration
#======================================================================
class CategoryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, categoryId ):
		try:	
			category = Category.objects.filter(id=categoryId)
			return category.first();
		except Category.DoesNotExist:
			raise ProcessingError("Category with id " + str(categoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, category):
		for model in serializers.deserialize("json", category):
			model.save()
			return model;

	def create(self, category):
		category.save()
		return category;

	def saveFromJson(self, category):
		for model in serializers.deserialize("json", category):
			model.save()
			return category;
	
	def save(self, category):
		category.save()
		return category;
	
	def delete(self, categoryId ):
		errMsg = "Failed to delete Category from db using id " + str(categoryId)
		
		try:
			category = Category.objects.get(id=categoryId)
			category.delete()
			return True
		except Category.DoesNotExist:
			raise ProcessingError("Category with id " + str(categoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Category.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Category from db")
		except Exception:
			return None;
		
	def assignCatalog( self, categoryId, catalogId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CatalogDelegate import CatalogDelegate

		errMsg = "Failed to assign element " + str(catalogId) + " for Catalog on Category"

		try:
			# get the Category from db
			category = self.get( categoryId ).first()	
			
			# get the Catalog from db
			catalog = CatalogDelegate().get(catalogId).first();
			
			# assign the Catalog		
			category.catalog = catalog
			
			#save it
			category.save()

			# reload and return the appropriate version					
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Catalog.DoesNotExist:
			raise ProcessingError(errMsg + " : Catalog with id " + str(catalogId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCatalog( self, categoryId ):
		errMsg = "Failed to unassign element " + str(catalogId) + " for Catalog on Category"

		try:
			# get the Category from db
			category = self.get( categoryId ).first()	
			
			# assign to None for unassignment
			category.catalog = None			

			#save it
			category.save()

			# reload and return the appropriate version					
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignParentCategory( self, categoryId, parentCategoryId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to assign element " + str(parentCategoryId) + " for ParentCategory on Category"

		try:
			# get the Category from db
			category = self.get( categoryId ).first()	
			
			# get the Category from db
			category = CategoryDelegate().get(parentCategoryId).first();
			
			# assign the ParentCategory		
			category.parentCategory = category
			
			#save it
			category.save()

			# reload and return the appropriate version					
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(parentCategoryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignParentCategory( self, categoryId ):
		errMsg = "Failed to unassign element " + str(parentCategoryId) + " for ParentCategory on Category"

		try:
			# get the Category from db
			category = self.get( categoryId ).first()	
			
			# assign to None for unassignment
			category.category = None			

			#save it
			category.save()

			# reload and return the appropriate version					
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Exception:
			return None;
		
	def addSubcategories( self, categoryId, subcategoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to add elements " + str(subcategoriesIds) + " for Subcategories on Category"

		try:
			# get the Category
			category = self.get( categoryId ).first()
				
			# split on a comma with no spaces
			idList = subcategoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				category.subcategories.add(category)
				
			# save it		
			category.save()
			
			# reload and return the appropriate version
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSubcategories( self, categoryId, subcategoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to remove elements " + str(subcategoriesIds) + " for Subcategories on Category"

		try:
			# get the Category
			category = self.get( categoryId ).first()
				
			# split on a comma with no spaces
			idList = subcategoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				category.subcategories.remove(category)
				
			# save it		
			category.save()
			
			# reload and return the appropriate version
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProducts( self, categoryId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to add elements " + str(productsIds) + " for Products on Category"

		try:
			# get the Category
			category = self.get( categoryId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				category.products.add(product)
				
			# save it		
			category.save()
			
			# reload and return the appropriate version
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProducts( self, categoryId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to remove elements " + str(productsIds) + " for Products on Category"

		try:
			# get the Category
			category = self.get( categoryId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				category.products.remove(product)
				
			# save it		
			category.save()
			
			# reload and return the appropriate version
			return self.get( categoryId );
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category with id " + str(categoryId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
