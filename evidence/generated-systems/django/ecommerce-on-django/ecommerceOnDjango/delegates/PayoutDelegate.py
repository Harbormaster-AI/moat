from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Payout import Payout
from ecommerceOnDjango.models.Seller import Seller
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Payout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayoutDelegate Declaration
#======================================================================
class PayoutDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, payoutId ):
		try:	
			payout = Payout.objects.filter(id=payoutId)
			return payout.first();
		except Payout.DoesNotExist:
			raise ProcessingError("Payout with id " + str(payoutId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, payout):
		for model in serializers.deserialize("json", payout):
			model.save()
			return model;

	def create(self, payout):
		payout.save()
		return payout;

	def saveFromJson(self, payout):
		for model in serializers.deserialize("json", payout):
			model.save()
			return payout;
	
	def save(self, payout):
		payout.save()
		return payout;
	
	def delete(self, payoutId ):
		errMsg = "Failed to delete Payout from db using id " + str(payoutId)
		
		try:
			payout = Payout.objects.get(id=payoutId)
			payout.delete()
			return True
		except Payout.DoesNotExist:
			raise ProcessingError("Payout with id " + str(payoutId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Payout.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Payout from db")
		except Exception:
			return None;
		
	def assignSeller( self, payoutId, sellerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SellerDelegate import SellerDelegate

		errMsg = "Failed to assign element " + str(sellerId) + " for Seller on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# get the Seller from db
			seller = SellerDelegate().get(sellerId).first();
			
			# assign the Seller		
			payout.seller = seller
			
			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSeller( self, payoutId ):
		errMsg = "Failed to unassign element " + str(sellerId) + " for Seller on Payout"

		try:
			# get the Payout from db
			payout = self.get( payoutId ).first()	
			
			# assign to None for unassignment
			payout.seller = None			

			#save it
			payout.save()

			# reload and return the appropriate version					
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Exception:
			return None;
		
	def addOrders( self, payoutId, ordersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Payout"

		try:
			# get the Payout
			payout = self.get( payoutId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				payout.orders.add(order)
				
			# save it		
			payout.save()
			
			# reload and return the appropriate version
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, payoutId, ordersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Payout"

		try:
			# get the Payout
			payout = self.get( payoutId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				payout.orders.remove(order)
				
			# save it		
			payout.save()
			
			# reload and return the appropriate version
			return self.get( payoutId );
		except Payout.DoesNotExist:
			raise ProcessingError(errMsg + " : Payout with id " + str(payoutId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
