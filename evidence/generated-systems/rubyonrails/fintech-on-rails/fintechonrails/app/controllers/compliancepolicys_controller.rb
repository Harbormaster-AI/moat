class CompliancePolicysController < ApplicationController
  def index
    @compliancePolicys = CompliancePolicy.all
  end
 
  def show
    @compliancePolicy = CompliancePolicy.find(params[:id])
  end
 
  def new
    @compliancePolicy = CompliancePolicy.new
  end
 
  def edit
    @compliancePolicy = CompliancePolicy.find(params[:id])
  end
 
  def create
    @compliancePolicy = CompliancePolicy.new(compliancePolicy_params)
 
    if @compliancePolicy.save
      redirect_to compliancePolicys_path
    else
      render 'new'
    end
  end
 
  def update
    @compliancePolicy = CompliancePolicy.find(params[:id])
 
    if @compliancePolicy.update(compliancePolicy_params)
      redirect_to compliancePolicys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @compliancePolicy = CompliancePolicy.find(params[:id])
    @compliancePolicy.destroy
    redirect_to compliancePolicys_path
  end

 
  private
    def compliancePolicy_params
      params.require(:compliancePolicy).permit(:name, :policyCode, :description, :Status)
    end
end