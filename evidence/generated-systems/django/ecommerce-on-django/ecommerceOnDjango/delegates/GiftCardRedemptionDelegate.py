from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.GiftCardRedemption import GiftCardRedemption
from ecommerceOnDjango.models.GiftCard import GiftCard
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model GiftCardRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardRedemptionDelegate Declaration
#======================================================================
class GiftCardRedemptionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, giftCardRedemptionId ):
		try:	
			giftCardRedemption = GiftCardRedemption.objects.filter(id=giftCardRedemptionId)
			return giftCardRedemption.first();
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError("GiftCardRedemption with id " + str(giftCardRedemptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, giftCardRedemption):
		for model in serializers.deserialize("json", giftCardRedemption):
			model.save()
			return model;

	def create(self, giftCardRedemption):
		giftCardRedemption.save()
		return giftCardRedemption;

	def saveFromJson(self, giftCardRedemption):
		for model in serializers.deserialize("json", giftCardRedemption):
			model.save()
			return giftCardRedemption;
	
	def save(self, giftCardRedemption):
		giftCardRedemption.save()
		return giftCardRedemption;
	
	def delete(self, giftCardRedemptionId ):
		errMsg = "Failed to delete GiftCardRedemption from db using id " + str(giftCardRedemptionId)
		
		try:
			giftCardRedemption = GiftCardRedemption.objects.get(id=giftCardRedemptionId)
			giftCardRedemption.delete()
			return True
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError("GiftCardRedemption with id " + str(giftCardRedemptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = GiftCardRedemption.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all GiftCardRedemption from db")
		except Exception:
			return None;
		
	def assignGiftCard( self, giftCardRedemptionId, giftCardId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.GiftCardDelegate import GiftCardDelegate

		errMsg = "Failed to assign element " + str(giftCardId) + " for GiftCard on GiftCardRedemption"

		try:
			# get the GiftCardRedemption from db
			giftCardRedemption = self.get( giftCardRedemptionId ).first()	
			
			# get the GiftCard from db
			giftCard = GiftCardDelegate().get(giftCardId).first();
			
			# assign the GiftCard		
			giftCardRedemption.giftCard = giftCard
			
			#save it
			giftCardRedemption.save()

			# reload and return the appropriate version					
			return self.get( giftCardRedemptionId );
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption with id " + str(giftCardRedemptionId) + " does not exist.")
		except GiftCard.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCard with id " + str(giftCardId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignGiftCard( self, giftCardRedemptionId ):
		errMsg = "Failed to unassign element " + str(giftCardId) + " for GiftCard on GiftCardRedemption"

		try:
			# get the GiftCardRedemption from db
			giftCardRedemption = self.get( giftCardRedemptionId ).first()	
			
			# assign to None for unassignment
			giftCardRedemption.giftCard = None			

			#save it
			giftCardRedemption.save()

			# reload and return the appropriate version					
			return self.get( giftCardRedemptionId );
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption with id " + str(giftCardRedemptionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrder( self, giftCardRedemptionId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on GiftCardRedemption"

		try:
			# get the GiftCardRedemption from db
			giftCardRedemption = self.get( giftCardRedemptionId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			giftCardRedemption.order = order
			
			#save it
			giftCardRedemption.save()

			# reload and return the appropriate version					
			return self.get( giftCardRedemptionId );
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption with id " + str(giftCardRedemptionId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, giftCardRedemptionId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on GiftCardRedemption"

		try:
			# get the GiftCardRedemption from db
			giftCardRedemption = self.get( giftCardRedemptionId ).first()	
			
			# assign to None for unassignment
			giftCardRedemption.order = None			

			#save it
			giftCardRedemption.save()

			# reload and return the appropriate version					
			return self.get( giftCardRedemptionId );
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption with id " + str(giftCardRedemptionId) + " does not exist.")
		except Exception:
			return None;
		
