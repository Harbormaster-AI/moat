class ComplianceAlertsController < ApplicationController
  def index
    @complianceAlerts = ComplianceAlert.all
  end
 
  def show
    @complianceAlert = ComplianceAlert.find(params[:id])
  end
 
  def new
    @complianceAlert = ComplianceAlert.new
  end
 
  def edit
    @complianceAlert = ComplianceAlert.find(params[:id])
  end
 
  def create
    @complianceAlert = ComplianceAlert.new(complianceAlert_params)
 
    if @complianceAlert.save
      redirect_to complianceAlerts_path
    else
      render 'new'
    end
  end
 
  def update
    @complianceAlert = ComplianceAlert.find(params[:id])
 
    if @complianceAlert.update(complianceAlert_params)
      redirect_to complianceAlerts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @complianceAlert = ComplianceAlert.find(params[:id])
    @complianceAlert.destroy
    redirect_to complianceAlerts_path
  end

 
  private
    def complianceAlert_params
      params.require(:complianceAlert).permit(:alertCode, :raisedAt, :notes, :Severity, :Status)
    end
end