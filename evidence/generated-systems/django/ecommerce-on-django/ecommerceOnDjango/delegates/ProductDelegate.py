from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.models.Brand import Brand
from ecommerceOnDjango.models.Category import Category
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.models.MediaAsset import MediaAsset
from ecommerceOnDjango.models.Review import Review
from ecommerceOnDjango.models.Seller import Seller
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Product
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductDelegate Declaration
#======================================================================
class ProductDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productId ):
		try:	
			product = Product.objects.filter(id=productId)
			return product.first();
		except Product.DoesNotExist:
			raise ProcessingError("Product with id " + str(productId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, product):
		for model in serializers.deserialize("json", product):
			model.save()
			return model;

	def create(self, product):
		product.save()
		return product;

	def saveFromJson(self, product):
		for model in serializers.deserialize("json", product):
			model.save()
			return product;
	
	def save(self, product):
		product.save()
		return product;
	
	def delete(self, productId ):
		errMsg = "Failed to delete Product from db using id " + str(productId)
		
		try:
			product = Product.objects.get(id=productId)
			product.delete()
			return True
		except Product.DoesNotExist:
			raise ProcessingError("Product with id " + str(productId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Product.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Product from db")
		except Exception:
			return None;
		
	def assignBrand( self, productId, brandId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.BrandDelegate import BrandDelegate

		errMsg = "Failed to assign element " + str(brandId) + " for Brand on Product"

		try:
			# get the Product from db
			product = self.get( productId ).first()	
			
			# get the Brand from db
			brand = BrandDelegate().get(brandId).first();
			
			# assign the Brand		
			product.brand = brand
			
			#save it
			product.save()

			# reload and return the appropriate version					
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Brand.DoesNotExist:
			raise ProcessingError(errMsg + " : Brand with id " + str(brandId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBrand( self, productId ):
		errMsg = "Failed to unassign element " + str(brandId) + " for Brand on Product"

		try:
			# get the Product from db
			product = self.get( productId ).first()	
			
			# assign to None for unassignment
			product.brand = None			

			#save it
			product.save()

			# reload and return the appropriate version					
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSeller( self, productId, sellerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SellerDelegate import SellerDelegate

		errMsg = "Failed to assign element " + str(sellerId) + " for Seller on Product"

		try:
			# get the Product from db
			product = self.get( productId ).first()	
			
			# get the Seller from db
			seller = SellerDelegate().get(sellerId).first();
			
			# assign the Seller		
			product.seller = seller
			
			#save it
			product.save()

			# reload and return the appropriate version					
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSeller( self, productId ):
		errMsg = "Failed to unassign element " + str(sellerId) + " for Seller on Product"

		try:
			# get the Product from db
			product = self.get( productId ).first()	
			
			# assign to None for unassignment
			product.seller = None			

			#save it
			product.save()

			# reload and return the appropriate version					
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
		
	def addCategories( self, productId, categoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to add elements " + str(categoriesIds) + " for Categories on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = categoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				product.categories.add(category)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCategories( self, productId, categoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to remove elements " + str(categoriesIds) + " for Categories on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = categoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				product.categories.remove(category)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addVariants( self, productId, variantsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProductVariant		
				productVariant = ProductVariantDelegate().get(id).first();	
				# add the ProductVariant
				product.variants.add(productVariant)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, productId, variantsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProductVariant		
				productVariant = ProductVariantDelegate().get(id).first();	
				# add the ProductVariant
				product.variants.remove(productVariant)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMediaAssets( self, productId, mediaAssetsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MediaAssetDelegate import MediaAssetDelegate

		errMsg = "Failed to add elements " + str(mediaAssetsIds) + " for MediaAssets on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = mediaAssetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MediaAsset		
				mediaAsset = MediaAssetDelegate().get(id).first();	
				# add the MediaAsset
				product.mediaAssets.add(mediaAsset)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMediaAssets( self, productId, mediaAssetsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MediaAssetDelegate import MediaAssetDelegate

		errMsg = "Failed to remove elements " + str(mediaAssetsIds) + " for MediaAssets on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = mediaAssetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MediaAsset		
				mediaAsset = MediaAssetDelegate().get(id).first();	
				# add the MediaAsset
				product.mediaAssets.remove(mediaAsset)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReviews( self, productId, reviewsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReviewDelegate import ReviewDelegate

		errMsg = "Failed to add elements " + str(reviewsIds) + " for Reviews on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = reviewsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Review		
				review = ReviewDelegate().get(id).first();	
				# add the Review
				product.reviews.add(review)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReviews( self, productId, reviewsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReviewDelegate import ReviewDelegate

		errMsg = "Failed to remove elements " + str(reviewsIds) + " for Reviews on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = reviewsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Review		
				review = ReviewDelegate().get(id).first();	
				# add the Review
				product.reviews.remove(review)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Review.DoesNotExist:
			raise ProcessingError(errMsg + " : Review does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
