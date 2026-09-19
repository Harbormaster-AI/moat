class BusinessGlossaryTermsController < ApplicationController
  def index
    @businessGlossaryTerms = BusinessGlossaryTerm.all
  end
 
  def show
    @businessGlossaryTerm = BusinessGlossaryTerm.find(params[:id])
  end
 
  def new
    @businessGlossaryTerm = BusinessGlossaryTerm.new
  end
 
  def edit
    @businessGlossaryTerm = BusinessGlossaryTerm.find(params[:id])
  end
 
  def create
    @businessGlossaryTerm = BusinessGlossaryTerm.new(businessGlossaryTerm_params)
 
    if @businessGlossaryTerm.save
      redirect_to businessGlossaryTerms_path
    else
      render 'new'
    end
  end
 
  def update
    @businessGlossaryTerm = BusinessGlossaryTerm.find(params[:id])
 
    if @businessGlossaryTerm.update(businessGlossaryTerm_params)
      redirect_to businessGlossaryTerms_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @businessGlossaryTerm = BusinessGlossaryTerm.find(params[:id])
    @businessGlossaryTerm.destroy
    redirect_to businessGlossaryTerms_path
  end

 
  private
    def businessGlossaryTerm_params
      params.require(:businessGlossaryTerm).permit(:term, :definition, :steward)
    end
end