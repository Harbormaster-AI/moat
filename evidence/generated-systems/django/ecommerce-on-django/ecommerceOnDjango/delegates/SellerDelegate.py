from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Seller import Seller
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.models.Payout import Payout
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Seller
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SellerDelegate Declaration
#======================================================================
class SellerDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, sellerId ):
		try:	
			seller = Seller.objects.filter(id=sellerId)
			return seller.first();
		except Seller.DoesNotExist:
			raise ProcessingError("Seller with id " + str(sellerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, seller):
		for model in serializers.deserialize("json", seller):
			model.save()
			return model;

	def create(self, seller):
		seller.save()
		return seller;

	def saveFromJson(self, seller):
		for model in serializers.deserialize("json", seller):
			model.save()
			return seller;
	
	def save(self, seller):
		seller.save()
		return seller;
	
	def delete(self, sellerId ):
		errMsg = "Failed to delete Seller from db using id " + str(sellerId)
		
		try:
			seller = Seller.objects.get(id=sellerId)
			seller.delete()
			return True
		except Seller.DoesNotExist:
			raise ProcessingError("Seller with id " + str(sellerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Seller.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Seller from db")
		except Exception:
			return None;
		
	def assignMerchant( self, sellerId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Seller"

		try:
			# get the Seller from db
			seller = self.get( sellerId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			seller.merchant = merchant
			
			#save it
			seller.save()

			# reload and return the appropriate version					
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, sellerId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Seller"

		try:
			# get the Seller from db
			seller = self.get( sellerId ).first()	
			
			# assign to None for unassignment
			seller.merchant = None			

			#save it
			seller.save()

			# reload and return the appropriate version					
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Exception:
			return None;
		
	def addProducts( self, sellerId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to add elements " + str(productsIds) + " for Products on Seller"

		try:
			# get the Seller
			seller = self.get( sellerId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				seller.products.add(product)
				
			# save it		
			seller.save()
			
			# reload and return the appropriate version
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProducts( self, sellerId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to remove elements " + str(productsIds) + " for Products on Seller"

		try:
			# get the Seller
			seller = self.get( sellerId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				seller.products.remove(product)
				
			# save it		
			seller.save()
			
			# reload and return the appropriate version
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayouts( self, sellerId, payoutsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PayoutDelegate import PayoutDelegate

		errMsg = "Failed to add elements " + str(payoutsIds) + " for Payouts on Seller"

		try:
			# get the Seller
			seller = self.get( sellerId ).first()
				
			# split on a comma with no spaces
			idList = payoutsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payout		
				payout = PayoutDelegate().get(id).first();	
				# add the Payout
				seller.payouts.add(payout)
				
			# save it		
			seller.save()
			
			# reload and return the appropriate version
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayouts( self, sellerId, payoutsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PayoutDelegate import PayoutDelegate

		errMsg = "Failed to remove elements " + str(payoutsIds) + " for Payouts on Seller"

		try:
			# get the Seller
			seller = self.get( sellerId ).first()
				
			# split on a comma with no spaces
			idList = payoutsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payout		
				payout = PayoutDelegate().get(id).first();	
				# add the Payout
				seller.payouts.remove(payout)
				
			# save it		
			seller.save()
			
			# reload and return the appropriate version
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, sellerId, ordersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Seller"

		try:
			# get the Seller
			seller = self.get( sellerId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				seller.orders.add(order)
				
			# save it		
			seller.save()
			
			# reload and return the appropriate version
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, sellerId, ordersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Seller"

		try:
			# get the Seller
			seller = self.get( sellerId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				seller.orders.remove(order)
				
			# save it		
			seller.save()
			
			# reload and return the appropriate version
			return self.get( sellerId );
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
