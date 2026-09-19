class TaxRulesController < ApplicationController
  def index
    @taxRules = TaxRule.all
  end
 
  def show
    @taxRule = TaxRule.find(params[:id])
  end
 
  def new
    @taxRule = TaxRule.new
  end
 
  def edit
    @taxRule = TaxRule.find(params[:id])
  end
 
  def create
    @taxRule = TaxRule.new(taxRule_params)
 
    if @taxRule.save
      redirect_to taxRules_path
    else
      render 'new'
    end
  end
 
  def update
    @taxRule = TaxRule.find(params[:id])
 
    if @taxRule.update(taxRule_params)
      redirect_to taxRules_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @taxRule = TaxRule.find(params[:id])
    @taxRule.destroy
    redirect_to taxRules_path
  end

 
  private
    def taxRule_params
      params.require(:taxRule).permit(:name, :country, :region, :rate, :taxInclusive, :TaxClass)
    end
end