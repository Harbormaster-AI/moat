from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.UnderwritingDecision import UnderwritingDecision
from insuranceOnDjango.models.Quote import Quote
from insuranceOnDjango.models.Underwriter import Underwriter
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model UnderwritingDecision
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwritingDecisionDelegate Declaration
#======================================================================
class UnderwritingDecisionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, underwritingDecisionId ):
		try:	
			underwritingDecision = UnderwritingDecision.objects.filter(id=underwritingDecisionId)
			return underwritingDecision.first();
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError("UnderwritingDecision with id " + str(underwritingDecisionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, underwritingDecision):
		for model in serializers.deserialize("json", underwritingDecision):
			model.save()
			return model;

	def create(self, underwritingDecision):
		underwritingDecision.save()
		return underwritingDecision;

	def saveFromJson(self, underwritingDecision):
		for model in serializers.deserialize("json", underwritingDecision):
			model.save()
			return underwritingDecision;
	
	def save(self, underwritingDecision):
		underwritingDecision.save()
		return underwritingDecision;
	
	def delete(self, underwritingDecisionId ):
		errMsg = "Failed to delete UnderwritingDecision from db using id " + str(underwritingDecisionId)
		
		try:
			underwritingDecision = UnderwritingDecision.objects.get(id=underwritingDecisionId)
			underwritingDecision.delete()
			return True
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError("UnderwritingDecision with id " + str(underwritingDecisionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = UnderwritingDecision.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all UnderwritingDecision from db")
		except Exception:
			return None;
		
	def assignQuote( self, underwritingDecisionId, quoteId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to assign element " + str(quoteId) + " for Quote on UnderwritingDecision"

		try:
			# get the UnderwritingDecision from db
			underwritingDecision = self.get( underwritingDecisionId ).first()	
			
			# get the Quote from db
			quote = QuoteDelegate().get(quoteId).first();
			
			# assign the Quote		
			underwritingDecision.quote = quote
			
			#save it
			underwritingDecision.save()

			# reload and return the appropriate version					
			return self.get( underwritingDecisionId );
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision with id " + str(underwritingDecisionId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignQuote( self, underwritingDecisionId ):
		errMsg = "Failed to unassign element " + str(quoteId) + " for Quote on UnderwritingDecision"

		try:
			# get the UnderwritingDecision from db
			underwritingDecision = self.get( underwritingDecisionId ).first()	
			
			# assign to None for unassignment
			underwritingDecision.quote = None			

			#save it
			underwritingDecision.save()

			# reload and return the appropriate version					
			return self.get( underwritingDecisionId );
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision with id " + str(underwritingDecisionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignUnderwriter( self, underwritingDecisionId, underwriterId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.UnderwriterDelegate import UnderwriterDelegate

		errMsg = "Failed to assign element " + str(underwriterId) + " for Underwriter on UnderwritingDecision"

		try:
			# get the UnderwritingDecision from db
			underwritingDecision = self.get( underwritingDecisionId ).first()	
			
			# get the Underwriter from db
			underwriter = UnderwriterDelegate().get(underwriterId).first();
			
			# assign the Underwriter		
			underwritingDecision.underwriter = underwriter
			
			#save it
			underwritingDecision.save()

			# reload and return the appropriate version					
			return self.get( underwritingDecisionId );
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision with id " + str(underwritingDecisionId) + " does not exist.")
		except Underwriter.DoesNotExist:
			raise ProcessingError(errMsg + " : Underwriter with id " + str(underwriterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignUnderwriter( self, underwritingDecisionId ):
		errMsg = "Failed to unassign element " + str(underwriterId) + " for Underwriter on UnderwritingDecision"

		try:
			# get the UnderwritingDecision from db
			underwritingDecision = self.get( underwritingDecisionId ).first()	
			
			# assign to None for unassignment
			underwritingDecision.underwriter = None			

			#save it
			underwritingDecision.save()

			# reload and return the appropriate version					
			return self.get( underwritingDecisionId );
		except UnderwritingDecision.DoesNotExist:
			raise ProcessingError(errMsg + " : UnderwritingDecision with id " + str(underwritingDecisionId) + " does not exist.")
		except Exception:
			return None;
		
