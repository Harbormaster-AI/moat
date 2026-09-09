from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.GiftCard import GiftCard
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.GiftCardRedemption import GiftCardRedemption
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model GiftCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardDelegate Declaration
#======================================================================
class GiftCardDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, giftCardId ):
		try:	
			giftCard = GiftCard.objects.filter(id=giftCardId)
			return giftCard.first();
		except GiftCard.DoesNotExist:
			raise ProcessingError("GiftCard with id " + str(giftCardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, giftCard):
		for model in serializers.deserialize("json", giftCard):
			model.save()
			return model;

	def create(self, giftCard):
		giftCard.save()
		return giftCard;

	def saveFromJson(self, giftCard):
		for model in serializers.deserialize("json", giftCard):
			model.save()
			return giftCard;
	
	def save(self, giftCard):
		giftCard.save()
		return giftCard;
	
	def delete(self, giftCardId ):
		errMsg = "Failed to delete GiftCard from db using id " + str(giftCardId)
		
		try:
			giftCard = GiftCard.objects.get(id=giftCardId)
			giftCard.delete()
			return True
		except GiftCard.DoesNotExist:
			raise ProcessingError("GiftCard with id " + str(giftCardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = GiftCard.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all GiftCard from db")
		except Exception:
			return None;
		
	def assignCustomer( self, giftCardId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on GiftCard"

		try:
			# get the GiftCard from db
			giftCard = self.get( giftCardId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			giftCard.customer = customer
			
			#save it
			giftCard.save()

			# reload and return the appropriate version					
			return self.get( giftCardId );
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard with id " + str(giftCardId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, giftCardId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on GiftCard"

		try:
			# get the GiftCard from db
			giftCard = self.get( giftCardId ).first()	
			
			# assign to None for unassignment
			giftCard.customer = None			

			#save it
			giftCard.save()

			# reload and return the appropriate version					
			return self.get( giftCardId );
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard with id " + str(giftCardId) + " does not exist.")
		except Exception:
			return None;
		
	def assignIssuedOrder( self, giftCardId, issuedOrderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(issuedOrderId) + " for IssuedOrder on GiftCard"

		try:
			# get the GiftCard from db
			giftCard = self.get( giftCardId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(issuedOrderId).first();
			
			# assign the IssuedOrder		
			giftCard.issuedOrder = order
			
			#save it
			giftCard.save()

			# reload and return the appropriate version					
			return self.get( giftCardId );
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard with id " + str(giftCardId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(issuedOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignIssuedOrder( self, giftCardId ):
		errMsg = "Failed to unassign element " + str(issuedOrderId) + " for IssuedOrder on GiftCard"

		try:
			# get the GiftCard from db
			giftCard = self.get( giftCardId ).first()	
			
			# assign to None for unassignment
			giftCard.order = None			

			#save it
			giftCard.save()

			# reload and return the appropriate version					
			return self.get( giftCardId );
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard with id " + str(giftCardId) + " does not exist.")
		except Exception:
			return None;
		
	def addRedemptions( self, giftCardId, redemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.GiftCardRedemptionDelegate import GiftCardRedemptionDelegate

		errMsg = "Failed to add elements " + str(redemptionsIds) + " for Redemptions on GiftCard"

		try:
			# get the GiftCard
			giftCard = self.get( giftCardId ).first()
				
			# split on a comma with no spaces
			idList = redemptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GiftCardRedemption		
				giftCardRedemption = GiftCardRedemptionDelegate().get(id).first();	
				# add the GiftCardRedemption
				giftCard.redemptions.add(giftCardRedemption)
				
			# save it		
			giftCard.save()
			
			# reload and return the appropriate version
			return self.get( giftCardId );
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard with id " + str(giftCardId) + " does not exist.")
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRedemptions( self, giftCardId, redemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.GiftCardRedemptionDelegate import GiftCardRedemptionDelegate

		errMsg = "Failed to remove elements " + str(redemptionsIds) + " for Redemptions on GiftCard"

		try:
			# get the GiftCard
			giftCard = self.get( giftCardId ).first()
				
			# split on a comma with no spaces
			idList = redemptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GiftCardRedemption		
				giftCardRedemption = GiftCardRedemptionDelegate().get(id).first();	
				# add the GiftCardRedemption
				giftCard.redemptions.remove(giftCardRedemption)
				
			# save it		
			giftCard.save()
			
			# reload and return the appropriate version
			return self.get( giftCardId );
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard with id " + str(giftCardId) + " does not exist.")
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
