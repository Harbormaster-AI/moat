from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Quote import Quote
from insuranceOnDjango.models.Application import Application
from insuranceOnDjango.models.UnderwritingDecision import UnderwritingDecision
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Quote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteDelegate Declaration
#======================================================================
class QuoteDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, quoteId ):
		try:	
			quote = Quote.objects.filter(id=quoteId)
			return quote.first();
		except Quote.DoesNotExist:
			raise ProcessingError("Quote with id " + str(quoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, quote):
		for model in serializers.deserialize("json", quote):
			model.save()
			return model;

	def create(self, quote):
		quote.save()
		return quote;

	def saveFromJson(self, quote):
		for model in serializers.deserialize("json", quote):
			model.save()
			return quote;
	
	def save(self, quote):
		quote.save()
		return quote;
	
	def delete(self, quoteId ):
		errMsg = "Failed to delete Quote from db using id " + str(quoteId)
		
		try:
			quote = Quote.objects.get(id=quoteId)
			quote.delete()
			return True
		except Quote.DoesNotExist:
			raise ProcessingError("Quote with id " + str(quoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Quote.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Quote from db")
		except Exception:
			return None;
		
	def assignApplication( self, quoteId, applicationId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ApplicationDelegate import ApplicationDelegate

		errMsg = "Failed to assign element " + str(applicationId) + " for Application on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the Application from db
			application = ApplicationDelegate().get(applicationId).first();
			
			# assign the Application		
			quote.application = application
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Application.DoesNotExist:
			raise ProcessingError(errMsg + " : Application with id " + str(applicationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignApplication( self, quoteId ):
		errMsg = "Failed to unassign element " + str(applicationId) + " for Application on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.application = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPolicy( self, quoteId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			quote.policy = policy
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, quoteId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.policy = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def addUnderwritingDecisions( self, quoteId, underwritingDecisionsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.UnderwritingDecisionDelegate import UnderwritingDecisionDelegate

		errMsg = "Failed to add elements " + str(underwritingDecisionsIds) + " for UnderwritingDecisions on Quote"

		try:
			# get the Quote
			quote = self.get( quoteId ).first()
				
			# split on a comma with no spaces
			idList = underwritingDecisionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the UnderwritingDecision		
				underwritingDecision = UnderwritingDecisionDelegate().get(id).first();	
				# add the UnderwritingDecision
				quote.underwritingDecisions.add(underwritingDecision)
				
			# save it		
			quote.save()
			
			# reload and return the appropriate version
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUnderwritingDecisions( self, quoteId, underwritingDecisionsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.UnderwritingDecisionDelegate import UnderwritingDecisionDelegate

		errMsg = "Failed to remove elements " + str(underwritingDecisionsIds) + " for UnderwritingDecisions on Quote"

		try:
			# get the Quote
			quote = self.get( quoteId ).first()
				
			# split on a comma with no spaces
			idList = underwritingDecisionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the UnderwritingDecision		
				underwritingDecision = UnderwritingDecisionDelegate().get(id).first();	
				# add the UnderwritingDecision
				quote.underwritingDecisions.remove(underwritingDecision)
				
			# save it		
			quote.save()
			
			# reload and return the appropriate version
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
