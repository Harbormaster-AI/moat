from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.Visualization import Visualization
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.models.BIQuery import BIQuery
from analyticsOnDjango.models.Tag import Tag
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Dashboard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DashboardDelegate Declaration
#======================================================================
class DashboardDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dashboardId ):
		try:	
			dashboard = Dashboard.objects.filter(id=dashboardId)
			return dashboard.first();
		except Dashboard.DoesNotExist:
			raise ProcessingError("Dashboard with id " + str(dashboardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dashboard):
		for model in serializers.deserialize("json", dashboard):
			model.save()
			return model;

	def create(self, dashboard):
		dashboard.save()
		return dashboard;

	def saveFromJson(self, dashboard):
		for model in serializers.deserialize("json", dashboard):
			model.save()
			return dashboard;
	
	def save(self, dashboard):
		dashboard.save()
		return dashboard;
	
	def delete(self, dashboardId ):
		errMsg = "Failed to delete Dashboard from db using id " + str(dashboardId)
		
		try:
			dashboard = Dashboard.objects.get(id=dashboardId)
			dashboard.delete()
			return True
		except Dashboard.DoesNotExist:
			raise ProcessingError("Dashboard with id " + str(dashboardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Dashboard.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Dashboard from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, dashboardId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on Dashboard"

		try:
			# get the Dashboard from db
			dashboard = self.get( dashboardId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			dashboard.workspace = analyticsWorkspace
			
			#save it
			dashboard.save()

			# reload and return the appropriate version					
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, dashboardId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on Dashboard"

		try:
			# get the Dashboard from db
			dashboard = self.get( dashboardId ).first()	
			
			# assign to None for unassignment
			dashboard.analyticsWorkspace = None			

			#save it
			dashboard.save()

			# reload and return the appropriate version					
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Exception:
			return None;
		
	def addVisualizations( self, dashboardId, visualizationsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

		errMsg = "Failed to add elements " + str(visualizationsIds) + " for Visualizations on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = visualizationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Visualization		
				visualization = VisualizationDelegate().get(id).first();	
				# add the Visualization
				dashboard.visualizations.add(visualization)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVisualizations( self, dashboardId, visualizationsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.VisualizationDelegate import VisualizationDelegate

		errMsg = "Failed to remove elements " + str(visualizationsIds) + " for Visualizations on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = visualizationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Visualization		
				visualization = VisualizationDelegate().get(id).first();	
				# add the Visualization
				dashboard.visualizations.remove(visualization)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReports( self, dashboardId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				dashboard.reports.add(report)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, dashboardId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				dashboard.reports.remove(report)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, dashboardId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dashboard.datasets.add(dataSet)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, dashboardId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dashboard.datasets.remove(dataSet)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAlerts( self, dashboardId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to add elements " + str(alertsIds) + " for Alerts on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				dashboard.alerts.add(alert)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAlerts( self, dashboardId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to remove elements " + str(alertsIds) + " for Alerts on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				dashboard.alerts.remove(alert)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQueries( self, dashboardId, queriesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

		errMsg = "Failed to add elements " + str(queriesIds) + " for Queries on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = queriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BIQuery		
				bIQuery = BIQueryDelegate().get(id).first();	
				# add the BIQuery
				dashboard.queries.add(bIQuery)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQueries( self, dashboardId, queriesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BIQueryDelegate import BIQueryDelegate

		errMsg = "Failed to remove elements " + str(queriesIds) + " for Queries on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = queriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BIQuery		
				bIQuery = BIQueryDelegate().get(id).first();	
				# add the BIQuery
				dashboard.queries.remove(bIQuery)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except BIQuery.DoesNotExist:
			raise ProcessingError(errMsg + " : BIQuery does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTags( self, dashboardId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to add elements " + str(tagsIds) + " for Tags on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				dashboard.tags.add(tag)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTags( self, dashboardId, tagsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TagDelegate import TagDelegate

		errMsg = "Failed to remove elements " + str(tagsIds) + " for Tags on Dashboard"

		try:
			# get the Dashboard
			dashboard = self.get( dashboardId ).first()
				
			# split on a comma with no spaces
			idList = tagsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Tag		
				tag = TagDelegate().get(id).first();	
				# add the Tag
				dashboard.tags.remove(tag)
				
			# save it		
			dashboard.save()
			
			# reload and return the appropriate version
			return self.get( dashboardId );
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
