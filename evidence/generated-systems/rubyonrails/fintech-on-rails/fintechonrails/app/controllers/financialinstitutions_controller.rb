class FinancialInstitutionsController < ApplicationController
  def index
    @financialInstitutions = FinancialInstitution.all
  end
 
  def show
    @financialInstitution = FinancialInstitution.find(params[:id])
  end
 
  def new
    @financialInstitution = FinancialInstitution.new
  end
 
  def edit
    @financialInstitution = FinancialInstitution.find(params[:id])
  end
 
  def create
    @financialInstitution = FinancialInstitution.new(financialInstitution_params)
 
    if @financialInstitution.save
      redirect_to financialInstitutions_path
    else
      render 'new'
    end
  end
 
  def update
    @financialInstitution = FinancialInstitution.find(params[:id])
 
    if @financialInstitution.update(financialInstitution_params)
      redirect_to financialInstitutions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @financialInstitution = FinancialInstitution.find(params[:id])
    @financialInstitution.destroy
    redirect_to financialInstitutions_path
  end

 
  private
    def financialInstitution_params
      params.require(:financialInstitution).permit(:name, :legalName, :countryOfIncorporation, :bic, :website)
    end
end