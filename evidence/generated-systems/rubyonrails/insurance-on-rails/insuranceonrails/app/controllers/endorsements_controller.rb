class EndorsementsController < ApplicationController
  def index
    @endorsements = Endorsement.all
  end
 
  def show
    @endorsement = Endorsement.find(params[:id])
  end
 
  def new
    @endorsement = Endorsement.new
  end
 
  def edit
    @endorsement = Endorsement.find(params[:id])
  end
 
  def create
    @endorsement = Endorsement.new(endorsement_params)
 
    if @endorsement.save
      redirect_to endorsements_path
    else
      render 'new'
    end
  end
 
  def update
    @endorsement = Endorsement.find(params[:id])
 
    if @endorsement.update(endorsement_params)
      redirect_to endorsements_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @endorsement = Endorsement.find(params[:id])
    @endorsement.destroy
    redirect_to endorsements_path
  end

 
  private
    def endorsement_params
      params.require(:endorsement).permit(:endorsementNumber, :effectiveDate, :description)
    end
end