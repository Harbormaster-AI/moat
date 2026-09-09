from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Position import Position
from fintechOnDjango.models.InvestmentPortfolio import InvestmentPortfolio
from fintechOnDjango.models.Security import Security
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Position
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PositionDelegate Declaration
#======================================================================
class PositionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, positionId ):
		try:	
			position = Position.objects.filter(id=positionId)
			return position.first();
		except Position.DoesNotExist:
			raise ProcessingError("Position with id " + str(positionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, position):
		for model in serializers.deserialize("json", position):
			model.save()
			return model;

	def create(self, position):
		position.save()
		return position;

	def saveFromJson(self, position):
		for model in serializers.deserialize("json", position):
			model.save()
			return position;
	
	def save(self, position):
		position.save()
		return position;
	
	def delete(self, positionId ):
		errMsg = "Failed to delete Position from db using id " + str(positionId)
		
		try:
			position = Position.objects.get(id=positionId)
			position.delete()
			return True
		except Position.DoesNotExist:
			raise ProcessingError("Position with id " + str(positionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Position.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Position from db")
		except Exception:
			return None;
		
	def assignPortfolio( self, positionId, portfolioId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentPortfolioDelegate import InvestmentPortfolioDelegate

		errMsg = "Failed to assign element " + str(portfolioId) + " for Portfolio on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# get the InvestmentPortfolio from db
			investmentPortfolio = InvestmentPortfolioDelegate().get(portfolioId).first();
			
			# assign the Portfolio		
			position.portfolio = investmentPortfolio
			
			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(portfolioId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPortfolio( self, positionId ):
		errMsg = "Failed to unassign element " + str(portfolioId) + " for Portfolio on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# assign to None for unassignment
			position.investmentPortfolio = None			

			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSecurity( self, positionId, securityId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SecurityDelegate import SecurityDelegate

		errMsg = "Failed to assign element " + str(securityId) + " for Security on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# get the Security from db
			security = SecurityDelegate().get(securityId).first();
			
			# assign the Security		
			position.security = security
			
			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSecurity( self, positionId ):
		errMsg = "Failed to unassign element " + str(securityId) + " for Security on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# assign to None for unassignment
			position.security = None			

			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
		
