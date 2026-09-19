class EnterprisesController < ApplicationController
  def index
    @enterprises = Enterprise.all
  end
 
  def show
    @enterprise = Enterprise.find(params[:id])
  end
 
  def new
    @enterprise = Enterprise.new
  end
 
  def edit
    @enterprise = Enterprise.find(params[:id])
  end
 
  def create
    @enterprise = Enterprise.new(enterprise_params)
 
    if @enterprise.save
      redirect_to enterprises_path
    else
      render 'new'
    end
  end
 
  def update
    @enterprise = Enterprise.find(params[:id])
 
    if @enterprise.update(enterprise_params)
      redirect_to enterprises_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @enterprise = Enterprise.find(params[:id])
    @enterprise.destroy
    redirect_to enterprises_path
  end

 
  private
    def enterprise_params
      params.require(:enterprise).permit(:name, :legalName, :registrationCountry, :website, :taxId)
    end
end