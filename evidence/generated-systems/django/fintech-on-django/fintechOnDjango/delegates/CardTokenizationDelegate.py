from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.CardTokenization import CardTokenization
from fintechOnDjango.models.PaymentCard import PaymentCard
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CardTokenization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CardTokenizationDelegate Declaration
#======================================================================
class CardTokenizationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, cardTokenizationId ):
		try:	
			cardTokenization = CardTokenization.objects.filter(id=cardTokenizationId)
			return cardTokenization.first();
		except CardTokenization.DoesNotExist:
			raise ProcessingError("CardTokenization with id " + str(cardTokenizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, cardTokenization):
		for model in serializers.deserialize("json", cardTokenization):
			model.save()
			return model;

	def create(self, cardTokenization):
		cardTokenization.save()
		return cardTokenization;

	def saveFromJson(self, cardTokenization):
		for model in serializers.deserialize("json", cardTokenization):
			model.save()
			return cardTokenization;
	
	def save(self, cardTokenization):
		cardTokenization.save()
		return cardTokenization;
	
	def delete(self, cardTokenizationId ):
		errMsg = "Failed to delete CardTokenization from db using id " + str(cardTokenizationId)
		
		try:
			cardTokenization = CardTokenization.objects.get(id=cardTokenizationId)
			cardTokenization.delete()
			return True
		except CardTokenization.DoesNotExist:
			raise ProcessingError("CardTokenization with id " + str(cardTokenizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CardTokenization.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CardTokenization from db")
		except Exception:
			return None;
		
	def assignCard( self, cardTokenizationId, cardId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

		errMsg = "Failed to assign element " + str(cardId) + " for Card on CardTokenization"

		try:
			# get the CardTokenization from db
			cardTokenization = self.get( cardTokenizationId ).first()	
			
			# get the PaymentCard from db
			paymentCard = PaymentCardDelegate().get(cardId).first();
			
			# assign the Card		
			cardTokenization.card = paymentCard
			
			#save it
			cardTokenization.save()

			# reload and return the appropriate version					
			return self.get( cardTokenizationId );
		except CardTokenization.DoesNotExist:
			raise ProcessingError(errMsg + " : CardTokenization with id " + str(cardTokenizationId) + " does not exist.")
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(cardId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCard( self, cardTokenizationId ):
		errMsg = "Failed to unassign element " + str(cardId) + " for Card on CardTokenization"

		try:
			# get the CardTokenization from db
			cardTokenization = self.get( cardTokenizationId ).first()	
			
			# assign to None for unassignment
			cardTokenization.paymentCard = None			

			#save it
			cardTokenization.save()

			# reload and return the appropriate version					
			return self.get( cardTokenizationId );
		except CardTokenization.DoesNotExist:
			raise ProcessingError(errMsg + " : CardTokenization with id " + str(cardTokenizationId) + " does not exist.")
		except Exception:
			return None;
		
