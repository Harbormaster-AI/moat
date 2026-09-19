class TerminationsController < ApplicationController
  def index
    @terminations = Termination.all
  end
 
  def show
    @termination = Termination.find(params[:id])
  end
 
  def new
    @termination = Termination.new
  end
 
  def edit
    @termination = Termination.find(params[:id])
  end
 
  def create
    @termination = Termination.new(termination_params)
 
    if @termination.save
      redirect_to terminations_path
    else
      render 'new'
    end
  end
 
  def update
    @termination = Termination.find(params[:id])
 
    if @termination.update(termination_params)
      redirect_to terminations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @termination = Termination.find(params[:id])
    @termination.destroy
    redirect_to terminations_path
  end

 
  private
    def termination_params
      params.require(:termination).permit(:terminationNumber, :terminationDate, :notes, :eligibleForRehire, :Reason, :Type)
    end
end