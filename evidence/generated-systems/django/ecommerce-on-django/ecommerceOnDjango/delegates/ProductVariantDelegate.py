from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.models.ProductPricing import ProductPricing
from ecommerceOnDjango.models.InventoryItem import InventoryItem
from ecommerceOnDjango.models.MediaAsset import MediaAsset
from ecommerceOnDjango.models.Subscription import Subscription
from ecommerceOnDjango.models.CartItem import CartItem
from ecommerceOnDjango.models.OrderLine import OrderLine
from ecommerceOnDjango.models.WishlistItem import WishlistItem
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProductVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductVariantDelegate Declaration
#======================================================================
class ProductVariantDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productVariantId ):
		try:	
			productVariant = ProductVariant.objects.filter(id=productVariantId)
			return productVariant.first();
		except ProductVariant.DoesNotExist:
			raise ProcessingError("ProductVariant with id " + str(productVariantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, productVariant):
		for model in serializers.deserialize("json", productVariant):
			model.save()
			return model;

	def create(self, productVariant):
		productVariant.save()
		return productVariant;

	def saveFromJson(self, productVariant):
		for model in serializers.deserialize("json", productVariant):
			model.save()
			return productVariant;
	
	def save(self, productVariant):
		productVariant.save()
		return productVariant;
	
	def delete(self, productVariantId ):
		errMsg = "Failed to delete ProductVariant from db using id " + str(productVariantId)
		
		try:
			productVariant = ProductVariant.objects.get(id=productVariantId)
			productVariant.delete()
			return True
		except ProductVariant.DoesNotExist:
			raise ProcessingError("ProductVariant with id " + str(productVariantId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProductVariant.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProductVariant from db")
		except Exception:
			return None;
		
	def assignProduct( self, productVariantId, productId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on ProductVariant"

		try:
			# get the ProductVariant from db
			productVariant = self.get( productVariantId ).first()	
			
			# get the Product from db
			product = ProductDelegate().get(productId).first();
			
			# assign the Product		
			productVariant.product = product
			
			#save it
			productVariant.save()

			# reload and return the appropriate version					
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, productVariantId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on ProductVariant"

		try:
			# get the ProductVariant from db
			productVariant = self.get( productVariantId ).first()	
			
			# assign to None for unassignment
			productVariant.product = None			

			#save it
			productVariant.save()

			# reload and return the appropriate version					
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except Exception:
			return None;
		
	def addPricing( self, productVariantId, pricingIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductPricingDelegate import ProductPricingDelegate

		errMsg = "Failed to add elements " + str(pricingIds) + " for Pricing on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = pricingIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ProductPricing		
				productPricing = ProductPricingDelegate().get(id).first();	
				# add the ProductPricing
				productVariant.pricing.add(productPricing)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except ProductPricing.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductPricing does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePricing( self, productVariantId, pricingIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductPricingDelegate import ProductPricingDelegate

		errMsg = "Failed to remove elements " + str(pricingIds) + " for Pricing on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = pricingIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ProductPricing		
				productPricing = ProductPricingDelegate().get(id).first();	
				# add the ProductPricing
				productVariant.pricing.remove(productPricing)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except ProductPricing.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductPricing does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInventoryItems( self, productVariantId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				productVariant.inventoryItems.add(inventoryItem)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, productVariantId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				productVariant.inventoryItems.remove(inventoryItem)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMediaAssets( self, productVariantId, mediaAssetsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MediaAssetDelegate import MediaAssetDelegate

		errMsg = "Failed to add elements " + str(mediaAssetsIds) + " for MediaAssets on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = mediaAssetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MediaAsset		
				mediaAsset = MediaAssetDelegate().get(id).first();	
				# add the MediaAsset
				productVariant.mediaAssets.add(mediaAsset)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMediaAssets( self, productVariantId, mediaAssetsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MediaAssetDelegate import MediaAssetDelegate

		errMsg = "Failed to remove elements " + str(mediaAssetsIds) + " for MediaAssets on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = mediaAssetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MediaAsset		
				mediaAsset = MediaAssetDelegate().get(id).first();	
				# add the MediaAsset
				productVariant.mediaAssets.remove(mediaAsset)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSubscriptions( self, productVariantId, subscriptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

		errMsg = "Failed to add elements " + str(subscriptionsIds) + " for Subscriptions on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = subscriptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Subscription		
				subscription = SubscriptionDelegate().get(id).first();	
				# add the Subscription
				productVariant.subscriptions.add(subscription)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSubscriptions( self, productVariantId, subscriptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

		errMsg = "Failed to remove elements " + str(subscriptionsIds) + " for Subscriptions on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = subscriptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Subscription		
				subscription = SubscriptionDelegate().get(id).first();	
				# add the Subscription
				productVariant.subscriptions.remove(subscription)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except Subscription.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscription does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCartItems( self, productVariantId, cartItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CartItemDelegate import CartItemDelegate

		errMsg = "Failed to add elements " + str(cartItemsIds) + " for CartItems on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = cartItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CartItem		
				cartItem = CartItemDelegate().get(id).first();	
				# add the CartItem
				productVariant.cartItems.add(cartItem)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCartItems( self, productVariantId, cartItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CartItemDelegate import CartItemDelegate

		errMsg = "Failed to remove elements " + str(cartItemsIds) + " for CartItems on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = cartItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CartItem		
				cartItem = CartItemDelegate().get(id).first();	
				# add the CartItem
				productVariant.cartItems.remove(cartItem)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrderLines( self, productVariantId, orderLinesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

		errMsg = "Failed to add elements " + str(orderLinesIds) + " for OrderLines on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = orderLinesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OrderLine		
				orderLine = OrderLineDelegate().get(id).first();	
				# add the OrderLine
				productVariant.orderLines.add(orderLine)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrderLines( self, productVariantId, orderLinesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

		errMsg = "Failed to remove elements " + str(orderLinesIds) + " for OrderLines on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = orderLinesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OrderLine		
				orderLine = OrderLineDelegate().get(id).first();	
				# add the OrderLine
				productVariant.orderLines.remove(orderLine)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWishlistItems( self, productVariantId, wishlistItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.WishlistItemDelegate import WishlistItemDelegate

		errMsg = "Failed to add elements " + str(wishlistItemsIds) + " for WishlistItems on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = wishlistItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WishlistItem		
				wishlistItem = WishlistItemDelegate().get(id).first();	
				# add the WishlistItem
				productVariant.wishlistItems.add(wishlistItem)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWishlistItems( self, productVariantId, wishlistItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.WishlistItemDelegate import WishlistItemDelegate

		errMsg = "Failed to remove elements " + str(wishlistItemsIds) + " for WishlistItems on ProductVariant"

		try:
			# get the ProductVariant
			productVariant = self.get( productVariantId ).first()
				
			# split on a comma with no spaces
			idList = wishlistItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WishlistItem		
				wishlistItem = WishlistItemDelegate().get(id).first();	
				# add the WishlistItem
				productVariant.wishlistItems.remove(wishlistItem)
				
			# save it		
			productVariant.save()
			
			# reload and return the appropriate version
			return self.get( productVariantId );
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(productVariantId) + " does not exist.")
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
