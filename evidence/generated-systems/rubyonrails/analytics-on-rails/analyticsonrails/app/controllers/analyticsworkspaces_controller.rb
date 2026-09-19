class AnalyticsWorkspacesController < ApplicationController
  def index
    @analyticsWorkspaces = AnalyticsWorkspace.all
  end
 
  def show
    @analyticsWorkspace = AnalyticsWorkspace.find(params[:id])
  end
 
  def new
    @analyticsWorkspace = AnalyticsWorkspace.new
  end
 
  def edit
    @analyticsWorkspace = AnalyticsWorkspace.find(params[:id])
  end
 
  def create
    @analyticsWorkspace = AnalyticsWorkspace.new(analyticsWorkspace_params)
 
    if @analyticsWorkspace.save
      redirect_to analyticsWorkspaces_path
    else
      render 'new'
    end
  end
 
  def update
    @analyticsWorkspace = AnalyticsWorkspace.find(params[:id])
 
    if @analyticsWorkspace.update(analyticsWorkspace_params)
      redirect_to analyticsWorkspaces_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @analyticsWorkspace = AnalyticsWorkspace.find(params[:id])
    @analyticsWorkspace.destroy
    redirect_to analyticsWorkspaces_path
  end

 
  private
    def analyticsWorkspace_params
      params.require(:analyticsWorkspace).permit(:name, :businessDomain, :ownerTeam, :GovernanceTier)
    end
end